package response

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/pkg/errors"
)

func JSON(code int, body interface{}) (events.APIGatewayProxyResponse, error) {
	bytes, err := json.Marshal(&body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, errors.Wrap(err, "response: JSON json.Marshal error")
	}

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		StatusCode: code,
		Body:       string(bytes),
	}, nil
}

func JSONErr(code int, err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		StatusCode: code,
		Body:       string(fmt.Sprintf(`{"error":%s}`, err)),
	}, nil

}

func Plain(code int, body string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Body:       body,
	}, nil
}
