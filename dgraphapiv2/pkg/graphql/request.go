package graphql

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// NewRequest generates an http.Request pointer populated with the provided
// headers and body values to be used in a GraphQL request to Dgraph.
func NewRequest(url, token, query string, variables map[string]string) (*http.Request, error) {
	payload := struct {
		Query     string      `json:"query"`
		Variables interface{} `json:"variables"`
	}{
		Query:     query,
		Variables: variables,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, ErrorRequestPayloadMarshal{err: err}
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, ErrorHTTPNewRequest{err: err}
	}

	req.Header.Set("X-Auth0-Token", token)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
