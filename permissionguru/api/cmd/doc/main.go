package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/permissionguru/api/pkg/contproc"
	"github.com/permissionguru/api/pkg/db"
	"github.com/permissionguru/api/pkg/docpars"
	"github.com/permissionguru/api/pkg/subscr"
)

func handler(dbClient db.Databaser, subscrClient subscr.Subscriber, docparsClient docpars.Parser, contprocClient contproc.Processor) func(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return func(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
		ctx := context.Background()

		contentType, ok := req.Headers["content-type"]
		if !ok || !checkSupportedContentType(contentType) {
			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       `{"error": "content type header not provided or not supported"}`,
			}, nil
		}

		if req.Body == "" {
			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       `{"error": "no request body provided"}`,
			}, nil
		}

		bodyBytes, err := base64.StdEncoding.DecodeString(req.Body)
		if err != nil {
			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       `{"error": "error processing request body"}`,
			}, nil
		}

		content, err := docparsClient.Parse(ctx, bodyBytes)
		if err != nil {
			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       `{"error": "error parsing document"}`,
			}, nil
		}

		result := &contproc.Result{}
		accountID, check := req.Headers["x-permission-guru-api-key"]
		if check {
			account, err := dbClient.GetAccount(ctx, accountID)
			if err != nil || account == nil {
				return &events.APIGatewayProxyResponse{
					StatusCode: http.StatusInternalServerError,
					Body:       `{"error": "error getting account"}`,
				}, nil
			}

			if err := subscrClient.AddUsage(ctx, account.ID); err != nil {
				return &events.APIGatewayProxyResponse{
					StatusCode: http.StatusInternalServerError,
					Body:       `{"error": "error adding account usage"}`,
				}, nil
			}

			result, err = contprocClient.Process(ctx, *content, account.Config)
			if err != nil {
				return &events.APIGatewayProxyResponse{
					StatusCode: http.StatusInternalServerError,
					Body:       `{"error": "error processing content"}`,
				}, nil
			}
		} else {
			var err error

			result, err = contprocClient.Process(ctx, *content, nil)
			if err != nil {
				return &events.APIGatewayProxyResponse{
					StatusCode: http.StatusInternalServerError,
					Body:       `{"error": "error processing content"}`,
				}, nil
			}
		}

		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       fmt.Sprintf(`{"message": "success", "is_permissioned": %t}`, result.IsPermissioned),
		}, nil
	}
}

func checkSupportedContentType(contentType string) bool {
	supportedTypes := []string{
		"image/jpeg",
		"image/png",
		"application/pdf",
	}

	for _, supportedType := range supportedTypes {
		if supportedType == contentType {
			return true
		}
	}

	return false
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
	docparsClient := docpars.New()
	contprocClient := contproc.New()

	lambda.Start(handler(dbClient, subscrClient, docparsClient, contprocClient))
}
