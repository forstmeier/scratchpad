package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewClient(t *testing.T) {
	testURL := "test.com"

	client := NewClient(testURL)

	if client.dgraphURL != testURL || client.httpClient == nil {
		t.Errorf("error creating client, received: %+v\n", client)
	}
}

func TestHandler(t *testing.T) {
	tests := []struct {
		description        string
		requestQuery       string
		requestVariables   map[string]string
		dgraphStatusCode   int
		responseStatusCode int
		responseBody       string
	}{
		{
			description:        "incorrect status code from backend",
			requestQuery:       "query test",
			requestVariables:   nil,
			dgraphStatusCode:   http.StatusInternalServerError,
			responseStatusCode: http.StatusInternalServerError,
			responseBody:       errorMessage(errorResponseStatusCode),
		},
		{
			description:        "successful handler invocation",
			requestQuery:       "query test",
			requestVariables:   nil,
			dgraphStatusCode:   http.StatusOK,
			responseStatusCode: http.StatusOK,
			responseBody:       `{"data":{"addUser":{"user":{"name":"username"}}}}`,
		},
	}

	// convert this to a full type with values for fields
	var responseBody struct {
		Data struct {
			AddUser struct {
				User struct {
					Name string `json:"name"`
				} `json:"user"`
			} `json:"addUser"`
		} `json:"data"`
	}
	responseBody.Data.AddUser.User.Name = "username"

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			dgraphMux := http.NewServeMux()
			dgraphMux.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(test.dgraphStatusCode)
				json.NewEncoder(w).Encode(responseBody)
			})

			dgraphServer := httptest.NewServer(dgraphMux)
			dgraphURL := dgraphServer.URL + "/graphql"

			payload := struct {
				Query     string      `json:"query"`
				Variables interface{} `json:"variables"`
			}{
				Query:     test.requestQuery,
				Variables: test.requestVariables,
			}

			payloadBytes, err := json.Marshal(payload)
			if err != nil {
				t.Fatalf("error creating test payload: %s", err.Error())
			}

			req, err := http.NewRequest(
				http.MethodPost,
				"/graphql",
				bytes.NewReader(payloadBytes),
			)

			rec := httptest.NewRecorder()

			c := Client{
				dgraphURL:  dgraphURL,
				httpClient: &http.Client{},
			}

			handler := http.HandlerFunc(c.Handler())

			handler.ServeHTTP(rec, req)

			if status := rec.Code; status != test.responseStatusCode {
				t.Errorf("status received: %d, expected: %d", status, test.responseStatusCode)
			}

			if body := strings.TrimSuffix(rec.Body.String(), "\n"); body != test.responseBody {
				t.Errorf("body received: %s, expected: %s", body, test.responseBody)
			}
		})
	}
}

func Test_errorMessage(t *testing.T) {
	input := "error message"
	received := errorMessage(input)
	expected := `{"message":"error message"}`
	if received != expected {
		t.Fatalf("incorrect message, received: %s, expected: %s", received, expected)
	}
}
