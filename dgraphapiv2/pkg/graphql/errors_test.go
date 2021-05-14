package graphql

import (
	"errors"
	"testing"
)

func TestErrorRequestPayloadMarshalError(t *testing.T) {
	input := errors.New("input error")

	err := ErrorRequestPayloadMarshal{err: input}

	received := err.Error()
	expected := "graphql: request payload marshal: input error"

	if received != expected {
		t.Errorf("error received: %s, expected: %s", received, expected)
	}
}

func TestErrorHTTPNewRequestError(t *testing.T) {
	input := errors.New("input error")

	err := ErrorHTTPNewRequest{err: input}

	received := err.Error()
	expected := "graphql: http new request: input error"

	if received != expected {
		t.Errorf("error received: %s, expected: %s", received, expected)
	}
}

func TestErrorMutateDoError(t *testing.T) {
	input := errors.New("input error")

	err := ErrorMutateDo{err: input}

	received := err.Error()
	expected := "graphql: http do: input error"

	if received != expected {
		t.Errorf("error received: %s, expected: %s", received, expected)
	}
}

func TestErrorBodyReadAllError(t *testing.T) {
	input := errors.New("input error")

	err := ErrorBodyReadAll{err: input}

	received := err.Error()
	expected := "graphql: response body read all: input error"

	if received != expected {
		t.Errorf("error received: %s, expected: %s", received, expected)
	}
}

func TestErrorBodyUnmarshalError(t *testing.T) {
	input := errors.New("input error")

	err := ErrorBodyUnmarshal{err: input}

	received := err.Error()
	expected := "graphql: response body unmarshal: input error"

	if received != expected {
		t.Errorf("error received: %s, expected: %s", received, expected)
	}
}

func TestErrorGraphQLResponseError(t *testing.T) {
	input := []graphQLError{
		graphQLError{
			Message: "first error",
		},
		graphQLError{
			Message: "second error",
		},
	}

	err := ErrorGraphQLResponse{errs: input}

	received := err.Error()
	expected := "graphql: response errors: message: first error, message: second error"

	if received != expected {
		t.Errorf("error received: %s, expected: %s", received, expected)
	}
}
