package specimens

import (
	"bytes"
	"errors"
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

func TestHandler(t *testing.T) {
	tests := []struct {
		description          string
		managementToken      string
		managementTokenError error
		requestBody          []byte
		dgraphBody           string
		responseStatusCode   int
		responseBody         string
	}{
		{
			description:          "error getting management token",
			managementToken:      "",
			managementTokenError: errors.New("mock management token error"),
			requestBody:          []byte("{}"),
			dgraphBody:           `{}`,
			responseStatusCode:   http.StatusBadRequest,
			responseBody:         errorGettingAppToken,
		},
		{
			description:          "invalid json body received in request to custom",
			managementToken:      "managementToken",
			managementTokenError: nil,
			requestBody:          []byte("---------"),
			dgraphBody:           `{}`,
			responseStatusCode:   http.StatusBadRequest,
			responseBody:         errorIncorrectRequestBody,
		},
		{
			description:          "successful add blood specimen update request to custom server",
			managementToken:      "managementToken",
			managementTokenError: nil,
			requestBody:          []byte(`{"orgID": "imp-med-cen", "bloodSpecimen": "vader", "bloodType": "A_NEG", "container": "VIAL", "volume": 10.0, "description": "sith lord blood sample"}`),
			dgraphBody:           `{"data": {"data": {"bloodSpecimenUpdate": [{"id": "midichlorian_0"}]}}}`,
			responseStatusCode:   http.StatusOK,
			responseBody:         `{"id":"midichlorian_0"}`,
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
			req, err := http.NewRequest(
				http.MethodPost,
				"/specimens",
				bytes.NewReader(test.requestBody),
			)
			if err != nil {
				t.Fatal("error creating request:", err.Error())
			}

			rec := httptest.NewRecorder()

			// custom specimens handler wrapper function
			handler := http.HandlerFunc(Handler(dgraphURL, &mockTokensClient{
				managementToken:      test.managementToken,
				managementTokenError: test.managementTokenError,
			}))

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
