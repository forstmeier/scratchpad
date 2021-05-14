package users

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockTokensClient struct {
	managementToken      string
	managementTokenError error
}

func (m *mockTokensClient) GetAppToken() (string, error) {
	return m.managementToken, m.managementTokenError
}

func (m *mockTokensClient) GetUserToken(user string) (string, error) {
	return "userToken", nil
}

func (m *mockTokensClient) UpdateUserToken(user, orgID, managementToken string) error {
	return nil
}

type mockUserClient struct {
	auth0ID string
}

func (m *mockUserClient) CreateUser(orgID, email, password, firstName, lastName string) (*string, error) {
	return &m.auth0ID, nil
}

func (m *mockUserClient) UpdateUser(password, auth0ID string) error {
	return nil
}

func (m *mockUserClient) DeleteUser(auth0ID string) error {
	return nil
}

func TestHandler(t *testing.T) {
	auth0ID := "auth0|0123456789"

	tests := []struct {
		description        string
		requestMethod      string
		requestBody        []byte
		dgraphBody         string
		responseStatusCode int
		responseBody       string
	}{
		{
			description:        "invalid json body received in request to custom",
			requestMethod:      http.MethodPost,
			requestBody:        []byte("---------"),
			dgraphBody:         "{}",
			responseStatusCode: http.StatusBadRequest,
			responseBody:       errorIncorrectRequestBody,
		},
		{
			description:        "unsupported http method in request to custom",
			requestMethod:      http.MethodPut,
			requestBody:        []byte(`{"email": "grandmaster@jeditemple.edu"}`),
			dgraphBody:         "{}",
			responseStatusCode: http.StatusBadRequest,
			responseBody:       errorIncorrectHTTPMethod,
		},
		{
			description:        "successful create user request to custom server",
			requestMethod:      http.MethodPost,
			requestBody:        []byte(`{"orgID":"jedi","email":"grandmaster@jeditemple.edu","password":"may-the-force-be-with-you","firstName":"yoda","lastName":"..."}`),
			dgraphBody:         fmt.Sprintf(`{"data": {"data": {"user": [{"id": "jedi_0", "org": {"id": "jedi"}, "email": "grandmaster@jeditemple.edu", "firstName": "yoda", "lastName": "...", "auth0ID": "%s"}]}}}`, auth0ID),
			responseStatusCode: http.StatusOK,
			responseBody:       fmt.Sprintf(`{"id":"jedi_0","org":{"id":"jedi"},"email":"grandmaster@jeditemple.edu","firstName":"yoda","lastName":"...","auth0ID":"%s"}`, auth0ID),
		},
		{
			description:        "successful update user request to custom server",
			requestMethod:      http.MethodPatch,
			requestBody:        []byte(fmt.Sprintf(`{"authZeroID":"%s","password":"new_password"}`, auth0ID)),
			dgraphBody:         fmt.Sprintf(`{"data": {"data": {"user": [{"id": "jedi_0", "org": {"id": "jedi"}, "email": "grandmaster@jeditemple.edu", "firstName": "yoda", "lastName": "...", "auth0ID": "%s"}]}}}`, auth0ID),
			responseStatusCode: http.StatusOK,
			responseBody:       fmt.Sprintf(`{"id":"jedi_0","org":{"id":"jedi"},"email":"grandmaster@jeditemple.edu","firstName":"yoda","lastName":"...","auth0ID":"%s"}`, auth0ID),
		},
		{
			description:        "successful delete user request to custom server",
			requestMethod:      http.MethodDelete,
			requestBody:        []byte(fmt.Sprintf(`{"authZeroID":"%s"}`, auth0ID)),
			dgraphBody:         fmt.Sprintf(`{"data": {"data": {"user": [{"id": "jedi_0", "org": {"id": "jedi"}, "email": "grandmaster@jeditemple.edu", "firstName": "yoda", "lastName": "...", "auth0ID": "%s"}]}}}`, auth0ID),
			responseStatusCode: http.StatusOK,
			responseBody:       fmt.Sprintf(`{"id":"jedi_0","org":{"id":"jedi"},"email":"grandmaster@jeditemple.edu","firstName":"yoda","lastName":"...","auth0ID":"%s"}`, auth0ID),
		},
	}

	for _, test := range tests {
		dgraphMux := http.NewServeMux()
		dgraphMux.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(test.dgraphBody))
		})

		dgraphServer := httptest.NewServer(dgraphMux)
		dgraphURL := dgraphServer.URL + "/graphql"

		t.Run(test.description, func(t *testing.T) {
			// mock request from dgraph @custom directive to custom server/handler
			req, err := http.NewRequest(
				test.requestMethod,
				"/users",
				bytes.NewReader(test.requestBody),
			)
			if err != nil {
				t.Fatal("error creating request:", err.Error())
			}

			rec := httptest.NewRecorder()

			// custom users handler wrapper function
			handler := http.HandlerFunc(Handler(dgraphURL, &mockTokensClient{}, &mockUserClient{auth0ID: auth0ID}))

			handler.ServeHTTP(rec, req)

			if status := rec.Code; status != test.responseStatusCode {
				t.Errorf("status received: %d, expected: %d", status, test.responseStatusCode)
			}

			if body := strings.TrimSuffix(rec.Body.String(), "\n"); body != test.responseBody {
				t.Errorf("body received: %s, expected: %s", body, test.responseBody)
			}
		})
	}
}
