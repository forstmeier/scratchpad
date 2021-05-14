package user

import (
	"net/http"
	"time"
)

// User defines methods for interacting with the external user
// management service.
//
// These methods are invoked by the @custom directive handlers.
//
// NOTE: this is not a generalized interface; ideally, it would be
// more flexibly defined to allow for other non-Auth0 implementations
// but this should be fine for the time being.
type User interface {
	CreateUser(orgID, email, password, firstName, lastName string) (*string, error)
	UpdateUser(password, auth0ID string) error
	DeleteUser(auth0ID string) error
}

// Client is the implementation of the Tokens interface.
//
// auth0URL: root Auth0 API URL - specific paths are joined on
//           this base for various API calls
// appToken: the Auth0 API management token
type Client struct {
	auth0URL   string
	appToken   string
	httpClient *http.Client
}

// New generates a pointer instance of the Client.
func New(auth0URL, appToken string) User {
	return &Client{
		auth0URL: auth0URL,
		appToken: appToken,
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}
