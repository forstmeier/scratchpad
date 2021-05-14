package blevedb

import (
	"os"

	"github.com/blevesearch/bleve/v2"
	"github.com/specimenguru/api/pkg/database"
)

var _ database.Database = &Client{}

// Client implements the Database interface.
type Client struct {
	index bleve.Index
}

// New returns a new instance of a Client pointer.
func New(filename string) (*Client, error) {
	var index bleve.Index
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		specimenMapping := bleve.NewDocumentMapping()

		// declaring field mapping individually for eventual customization
		idFieldMapping := bleve.NewTextFieldMapping()
		timestampFieldMapping := bleve.NewDateTimeFieldMapping()
		bodyFieldMapping := bleve.NewTextFieldMapping()

		specimenMapping.AddFieldMappingsAt("id", idFieldMapping)
		specimenMapping.AddFieldMappingsAt("timestamp", timestampFieldMapping)
		specimenMapping.AddFieldMappingsAt("body", bodyFieldMapping)

		bleveSpecimenMapping := bleve.NewDocumentMapping()
		bleveSpecimenMapping.AddSubDocumentMapping("specimen", specimenMapping)

		indexMapping := bleve.NewIndexMapping()
		indexMapping.AddDocumentMapping("bleve_specimen", bleveSpecimenMapping)

		index, err = bleve.New(filename, indexMapping)
		if err != nil {
			return nil, &ErrorNewBleveIndex{err: err}
		}
	} else {
		index, err = bleve.Open(filename)
		return nil, &ErrorOpenBleveIndex{err: err}
	}

	return &Client{
		index: index,
	}, nil
}
