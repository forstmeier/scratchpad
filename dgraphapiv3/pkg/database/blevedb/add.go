package blevedb

import "github.com/specimenguru/api/pkg/data"

// Add is the implementation of the database.Database.Add method.
func (c *Client) Add(input []data.Specimen) error {
	if len(input) < 1 {
		return &ErrorNoAddData{}
	}

	batch := c.index.NewBatch()
	for _, specimen := range input {
		batch.Index(specimen.ID, &BleveSpecimen{
			Specimen: specimen,
		})
	}

	return c.index.Batch(batch)
}
