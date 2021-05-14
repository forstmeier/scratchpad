package user

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	c := New("auth0URL", "appToken")
	if c == nil {
		t.Error("error calling new client function")
	}
}

func TestCreateUser(t *testing.T) {
	auth0ID := "auth0|id"

	tests := []struct {
		description             string
		inputs                  []string
		auth0ReceivedBody       string
		auth0ResponseStatusCode int
		auth0ResponseBody       string
		response                *string
		error                   error
	}{
		{
			description:             "failed call create auth0 api",
			inputs:                  []string{"orgID", "test@email.com", "password", "first", "last"},
			auth0ReceivedBody:       `{"email":"test@email.com","password":"password","app_metadata":{"orgID":"orgID"},"given_name":"first","family_name":"last","connection":"Username-Password-Authentication"}`,
			auth0ResponseStatusCode: http.StatusBadRequest,
			auth0ResponseBody:       "",
			response:                nil,
			error:                   ErrorCheckStatus{},
		},
		{
			description:             "successful create user invocation",
			inputs:                  []string{"orgID", "test@email.com", "password", "first", "last"},
			auth0ReceivedBody:       `{"email":"test@email.com","password":"password","app_metadata":{"orgID":"orgID"},"given_name":"first","family_name":"last","connection":"Username-Password-Authentication"}`,
			auth0ResponseStatusCode: http.StatusOK,
			auth0ResponseBody:       fmt.Sprintf(`{"user_id": "%s"}`, auth0ID),
			response:                &auth0ID,
			error:                   nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var auth0ReceivedBody []byte

			auth0Mux := http.NewServeMux()
			auth0Mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
				receivedBytes, err := ioutil.ReadAll(r.Body)
				if err != nil {
					http.Error(w, "error parsing received body", http.StatusBadRequest)
					return
				}

				auth0ReceivedBody = receivedBytes
				w.WriteHeader(test.auth0ResponseStatusCode)
				w.Write([]byte(test.auth0ResponseBody))
			})

			auth0Server := httptest.NewServer(auth0Mux)
			auth0URL := auth0Server.URL + "/"

			client := New(auth0URL, "appToken")

			response, err := client.CreateUser(
				test.inputs[0],
				test.inputs[1],
				test.inputs[2],
				test.inputs[3],
				test.inputs[4],
			)

			if (err != nil && test.error != nil) && !errors.As(err, &test.error) {
				t.Errorf("error received: %v, expected: %v\n", err, test.error)
			}

			if !reflect.DeepEqual(response, test.response) {
				t.Errorf("response received: %+v, expected: %+v\n", response, test.response)
			}

			if auth0ReceivedBody != nil {
				receivedString := string(auth0ReceivedBody)
				if receivedString != test.auth0ReceivedBody {
					t.Errorf("auth0 body received: %s, expected: %s", receivedString, test.auth0ReceivedBody)
				}
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	auth0ID := "auth0|id"

	tests := []struct {
		description             string
		inputs                  []string
		auth0ReceivedBody       string
		auth0ResponseStatusCode int
		auth0ResponseBody       string
		error                   error
	}{
		{
			description:             "error invoking update auth0 api",
			inputs:                  []string{"password", auth0ID},
			auth0ReceivedBody:       `{"password":"password"}`,
			auth0ResponseStatusCode: http.StatusBadRequest,
			auth0ResponseBody:       "",
			error:                   ErrorCheckStatus{},
		},
		{
			description:             "successful update user password invocation",
			inputs:                  []string{"password", auth0ID},
			auth0ReceivedBody:       `{"password":"password"}`,
			auth0ResponseStatusCode: http.StatusOK,
			auth0ResponseBody:       fmt.Sprintf(`{"user_id": "%s"}`, auth0ID),
			error:                   nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var auth0ReceivedBody []byte

			auth0Mux := http.NewServeMux()
			auth0Mux.HandleFunc("/users/"+auth0ID, func(w http.ResponseWriter, r *http.Request) {
				receivedBytes, err := ioutil.ReadAll(r.Body)
				if err != nil {
					http.Error(w, "error parsing received body", http.StatusBadRequest)
					return
				}

				auth0ReceivedBody = receivedBytes
				w.WriteHeader(test.auth0ResponseStatusCode)
				w.Write([]byte(test.auth0ResponseBody))
			})

			auth0Server := httptest.NewServer(auth0Mux)
			auth0URL := auth0Server.URL + "/"

			client := New(auth0URL, "appToken")

			err := client.UpdateUser(test.inputs[0], test.inputs[1])

			if (err != nil && test.error != nil) && !errors.As(err, &test.error) {
				t.Errorf("error received: %v, expected: %v\n", err, test.error)
			}

			if auth0ReceivedBody != nil {
				receivedString := string(auth0ReceivedBody)
				if receivedString != test.auth0ReceivedBody {
					t.Errorf("auth0 body received: %s, expected: %s", receivedString, test.auth0ReceivedBody)
				}
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	auth0ID := "auth0|id"

	tests := []struct {
		description             string
		input                   string
		auth0ReceivedBody       string
		auth0ResponseStatusCode int
		auth0ResponseBody       string
		error                   error
	}{
		{
			description:             "failed call delete auth0 api",
			input:                   auth0ID,
			auth0ReceivedBody:       "",
			auth0ResponseStatusCode: http.StatusBadRequest,
			auth0ResponseBody:       "",
			error:                   ErrorCheckStatus{},
		},
		{
			description:             "successful delete user invocation",
			input:                   auth0ID,
			auth0ReceivedBody:       "",
			auth0ResponseStatusCode: http.StatusOK,
			auth0ResponseBody:       "",
			error:                   nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var auth0ReceivedBody []byte

			auth0Mux := http.NewServeMux()
			auth0Mux.HandleFunc("/users/"+auth0ID, func(w http.ResponseWriter, r *http.Request) {
				receivedBytes, err := ioutil.ReadAll(r.Body)
				if err != nil {
					http.Error(w, "error parsing received body", http.StatusBadRequest)
					return
				}

				auth0ReceivedBody = receivedBytes
				w.WriteHeader(test.auth0ResponseStatusCode)
				w.Write([]byte(test.auth0ResponseBody))
			})

			auth0Server := httptest.NewServer(auth0Mux)
			auth0URL := auth0Server.URL + "/"

			client := New(auth0URL, "appToken")

			err := client.DeleteUser(test.input)

			if (err != nil && test.error != nil) && !errors.As(err, &test.error) {
				t.Errorf("error received: %v, expected: %v\n", err, test.error)
			}

			if auth0ReceivedBody != nil {
				receivedString := string(auth0ReceivedBody)
				if receivedString != test.auth0ReceivedBody {
					t.Errorf("auth0 body received: %s, expected: %s", receivedString, test.auth0ReceivedBody)
				}
			}
		})
	}
}
