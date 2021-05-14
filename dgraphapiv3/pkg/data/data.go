package data

import (
	"time"

	"github.com/google/uuid"
)

// Specimen represents a specimen stored in the database.
type Specimen struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Body      string    `json:"body"`
}

// NewSpecimen generates a new Specimen instance with the provided body value.
func NewSpecimen(body string) Specimen {
	return Specimen{
		ID:        uuid.NewString(),
		Timestamp: time.Now(),
		Body:      body,
	}
}
