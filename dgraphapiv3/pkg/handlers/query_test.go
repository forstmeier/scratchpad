package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/specimenguru/api/pkg/data"
)

func TestQueryHandler(t *testing.T) {
	tests := []struct {
		description string
		params      string
		queryError  error
		statusCode  int
		body        string
	}{
		{
			description: "no query param received in url",
			params:      "",
			queryError:  nil,
			statusCode:  http.StatusBadRequest,
			body:        `{"error":"no query url parameter provided"}`,
		},
		{
			description: "database query error",
			params:      "?full_text=example",
			queryError:  errors.New("query error"),
			statusCode:  http.StatusInternalServerError,
			body:        `{"error":"error querying specimens in database"}`,
		},
		{
			description: "successful query handler invocation",
			params:      "?full_text=orange%20red",
			queryError:  nil,
			statusCode:  http.StatusOK,
			body:        `[{"content":"orange specimen here"},{"content":"very red specimen"}]`,
		},
	}

	for _, test := range tests {
		now := time.Now()

		client := &Client{
			db: &testDatabase{
				querySpecimens: []data.Specimen{
					{
						ID:        "first_id",
						Timestamp: now,
						Body:      `{"content":"orange specimen here"}`,
					},
					{
						ID:        "second_id",
						Timestamp: now,
						Body:      `{"content":"very red specimen"}`,
					},
				},
				queryError: test.queryError,
			},
		}

		t.Run(test.description, func(t *testing.T) {
			req, err := http.NewRequest(
				http.MethodGet,
				"/database"+test.params,
				nil,
			)
			if err != nil {
				t.Fatalf("error creating request: %s", err.Error())
			}

			rec := httptest.NewRecorder()

			handler := http.HandlerFunc(client.QueryHandler())

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
