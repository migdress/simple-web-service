package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/back/example/pkg/model"
	"github.com/pkg/errors"
)

const (
	keyUserID          = "id"
	keyUserName        = "name"
	keyUserPhoneNumber = "phone_number"
	keyUserEmail       = "email"
)

type UserRepository struct {
	dynamodbClient dynamodbClient
	table          string
}

func NewUserRepository(dynamodbClient dynamodbClient, tableName string) *UserRepository {
	return &UserRepository{
		dynamodbClient,
		tableName,
	}
}

type dynamodbClient interface {
	PutItem(item *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
}

func (r *UserRepository) Create(user model.User) error {
	_, err := r.dynamodbClient.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(r.table),
		Item: map[string]*dynamodb.AttributeValue{
			keyUserID: {
				S: aws.String(user.ID),
			},
			keyUserName: {
				S: aws.String(user.Name),
			},
			keyUserPhoneNumber: {
				S: aws.String(user.PhoneNumber),
			},
			keyUserEmail: {
				S: aws.String(user.Email),
			},
		},
	})
	if err != nil {
		return errors.Wrap(err, "repository: userRepository.Create dynamodbClient.PutItem error")
	}

	return nil
}
