package blevedb

import (
	"log"
	"time"

	"github.com/blevesearch/bleve/v2"

	"github.com/specimenguru/api/pkg/data"
	"github.com/specimenguru/api/pkg/database"
)

// Bleve field names
const (
	idKey        = "specimen.id"
	timestampKey = "specimen.timestamp"
	bodyKey      = "specimen.body"
)

// Query is the implementation of the database.Database.Query method.
func (c *Client) Query(query database.Query) ([]data.Specimen, error) {
	specimens := []data.Specimen{}

	if query.FullText == "" {
		return specimens, &ErrorNoQuery{}
	}

	qsQuery := bleve.NewQueryStringQuery(query.FullText)
	qsSearch := bleve.NewSearchRequest(qsQuery)
	qsSearch.Fields = []string{
		idKey,
		timestampKey,
		bodyKey,
	}

	qsResults, err := c.index.Search(qsSearch)
	if err != nil {
		return specimens, &ErrorRunBleveQuery{err: err}
	}

	for _, hit := range qsResults.Hits {
		fields := hit.Fields

		timestampField := fields[timestampKey].(string)
		timestamp, err := time.Parse(time.RFC3339, timestampField)
		if err != nil {
			log.Fatalf("error parsing timestamp: %s", err.Error())
		}

		specimen := data.Specimen{
			ID:        fields[idKey].(string),
			Timestamp: timestamp,
			Body:      fields[bodyKey].(string),
		}

		specimens = append(specimens, specimen)
	}

	return specimens, nil
}
