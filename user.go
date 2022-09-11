package op

import (
	"strings"
	"time"
)

type User struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Type       UserType  `json:"type"`
	State      UserState `json:"state"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	LastAuthAt time.Time `json:"last_auth_at"`
}

type UserType string

const (
	UserTypeMember         = "MEMBER"
	UserTypeGuest          = "GUEST"
	UserTypeServiceAccount = "SERVICE_ACCOUNT"
	UserTypeUnknown        = "UNKNOWN"
)

type UserState string

const (
	UserStateActive                                 = "ACTIVE"
	UserStatePending                                = "PENDING"
	UserStateDeleted                                = "DELETED"
	UserStateSuspended                              = "SUSPENDED"
	UserStateRecoveryStarted                        = "RECOVERY_STARTED"
	UserStateRecoveryAccepted                       = "RECOVERY_ACCEPTED"
	UserStateTransferPending                        = "TRANSFER_PENDING"
	UserStateTransferStarted                        = "TRANSFER_STARTED"
	UserStateTransferAccepted                       = "TRANSFER_ACCEPTED"
	UserStateEmailVerifiedButRegistrationIncomplete = "EMAIL_VERIFIED_BUT_REGISTRATION_INCOMPLETE"
	UserStateTeamRegistrationInitiated              = "TEAM_REGISTRATION_INITIATED"
	UserStateUnknown                                = "UNKNOWN"
)

// ListUsers returns the list of users in the current account.
//
// Supported filters:
//
//   - WithGroup()   List users who belong to a group.
//   - WithVault()   List users who have direct access to vault.
func (c *CLI) ListUsers(filters ...Filter) ([]User, error) {
	var val []User
	err := c.execJSON(applyFilters([]string{"user", "list"}, filters), nil, &val)
	return val, err
}

// GetCurrentUser returns the details of the currently authenticated user.
func (c *CLI) GetCurrentUser() (*User, error) {
	var val *User
	err := c.execJSON([]string{"user", "get", "--me"}, nil, &val)
	return val, err
}

// GetUser returns the details of a user specified by their e-mail address, name, or ID.
func (c *CLI) GetUser(name string) (*User, error) {
	var val *User
	err := c.execJSON([]string{"user", "get", sanitize(name)}, nil, &val)
	return val, err
}

// GetUserFingerprint returns the user's public key fingerprint.
func (c *CLI) GetUserFingerprint(name string) (string, error) {
	b, err := c.execRaw([]string{"user", "get", sanitize(name), "--fingerprint"}, nil)
	return strings.TrimSpace(string(b)), err
}

// GetUserPublicKey the user's public key.
func (c *CLI) GetUserPublicKey(name string) (string, error) {
	b, err := c.execRaw([]string{"user", "get", sanitize(name), "--public-key"}, nil)
	return strings.TrimSpace(string(b)), err
}
