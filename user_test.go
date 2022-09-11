package op

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListUsers(t *testing.T) {
	tests := []struct {
		name string
		call func(cli *CLI) (any, error)
		resp []User
		err  string
	}{
		{
			name: "All",
			call: func(cli *CLI) (any, error) {
				return cli.ListUsers()
			},
			resp: []User{
				{
					ID:    "F7GSLUVENFGZVF2HVACL3IAS7F",
					Name:  "Foo",
					Email: "foo@example.com",
					Type:  UserTypeMember,
					State: UserStateActive,
				},
				{
					ID:    "F7GSLUVENFGZVF2HVACL3IAS7G",
					Name:  "Qux",
					Email: "qux@example.com",
					Type:  UserTypeGuest,
					State: UserStateSuspended,
				},
			},
		},
		{
			name: "Group",
			call: func(cli *CLI) (any, error) {
				return cli.ListUsers(WithGroup("Bar"))
			},
			resp: []User{
				{
					ID:    "F7GSLUVENFGZVF2HVACL3IAS7F",
					Name:  "Foo",
					Email: "foo@example.com",
					Type:  UserTypeMember,
					State: UserStateActive,
				},
			},
		},
		{
			name: "Vault",
			call: func(cli *CLI) (any, error) {
				return cli.ListUsers(WithVault("Qux"))
			},
			resp: []User{
				{
					ID:    "F7GSLUVENFGZVF2HVACL3IAS7G",
					Name:  "Qux",
					Email: "qux@example.com",
					Type:  UserTypeGuest,
					State: UserStateSuspended,
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cli := &CLI{Path: mockOp(t)}
			resp, err := test.call(cli)
			if test.err == "" {
				assert.Equal(t, test.resp, resp)
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.err)
			}
		})
	}
}

func TestGetCurrentUser(t *testing.T) {
	tests := []struct {
		name string
		call func(cli *CLI) (any, error)
		resp *User
		err  string
	}{
		{
			name: "Success",
			call: func(cli *CLI) (any, error) {
				return cli.GetCurrentUser()
			},
			resp: &User{
				ID:         "F7GSLUVENFGZVF2HVACL3IAS7F",
				Name:       "Foo",
				Email:      "foo@example.com",
				Type:       UserTypeMember,
				State:      UserStateActive,
				CreatedAt:  time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
				UpdatedAt:  time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
				LastAuthAt: time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cli := &CLI{Path: mockOp(t)}
			resp, err := test.call(cli)
			if test.err == "" {
				assert.Equal(t, test.resp, resp)
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.err)
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	tests := []struct {
		name string
		call func(cli *CLI) (any, error)
		resp *User
		err  string
	}{
		{
			name: "Success",
			call: func(cli *CLI) (any, error) {
				return cli.GetUser("foo@example.com")
			},
			resp: &User{
				ID:         "F7GSLUVENFGZVF2HVACL3IAS7F",
				Name:       "Foo",
				Email:      "foo@example.com",
				Type:       UserTypeMember,
				State:      UserStateActive,
				CreatedAt:  time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
				UpdatedAt:  time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
				LastAuthAt: time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cli := &CLI{Path: mockOp(t)}
			resp, err := test.call(cli)
			if test.err == "" {
				assert.Equal(t, test.resp, resp)
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.err)
			}
		})
	}
}

func TestGetUserFingerprint(t *testing.T) {
	tests := []struct {
		name string
		call func(cli *CLI) (any, error)
		resp string
		err  string
	}{
		{
			name: "Success",
			call: func(cli *CLI) (any, error) {
				return cli.GetUserFingerprint("foo@example.com")
			},
			resp: "ad34e5 384737 4c77fa 234ac2 193438 341431 ff782f cf232a",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cli := &CLI{Path: mockOp(t)}
			resp, err := test.call(cli)
			if test.err == "" {
				assert.Equal(t, test.resp, resp)
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.err)
			}
		})
	}
}

func TestGetUserPublicKey(t *testing.T) {
	tests := []struct {
		name string
		call func(cli *CLI) (any, error)
		resp string
		err  string
	}{
		{
			name: "Success",
			call: func(cli *CLI) (any, error) {
				return cli.GetUserPublicKey("foo@example.com")
			},
			resp: `{"alg":"RSA-OAEP","kid":"rh2vnbjhdnekdpvgoqivu57ahm","ext":true,"e":"AVAD","n":"wrpn5bDMotUqMv6H57kGMVwkfdDLqlZJ83746asOUhnrbX5XF7SIKJ747d33ShApCl-SW6_a4MW6gjtc1fiGLpWsZ8pPEUJi4NBoIbJ6sDMWqX5_v7YY1PjmSXaMq-qg5_vkcPkqdPziSvWuOXtgSxmuseAA0WZiKI_Wyh5MAso4-DDTAJEyeSiVPpjAhgQY-wJJ-HlYoJ6DrypGH26SMjZyoObpT3-hKi2fQzCy2x0KLU2NqW1bsBNzSW-ejuga9feg1PtepmBxTGKua3Z1gFZsJ_uDKjJ-x7DKDA8jkU2XTXn3MnqL031Hfs8-d8GZVOsRs-rF453RpVO-6MXeQS","key_ops":["encrypt"],"kty":"RSA"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cli := &CLI{Path: mockOp(t)}
			resp, err := test.call(cli)
			if test.err == "" {
				assert.Equal(t, test.resp, resp)
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.err)
			}
		})
	}
}
