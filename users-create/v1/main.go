package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/back/example/pkg/model"
	"github.com/back/example/pkg/response"
	"github.com/back/example/pkg/uuid"
	"github.com/back/example/users-create/v1/repository"
)

type lambdaHandler func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

type usersRepository interface {
	Create(user model.User) error
}

type uuidWrapper interface {
	New() string
}

type request struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_nuber"`
}

func Adapter(usersRepository usersRepository, uuidWrapper uuidWrapper) lambdaHandler {
	return func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		body := request{}
		err := json.Unmarshal([]byte(req.Body), &body)
		if err != nil {
			return response.JSONErr(http.StatusBadRequest, err)
		}

		err = usersRepository.Create(model.User{
			ID:          uuidWrapper.New(),
			Name:        body.Name,
			Email:       body.Email,
			PhoneNumber: body.PhoneNumber,
		})
		if err != nil {
			return response.JSONErr(http.StatusInternalServerError, err)
		}

		return response.Plain(http.StatusOK, "sucess")
	}
}

func main() {
	session := session.New()
	dynamodb := dynamodb.New(session)

	usersRepository := repository.NewUserRepository(dynamodb, "users")
	uuidWrapper := uuid.NewUUID()

	handler := Adapter(usersRepository, uuidWrapper)

	lambda.Start(handler)
}
