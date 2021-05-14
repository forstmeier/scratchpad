package query

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"cloud.google.com/go/bigquery"
	"github.com/aws/aws-lambda-go/events"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	"github.com/forstmeier/quehook/storage"
	"github.com/forstmeier/quehook/table"
)

func createResponse(code int, msg string) (events.APIGatewayProxyResponse, error) {
	log.Printf("response code: %d, message: %s\n", code, msg)
	resp := events.APIGatewayProxyResponse{
		StatusCode:      code,
		Body:            msg,
		IsBase64Encoded: false,
	}

	if msg == "success" || msg == "query already exists" {
		return resp, nil
	}

	return resp, errors.New(msg)
}

// Create adds a query to S3 for periodic execution
func Create(request events.APIGatewayProxyRequest, t table.Table, s storage.Storage) (events.APIGatewayProxyResponse, error) {
	headers := request.Headers
	log.Printf("request: %+v\n", request)
	queryName := headers["query-name"]
	queryDescription := headers["query-description"]
	queryAuthor := headers["query-author"]
	queryAuthorEmail := headers["query-author-email"]
	log.Printf("query name: %s, description: %s, author: %s, author email: %s\n", queryName, queryDescription, queryAuthor, queryAuthorEmail)

	output, err := t.Get("queries", queryName, "query_name")
	if err != nil {
		return createResponse(500, "error getting query table: "+err.Error())
	}
	log.Println("get output:", output)

	if len(output) == 0 {
		if err := t.Add("queries", queryName, queryAuthor, queryAuthorEmail); err != nil {
			return createResponse(500, "error creating query: "+err.Error())
		}

		if err := s.PutFile("queries/"+queryName, strings.NewReader(request.Body)); err != nil {
			return createResponse(500, "error putting query file: "+err.Error())
		}
	} else {
		log.Println("query already exists")
		return createResponse(200, "query already exists")
	}

	log.Println("create success")
	return createResponse(200, "success")
}

var query = func(s storage.Storage, q string, rows *[][]bigquery.Value) error {
	credentials, err := s.GetFile("quehook_credentials.json")
	if err != nil {
		return errors.New("error getting credentials file: " + err.Error())
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(credentials)
	opts := option.WithCredentialsJSON(buf.Bytes())

	client, err := bigquery.NewClient(context.Background(), "quehook", opts)
	if err != nil {
		return errors.New("error creating bigquery client: " + err.Error())
	}

	qry := client.Query(q)
	itr, err := qry.Read(context.Background())
	if err != nil {
		return errors.New("error reading query: " + err.Error())
	}

	for {
		var row []bigquery.Value
		err := itr.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return errors.New("error iterating query results: " + err.Error())
		}

		*rows = append(*rows, row)
	}

	return nil
}

// Run executes all stored queries and returns results to subscribers
func Run(s storage.Storage, t table.Table) (events.APIGatewayProxyResponse, error) {
	log.Printf("query request triggered")
	queries, err := s.GetPaths()
	if err != nil {
		return createResponse(500, "error listing query files: "+err.Error())
	}
	log.Printf("queries: %s\n", queries)

	for _, q := range queries {
		file, err := s.GetFile(q)
		if err != nil {
			return createResponse(500, "error getting query file: "+err.Error())
		}

		buf := new(bytes.Buffer)
		buf.ReadFrom(file)
		log.Printf("query content: %s\n", buf.String())

		rows := [][]bigquery.Value{}
		if err := query(s, buf.String(), &rows); err != nil {
			return createResponse(500, err.Error())
		}

		output, err := json.Marshal(rows)
		if err != nil {
			return createResponse(500, "error marshalling output: "+err.Error())
		}
		log.Printf("query output: %s\n", string(output))

		subscribers, err := t.Get("subscribers", q, "subscriber_target")
		if err != nil {
			return createResponse(500, "error getting subscribers: "+err.Error())
		}
		log.Printf("subscribers: %s\n", subscribers)

		client := &http.Client{}
		for _, subscriber := range subscribers {
			req, err := http.NewRequest("POST", subscriber, bytes.NewBuffer(output))
			req.Header.Set("Content-Type", "application/json")
			resp, err := client.Do(req)
			if err != nil {
				return createResponse(500, "error posting results: "+err.Error())
			}
			_ = resp // TEMP
		}
	}

	return createResponse(200, "success")
}

// Delete removes a query from S3 - internal use only
func Delete(request events.APIGatewayProxyRequest, t table.Table, s storage.Storage) (events.APIGatewayProxyResponse, error) {
	log.Printf("request: %+v\n", request)
	if request.Headers["quehook_secret"] != os.Getenv("quehook_secret") {
		return createResponse(500, "incorrect secret received: "+request.Headers["quehook_secret"])
	}

	body := struct {
		QueryName string `json:"query_name"`
	}{}

	if err := json.Unmarshal([]byte(request.Body), &body); err != nil {
		return createResponse(500, "error parsing request body: "+err.Error())
	}
	log.Printf("query name: %s\n", body.QueryName)

	output, err := t.Get("queries", body.QueryName, "query_name")
	if err != nil {
		return createResponse(500, "error getting query: "+err.Error())
	}

	if len(output) > 0 {
		if err := s.DeleteFile("queries/" + body.QueryName); err != nil {
			return createResponse(500, "error deleting query file: "+err.Error())
		}
		log.Println("deleted query file")

		if err := t.Remove("queries", body.QueryName, ""); err != nil {
			return createResponse(500, "error removing query item: "+err.Error())
		}
		log.Println("removed query entry")

		if err := t.Remove("subscribers", body.QueryName, ""); err != nil {
			return createResponse(500, "error removing subscribers items: "+err.Error())
		}
		log.Println("removed subscriber entries")
	}

	return createResponse(200, "success")
}
