package op

type User struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Type       UserType  `json:"type"`
	State      UserState `json:"state"`
	CreatedAt  string    `json:"created_at"`
	UpdatedAt  string    `json:"updated_at"`
	LastAuthAt string    `json:"last_auth_at"`
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
//	--group group   List users who belong to a group.
//	--vault vault   List users who have direct access to vault.
func (c *CLI) ListUsers(filters ...Filter) ([]User, error) {
	var u []User
	err := c.execJSON(applyFilters([]string{"user", "list"}, filters), nil, u)
	return u, err
}

// GetCurrentUser returns the details of the currently authenticated user.
func (c *CLI) GetCurrentUser() (*User, error) {
	var u *User
	err := c.execJSON([]string{"user", "get", "--me"}, nil, u)
	return u, err
}

// GetUser returns the details of a user specified by their e-mail address, name, or ID.
func (c *CLI) GetUser(name string) (*User, error) {
	var u *User
	err := c.execJSON([]string{"user", "get", name}, nil, u)
	return u, err
}

// GetUserFingerprint returns the user's public key fingerprint.
func (c *CLI) GetUserFingerprint(name string) (string, error) {
	b, err := c.execRaw([]string{"user", "get", name, "--fingerprint"}, nil)
	return string(b), err
}

// GetUserPublicKey the user's public key.
func (c *CLI) GetUserPublicKey(name string) (string, error) {
	b, err := c.execRaw([]string{"user", "get", name, "--public-key"}, nil)
	return string(b), err
}
