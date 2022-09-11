package op

type Vault struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ListVaults returns all vaults the account has access to.
//
// Supported filters:
//
//   - WithUser()    List vaults a group has access to.
//   - WithGroup()   List vaults that a given user has access to.
func (c *CLI) ListVaults(filters ...Filter) ([]Vault, error) {
	var val []Vault
	err := c.execJSON(applyFilters([]string{"vault", "list"}, filters), nil, &val)
	return val, err
}

// GetVault returns the details of a vault specified by its name or ID.
func (c *CLI) GetVault(name string) (*Vault, error) {
	var val *Vault
	err := c.execJSON([]string{"vault", "get", sanitize(name)}, nil, &val)
	return val, err
}
