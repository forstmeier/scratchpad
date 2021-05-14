package contproc

import (
	"context"

	"github.com/permissionguru/api/pkg/config"
	"github.com/permissionguru/api/pkg/docpars"
)

// Processor defines methods for determining the selection status
// of the provided content and optional user-provided configuration.
type Processor interface {
	Process(ctx context.Context, content docpars.Content, cfg *config.Config) (*Result, error)
}

// Result is the output from processing a given file's content.
type Result struct {
	ID             string `json:"id"`
	IsPermissioned bool   `json:"is_permissioned"`
}
