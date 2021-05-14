package handlers

import (
	"testing"

	"github.com/specimenguru/api/pkg/data"
	"github.com/specimenguru/api/pkg/database"
)

type testDatabase struct {
	addError       error
	querySpecimens []data.Specimen
	queryError     error
}

func (db *testDatabase) Add(input []data.Specimen) error {
	return db.addError
}

func (db *testDatabase) Query(query database.Query) ([]data.Specimen, error) {
	return db.querySpecimens, db.queryError
}

func TestNewClient(t *testing.T) {
	db := &testDatabase{}

	client, err := NewClient(db)
	if err != nil {
		t.Errorf("error creating new client: %s", err.Error())
	}

	if client == nil {
		t.Errorf("incorrect client, received: %+v", client)
	}
}
