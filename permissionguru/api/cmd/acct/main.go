package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"

	"github.com/permissionguru/api/pkg/db"
	"github.com/permissionguru/api/pkg/subscr"
)

func handler(dbClient db.Databaser, subscrClient subscr.Subscriber) func(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return func(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
		httpMethod := strings.ToLower(req.HTTPMethod)

		if httpMethod == "post" {
			ctx := context.Background()
			accountID := uuid.NewString()

			subscriberInfo := subscr.SubscriberInfo{}
			if err := json.Unmarshal([]byte(req.Body), &subscriberInfo); err != nil {
				return &events.APIGatewayProxyResponse{
					StatusCode: http.StatusBadRequest,
					Body:       `{"error": "error processing request body"}`,
				}, nil
			}
			subscriberInfo.ID = uuid.NewString()

			if err := dbClient.AddAccount(ctx, accountID); err != nil {
				return &events.APIGatewayProxyResponse{
					StatusCode: http.StatusBadRequest,
					Body:       `{"error": "error adding account id"}`,
				}, nil
			}

			subscription, err := subscrClient.CreateSubscription(ctx, subscriberInfo)
			if err != nil {
				return &events.APIGatewayProxyResponse{
					StatusCode: http.StatusInternalServerError,
					Body:       `{"error": "error creating subscription"}`,
				}, nil
			}

			if err := dbClient.AddSubscription(ctx, accountID, *subscription); err != nil {
				return &events.APIGatewayProxyResponse{
					StatusCode: http.StatusInternalServerError,
					Body:       `{"error": "error adding subscription"}`,
				}, nil
			}

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusOK,
				Body:       fmt.Sprintf(`{"message": "success", "account_id": "%s"}`, accountID),
			}, nil

		} else if httpMethod == "delete" {
			ctx := context.Background()

			accountID, ok := req.Headers["x-permission-guru-api-key"]
			if !ok {
				return &events.APIGatewayProxyResponse{
					StatusCode: http.StatusBadRequest,
					Body:       `{"error": "no api key header received"}`,
				}, nil
			}

			if err := subscrClient.RemoveSubscription(ctx, accountID); err != nil {
				return &events.APIGatewayProxyResponse{
					StatusCode: http.StatusInternalServerError,
					Body:       `{"error": "error removing subscription"}`,
				}, nil
			}

			if err := dbClient.RemoveAccount(ctx, accountID); err != nil {
				return &events.APIGatewayProxyResponse{
					StatusCode: http.StatusInternalServerError,
					Body:       `{"error": "error removing account"}`,
				}, nil
			}

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusOK,
				Body:       fmt.Sprintf(`{"message": "success", "account_id": "%s"}`, accountID),
			}, nil
		}

		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       fmt.Sprintf(`{"error": "http method %s not supported"}`, httpMethod),
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
	subscrClient := subscr.New()

	lambda.Start(handler(dbClient, subscrClient))
}
