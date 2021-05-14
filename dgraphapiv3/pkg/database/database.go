package database

import "github.com/specimenguru/api/pkg/data"

// Database defines the methods available for interacting with the
// API backend storage.
type Database interface {
	Add(input []data.Specimen) error
	Query(query Query) ([]data.Specimen, error)
}

// Query holds the content required to query the database.
type Query struct {
	FullText string // required
}
