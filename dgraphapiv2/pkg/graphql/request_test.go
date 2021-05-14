package graphql

import "testing"

func TestNewRequest(t *testing.T) {
	url := "example.com"
	token := "jwt"
	query := "graphql query"
	variables := map[string]string{
		"key": "value",
	}

	req, err := NewRequest(url, token, query, variables)
	if err != nil {
		t.Errorf("error received: %s", err.Error())
	}

	receivedToken := req.Header.Get("X-Auth0-Token")
	if receivedToken != token {
		t.Errorf("incorrect token, received: %s, expected: %s", receivedToken, token)
	}

	receivedURL := req.URL.String()
	if receivedURL != url {
		t.Errorf("incorrect url, received: %s, expected: %s", receivedURL, url)
	}
}
