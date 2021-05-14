package contproc

import (
	"context"
	"testing"

	"github.com/permissionguru/api/pkg/config"
	"github.com/permissionguru/api/pkg/docpars"
)

func TestNew(t *testing.T) {
	client := New()
	if client == nil {
		t.Error("error creating processor client")
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		description string
		content     docpars.Content
		config      *config.Config
		result      *Result // keep as full Result for expansion
		error       error   // keep for error assertions
	}{
		{
			description: "no config provided and is not permissioned",
			content: docpars.Content{
				ID: "test_content_id",
				Selections: []docpars.Selection{
					{
						Selected: false,
					},
				},
			},
			config: nil,
			result: &Result{
				IsPermissioned: false,
			},
			error: nil,
		},
		{
			description: "no config provided and is permissioned",
			content: docpars.Content{
				ID: "test_content_id",
				Selections: []docpars.Selection{
					{
						Selected: true,
					},
				},
			},
			config: nil,
			result: &Result{
				IsPermissioned: true,
			},
			error: nil,
		},
		{
			description: "config provided and mismatched selections",
			content: docpars.Content{
				ID: "test_content_id",
				Selections: []docpars.Selection{
					{
						Text:     "example selection text",
						Selected: true,
					},
				},
			},
			config: &config.Config{
				Selections: []config.Selection{
					{
						Text:     "example selection text",
						Selected: false,
					},
				},
			},
			result: &Result{
				IsPermissioned: false,
			},
			error: nil,
		},
		{
			description: "config provided and matching selections",
			content: docpars.Content{
				ID: "test_content_id",
				Selections: []docpars.Selection{
					{
						Text:     "example selection text",
						Selected: true,
					},
				},
			},
			config: &config.Config{
				Selections: []config.Selection{
					{
						Text:     "example selection text",
						Selected: true,
					},
				},
			},
			result: &Result{
				IsPermissioned: true,
			},
			error: nil,
		},
	}

	for _, test := range tests {
		client := New()

		result, err := client.Process(context.Background(), test.content, test.config)

		if result != nil {
			if result.IsPermissioned != test.result.IsPermissioned {
				t.Errorf("incorrect result, received: %+v, expected: %+v", result, test.result)
			}
		}

		if err != test.error {
			t.Errorf("incorrect nil error, received: %v, expected: %v", err, test.error)
		}
	}
}
