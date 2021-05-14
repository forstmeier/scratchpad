package user

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func checkCreateUserValues(orgID, email, password, firstName, lastName string) (bool, string) {
	fields := []string{}
	if orgID == "" {
		fields = append(fields, "orgID")
	}
	if email == "" {
		fields = append(fields, "email")
	}
	if password == "" {
		fields = append(fields, "password")
	}
	if firstName == "" {
		fields = append(fields, "firstName")
	}
	if lastName == "" {
		fields = append(fields, "lastName")
	}

	if len(fields) > 0 {
		return false, strings.Join(fields, ", ")
	}

	return true, ""
}

func checkUpdateUserValues(password, auth0ID string) (bool, string) {
	fields := []string{}
	if password == "" {
		fields = append(fields, "password")
	}
	if auth0ID == "" {
		fields = append(fields, "auth0ID")
	}

	if len(fields) > 0 {
		return false, strings.Join(fields, ", ")
	}

	return true, ""
}

func (c *Client) newRequest(payload interface{}, method, url string) (*http.Request, error) {
	var req *http.Request
	var err error

	if payload == nil {
		req, err = http.NewRequest(method, url, nil)
	} else {
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			return nil, newErrorMarshalRequest(err)
		}

		req, err = http.NewRequest(method, url, bytes.NewReader(payloadBytes))
	}
	if err != nil {
		return nil, newErrorCreateRequest(err)
	}

	req.Header.Set("Authorization", "Bearer "+c.appToken)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *Client) sendRequest(req *http.Request) ([]byte, error) {
	auth0Resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, newErrorClientDo(err)
	}

	if !checkSuccess(auth0Resp.StatusCode) {
		return nil, newErrorCheckStatus(auth0Resp.StatusCode)
	}
	defer auth0Resp.Body.Close()

	auth0RespBytes, err := ioutil.ReadAll(auth0Resp.Body)
	if err != nil {
		return nil, newErrorResponseReadBody(err)
	}
	return auth0RespBytes, nil
}

func checkSuccess(status int) bool {
	return status == http.StatusOK || status == http.StatusCreated || status == http.StatusNoContent
}
