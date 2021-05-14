package graphql

import "testing"

func TestNew(t *testing.T) {
	graphqlURL := "example.com"
	output := NewClient(graphqlURL)
	client := output.(*Client)
	if client.graphqlURL != graphqlURL || client.httpClient == nil {
		t.Errorf("erorr creating client, received: %+v\n", client)
	}
}
