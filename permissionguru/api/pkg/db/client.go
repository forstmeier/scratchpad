package db

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/permissionguru/api/pkg/acct"
	"github.com/permissionguru/api/pkg/config"
	"github.com/permissionguru/api/pkg/subscr"
)

const (
	tableName                   = "accounts"
	accountIDKey                = "id"
	subscriptionIDKey           = "subscription_id"
	stripePaymentMethodIDKey    = "stripe_payment_method_id"
	stripeCustomerIDKey         = "stripe_customer_id"
	stripeSubscriptionIDKey     = "stripe_subscription_id"
	stripeSubscriptionItemIDKey = "stripe_subscription_item_id"
	configKey                   = "config"
)

var _ Databaser = &Client{}

// Client implements the Databaser methods using AWS DynamoDB.
type Client struct {
	dynamoDBClient dynamoDBClient
}

type dynamoDBClient interface {
	PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
	UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error)
	GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
	DeleteItem(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error)
}

// New generates a Client pointer instance with an AWS
// DynamoDB session client.
func New(service *dynamodb.DynamoDB) *Client {
	return &Client{
		dynamoDBClient: service,
	}
}

// AddAccount implements the Databaser.AddAccountID method and
// adds the provided UUID id to the DynamoDB table as a new row.
func (c *Client) AddAccount(ctx context.Context, id string) error {
	if id == "" {
		return &ErrorMissingID{}
	}

	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			accountIDKey: {
				S: aws.String(id),
			},
		},
		TableName: aws.String(tableName),
	}

	_, err := c.dynamoDBClient.PutItem(input)
	if err != nil {
		return &ErrorPutItem{err: err}
	}

	return nil
}

// AddSubscription implements the Databaser.AddSubscription method
// and adds the provided Subscription values to the DynamoDB table
// on the provided id key.
func (c *Client) AddSubscription(ctx context.Context, id string, subscription subscr.Subscription) error {
	fields := checkSubscriptionFields(subscription)
	if len(fields) > 0 {
		return &ErrorMissingFields{fields: fields}
	}

	expression := fmt.Sprintf("SET %s = :s_id, %s = :spm_id, %s = :sc_id, %s = :ss_id, %s = :ssi_id",
		subscriptionIDKey,
		stripePaymentMethodIDKey,
		stripeCustomerIDKey,
		stripeSubscriptionIDKey,
		stripeSubscriptionItemIDKey,
	)

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			accountIDKey: {
				S: aws.String(id),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":s_id": {
				S: aws.String(subscription.ID),
			},
			":spm_id": {
				S: aws.String(subscription.StripePaymentMethodID),
			},
			":sc_id": {
				S: aws.String(subscription.StripeCustomerID),
			},
			":ss_id": {
				S: aws.String(subscription.StripeSubscriptionID),
			},
			":ssi_id": {
				S: aws.String(subscription.StripeSubscriptionItemID),
			},
		},
		UpdateExpression:    aws.String(expression),
		ConditionExpression: aws.String("attribute_exists(id)"),
	}

	_, err := c.dynamoDBClient.UpdateItem(input)
	if err != nil {
		return &ErrorUpdateItem{
			item: "subscription",
			err:  err,
		}
	}

	return nil
}

func checkSubscriptionFields(subscription subscr.Subscription) []string {
	fields := []string{}
	if subscription.ID == "" {
		fields = append(fields, "subscription id")
	}
	if subscription.StripePaymentMethodID == "" {
		fields = append(fields, "stripe payment method id")
	}
	if subscription.StripeCustomerID == "" {
		fields = append(fields, "stripe customer id")
	}
	if subscription.StripeSubscriptionID == "" {
		fields = append(fields, "stripe subscription id")
	}
	if subscription.StripeSubscriptionItemID == "" {
		fields = append(fields, "stripe subscription item id")
	}

	return fields
}

// AddConfig implements the Databaser.AddConfig method and adds
// the provided Config values to the DynamoDB table on the provided id key.
func (c *Client) AddConfig(ctx context.Context, id string, cfg config.Config) error {
	if len(cfg.Selections) == 0 {
		return &ErrorNoConfigSelections{}
	}

	configBytes, err := json.Marshal(cfg)
	if err != nil {
		return &ErrorMarshalConfig{err: err}
	}

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			accountIDKey: {
				S: aws.String(id),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":c": {
				B: configBytes,
			},
		},
		UpdateExpression:    aws.String(fmt.Sprintf("SET %s = :c", configKey)),
		ConditionExpression: aws.String("attribute_exists(id)"),
	}

	_, err = c.dynamoDBClient.UpdateItem(input)
	if err != nil {
		return &ErrorUpdateItem{err: err}
	}

	return nil
}

// GetAccount implements the Databaser.GetAccount method and
// returns an Account pointer if id exists in the DynamoDB table.
func (c *Client) GetAccount(ctx context.Context, id string) (*acct.Account, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			accountIDKey: {
				S: aws.String(id),
			},
		},
		TableName: aws.String(tableName),
	}

	output, err := c.dynamoDBClient.GetItem(input)
	if err != nil {
		return nil, &ErrorGetItem{err: err}
	}

	account := &acct.Account{}
	if _, accountCheck := output.Item[accountIDKey]; accountCheck {
		item := output.Item

		if _, subscriptionCheck := item[subscriptionIDKey]; !subscriptionCheck {
			return nil, &ErrorNoSubscription{}
		}

		account.ID = *item[accountIDKey].S

		account.Subscription = subscr.Subscription{
			ID:                       *item[subscriptionIDKey].S,
			StripePaymentMethodID:    *item[stripePaymentMethodIDKey].S,
			StripeCustomerID:         *item[stripeCustomerIDKey].S,
			StripeSubscriptionID:     *item[stripeSubscriptionIDKey].S,
			StripeSubscriptionItemID: *item[stripeSubscriptionItemIDKey].S,
		}

		if _, configCheck := item[configKey]; configCheck {
			cfg := config.Config{}
			if err := json.Unmarshal(item[configKey].B, &cfg); err != nil {
				return nil, &ErrorUnmarshalConfig{err: err}
			}
			account.Config = &cfg
		} else {
			account.Config = nil
		}

	} else {
		return nil, &ErrorNoAccount{}
	}

	return account, nil
}

// RemoveAccount implements the Databaser.RemoveAccount method and
// removes the Account values by id from the DynamoDB table.
func (c *Client) RemoveAccount(ctx context.Context, id string) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			accountIDKey: {
				S: aws.String(id),
			},
		},
		TableName: aws.String(tableName),
	}

	_, err := c.dynamoDBClient.DeleteItem(input)
	if err != nil {
		return &ErrorDeleteItem{err: err}
	}

	return nil
}
