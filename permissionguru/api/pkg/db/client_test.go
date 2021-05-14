package db

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/permissionguru/api/pkg/acct"
	"github.com/permissionguru/api/pkg/config"
	"github.com/permissionguru/api/pkg/subscr"
)

func TestNew(t *testing.T) {
	service := dynamodb.New(session.New())

	client := New(service)
	if client == nil {
		t.Error("error creating db client")
	}
}

type mockDynamoDBClient struct {
	putItemError    error
	updateItemError error
	getItemOutput   *dynamodb.GetItemOutput
	getItemError    error
	deleteItemError error
}

func (m *mockDynamoDBClient) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return nil, m.putItemError
}

func (m *mockDynamoDBClient) UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	return nil, m.updateItemError
}

func (m *mockDynamoDBClient) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return m.getItemOutput, m.getItemError
}

func (m *mockDynamoDBClient) DeleteItem(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	return nil, m.deleteItemError
}

func TestAddAccount(t *testing.T) {
	tests := []struct {
		description  string
		accountID    string
		putItemError error
		error        error
	}{
		{
			description:  "missing account id",
			accountID:    "",
			putItemError: nil,
			error:        &ErrorMissingID{},
		},
		{
			description:  "dynamodb put item error",
			accountID:    "test_account_id",
			putItemError: errors.New("mock put error"),
			error:        &ErrorPutItem{},
		},
		{
			description:  "successful add account to dynamodb",
			accountID:    "test_account_id",
			putItemError: nil,
			error:        nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			client := &Client{
				dynamoDBClient: &mockDynamoDBClient{
					putItemError: test.putItemError,
				},
			}

			err := client.AddAccount(context.Background(), test.accountID)

			if err != nil {
				switch test.error.(type) {
				case *ErrorMissingID:
					var testError *ErrorMissingID
					if !errors.As(err, &testError) {
						t.Errorf("incorrect error, received: %v, expected: %v", err, testError)
					}
				case *ErrorPutItem:
					var testError *ErrorPutItem
					if !errors.As(err, &testError) {
						t.Errorf("incorrect error, received: %v, expected: %v", err, testError)
					}
				default:
					t.Fatalf("unexpected error type: %v", err)
				}
			} else {
				if err != test.error {
					t.Errorf("incorrect nil error, received: %v, expected: %v", err, test.error)
				}
			}
		})
	}
}

func TestAddSubscription(t *testing.T) {
	tests := []struct {
		description     string
		subscription    subscr.Subscription
		updateItemError error
		error           error
	}{
		{
			description:     "missing field values from subscription",
			updateItemError: nil,
			subscription:    subscr.Subscription{},
			error:           &ErrorMissingFields{},
		},
		{
			description: "dynamodb update item error",
			subscription: subscr.Subscription{
				ID:                       "test_subscription_id",
				StripePaymentMethodID:    "test_stripe_payment_method_id",
				StripeCustomerID:         "test_stripe_customer_id",
				StripeSubscriptionID:     "test_stripe_subscription_id",
				StripeSubscriptionItemID: "test_stripe_subscription_item_id",
			},
			updateItemError: errors.New("mock update item error"),
			error:           &ErrorUpdateItem{},
		},
		{
			description: "successful add subscription call",
			subscription: subscr.Subscription{
				ID:                       "test_subscription_id",
				StripePaymentMethodID:    "test_stripe_payment_method_id",
				StripeCustomerID:         "test_stripe_customer_id",
				StripeSubscriptionID:     "test_stripe_subscription_id",
				StripeSubscriptionItemID: "test_stripe_subscription_item_id",
			},
			updateItemError: nil,
			error:           nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			client := &Client{
				dynamoDBClient: &mockDynamoDBClient{
					updateItemError: test.updateItemError,
				},
			}

			err := client.AddSubscription(context.Background(), "account_id", test.subscription)

			if err != nil {
				switch test.error.(type) {
				case *ErrorMissingFields:
					var testError *ErrorMissingFields
					if !errors.As(err, &testError) {
						t.Errorf("incorrect error, received: %v, expected: %v", err, testError)
					}
				case *ErrorUpdateItem:
					var testError *ErrorUpdateItem
					if !errors.As(err, &testError) {
						t.Errorf("incorrect error, received: %v, expected: %v", err, testError)
					}
				default:
					t.Fatalf("unexpected error type: %v", err)
				}
			} else {
				if err != test.error {
					t.Errorf("incorrect nil error, received: %v, expected: %v", err, test.error)
				}
			}
		})
	}
}

