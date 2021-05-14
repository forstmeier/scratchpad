package docpars

import "context"

// Parser defines methods for converting raw files into object
// for processing.
type Parser interface {
	Parse(ctx context.Context, doc []byte) (*Content, error)
}

// Content is a processed raw file containing any user selections
// found in the document.
type Content struct {
	ID         string      `json:"id"`
	Selections []Selection `json:"selections"`
}

// Selection is a selection option in a document and what state
// the document signer submitted it in.
type Selection struct {
	ID         string  `json:"id"`
	Text       string  `json:"text"`
	Selected   bool    `json:"selected"`
	Confidence float64 `json:"confidence"`
}
