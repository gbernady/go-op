package op

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListAccounts(t *testing.T) {
	tests := []struct {
		name string
		call func(cli *CLI) (any, error)
		resp []Account
		err  string
	}{
		{
			name: "All",
			call: func(cli *CLI) (any, error) { return cli.ListAccounts() },
			resp: []Account{
				{
					URL:         "my.1password.eu",
					Email:       "foo@example.com",
					UserUUID:    "F7GSLUVENFGZVF2HVACL3IAS7F",
					AccountUUID: "3TBQIJZ6TVAZ5DM56FZCF653BW",
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

func TestGetAccount(t *testing.T) {
	tests := []struct {
		name string
		call func(cli *CLI) (any, error)
		resp *AccountDetails
		err  string
	}{
		{
			name: "Current",
			call: func(cli *CLI) (any, error) { return cli.GetAccount() },
			resp: &AccountDetails{
				ID:        "3TBQIJZ6TVAZ5DM56FZCF653BW",
				Name:      "Foo",
				Domain:    "my",
				Type:      AccountTypeIndividual,
				State:     AccountStateActive,
				CreatedAt: time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
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
