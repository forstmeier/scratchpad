package handlers

import "github.com/specimenguru/api/pkg/database"

// Client holds external resources for use by the various HTTP
// handler methods defined on it.
type Client struct {
	db database.Database
}

// NewClient generates a Client pointer with the provided db instance.
func NewClient(db database.Database) (*Client, error) {
	return &Client{
		db: db,
	}, nil
}
