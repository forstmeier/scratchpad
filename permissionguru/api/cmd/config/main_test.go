package main

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/permissionguru/api/pkg/acct"
	"github.com/permissionguru/api/pkg/config"
	"github.com/permissionguru/api/pkg/subscr"
)

type mockDatabaser struct {
	addConfigError error
}

func (m *mockDatabaser) AddAccount(ctx context.Context, id string) error {
	return nil
}

func (m *mockDatabaser) AddSubscription(ctx context.Context, id string, sub subscr.Subscription) error {
	return nil
}

func (m *mockDatabaser) AddConfig(ctx context.Context, id string, cfg config.Config) error {
	return m.addConfigError
}

func (m *mockDatabaser) GetAccount(ctx context.Context, id string) (*acct.Account, error) {
	return nil, nil
}

func (m *mockDatabaser) RemoveAccount(ctx context.Context, id string) error {
	return nil
}

func Test_handler(t *testing.T) {
	tests := []struct {
		description    string
		request        events.APIGatewayProxyRequest
		addConfigError error
		statusCode     int
		body           string
	}{
		{
			description: "no api key header received",
			request: events.APIGatewayProxyRequest{
				Headers: map[string]string{},
			},
			addConfigError: nil,
			statusCode:     http.StatusBadRequest,
			body:           `{"error": "no api key header received"}`,
		},
		{
			description: "error unmarshalling request body",
			request: events.APIGatewayProxyRequest{
				Headers: map[string]string{
					"x-permission-guru-api-key": "api_key",
				},
				Body: "---------",
			},
			addConfigError: nil,
			statusCode:     http.StatusInternalServerError,
			body:           `{"error": "error processing request body"}`,
		},
		{
			description: "error unmarshalling request body",
			request: events.APIGatewayProxyRequest{
				Headers: map[string]string{
					"x-permission-guru-api-key": "api_key",
				},
				Body: `{
					"selections": [
						{
							"text": "example text",
							"selected": true
						}
					]
				}`,
			},
			addConfigError: errors.New("mock add config error"),
			statusCode:     http.StatusInternalServerError,
			body:           `{"error": "error adding config"}`,
		},
		{
			description: "successful add config handler call",
			request: events.APIGatewayProxyRequest{
				Headers: map[string]string{
					"x-permission-guru-api-key": "api_key",
				},
				Body: `{
					"selections": [
						{
							"text": "example text",
							"selected": true
						}
					]
				}`,
			},
			addConfigError: nil,
			statusCode:     http.StatusOK,
			body:           `{"message": "success"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			dbClient := &mockDatabaser{
				addConfigError: test.addConfigError,
			}

			h := handler(dbClient)

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
