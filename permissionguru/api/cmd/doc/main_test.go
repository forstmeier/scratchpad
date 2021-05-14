package main

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"

	"github.com/permissionguru/api/pkg/acct"
	"github.com/permissionguru/api/pkg/config"
	"github.com/permissionguru/api/pkg/contproc"
	"github.com/permissionguru/api/pkg/docpars"
	"github.com/permissionguru/api/pkg/subscr"
)

type mockDatabaser struct {
	getAccountOutput *acct.Account
	getAccountError  error
}

func (m *mockDatabaser) AddAccount(ctx context.Context, id string) error {
	return nil
}

func (m *mockDatabaser) AddSubscription(ctx context.Context, id string, sub subscr.Subscription) error {
	return nil
}

func (m *mockDatabaser) AddConfig(ctx context.Context, id string, cfg config.Config) error {
	return nil
}

func (m *mockDatabaser) GetAccount(ctx context.Context, id string) (*acct.Account, error) {
	return m.getAccountOutput, m.getAccountError
}

func (m *mockDatabaser) RemoveAccount(ctx context.Context, id string) error {
	return nil
}

type mockSubscriber struct {
	addUsageError error
}

func (m *mockSubscriber) CreateSubscription(ctx context.Context, info subscr.SubscriberInfo) (*subscr.Subscription, error) {
	return nil, nil
}

func (m *mockSubscriber) RemoveSubscription(ctx context.Context, id string) error {
	return nil
}

func (m *mockSubscriber) AddUsage(ctx context.Context, id string) error {
	return m.addUsageError
}

type mockParser struct {
	parseOutput *docpars.Content
	parseError  error
}

func (m *mockParser) Parse(ctx context.Context, doc []byte) (*docpars.Content, error) {
	return m.parseOutput, m.parseError
}

type mockProcessor struct {
	processOutput *contproc.Result
	processError  error
}

func (m *mockProcessor) Process(ctx context.Context, content docpars.Content, cfg *config.Config) (*contproc.Result, error) {
	return m.processOutput, m.processError
}

