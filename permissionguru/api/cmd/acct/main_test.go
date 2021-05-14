package main

import (
	"context"
	"errors"
	"net/http"
	"regexp"
	"testing"

	"github.com/aws/aws-lambda-go/events"

	"github.com/permissionguru/api/pkg/acct"
	"github.com/permissionguru/api/pkg/config"
	"github.com/permissionguru/api/pkg/subscr"
)

type mockDatabaser struct {
	addAccountError      error
	addSubscriptionError error
	removeAccountError   error
}

func (m *mockDatabaser) AddAccount(ctx context.Context, id string) error {
	return m.addAccountError
}

func (m *mockDatabaser) AddSubscription(ctx context.Context, id string, sub subscr.Subscription) error {
	return m.addSubscriptionError
}

func (m *mockDatabaser) AddConfig(ctx context.Context, id string, cfg config.Config) error {
	return nil
}

func (m *mockDatabaser) GetAccount(ctx context.Context, id string) (*acct.Account, error) {
	return nil, nil
}

func (m *mockDatabaser) RemoveAccount(ctx context.Context, id string) error {
	return m.removeAccountError
}

type mockSubscriber struct {
	createSubscriptionOutput *subscr.Subscription
	createSubscriptionError  error
	removeSubscriptionError  error
}

func (m *mockSubscriber) CreateSubscription(ctx context.Context, info subscr.SubscriberInfo) (*subscr.Subscription, error) {
	return m.createSubscriptionOutput, m.createSubscriptionError
}

func (m *mockSubscriber) RemoveSubscription(ctx context.Context, id string) error {
	return m.removeSubscriptionError
}

func (m *mockSubscriber) AddUsage(ctx context.Context, id string) error {
	return nil
}

