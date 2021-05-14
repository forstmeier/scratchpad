package graphql

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMutate(t *testing.T) {
	var resp struct {
		Add struct {
			User struct {
				Name string `json:"name"`
			} `json:"user"`
		} `json:"add"`
	}

	tests := []struct {
		description    string
		response       interface{}
		dgraphResponse string
		error          error
		output         interface{}
	}{
		{
			description:    "error unmarshalling dgraph response body",
			response:       nil,
			dgraphResponse: "---------",
			error: ErrorBodyUnmarshal{
				err: errors.New("wrapped error"),
			},
			output: nil,
		},
		{
			description:    "errors returned from dgraph",
			response:       nil,
			dgraphResponse: `{"errors":[{"message":"error message"}]}`,
			error: ErrorGraphQLResponse{
				errs: []graphQLError{},
			},
			output: nil,
		},
		{
			description:    "successful mutate invocation",
			response:       resp,
			dgraphResponse: `{"data":{"add":{"user":{"name":"test"}}}}`,
			error:          nil,
			output:         nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			dgraphMux := http.NewServeMux()

			dgraphMux.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(test.dgraphResponse))
			})

			dgraphServer := httptest.NewServer(dgraphMux)
			dgraphURL := dgraphServer.URL + "/graphql"

			client := NewClient(dgraphURL)

			request, err := NewRequest(dgraphURL, "token", "query example", nil)
			if err != nil {
				t.Fatalf("error creating test request: %s", err.Error())
			}

			err = client.Mutate(request, &test.response)
			if err != nil && !errors.As(err, &test.error) {
				t.Errorf("error received: %s, expected: %s", err.Error(), test.error.Error())
			}

			if err == nil {
				// this should be updated to be a json unmarshal with an explicit response type
				response := test.response.(map[string]interface{})
				add := response["add"].(map[string]interface{})
				user := add["user"].(map[string]interface{})
				name := user["name"].(string)

				expected := "test"
				if name != expected {
					t.Errorf("response received: %s, expected: %s", name, expected)
				}
			}
		})
	}
}
