package graphql

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Mutate executes a GraphQL mutation against the Dgraph database and populates
// received data into the provided response object or returns an error.
func (c *Client) Mutate(request *http.Request, response interface{}) error {
	resp, err := c.httpClient.Do(request)
	if err != nil {
		return ErrorMutateDo{err: err}
	}
	defer resp.Body.Close()

	graphQLResp := graphQLResponse{
		Data: response,
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ErrorBodyReadAll{err: err}
	}

	if err := json.Unmarshal(bodyBytes, &graphQLResp); err != nil {
		return ErrorBodyUnmarshal{err: err}
	}

	errs := graphQLResp.Errors
	if len(errs) > 0 {
		return ErrorGraphQLResponse{errs: errs}
	}

	return nil
}

type graphQLError struct {
	Message string `json:"message"`
}

type graphQLResponse struct {
	Data   interface{}    `json:"data"`
	Errors []graphQLError `json:"errors"`
}
