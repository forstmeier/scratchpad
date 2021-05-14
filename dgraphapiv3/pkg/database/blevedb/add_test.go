package blevedb

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/blevesearch/bleve/v2"
	"github.com/specimenguru/api/pkg/data"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		description string
		input       []data.Specimen
		error       error
	}{
		{
			description: "received empty input slice",
			input:       []data.Specimen{},
			error:       &ErrorNoAddData{},
		},
		{
			description: "successful invocation with one specimen received",
			input: []data.Specimen{
				{
					ID:        "first_id",
					Timestamp: time.Now(),
					Body:      `{"key":"value"}`,
				},
			},
			error: nil,
		},
		{
			description: "successful invocation with three specimens received",
			input: []data.Specimen{
				{
					ID:        "first_id",
					Timestamp: time.Now(),
					Body:      `{"key":"value"}`,
				},
				{
					ID:        "second_id",
					Timestamp: time.Now(),
					Body:      `{"key":"value"}`,
				},
				{
					ID:        "third_id",
					Timestamp: time.Now(),
					Body:      `{"key":"value"}`,
				},
			},
			error: nil,
		},
	}

	indexName := "test_add.bleve"
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New(indexName, mapping)
	if err != nil {
		t.Fatalf("error creating test index: %s", err.Error())
	}
	defer os.RemoveAll(indexName)

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			client := &Client{
				index: index,
			}

			err := client.Add(test.input)

			if err != nil {
				// switch included for future additional errors
				switch test.error.(type) {
				case *ErrorNoAddData:
					var testError *ErrorNoAddData
					if !errors.As(err, &testError) {
						t.Errorf("incorrect error, received: %v, expected: %v", err, testError)
					}
				default:
					t.Fatalf("unexpected error type: %v", err)
				}
			}

			if err == nil {
				for _, specimen := range test.input {
					doc, err := client.index.Document(specimen.ID)
					if err != nil {
						t.Errorf("error getting document: %s", err.Error())
					}

					received := doc.ID()
					expected := specimen.ID
					if received != expected {
						t.Errorf("incorrect id, received: %s, expected: %s", received, expected)
					}
				}
			}
		})
	}
}
