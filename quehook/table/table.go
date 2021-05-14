package table

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type dynamoDBClient interface {
	PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
	Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
	DeleteItem(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error)
}

// Table provides helper methods for persisting/retrieving/deleting items
type Table interface {
	Add(table string, items ...string) error
	Get(table, key, attribute string) ([]string, error)
	Remove(table, key, attribute string) error
}

// Client implements the Table interface
type Client struct {
	dynamodb dynamoDBClient
}

// New generates a Table implementation with an active client
func New() Table {
	return &Client{
		dynamodb: dynamodb.New(session.New()),
	}
}

// Add puts a new item into DynamoDB
func (c *Client) Add(table string, items ...string) error {
	log.Printf("add table: %s, items: %v\n", table, items)

	input := &dynamodb.PutItemInput{
		TableName: aws.String("quehook-" + table),
		Item: map[string]*dynamodb.AttributeValue{
			"query_name": {
				S: aws.String(items[0]),
			},
		},
	}

	if table == "subscribers" {
		input.Item["subscriber_email"] = &dynamodb.AttributeValue{
			S: aws.String(items[1]),
		}
		input.Item["subscriber_name"] = &dynamodb.AttributeValue{
			S: aws.String(items[2]),
		}
		input.Item["subscriber_target"] = &dynamodb.AttributeValue{
			S: aws.String(items[3]),
		}
	} else if table == "queries" {
		input.Item["author_name"] = &dynamodb.AttributeValue{
			S: aws.String(items[1]),
		}
		input.Item["author_email"] = &dynamodb.AttributeValue{
			S: aws.String(items[2]),
		}
	}

	_, err := c.dynamodb.PutItem(input)
	if err != nil {
		return fmt.Errorf("put item error: %s", err.Error())
	}
	return nil
}

// Get retrieves an item or items from DynamoDB
//
// Functionality: for a given table and primary key element, return
// all attributes that match the provided attribute name
func (c *Client) Get(table, key, attribute string) ([]string, error) {
	log.Printf("get table: %s, key: %s, attribute: %s\n", table, key, attribute)

	output := []string{}

	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":queryName": {
				S: aws.String(key),
			},
		},
		KeyConditionExpression: aws.String("query_name = :queryName"),
		TableName:              aws.String("quehook-" + table),
	}

	results, err := c.dynamodb.Query(input)
	if err != nil {
		return nil, fmt.Errorf("query items error: %s", err.Error())
	}
	log.Printf("query results: %+v\n", results)

	for _, item := range results.Items {
		output = append(output, *item[attribute].S)
	}

	log.Println("output: ", output)
	return output, nil
}

// Remove deletes an item from DynamoDB
func (c *Client) Remove(table, key, attribute string) error {
	log.Printf("remove table: %s, key: %s, attribute: %s\n", table, key, attribute)

	if table == "subscribers" {
		input := &dynamodb.QueryInput{
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":queryName": {
					S: aws.String(key),
				},
			},
			KeyConditionExpression: aws.String("query_name = :queryName"),
			TableName:              aws.String("quehook-" + table),
		}

		results, err := c.dynamodb.Query(input)
		if err != nil {
			return fmt.Errorf("query items error: %s", err.Error())
		}
		log.Printf("query results: %+v\n", results)

		for _, item := range results.Items {
			sortKey := *item["subscriber_email"].S
			log.Printf("item primary key: %s, sort key: %s", key, sortKey)
			response, err := c.dynamodb.DeleteItem(&dynamodb.DeleteItemInput{
				TableName: aws.String("quehook-" + table),
				Key: map[string]*dynamodb.AttributeValue{
					"query_name": {
						S: aws.String(key),
					},
					"subscriber_email": {
						S: aws.String(sortKey),
					},
				},
				ReturnValues: aws.String("ALL_OLD"),
			})
			if err != nil {
				return fmt.Errorf("delete item error: %s", err.Error())
			}
			log.Printf("item removed: %+v\n", response.Attributes)
		}

	} else if table == "queries" {
		response, err := c.dynamodb.DeleteItem(&dynamodb.DeleteItemInput{
			TableName: aws.String("quehook-" + table),
			Key: map[string]*dynamodb.AttributeValue{
				"query_name": {
					S: aws.String(key),
				},
			},
			ReturnValues: aws.String("ALL_OLD"),
		})
		if err != nil {
			return fmt.Errorf("delete item error: %s", err.Error())
		}
		log.Printf("item removed: %+v\n", response.Attributes)
	}

	return nil
}
