package blevedb

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/specimenguru/api/pkg/data"
	"github.com/specimenguru/api/pkg/database"
)

func TestQuery(t *testing.T) {
	tests := []struct {
		description string
		query       database.Query
		output      []string
		error       error
	}{
		{
			description: "no full text query received",
			query:       database.Query{},
			output:      nil,
			error:       &ErrorNoQuery{},
		},
		{
			description: "no results returned from database",
			query: database.Query{
				FullText: "nothing",
			},
			output: []string{},
			error:  nil,
		},
		{
			description: "single result returned from database",
			query: database.Query{
				FullText: "good yellow",
			},
			output: []string{"third_id"},
			error:  nil,
		},
	}

	indexName := "test_query.bleve"

	client, err := New(indexName)

	if err != nil {
		t.Fatalf("error creating test client: %s", err.Error())
	}

	now := time.Now()
	testSpecimens := []data.Specimen{
		{ID: "first_id", Timestamp: now, Body: `{"content":"a very blue specimen"}`},
		{ID: "second_id", Timestamp: now, Body: `{"content":"specimen that is red"}`},
		{ID: "third_id", Timestamp: now, Body: `{"content":"yellow specimen"}`},
	}

	if err := client.Add(testSpecimens); err != nil {
		t.Fatalf("error adding test specimens: %s", err.Error())
	}

	defer os.RemoveAll(indexName)

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			specimens, err := client.Query(test.query)

			if err != nil {
				// switch included for future additional errors
				switch test.error.(type) {
				case *ErrorNoQuery:
					var testError *ErrorNoQuery
					if !errors.As(err, &testError) {
						t.Errorf("incorrect error, received: %v, expected: %v", err, testError)
					}
				default:
					t.Fatalf("unexpected error type: %v", err)
				}
			}

			if len(test.output) > 0 {
				if len(specimens) != len(test.output) {
					t.Errorf("incorrect specimens count, received: %d, expected: %d", len(specimens), len(test.output))
				}

				for i, specimen := range specimens {
					received := specimen.ID
					expected := test.output[i]
					if received != expected {
						t.Errorf("incorrect id, received: %s, expected: %s", received, expected)
					}
				}
			}
		})
	}
}
