package handlers

import (
	"errors"
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

func Test_specimensFromCSV(t *testing.T) {
	tests := []struct {
		description string
		input       io.ReadCloser
		bodies      []string
		error       error
	}{
		{
			description: "non-csv input received",
			input:       ioutil.NopCloser(strings.NewReader(`{"not":"csv"}`)),
			bodies:      nil,
			error:       &ErrorCSVReadAll{},
		},
		{
			description: "csv headers with no data rows provided",
			input:       ioutil.NopCloser(strings.NewReader("first,last")),
			bodies:      nil,
			error:       &ErrorNotEnoughRows{},
		},
		{
			description: "successful invocation",
			input:       ioutil.NopCloser(strings.NewReader("first,last\ntest_first_1,test_last_1\ntest_first_2,test_last_2")),
			bodies: []string{
				`{"first":"test_first_1","last":"test_last_1"}`,
				`{"first":"test_first_2","last":"test_last_2"}`,
			},
			error: nil,
		},
	}

	for _, test := range tests {
		specimens, err := specimensFromCSV(test.input)

		if err != nil {
			switch test.error.(type) {
			case *ErrorCSVReadAll:
				var testError *ErrorCSVReadAll
				if !errors.As(err, &testError) {
					t.Errorf("incorrect error, received: %v, expected: %v", err, testError)
				}
			case *ErrorNotEnoughRows:
				var testError *ErrorNotEnoughRows
				if !errors.As(err, &testError) {
					t.Errorf("incorrect error, received: %v, expected: %v", err, testError)
				}
			default:
				t.Fatalf("unexpected error type: %v", err)
			}
		}

		for i, specimen := range specimens {
			received := specimen.Body
			expected := test.bodies[i]
			if received != expected {
				t.Errorf("incorrect body, received: %s, expected: %s", received, expected)
			}
		}
	}
}

func Test_specimensFromJSON(t *testing.T) {
	tests := []struct {
		description string
		input       io.ReadCloser
		bodies      []string
	}{
		{
			description: "non-json input received",
			input:       ioutil.NopCloser(strings.NewReader(`------------`)),
			bodies:      nil,
		},
		{
			description: "successful invocation",
			input:       ioutil.NopCloser(strings.NewReader(`[{"first":"test_first_1","last":"test_last_1"},{"first":"test_first_2","last":"test_last_2"}]`)),
			bodies: []string{
				`{"first":"test_first_1","last":"test_last_1"}`,
				`{"first":"test_first_2","last":"test_last_2"}`,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			specimens, err := specimensFromJSON(test.input)

			if err != nil {
				var testError *ErrorJSONDecode
				if !errors.As(err, &testError) {
					t.Errorf("incorrect error, received: %v, expected: %v", err, testError)
				}
			}

			for i, specimen := range specimens {
				received := specimen.Body
				expected := test.bodies[i]
				if received != expected {
					t.Errorf("incorrect body, received: %s, expected: %s", received, expected)
				}
			}
		})
	}
}

func Test_errorMessage(t *testing.T) {
	expected := `{"error":"error message"}`
	received := errorMessage("error message")
	if expected != received {
		t.Errorf("incorrect error, received: %s, expected: %s", received, expected)
	}
}
