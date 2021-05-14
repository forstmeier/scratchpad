package acct

import (
	"github.com/permissionguru/api/pkg/config"
	"github.com/permissionguru/api/pkg/subscr"
)

// Account holds account-level values.
type Account struct {
	ID           string
	Subscription subscr.Subscription
	Config       *config.Config
}
