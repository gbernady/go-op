package op

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListVaults(t *testing.T) {
	tests := []struct {
		name string
		call func(cli *CLI) (any, error)
		resp []Vault
		err  string
	}{
		{
			name: "All",
			call: func(cli *CLI) (any, error) {
				return cli.ListVaults()
			},
			resp: []Vault{
				{
					ID:   "ynghx4vwntpezvhqyeglcp7v7f",
					Name: "Personal",
				},
				{
					ID:   "ynghx4vwntpezvhqyeglcp7v7g",
					Name: "Bar",
				},
				{
					ID:   "ynghx4vwntpezvhqyeglcp7v7h",
					Name: "Qux",
				},
			},
		},
		{
			name: "Group",
			call: func(cli *CLI) (any, error) {
				return cli.ListVaults(WithGroup("Bar"))
			},
			resp: []Vault{
				{
					ID:   "ynghx4vwntpezvhqyeglcp7v7g",
					Name: "Bar",
				},
			},
		},
		{
			name: "User",
			call: func(cli *CLI) (any, error) {
				return cli.ListVaults(WithUser("foo@example.com"))
			},
			resp: []Vault{
				{
					ID:   "ynghx4vwntpezvhqyeglcp7v7f",
					Name: "Personal",
				},
				{
					ID:   "ynghx4vwntpezvhqyeglcp7v7g",
					Name: "Bar",
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

func TestGetVault(t *testing.T) {
	tests := []struct {
		name string
		call func(cli *CLI) (any, error)
		resp *Vault
		err  string
	}{
		{
			name: "Personal",
			call: func(cli *CLI) (any, error) { return cli.GetVault("Personal") },
			resp: &Vault{
				ID:   "ynghx4vwntpezvhqyeglcp7v7f",
				Name: "Personal",
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
