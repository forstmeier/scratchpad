package handlers

import (
	"fmt"
	"net/http"

	"github.com/specimenguru/api/pkg/data"
)

const (
	csvType  = "text/csv"
	jsonType = "application/json"
)

// AddHandler returns the http.HandlerFunc responsible for handling "add" requests.
func (c *Client) AddHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		specimens := []data.Specimen{}
		var err error
		fileType := ""
		if contentType == csvType {
			specimens, err = specimensFromCSV(r.Body)
			fileType = "csv"
		} else if contentType == jsonType {
			specimens, err = specimensFromJSON(r.Body)
			fileType = "json"
		} else {
			message := "no content type header provided"
			if contentType != "" {
				message = fmt.Sprintf("content type %s not supported", contentType)
			}

			http.Error(w, errorMessage(message), http.StatusBadRequest)
			return
		}

		if err != nil {
			message := fmt.Sprintf("error parsing %s file", fileType)
			http.Error(w, errorMessage(message), http.StatusInternalServerError)
			return
		}

		if err := c.db.Add(specimens); err != nil {
			http.Error(w, errorMessage("error adding specimens to database"), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, `{"message":"success"}`)
	})
}
