package graphql

import "net/http"

// GraphQL provides wrapper methods around HTTP requests sent to the
// configured Dgraph URL endpoint.
type GraphQL interface {
	Mutate(request *http.Request, response interface{}) error
}

// Client implements the GraphQL interface.
type Client struct {
	graphqlURL string
	httpClient *http.Client
}

// NewClient generates a pointer Client instance.
func NewClient(graphqlURL string) GraphQL {
	return &Client{
		graphqlURL: graphqlURL,
		httpClient: &http.Client{},
	}
}
