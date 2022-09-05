package op

import (
	"time"
)

type Account struct {
	URL         string `json:"url"`
	Email       string `json:"email"`
	UserUUID    string `json:"user_uuid"`
	AccountUUID string `json:"account_uuid"`
	Shorthand   string `json:"shorthand,omitempty"`
}

type AccountDetails struct {
	ID        string       `json:"id"`
	Name      string       `json:"name"`
	Domain    string       `json:"domain"`
	Type      AccountType  `json:"type"`
	State     AccountState `json:"state"`
	CreatedAt time.Time    `json:"created_at"`
}

type AccountType string

const (
	AccountTypeBusiness   = "BUSINESS"
	AccountTypeTeam       = "TEAM"
	AccountTypeFamily     = "FAMILY"
	AccountTypeIndividual = "INDIVIDUAL"
	AccountTypeUnknown    = "UNKNOWN"
)

type AccountState string

const (
	AccountStateRegistered = "REGISTERED"
	AccountStateActive     = "ACTIVE"
	AccountStateSuspended  = "SUSPENDED"
	AccountStateDeleted    = "DELETED"
	AccountStatePurging    = "PURGING"
	AccountStatePurged     = "PURGED"
	AccountStateUnknown    = "UNKNOWN"
)

// ListAccount returns the accounts set up on this device.
func (c *CLI) ListAccounts() ([]Account, error) {
	var val []Account
	err := c.execJSON([]string{"account", "list"}, nil, &val)
	return val, err
}

// GetAccount returns details of your account.
func (c *CLI) GetAccount() (*AccountDetails, error) {
	var val *AccountDetails
	err := c.execJSON([]string{"account", "get"}, nil, &val)
	return val, err
}
