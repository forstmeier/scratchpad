package graphql

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// GraphQL provides wrapper methods around HTTP requests sent to the
// configured Dgraph URL endpoint.
type GraphQL interface {
	Mutate(query string, variables map[string]interface{}, response interface{}) error
}

// Client implements the GraphQL interface.
type Client struct {
	client *http.Client
	url    string
	token  string
}

// New generates a pointer Dgraph client instance.
func New(url, token string) GraphQL {
	return &Client{
		client: &http.Client{},
		url:    url,
		token:  token,
	}
}

// Mutate executes the provided mutation and optional variables against the
// Dgraph GraphQL database endpoint and populates the returned data into
// the response parameter.
func (c *Client) Mutate(mutation string, variables map[string]interface{}, response interface{}) error {
	payload := struct {
		Query     string      `json:"query"`
		Variables interface{} `json:"variables"`
	}{
		Query:     mutation,
		Variables: variables,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return newErrorMarshalPayload(err)
	}

	req, err := http.NewRequest(http.MethodPost, c.url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return newErrorNewRequest(err)
	}
	req.Header.Set("X-Auth0-Token", c.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return newErrorClientDo(err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return newErrorJSONDecode(err)
	}

	return nil
}
