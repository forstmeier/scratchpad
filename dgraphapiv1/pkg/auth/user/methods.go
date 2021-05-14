package user

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// CreateUser adds a new user to the Auth0 account and the Dgraph database.
//
// This method is invoked by the users handler responding to the @custom
// customAddUser directive.
func (c *Client) CreateUser(orgID, email, password, firstName, lastName string) (*string, error) {
	if set, fields := checkCreateUserValues(orgID, email, password, firstName, lastName); !set {
		return nil, newErrorCreateUserFieldsNotSet(fields)
	}

	createUserJSON := CreateUserRequest{
		Email:    email,
		Password: password,
		AppMetadata: AppMetadata{
			OrgID: orgID,
		},
		FirstName:  firstName,
		LastName:   lastName,
		Connection: "Username-Password-Authentication",
	}

	auth0CreateURL := c.auth0URL + "users"
	auth0Req, err := c.newRequest(createUserJSON, http.MethodPost, auth0CreateURL)
	if err != nil {
		return nil, err
	}

	auth0RespBytes, err := c.sendRequest(auth0Req)

	resp := CreateUserResponse{}
	if err := json.Unmarshal(auth0RespBytes, &resp); err != nil {
		return nil, newErrorResponseUnmarshal(err)
	}

	return &resp.Auth0ID, nil
}

// UpdateUser updates an existing user in the Auth0 account and the Dgraph database.
//
// This method is invoked by the users handler responding to the @custom
// customUpdateUser directive.
func (c *Client) UpdateUser(password, auth0ID string) error {
	if set, fields := checkUpdateUserValues(password, auth0ID); !set {
		return newErrorUpdateUserFieldsNotSet(fields)
	}

	updateUserJSON := UpdateUserRequest{
		Password: password,
	}

	auth0UpdateURL := c.auth0URL + "users/" + url.PathEscape(auth0ID)
	auth0Req, err := c.newRequest(updateUserJSON, http.MethodPatch, auth0UpdateURL)
	if err != nil {
		return newErrorCreateRequest(err)
	}

	_, err = c.sendRequest(auth0Req)
	return err
}

// DeleteUser deletes an existing user from the Auth0 account and the Dgraph database.
//
// This method is invoked by the users handler responding to the @custom
// customDeleteUser directive.
func (c *Client) DeleteUser(auth0ID string) error {
	if auth0ID == "" {
		return newErrorDeleteUserFieldNotSet()
	}

	auth0DeleteURL := c.auth0URL + "users/" + url.PathEscape(auth0ID)
	auth0Req, err := c.newRequest(nil, http.MethodDelete, auth0DeleteURL)
	if err != nil {
		return newErrorCreateRequest(err)
	}

	_, err = c.sendRequest(auth0Req)
	return err
}
