package main

import (
	"log"
	"os"
	"time"

	"github.com/blevesearch/bleve/v2"
)

// Specimen is the root specimen.
type Specimen struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Body      string    `json:"body"`
}

// BleveSpecimen wraps the root specimen.
type BleveSpecimen struct {
	// Specimen `json:"specimen"`
	Specimen Specimen `json:"specimen"`
}

// Type implements the bleve.Classifier interface.
func (bs *BleveSpecimen) Type() string {
	return "bleve_specimen"
}

var now = time.Now()

var specimens = []Specimen{
	{ID: "one", Timestamp: now, Body: `{"text":"the quick brown fox jumped over the lazy dog"}`},
	{ID: "two", Timestamp: now, Body: `{"text":"jump over the brick wall"}`},
	{ID: "three", Timestamp: now, Body: `{"text":"carnivours are delicious"}`},
}

const (
	idKey        = "specimen.id"
	timestampKey = "specimen.timestamp"
	bodyKey      = "specimen.body"
)

func main() {
	specimenMapping := bleve.NewDocumentMapping()

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

	name := "testing.bleve"
	index, err := bleve.New(name, indexMapping)
	if err != nil {
		log.Fatalf("error creating index: %s", err.Error())
	}
	defer os.RemoveAll(name)

	batch := index.NewBatch()
	for _, specimen := range specimens {
		batch.Index(specimen.ID, &BleveSpecimen{
			Specimen: specimen,
		})
	}

	if err := index.Batch(batch); err != nil {
		log.Fatalf("error calling batch: %s", err.Error())
	}

	query := bleve.NewQueryStringQuery("carnivours")
	search := bleve.NewSearchRequest(query)
	search.Fields = []string{
		idKey,
		timestampKey,
		bodyKey,
	}
	searchResults, err := index.Search(search)
	if err != nil {
		log.Fatalf("error running query: %s", err.Error())
	}

	// log.Printf("results: %+v", searchResults.Hits[0].Fields)

	hits := searchResults.Hits

	for _, hit := range hits {
		fields := hit.Fields
		idValue := fields[idKey].(string)
		timestampValue := fields[timestampKey].(string)

		timestamp, err := time.Parse(time.RFC3339, timestampValue)
		if err != nil {
			log.Fatalf("error parsing timestamp: %s", err.Error())
		}

		bodyValue := fields[bodyKey].(string)

		log.Println("idValue:", idValue)
		log.Println("timestampValue:", timestamp)
		log.Println("bodyValue:", bodyValue)
	}
}
