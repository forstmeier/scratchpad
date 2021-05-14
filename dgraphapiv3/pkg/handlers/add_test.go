package handlers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddHandler(t *testing.T) {
	tests := []struct {
		description string
		payload     []byte
		contentType string
		addError    error
		statusCode  int
		body        string
	}{
		{
			description: "incorrect content type header received",
			payload:     []byte(""),
			contentType: "incorrect/type",
			addError:    nil,
			statusCode:  http.StatusBadRequest,
			body:        `{"error":"content type incorrect/type not supported"}`,
		},
		{
			description: "database add error",
			payload:     []byte(`[{"key":"value"}]`),
			contentType: "application/json",
			addError:    errors.New("add error"),
			statusCode:  http.StatusInternalServerError,
			body:        `{"error":"error adding specimens to database"}`,
		},
		{
			description: "successful add handler invocation",
			payload:     []byte(`[{"key":"value"}]`),
			contentType: "application/json",
			addError:    nil,
			statusCode:  http.StatusOK,
			body:        `{"message":"success"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			req, err := http.NewRequest(
				http.MethodPost,
				"/database",
				bytes.NewReader(test.payload),
			)
			if err != nil {
				t.Fatalf("error creating request: %s", err.Error())
			}
			req.Header.Add("Content-Type", test.contentType)

			rec := httptest.NewRecorder()

			client := &Client{
				db: &testDatabase{
					addError: test.addError,
				},
			}

			handler := http.HandlerFunc(client.AddHandler())

			handler.ServeHTTP(rec, req)

			if rec.Code != test.statusCode {
				t.Errorf("incorrect response code, received: %d, expected: %d", rec.Code, test.statusCode)
			}

			received := strings.TrimSuffix(rec.Body.String(), "\n")
			expected := test.body
			if received != expected {
				t.Errorf("incorrect response body, received: %s, expected: %s", received, expected)
			}
		})
	}
}
