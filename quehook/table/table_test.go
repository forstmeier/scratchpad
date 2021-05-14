package table

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestNew(t *testing.T) {
	tbl := New()
	if tbl == nil {
		t.Errorf("description: error creating table client, received: %+v", tbl)
	}
}

type tableMock struct {
	putItemOutput    *dynamodb.PutItemOutput
	putItemError     error
	queryItemOutput  *dynamodb.QueryOutput
	queryItemError   error
	deleteItemOutput *dynamodb.DeleteItemOutput
	deleteItemError  error
}

func (mock *tableMock) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return mock.putItemOutput, mock.putItemError
}

func (mock *tableMock) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	return mock.queryItemOutput, mock.queryItemError
}

func (mock *tableMock) DeleteItem(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	return mock.deleteItemOutput, mock.deleteItemError
}

func TestAdd(t *testing.T) {
	tests := []struct {
		desc          string
		table         string
		items         []string
		putItemOutput *dynamodb.PutItemOutput // kept for possible method expansion
		putItemError  error
		err           string
	}{
		{
			desc:  "put item error",
			table: "queries",
			items: []string{
				"newPowerConverters",
				"luke",
				"luke@lars.homestead",
			},
			putItemOutput: nil,
			putItemError:  errors.New("mock put error"),
			err:           "put item error: mock put error",
		},
		{
			desc:  "successful subscribers invocation",
			table: "subscribers",
			items: []string{
				"newChores",
				"luke@lars.homstead",
				"luke",
				"https://holonet.com/skywalker",
			},
			putItemOutput: nil,
			putItemError:  nil,
			err:           "",
		},
		{
			desc:  "successful queries invocation",
			table: "queries",
			items: []string{
				"lotsOfTrouble",
				"r2-d2",
				"blue-and-white@royalengineers.nb",
			},
			putItemOutput: nil,
			putItemError:  nil,
			err:           "",
		},
	}

	for _, test := range tests {
		c := &Client{
			dynamodb: &tableMock{
				putItemOutput: test.putItemOutput,
				putItemError:  test.putItemError,
			},
		}

		if err := c.Add(test.table, test.items...); err != nil && err.Error() != test.err {
			t.Errorf("description: %s, error received: %s, expected: %s", test.desc, err.Error(), test.err)
		}
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		desc            string
		table           string
		key             string
		attribute       string
		queryItemOutput *dynamodb.QueryOutput
		queryItemError  error
		output          []string
		err             string
	}{
		{
			desc:            "get item error",
			table:           "queries",
			key:             "query",
			attribute:       "",
			queryItemOutput: nil,
			queryItemError:  errors.New("mock query error"),
			output:          nil,
			err:             "query items error: mock query error",
		},
		{
			desc:      "successful subscribers invocation",
			table:     "subscribers",
			key:       "query",
			attribute: "subscriber_target",
			queryItemOutput: &dynamodb.QueryOutput{
				Items: []map[string]*dynamodb.AttributeValue{
					map[string]*dynamodb.AttributeValue{
						"subscriber_target": {
							S: aws.String("test-target"),
						},
					},
				},
			},
			queryItemError: nil,
			output: []string{
				"test-query",
			},
			err: "",
		},
		{
			desc:      "successful queries invocation",
			table:     "queries",
			key:       "key",
			attribute: "query_name",
			queryItemOutput: &dynamodb.QueryOutput{
				Items: []map[string]*dynamodb.AttributeValue{
					map[string]*dynamodb.AttributeValue{
						"query_name": {
							S: aws.String("test-query"),
						},
					},
				},
			},
			queryItemError: nil,
			output: []string{
				"test-query",
			},
			err: "",
		},
	}

	for _, test := range tests {
		c := &Client{
			dynamodb: &tableMock{
				queryItemOutput: test.queryItemOutput,
				queryItemError:  test.queryItemError,
			},
		}

		output, err := c.Get(test.table, test.key, test.attribute)

		if len(output) != len(test.output) {
			t.Errorf("description: %s, output received: %d, expected: %d", test.desc, len(output), len(test.output))
		}

		if err != nil && err.Error() != test.err {
			t.Errorf("description: %s, error received: %s, expected: %s", test.desc, err.Error(), test.err)
		}
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		desc             string
		table            string
		key              string
		attribute        string
		queryItemOutput  *dynamodb.QueryOutput
		queryItemError   error
		deleteItemOutput *dynamodb.DeleteItemOutput
		deleteItemError  error
		err              string
	}{
		{
			desc:             "delete queries error",
			table:            "queries",
			key:              "query",
			attribute:        "",
			queryItemOutput:  nil,
			queryItemError:   nil,
			deleteItemOutput: nil,
			deleteItemError:  errors.New("mock delete error"),
			err:              "delete item error: mock delete error",
		},
		{
			desc:            "delete queries successful invocation",
			table:           "queries",
			key:             "query",
			attribute:       "",
			queryItemOutput: nil,
			queryItemError:  nil,
			deleteItemOutput: &dynamodb.DeleteItemOutput{
				Attributes: map[string]*dynamodb.AttributeValue{
					"query_name": {
						S: aws.String("test-query"),
					},
				},
			},
			deleteItemError: nil,
			err:             "",
		},
		{
			desc:             "delete subscribers query error",
			table:            "subscribers",
			key:              "query",
			attribute:        "",
			queryItemOutput:  nil,
			queryItemError:   errors.New("mock query error"),
			deleteItemOutput: nil,
			deleteItemError:  nil,
			err:              "query items error: mock query error",
		},
		{
			desc:      "delete subscribers delete error",
			table:     "subscribers",
			key:       "query",
			attribute: "query_name",
			queryItemOutput: &dynamodb.QueryOutput{
				Items: []map[string]*dynamodb.AttributeValue{
					map[string]*dynamodb.AttributeValue{
						"query_name": {
							S: aws.String("test-query"),
						},
						"subscriber_email": {
							S: aws.String("test-subscriber-email"),
						},
					},
				},
			},
			queryItemError:   nil,
			deleteItemOutput: nil,
			deleteItemError:  errors.New("mock delete error"),
			err:              "delete item error: mock delete error",
		},
		{
			desc:      "delete subscribers successful invocation",
			table:     "subscribers",
			key:       "query",
			attribute: "query_name",
			queryItemOutput: &dynamodb.QueryOutput{
				Items: []map[string]*dynamodb.AttributeValue{
					map[string]*dynamodb.AttributeValue{
						"query_name": {
							S: aws.String("test-query"),
						},
						"subscriber_email": {
							S: aws.String("test-subscriber-email"),
						},
					},
				},
			},
			queryItemError: nil,
			deleteItemOutput: &dynamodb.DeleteItemOutput{
				Attributes: map[string]*dynamodb.AttributeValue{
					"query_name": {
						S: aws.String("test-query"),
					},
					"subscriber_email": {
						S: aws.String("test-subscriber-email"),
					},
				},
			},
			deleteItemError: nil,
			err:             "",
		},
	}

	for _, test := range tests {
		c := &Client{
			dynamodb: &tableMock{
				queryItemOutput:  test.queryItemOutput,
				queryItemError:   test.queryItemError,
				deleteItemOutput: test.deleteItemOutput,
				deleteItemError:  test.deleteItemError,
			},
		}

		if err := c.Remove(test.table, test.key, test.attribute); err != nil && err.Error() != test.err {
			t.Errorf("description: %s, error received: %s, expected: %s", test.desc, err.Error(), test.err)
		}
	}
}
