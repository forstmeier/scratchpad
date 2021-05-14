package tokens

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

type responseJSON struct {
	AccessToken string `json:"access_token"`
}

// GetAppToken fetches a token from the Auth0 Application with Management
// API access.
func (c *Client) GetAppToken() (string, error) {
	api := c.config.Auth0.API
	payloadString := fmt.Sprintf(
		"grant_type=client_credentials&client_id=%s&client_secret=%s&audience=%s",
		api.ClientID,
		api.ClientSecret,
		c.config.Auth0.AudienceURL,
	)

	payload := strings.NewReader(payloadString)
	req, err := http.NewRequest("POST", c.config.Auth0.TokenURL, payload)
	if err != nil {
		return "", newErrorNewRequest(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.client.Do(req)
	if err != nil {
		return "", newErrorClientDo(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", newErrorReadResponseBody(err)
	}

	appToken := gjson.Get(string(bodyBytes), "access_token")
	return appToken.String(), nil
}