func Test_checkSubscriptionFields(t *testing.T) {
	tests := []struct {
		description  string
		subscription subscr.Subscription
		fields       []string
	}{
		{
			description:  "missing all fields",
			subscription: subscr.Subscription{},
			fields: []string{
				"subscription id",
				"stripe payment method id",
				"stripe customer id",
				"stripe subscription id",
				"stripe subscription item id",
			},
		},
		{
			description: "missing no fields",
			subscription: subscr.Subscription{
				ID:                       "account_id",
				StripePaymentMethodID:    "stripe_payment_id",
				StripeCustomerID:         "stripe_customer_id",
				StripeSubscriptionID:     "stripe_subscription_id",
				StripeSubscriptionItemID: "stripe_subscription_item_id",
			},
			fields: []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			fields := checkSubscriptionFields(test.subscription)

			for i, field := range fields {
				if field != test.fields[i] {
					t.Errorf("incorrect fields, received: %v, expected: %v", fields, test.fields)
					break // keep to ensure all test scenarios run
				}
			}
		})
	}
}

func TestAddConfig(t *testing.T) {
	tests := []struct {
		description     string
		config          config.Config
		updateItemError error
		error           error
	}{
		{
			description:     "missing config selections list",
			config:          config.Config{},
			updateItemError: nil,
			error:           &ErrorNoConfigSelections{},
		},
		{
			description: "dynamodb update item error",
			config: config.Config{
				ID: "test_config_id",
				Selections: []config.Selection{
					{
						Text:     "selection text",
						Selected: true,
					},
				},
			},
			updateItemError: errors.New("mock update error"),
			error:           &ErrorUpdateItem{},
		},
		{
			description: "successful add config call",
			config: config.Config{
				ID: "test_config_id",
				Selections: []config.Selection{
					{
						Text:     "selection text",
						Selected: true,
					},
				},
			},
			updateItemError: nil,
			error:           nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			client := &Client{
				dynamoDBClient: &mockDynamoDBClient{
					updateItemError: test.updateItemError,
				},
			}

			err := client.AddConfig(context.Background(), "account_id", test.config)

			if err != nil {
				switch test.error.(type) {
				case *ErrorNoConfigSelections:
					var testError *ErrorNoConfigSelections
					if !errors.As(err, &testError) {
						t.Errorf("incorrect error, received: %v, expected: %v", err, testError)
					}
				case *ErrorUpdateItem:
					var testError *ErrorUpdateItem
					if !errors.As(err, &testError) {
						t.Errorf("incorrect error, received: %v, expected: %v", err, testError)
					}
				default:
					t.Fatalf("unexpected error type: %v", err)
				}
			} else {
				if err != test.error {
					t.Errorf("incorrect nil error, received: %v, expected: %v", err, test.error)
				}
			}
		})
	}
}

