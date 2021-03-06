package subscription

import (
	"encoding/json"
	"errors"
	"log"
	"net/url"

	"github.com/aws/aws-lambda-go/events"

	"github.com/forstmeier/quehook/table"
)

func createResponse(code int, msg string) (events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{
		StatusCode:      code,
		Body:            msg,
		IsBase64Encoded: false,
	}

	if msg == "success" {
		return resp, nil
	}

	return resp, errors.New(msg)
}

type sub struct {
	QueryName        string `json:"query_name"`
	SubscriberEmail  string `json:"subscriber_email"`
	SubscriberName   string `json:"subscriber_name"`
	SubscriberTarget string `json:"subscriber_target"`
}

// Subscribe adds a new subscriber to webhook query events
func Subscribe(request events.APIGatewayProxyRequest, t table.Table) (events.APIGatewayProxyResponse, error) {
	log.Printf("request: %+v\n", request)
	s := sub{}
	if err := json.Unmarshal([]byte(request.Body), &s); err != nil {
		return createResponse(400, "error parsing request: "+err.Error())
	}

	if _, err := url.ParseRequestURI(s.SubscriberTarget); err != nil {
		return createResponse(400, "error parsing url: "+err.Error())
	}

	output, err := t.Get("queries", s.QueryName, "query_name")
	if err != nil {
		return createResponse(500, "error getting query: "+err.Error())
	} else if len(output) == 0 {
		return createResponse(500, "no output returned")
	}
	log.Println("get output:", output)

	if err := t.Add("subscribers", s.QueryName, s.SubscriberEmail, s.SubscriberName, s.SubscriberTarget); err != nil {
		return createResponse(500, "error adding subscriber: "+err.Error())
	}

	return createResponse(200, "success")
}

// Unsubscribe removes a subscriber from webhook query events
func Unsubscribe(request events.APIGatewayProxyRequest, t table.Table) (events.APIGatewayProxyResponse, error) {
	log.Printf("request: %+v\n", request)
	s := &sub{}
	if err := json.Unmarshal([]byte(request.Body), &s); err != nil {
		return createResponse(400, "error parsing request: "+err.Error())
	}

	if err := t.Remove("subscribers", s.QueryName, s.SubscriberEmail); err != nil {
		return createResponse(500, "error removing subscriber: "+err.Error())
	}

	return createResponse(200, "success")
}
