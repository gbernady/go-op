package op

import "time"

type Group struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	State       GroupState        `json:"state"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	Permissions []GroupPermission `json:"permissions"`
	Type        GroupType         `json:"type"`
}

type GroupState string

const (
	GroupStateActive   = "ACTIVE"
	GroupStateDeleted  = "DELETED"
	GroupStateInactive = "INACTIVE"
)

type GroupPermission string

const (
	GroupPermissionAddPerson                 = "ADD_PERSON"
	GroupPermissionChangePersonName          = "CHANGE_PERSON_NAME"
	GroupPermissionChangeTeamAttributes      = "CHANGE_TEAM_ATTRIBUTES"
	GroupPermissionChangeTeamDomain          = "CHANGE_TEAM_DOMAIN"
	GroupPermissionChangeTeamSettings        = "CHANGE_TEAM_SETTINGS"
	GroupPermissionCreateVaults              = "CREATE_VAULTS"
	GroupPermissionDeletePerson              = "DELETE_PERSON"
	GroupPermissionDeleteTeam                = "DELETE_TEAM"
	GroupPermissionManageBilling             = "MANAGE_BILLING"
	GroupPermissionManageGroups              = "MANAGE_GROUPS"
	GroupPermissionManageTemplates           = "MANAGE_TEMPLATES"
	GroupPermissionManageVaults              = "MANAGE_VAULTS"
	GroupPermissionRecoverAccounts           = "RECOVER_ACCOUNTS"
	GroupPermissionSuspendPerson             = "SUSPEND_PERSON"
	GroupPermissionSuspendTeam               = "SUSPEND_TEAM"
	GroupPermissionViewActivitiesLog         = "VIEW_ACTIVITIES_LOG"
	GroupPermissionViewAdministrativeSidebar = "VIEW_ADMINISTRATIVE_SIDEBAR"
	GroupPermissionViewBilling               = "VIEW_BILLING"
	GroupPermissionViewPeople                = "VIEW_PEOPLE"
	GroupPermissionViewSecurityReports       = "VIEW_SECURITY_REPORTS"
	GroupPermissionViewTeamSettings          = "VIEW_TEAM_SETTINGS"
	GroupPermissionViewTemplates             = "VIEW_TEMPLATES"
	GroupPermissionViewVaults                = "VIEW_VAULTS"
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
//
//   - WithUser()    List groups that a user belongs to.
//   - WithVault()   List groups that have direct access to a vault.
func (c *CLI) ListGroups(filters ...Filter) ([]Group, error) {
	var val []Group
	err := c.execJSON(applyFilters([]string{"group", "list"}, filters), nil, &val)
	return val, err
}

// GetGroup returns the details of a group specified by its name or ID.
func (c *CLI) GetGroup(name string) (*Group, error) {
	var val *Group
	err := c.execJSON([]string{"group", "get", sanitize(name)}, nil, &val)
	return val, err
}
