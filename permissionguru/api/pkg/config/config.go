package config

// Config is a user-provided specificactions for how specific
// user selection options should be handled.
type Config struct {
	ID         string      `json:"id"`
	Selections []Selection `json:"selections"`
}

// Selection defines a specific section of document text and whether
// or not it should be selected by the document signer.
type Selection struct {
	Text     string `json:"text"`
	Selected bool   `json:"selected"`
}
