package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/specimenguru/api/pkg/database"
)

const fullText = "full_text"

// QueryHandler returns the http.HandlerFunc responsible for handling "query" requests.
func (c *Client) QueryHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		queryText, ok := r.URL.Query()[fullText]
		if !ok {
			message := "no query url parameter provided"
			http.Error(w, errorMessage(message), http.StatusBadRequest)
			return
		}

		query := database.Query{
			FullText: queryText[0],
		}

		specimens, err := c.db.Query(query)
		if err != nil {
			message := "error querying specimens in database"
			http.Error(w, errorMessage(message), http.StatusInternalServerError)
			return
		}

		bodies := []string{}
		for _, specimen := range specimens {
			bodies = append(bodies, specimen.Body)
		}

		response := fmt.Sprintf("[%s]", strings.Join(bodies, ","))

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(response))
	})
}
