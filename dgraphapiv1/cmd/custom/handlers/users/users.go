package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/folivoralabs/api/pkg/auth/tokens"
	"github.com/folivoralabs/api/pkg/auth/user"
	"github.com/folivoralabs/api/pkg/graphql"
)

// Handler is an HTTP listener for User type @custom directive events.
//
// createUser: adds a user to Auth0 and to Dgraph with the Auth0 ID field
// editUser: updates an Auth0 user role or password in Auth0 and Dgraph
// removeUser: deletes an Auth0 user from Auth0 and Dgraph
func Handler(dgraphURL string, tokensClient tokens.Tokens, userClient user.User) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		appToken, err := tokensClient.GetAppToken()
		if err != nil {
			http.Error(w, errorGettingAppToken, http.StatusBadRequest)
			return
		}

		dgraphAppClient := graphql.New(
			dgraphURL,
			appToken,
		)

		dgraphUserClient := graphql.New(
			dgraphURL,
			r.Header.Get("X-Auth0-Token"),
		)

		response := &Response{}

		var buf bytes.Buffer

		if r.Method == http.MethodPost {
			var createUserReqJSON Auth0CreateUserRequest
			if err := json.NewDecoder(r.Body).Decode(&createUserReqJSON); err != nil {
				http.Error(w, errorIncorrectRequestBody, http.StatusBadRequest)
				return
			}

			auth0ID, err := userClient.CreateUser(
				createUserReqJSON.OrgID,
				createUserReqJSON.Email,
				createUserReqJSON.Password,
				createUserReqJSON.FirstName,
				createUserReqJSON.LastName,
			)
			if err != nil {
				http.Error(w, errorCreateAuth0User, http.StatusInternalServerError)
				return
			}

			if err := dgraphAppClient.Mutate(
				addUser,
				createUserDgraphReq(createUserReqJSON, *auth0ID),
				&response,
			); err != nil {
				http.Error(w, errorDgraphMutation, http.StatusInternalServerError)
				return
			}

			if err := json.NewEncoder(&buf).Encode(response.Body.Data.User[0]); err != nil {
				http.Error(w, errorDgraphJSONEncode, http.StatusBadRequest)
				return
			}

		} else if r.Method == http.MethodPatch {
			var updateUserReqJSON Auth0UpdateUserRequest
			if err := json.NewDecoder(r.Body).Decode(&updateUserReqJSON); err != nil {
				http.Error(w, errorIncorrectRequestBody, http.StatusBadRequest)
				return
			}

			if err := userClient.UpdateUser(
				updateUserReqJSON.Password,
				updateUserReqJSON.Auth0ID,
			); err != nil {
				http.Error(w, errorUpdateAuth0User, http.StatusInternalServerError)
				return
			}

			if err := dgraphUserClient.Mutate(
				updateUser,
				updateUserDgraphReq(updateUserReqJSON),
				&response,
			); err != nil {
				http.Error(w, errorDgraphMutation, http.StatusInternalServerError)
				return
			}

			if err := json.NewEncoder(&buf).Encode(response.Body.Data.User[0]); err != nil {
				http.Error(w, errorDgraphJSONEncode, http.StatusBadRequest)
				return
			}

		} else if r.Method == http.MethodDelete {
			var deleteUserReqJSON Auth0DeleteUserRequest
			if err := json.NewDecoder(r.Body).Decode(&deleteUserReqJSON); err != nil {
				http.Error(w, errorIncorrectRequestBody, http.StatusBadRequest)
				return
			}

			if err := userClient.DeleteUser(deleteUserReqJSON.Auth0ID); err != nil {
				http.Error(w, errorDeleteAuth0User, http.StatusInternalServerError)
				return
			}

			deleteGraphReq := map[string]interface{}{
				"filter": map[string]interface{}{
					"auth0ID": map[string]string{
						"eq": deleteUserReqJSON.Auth0ID,
					},
				},
			}
			if err := dgraphAppClient.Mutate(
				deleteUser,
				deleteGraphReq,
				&response,
			); err != nil {
				http.Error(w, errorDgraphMutation, http.StatusInternalServerError)
				return
			}

			if err := json.NewEncoder(&buf).Encode(response.Body.Data.User[0]); err != nil {
				http.Error(w, errorDgraphJSONEncode, http.StatusBadRequest)
				return
			}

		} else {
			http.Error(w, errorIncorrectHTTPMethod, http.StatusBadRequest)
			return
		}

		fmt.Fprint(w, buf.String())
	})
}
