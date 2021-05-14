package tokens

import (
	"net/http"

	"github.com/folivoralabs/api/pkg/config"
)

// Tokens exposes helper methods for retrieving and updating tokens
// via the Auth0 API.
type Tokens interface {
	GetAppToken() (string, error)
	GetUserToken(user string) (string, error)
	UpdateUserToken(user, orgID, managementToken string) error
}

// Client is the implementation of the Tokens interface.
type Client struct {
	client *http.Client
	config *config.Config
}

// New generates a pointer instance of the Client.
func New(cfg *config.Config) Tokens {
	return &Client{
		client: &http.Client{},
		config: cfg,
	}
}
