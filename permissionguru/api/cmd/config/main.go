package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/permissionguru/api/pkg/config"
	"github.com/permissionguru/api/pkg/db"
)

func handler(dbClient db.Databaser) func(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return func(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
		ctx := context.Background()

		accountID, ok := req.Headers["x-permission-guru-api-key"]
		if !ok {
			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       `{"error": "no api key header received"}`,
			}, nil
		}

		cfg := config.Config{}
		if err := json.Unmarshal([]byte(req.Body), &cfg); err != nil {
			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       `{"error": "error processing request body"}`,
			}, nil
		}

		if err := dbClient.AddConfig(ctx, accountID, cfg); err != nil {
			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       `{"error": "error adding config"}`,
			}, nil
		}

		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       `{"message": "success"}`,
		}, nil
	}
}

func main() {
	dbService := &dynamodb.DynamoDB{}
	if os.Getenv("ENVIRONMENT_NAME") == "prod" {
		dbService = dynamodb.New(session.New())
	} else {
		dbServiceConfig := &aws.Config{
			Endpoint: aws.String("http://localhost:8000"),
		}
		dbService = dynamodb.New(session.New(), dbServiceConfig)
	}

	dbClient := db.New(dbService)

	lambda.Start(handler(dbClient))
}