func Test_handler(t *testing.T) {
	tests := []struct {
		description              string
		request                  events.APIGatewayProxyRequest
		addAccountError          error
		createSubscriptionOutput *subscr.Subscription
		createSubscriptionError  error
		addSubscriptionError     error
		removeSubscriptionError  error
		removeAccountError       error
		statusCode               int
		bodyRegexp               string
	}{
		{
			description: "error unmarshalling request body",
			request: events.APIGatewayProxyRequest{
				HTTPMethod: "POST",
				Body:       "---------",
			},
			addAccountError:          nil,
			createSubscriptionOutput: nil,
			createSubscriptionError:  nil,
			addSubscriptionError:     nil,
			removeSubscriptionError:  nil,
			removeAccountError:       nil,
			statusCode:               http.StatusBadRequest,
			bodyRegexp:               `{"error": "error processing request body"}`,
		},
		{
			description: "error adding id to dynamodb",
			request: events.APIGatewayProxyRequest{
				HTTPMethod: "POST",
				Body: `{
					"email": "test@email.com",
					"zip": "12345",
					"expiration_month": "05",
					"expiration_year": "1977",
					"card_number": "000000000000",
					"card_security_code": "000"
				}`,
			},
			addAccountError:          errors.New("mock add account error"),
			createSubscriptionOutput: nil,
			createSubscriptionError:  nil,
			addSubscriptionError:     nil,
			removeSubscriptionError:  nil,
			removeAccountError:       nil,
			statusCode:               http.StatusBadRequest,
			bodyRegexp:               `{"error": "error adding account id"}`,
		},
		{
			description: "error creating subscription",
			request: events.APIGatewayProxyRequest{
				HTTPMethod: "POST",
				Body: `{
					"email": "test@email.com",
					"zip": "12345",
					"expiration_month": "05",
					"expiration_year": "1977",
					"card_number": "000000000000",
					"card_security_code": "000"
				}`,
			},
			addAccountError:          nil,
			createSubscriptionOutput: nil,
			createSubscriptionError:  errors.New("mock create subscription error"),
			addSubscriptionError:     nil,
			removeSubscriptionError:  nil,
			removeAccountError:       nil,
			statusCode:               http.StatusInternalServerError,
			bodyRegexp:               `{"error": "error creating subscription"}`,
		},
		{
			description: "error adding subscription to dynamodb",
			request: events.APIGatewayProxyRequest{
				HTTPMethod: "POST",
				Body: `{
					"email": "test@email.com",
					"zip": "12345",
					"expiration_month": "05",
					"expiration_year": "1977",
					"card_number": "000000000000",
					"card_security_code": "000"
				}`,
			},
			addAccountError:          nil,
			createSubscriptionOutput: &subscr.Subscription{},
			createSubscriptionError:  nil,
			addSubscriptionError:     errors.New("mock add subscription error"),
			removeSubscriptionError:  nil,
			removeAccountError:       nil,
			statusCode:               http.StatusInternalServerError,
			bodyRegexp:               `{"error": "error adding subscription"}`,
		},
		{
			description: "successful new account handler call",
			request: events.APIGatewayProxyRequest{
				HTTPMethod: "POST",
				Body: `{
					"email": "test@email.com",
					"zip": "12345",
					"expiration_month": "05",
					"expiration_year": "1977",
					"card_number": "000000000000",
					"card_security_code": "000"
				}`,
			},
			addAccountError:          nil,
			createSubscriptionOutput: &subscr.Subscription{},
			createSubscriptionError:  nil,
			addSubscriptionError:     nil,
			removeSubscriptionError:  nil,
			removeAccountError:       nil,
			statusCode:               http.StatusOK,
			bodyRegexp:               `{"message": "success", "account_id": ".*"}`,
		},
		{
			description: "error missing api key header",
			request: events.APIGatewayProxyRequest{
				HTTPMethod: "DELETE",
			},
			addAccountError:          nil,
			createSubscriptionOutput: nil,
			createSubscriptionError:  nil,
			addSubscriptionError:     nil,
			removeSubscriptionError:  nil,
			removeAccountError:       nil,
			statusCode:               http.StatusBadRequest,
			bodyRegexp:               `{"error": "no api key header received"}`,
		},
		{
			description: "error removing subscription",
			request: events.APIGatewayProxyRequest{
				HTTPMethod: "DELETE",
				Headers: map[string]string{
					"x-permission-guru-api-key": "test-key",
				},
			},
			addAccountError:          nil,
			createSubscriptionOutput: nil,
			createSubscriptionError:  nil,
			addSubscriptionError:     nil,
			removeSubscriptionError:  errors.New("mock remove subscription error"),
			removeAccountError:       nil,
			statusCode:               http.StatusInternalServerError,
			bodyRegexp:               `{"error": "error removing subscription"}`,
		},
		{
			description: "error removing account from dynamodb",
			request: events.APIGatewayProxyRequest{
				HTTPMethod: "DELETE",
				Headers: map[string]string{
					"x-permission-guru-api-key": "test-key",
				},
			},
			addAccountError:          nil,
			createSubscriptionOutput: nil,
			createSubscriptionError:  nil,
			addSubscriptionError:     nil,
			removeSubscriptionError:  nil,
			removeAccountError:       errors.New("mock remove account error"),
			statusCode:               http.StatusInternalServerError,
			bodyRegexp:               `{"error": "error removing account"}`,
		},
		{
			description: "successful remove account handler call",
			request: events.APIGatewayProxyRequest{
				HTTPMethod: "DELETE",
				Headers: map[string]string{
					"x-permission-guru-api-key": "test-key",
				},
			},
			addAccountError:          nil,
			createSubscriptionOutput: nil,
			createSubscriptionError:  nil,
			addSubscriptionError:     nil,
			removeSubscriptionError:  nil,
			removeAccountError:       nil,
			statusCode:               http.StatusOK,
			bodyRegexp:               `{"message": "success", "account_id": "test-key"}`,
		},
		{
			description: "error http method not supported",
			request: events.APIGatewayProxyRequest{
				HTTPMethod: "GET",
			},
			addAccountError:          nil,
			createSubscriptionOutput: nil,
			createSubscriptionError:  nil,
			addSubscriptionError:     nil,
			removeSubscriptionError:  nil,
			removeAccountError:       nil,
			statusCode:               http.StatusBadRequest,
			bodyRegexp:               `{"error": "http method get not supported"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			dbClient := &mockDatabaser{
				addAccountError:      test.addAccountError,
				addSubscriptionError: test.addSubscriptionError,
				removeAccountError:   test.removeAccountError,
			}

			subscrClient := &mockSubscriber{
				createSubscriptionOutput: test.createSubscriptionOutput,
				createSubscriptionError:  test.createSubscriptionError,
				removeSubscriptionError:  test.removeSubscriptionError,
			}

			h := handler(dbClient, subscrClient)

			response, _ := h(test.request)

			if response.StatusCode != test.statusCode {
				t.Errorf("incorrect status code, received: %d, expected: %d", response.StatusCode, test.statusCode)
			}

			matched, err := regexp.MatchString(test.bodyRegexp, response.Body)
			if err != nil {
				t.Fatalf("error matching body regexp: %s", err.Error())
			}

			if !matched {
				t.Errorf("incorrect body, received: %s", response.Body)
			}
		})
	}
}
