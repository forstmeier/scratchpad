package contproc

import (
	"context"

	"github.com/google/uuid"
	"github.com/permissionguru/api/pkg/config"
	"github.com/permissionguru/api/pkg/docpars"
)

var _ Processor = &Client{}

// Client implements the contproc.Processor.Process methods.
type Client struct {
	matchTexts func([]string, string) string
}

// New generates a Client pointer instance using the closestmatch
// fuzzy matching package.
func New() *Client {
	return &Client{
		matchTexts: matchTexts,
	}
}

// Process implements the contproc.Processor.Process interface method
// and expects content and cfg to be populated.
func (c *Client) Process(ctx context.Context, content docpars.Content, cfg *config.Config) (*Result, error) {
	result := &Result{
		ID:             uuid.NewString(),
		IsPermissioned: true,
	}

	if cfg == nil {
		for _, selection := range content.Selections {
			if selection.Selected == false {
				result.IsPermissioned = false
				return result, nil
			}
		}

		return result, nil
	}

	contentSelecteds := make(map[string]bool, len(content.Selections))
	contentTexts := make([]string, len(content.Selections))
	for _, selection := range content.Selections {
		contentSelecteds[selection.Text] = selection.Selected
		contentTexts = append(contentTexts, selection.Text)
	}

	for _, configSelection := range cfg.Selections {
		contentText := c.matchTexts(contentTexts, configSelection.Text)

		contentSelected := contentSelecteds[contentText]
		configSelected := configSelection.Selected

		if contentSelected != configSelected {
			result.IsPermissioned = false
			return result, nil
		}
	}

	return result, nil
}
