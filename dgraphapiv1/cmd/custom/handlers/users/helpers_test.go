package users

import (
	"reflect"
	"testing"
)

func Test_createUserDgraphReq(t *testing.T) {
	orgID := "id"
	email := "test@email.com"
	password := "password"
	firstName := "firstName"
	lastName := "lastName"
	auth0ID := "auth0|id"

	req := Auth0CreateUserRequest{
		OrgID:     orgID,
		Email:     email,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
	}

	received := createUserDgraphReq(req, auth0ID)
	expected := map[string]interface{}{
		"input": []map[string]interface{}{
			map[string]interface{}{
				"org": map[string]string{
					"id": orgID,
				},
				"email":     email,
				"firstName": firstName,
				"lastName":  lastName,
				"auth0ID":   auth0ID,
			},
		},
	}

	if !reflect.DeepEqual(received, expected) {
		t.Errorf("create user request received: %v, expected: %v", received, expected)
	}
}

func Test_updateUserDgraphReq(t *testing.T) {
	password := "password"
	auth0ID := "auth0|id"

	req := Auth0UpdateUserRequest{
		Password: password,
		Auth0ID:  auth0ID,
	}

	received := updateUserDgraphReq(req)
	expected := map[string]interface{}{
		"input": map[string]interface{}{
			"filter": map[string]interface{}{
				"auth0ID": map[string]string{
					"eq": auth0ID,
				},
			},
			"set": map[string]string{
				"auth0ID": auth0ID,
			},
		},
	}

	if !reflect.DeepEqual(received, expected) {
		t.Errorf("update user request received: %v, expected: %v", received, expected)
	}
}
