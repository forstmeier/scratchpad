package handler

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	errorParseURL           = "error calling backend api: configure api url"
	errorHTTPDo             = "error calling backend api: http do method"
	errorResponseStatusCode = "error calling backend api: incorrect status code"
)

// Client holds a reusable http.Client and the configured URL for the
// Dgraph database.
//
// Additionally, Client exposes the Handler method for generating the root
// http.HandlerFunc.
type Client struct {
	dgraphURL  string
	httpClient *http.Client
}

// NewClient generates a pointer instance of Client.
func NewClient(dgraphURL string) *Client {
	return &Client{
		dgraphURL:  dgraphURL,
		httpClient: &http.Client{},
	}
}

// Handler returns the root http.HandlerFunc "wrapper" around the Dgraph
// database endpoint.
func (c *Client) Handler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dgraphURL, err := url.Parse(c.dgraphURL)
		if err != nil {
			http.Error(w, errorMessage(errorParseURL), http.StatusInternalServerError)
			return
		}

		if strings.Contains(r.URL.Path, "graphql") {
			dgraphURL.Path = "/graphql"
		} else if strings.Contains(r.URL.Path, "admin/schema") {
			dgraphURL.Path = "/admin/schema"
		}

		r.URL = dgraphURL

		resp, err := c.httpClient.Do(r)

		if err != nil {
			http.Error(w, errorMessage(errorHTTPDo), http.StatusInternalServerError)
			return
		}

		statusCode := resp.StatusCode
		if statusCode != http.StatusOK {
			http.Error(w, errorMessage(errorResponseStatusCode), http.StatusInternalServerError)
			return
		}

		io.Copy(w, resp.Body)
	})
}

func errorMessage(msg string) string {
	return fmt.Sprintf(`{"message":"%s"}`, msg)
}

// notes:
// [ ] https://golang.org/pkg/net/url/#Parse
// [ ] https://medium.com/@xoen/golang-read-from-an-io-readwriter-without-loosing-its-content-2c6911805361
