package db

import (
	"context"

	"github.com/permissionguru/api/pkg/acct"
	"github.com/permissionguru/api/pkg/config"
	"github.com/permissionguru/api/pkg/subscr"
)

// Databaser defines the methods for interacting with
// values in the underlying data persistence system.
type Databaser interface {
	AddAccount(ctx context.Context, id string) error
	AddSubscription(ctx context.Context, id string, sub subscr.Subscription) error
	AddConfig(ctx context.Context, id string, cfg config.Config) error
	GetAccount(ctx context.Context, id string) (*acct.Account, error)
	RemoveAccount(ctx context.Context, id string) error
}
