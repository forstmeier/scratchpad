package blevedb

import "github.com/specimenguru/api/pkg/data"

// BleveSpecimen wraps the general specimen type for interface implementation.
type BleveSpecimen struct {
	Specimen data.Specimen `json:"specimen"`
}

// Type implements the bleve.Classifier interface for document mapping.
func (bs *BleveSpecimen) Type() string {
	return "bleve_specimen"
}