func Test_handler(t *testing.T) {
	tests := []struct {
		description      string
		request          events.APIGatewayProxyRequest
		parseOutput      *docpars.Content
		parseError       error
		getAccountOutput *acct.Account
		getAccountError  error
		addUsageError    error
		processOutput    *contproc.Result
		processError     error
		statusCode       int
		body             string
	}{
		{
			description: "incorrect content type header",
			request: events.APIGatewayProxyRequest{
				Headers: map[string]string{
					"x-permission-guru-api-key": "api_key",
				},
			},
			parseOutput:      nil,
			parseError:       nil,
			getAccountOutput: nil,
			getAccountError:  nil,
			addUsageError:    nil,
			processOutput:    nil,
			processError:     nil,
			statusCode:       http.StatusBadRequest,
			body:             `{"error": "content type header not provided or not supported"}`,
		},
		{
			description: "empty request body",
			request: events.APIGatewayProxyRequest{
				Headers: map[string]string{
					"x-permission-guru-api-key": "api_key",
					"content-type":              "image/png",
				},
				Body: "",
			},
			parseOutput:      nil,
			parseError:       nil,
			getAccountOutput: nil,
			getAccountError:  nil,
			addUsageError:    nil,
			processOutput:    nil,
			processError:     nil,
			statusCode:       http.StatusBadRequest,
			body:             `{"error": "no request body provided"}`,
		},
		{
			description: "non-base64 encoded request body",
			request: events.APIGatewayProxyRequest{
				Headers: map[string]string{
					"x-permission-guru-api-key": "api_key",
					"content-type":              "image/png",
				},
				Body: "not-base64",
			},
			parseOutput:      nil,
			parseError:       nil,
			getAccountOutput: nil,
			getAccountError:  nil,
			addUsageError:    nil,
			processOutput:    nil,
			processError:     nil,
			statusCode:       http.StatusInternalServerError,
			body:             `{"error": "error processing request body"}`,
		},
		{
			description: "error parsing document",
			request: events.APIGatewayProxyRequest{
				Headers: map[string]string{
					"x-permission-guru-api-key": "api_key",
					"content-type":              "image/png",
				},
				Body: "bm90IGFuIGltYWdlIGZpbGU=",
			},
			parseOutput:      nil,
			parseError:       errors.New("mock parse error"),
			getAccountOutput: nil,
			getAccountError:  nil,
			addUsageError:    nil,
			processOutput:    nil,
			processError:     nil,
			statusCode:       http.StatusInternalServerError,
			body:             `{"error": "error parsing document"}`,
		},
		{
			description: "no api key and error processing content",
			request: events.APIGatewayProxyRequest{
				Headers: map[string]string{
					"content-type": "image/png",
				},
				Body: "bm90IGFuIGltYWdlIGZpbGU=",
			},
			parseOutput:      &docpars.Content{},
			parseError:       nil,
			getAccountOutput: nil,
			getAccountError:  nil,
			processOutput:    nil,
			processError:     errors.New("mock process error"),
			statusCode:       http.StatusInternalServerError,
			body:             `{"error": "error processing content"}`,
		},
		{
			description: "no api key and successful doc parsing/processing handler call",
			request: events.APIGatewayProxyRequest{
				Headers: map[string]string{
					"content-type": "image/png",
				},
				Body: "bm90IGFuIGltYWdlIGZpbGU=",
			},
			parseOutput:      &docpars.Content{},
			parseError:       nil,
			getAccountOutput: nil,
			getAccountError:  nil,
			processOutput: &contproc.Result{
				IsPermissioned: true,
			},
			processError: nil,
			statusCode:   http.StatusOK,
			body:         `{"message": "success", "is_permissioned": true}`,
		},
		{
			description: "api key and error getting account",
			request: events.APIGatewayProxyRequest{
				Headers: map[string]string{
					"x-permission-guru-api-key": "api_key",
					"content-type":              "image/png",
				},
				Body: "bm90IGFuIGltYWdlIGZpbGU=",
			},
			parseOutput:      &docpars.Content{},
			parseError:       nil,
			getAccountOutput: nil,
			getAccountError:  errors.New("mock get account error"),
			addUsageError:    nil,
			processOutput:    nil,
			processError:     nil,
			statusCode:       http.StatusInternalServerError,
			body:             `{"error": "error getting account"}`,
		},
		{
			description: "api key and error adding account usage",
			request: events.APIGatewayProxyRequest{
				Headers: map[string]string{
					"x-permission-guru-api-key": "api_key",
					"content-type":              "image/png",
				},
				Body: "bm90IGFuIGltYWdlIGZpbGU=",
			},
			parseOutput:      &docpars.Content{},
			parseError:       nil,
			getAccountOutput: &acct.Account{},
			getAccountError:  nil,
			addUsageError:    errors.New("mock add account usage error"),
			processOutput:    nil,
			processError:     nil,
			statusCode:       http.StatusInternalServerError,
			body:             `{"error": "error adding account usage"}`,
		},
		{
			description: "api key and error processing content",
			request: events.APIGatewayProxyRequest{
				Headers: map[string]string{
					"x-permission-guru-api-key": "api_key",
					"content-type":              "image/png",
				},
				Body: "bm90IGFuIGltYWdlIGZpbGU=",
			},
			parseOutput:      &docpars.Content{},
			parseError:       nil,
			getAccountOutput: &acct.Account{},
			getAccountError:  nil,
			addUsageError:    nil,
			processOutput:    nil,
			processError:     errors.New("mock process error"),
			statusCode:       http.StatusInternalServerError,
			body:             `{"error": "error processing content"}`,
		},
		{
			description: "api key and successful doc parsing/processing handler call",
			request: events.APIGatewayProxyRequest{
				Headers: map[string]string{
					"x-permission-guru-api-key": "api_key",
					"content-type":              "image/png",
				},
				Body: "bm90IGFuIGltYWdlIGZpbGU=",
			},
			parseOutput:      &docpars.Content{},
			parseError:       nil,
			getAccountOutput: &acct.Account{},
			getAccountError:  nil,
			addUsageError:    nil,
			processOutput: &contproc.Result{
				IsPermissioned: true,
			},
			processError: nil,
			statusCode:   http.StatusOK,
			body:         `{"message": "success", "is_permissioned": true}`,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			dbClient := &mockDatabaser{
				getAccountOutput: test.getAccountOutput,
				getAccountError:  test.getAccountError,
			}
			subscrClient := &mockSubscriber{
				addUsageError: test.addUsageError,
			}
			docparsClient := &mockParser{
				parseOutput: test.parseOutput,
				parseError:  test.parseError,
			}
			contprocClient := &mockProcessor{
				processOutput: test.processOutput,
				processError:  test.processError,
			}

			h := handler(dbClient, subscrClient, docparsClient, contprocClient)

			response, _ := h(test.request)

			if response.StatusCode != test.statusCode {
				t.Errorf("incorrect status code, received: %d, expected: %d", response.StatusCode, test.statusCode)
			}

			if response.Body != test.body {
				t.Errorf("incorrect body, received: %s, expected: %s", response.Body, test.body)
			}
		})
	}
}

func Test_checkSupportedContentType(t *testing.T) {
	tests := []struct {
		description string
		contentType string
		output      bool
	}{
		{
			description: "type not supported",
			contentType: "not-supported",
			output:      false,
		},
		{
			description: "type supported",
			contentType: "image/jpeg",
			output:      true,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			if output := checkSupportedContentType(test.contentType); output != test.output {
				t.Errorf("incorrect output, received: %t, expected: %t", output, test.output)
			}
		})
	}
}
