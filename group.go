package op

type Group struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	State       GroupState `json:"state"`
	CreatedAt   string     `json:"created_at"`
	UpdatedAt   string     `json:"updated_at"`
	Type        GroupType  `json:"type"`
}

type GroupState string

const (
	GroupStateActive   = "ACTIVE"
	GroupStateDeleted  = "DELETED"
	GroupStateInactive = "INACTIVE"
)

type GroupType string

const (
	GroupTypeAdministrators = "ADMINISTRATORS"
	GroupTypeOwners         = "OWNERS"
	GroupTypeRecovery       = "RECOVERY"
	GroupTypeSecurity       = "SECURITY"
	GroupTypeTeamMembers    = "TEAM_MEMBERS"
	GroupTypeUnknown        = "UNKNOWN_TYPE"
	GroupTypeUserDefined    = "USER_DEFINED"
)

// ListGroups returns the list of groups in the current account.
//
// Supported filters:
// --user user     List groups that a user belongs to.
// --vault vault   List groups that have direct access to a vault.
func (c *CLI) ListGroups(filters ...Filter) ([]Group, error) {
	var val []Group
	err := c.execJSON(applyFilters([]string{"group", "list"}, filters), nil, &val)
	return val, err
}

// GetGroup returns the details of a group specified by its name or ID.
func (c *CLI) GetGroup(name string) (*Group, error) {
	var val *Group
	err := c.execJSON([]string{"group", "get"}, nil, &val)
	return val, err
}