func TestGetAccount(t *testing.T) {
	tests := []struct {
		description   string
		getItemOutput *dynamodb.GetItemOutput
		getItemError  error
		account       *acct.Account
		error         error
	}{
		{
			description:   "dynamodb get item error",
			getItemOutput: nil,
			getItemError:  errors.New("mock get error"),
			account:       nil,
			error:         &ErrorGetItem{},
		},
		{
			description: "no account value in dynamodb",
			getItemOutput: &dynamodb.GetItemOutput{
				Item: map[string]*dynamodb.AttributeValue{},
			},
			getItemError: nil,
			account:      nil,
			error:        &ErrorNoAccount{},
		},
		{
			description: "no subscription values in dynamodb",
			getItemOutput: &dynamodb.GetItemOutput{
				Item: map[string]*dynamodb.AttributeValue{
					accountIDKey: {
						S: aws.String("test_account_id"),
					},
				},
			},
			getItemError: nil,
			account:      nil,
			error:        &ErrorNoSubscription{},
		},
		{
			description: "unmarshal config object error",
			getItemOutput: &dynamodb.GetItemOutput{
				Item: map[string]*dynamodb.AttributeValue{
					accountIDKey: {
						S: aws.String("test_account_id"),
					},
					subscriptionIDKey: {
						S: aws.String("test_subscription_id"),
					},
					stripePaymentMethodIDKey: {
						S: aws.String("test_stripe_payment_method_id"),
					},
					stripeCustomerIDKey: {
						S: aws.String("test_stripe_customer_id"),
					},
					stripeSubscriptionIDKey: {
						S: aws.String("test_stripe_subscription_id"),
					},
					stripeSubscriptionItemIDKey: {
						S: aws.String("test_stripe_subscription_item_id"),
					},
					configKey: {
						B: []byte("------------"),
					},
				},
			},
			getItemError: nil,
			account:      nil,
			error:        &ErrorUnmarshalConfig{},
		},
		{
			description: "successful get account call",
			getItemOutput: &dynamodb.GetItemOutput{
				Item: map[string]*dynamodb.AttributeValue{
					accountIDKey: {
						S: aws.String("test_account_id"),
					},
					subscriptionIDKey: {
						S: aws.String("test_subscription_id"),
					},
					stripePaymentMethodIDKey: {
						S: aws.String("test_stripe_payment_method_id"),
					},
					stripeCustomerIDKey: {
						S: aws.String("test_stripe_customer_id"),
					},
					stripeSubscriptionIDKey: {
						S: aws.String("test_stripe_subscription_id"),
					},
					stripeSubscriptionItemIDKey: {
						S: aws.String("test_stripe_subscription_item_id"),
					},
					configKey: {
						B: []byte(`{"id":"config_id","selections":[{"text":"selection text","selected":true}]}`),
					},
				},
			},
			getItemError: nil,
			account: &acct.Account{
				ID: "test_account_id",
				Subscription: subscr.Subscription{
					ID:                       "test_subscription_id",
					StripePaymentMethodID:    "test_stripe_payment_method_id",
					StripeCustomerID:         "test_stripe_customer_id",
					StripeSubscriptionID:     "test_stripe_subscription_id",
					StripeSubscriptionItemID: "test_stripe_subscription_item_id",
				},
				Config: &config.Config{
					ID: "config_id",
					Selections: []config.Selection{
						{
							Text:     "selection text",
							Selected: true,
						},
					},
				},
			},
			error: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			client := &Client{
				dynamoDBClient: &mockDynamoDBClient{
					getItemOutput: test.getItemOutput,
					getItemError:  test.getItemError,
				},
			}

			account, err := client.GetAccount(context.Background(), "test_account_id")

			if err != nil {
				switch test.error.(type) {
				case *ErrorGetItem:
					var testError *ErrorGetItem
					if !errors.As(err, &testError) {
						t.Errorf("incorrect error, received: %v, expected: %v", err, testError)
					}
				case *ErrorNoAccount:
					var testError *ErrorNoAccount
					if !errors.As(err, &testError) {
						t.Errorf("incorrect error, received: %v, expected: %v", err, testError)
					}
				case *ErrorNoSubscription:
					var testError *ErrorNoSubscription
					if !errors.As(err, &testError) {
						t.Errorf("incorrect error, received: %v, expected: %v", err, testError)
					}
				case *ErrorUnmarshalConfig:
					var testError *ErrorUnmarshalConfig
					if !errors.As(err, &testError) {
						t.Errorf("incorrect error, received: %v, expected: %v", err, testError)
					}
				default:
					t.Fatalf("unexpected error type: %v", err)
				}
			} else {
				if err != test.error {
					t.Errorf("incorrect nil error, received: %v, expected: %v", err, test.error)
				}
			}

			if account != nil {
				if !reflect.DeepEqual(account, test.account) {
					t.Errorf("incorrect account, received: %+v, expected: %+v", *account, *test.account)
				}
			} else {
				if account != test.account {
					t.Errorf("incorrect nil check, received: %+v, expected: %+v", account, test.account)
				}
			}
		})
	}
}

func TestRemoveAccount(t *testing.T) {
	tests := []struct {
		description     string
		deleteItemError error
		error           error
	}{
		{
			description:     "receive dynamodb delete item error",
			deleteItemError: errors.New("mock delete error"),
			error:           &ErrorDeleteItem{},
		},
		{
			description:     "successful remove account call",
			deleteItemError: nil,
			error:           nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			client := &Client{
				dynamoDBClient: &mockDynamoDBClient{
					deleteItemError: test.deleteItemError,
				},
			}

			err := client.RemoveAccount(context.Background(), "test_account_id")
			if err != nil {
				switch test.error.(type) {
				case *ErrorDeleteItem:
					var testError *ErrorDeleteItem
					if !errors.As(err, &testError) {
						t.Errorf("incorrect error, received: %v, expected: %v", err, testError)
					}
				default:
					t.Fatalf("unexpected error type: %v", err)
				}
			} else {
				if err != test.error {
					t.Errorf("incorrect nil error, received: %v, expected: %v", err, test.error)
				}
			}
		})
	}
}
