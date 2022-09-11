package op

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListGroups(t *testing.T) {
	tests := []struct {
		name string
		call func(cli *CLI) (any, error)
		resp []Group
		err  string
	}{
		{
			name: "All",
			call: func(cli *CLI) (any, error) {
				return cli.ListGroups()
			},
			resp: []Group{
				{
					ID:          "syn5ckyj3t5nff6wng2n2waeen",
					Name:        "Recovery",
					Description: "Can reset user passwords if account recovery is enabled.",
					State:       GroupStateActive,
					CreatedAt:   time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
				},
				{
					ID:          "syn5ckyj3t5nff6wng2n2waeeo",
					Name:        "Owners",
					Description: "Access to billing and account administration.",
					State:       GroupStateActive,
					CreatedAt:   time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
				},
				{
					ID:          "syn5ckyj3t5nff6wng2n2waeep",
					Name:        "Administrators",
					Description: "Administration of users, groups, and vaults.",
					State:       GroupStateActive,
					CreatedAt:   time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
				},
				{
					ID:          "syn5ckyj3t5nff6wng2n2waeeq",
					Name:        "Team Members",
					Description: "All team members.",
					State:       GroupStateActive,
					CreatedAt:   time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "User",
			call: func(cli *CLI) (any, error) {
				return cli.ListGroups(WithUser("foo@example.com"))
			},
			resp: []Group{
				{
					ID:          "syn5ckyj3t5nff6wng2n2waeeq",
					Name:        "Team Members",
					Description: "All team members.",
					State:       GroupStateActive,
					CreatedAt:   time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "Vault",
			call: func(cli *CLI) (any, error) {
				return cli.ListGroups(WithVault("Bar"))
			},
			resp: []Group{
				{
					ID:          "syn5ckyj3t5nff6wng2n2waeep",
					Name:        "Administrators",
					Description: "Administration of users, groups, and vaults.",
					State:       GroupStateActive,
					CreatedAt:   time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
				},
				{
					ID:          "syn5ckyj3t5nff6wng2n2waeeq",
					Name:        "Team Members",
					Description: "All team members.",
					State:       GroupStateActive,
					CreatedAt:   time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
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

func TestGetGroup(t *testing.T) {
	tests := []struct {
		name string
		call func(cli *CLI) (any, error)
		resp *Group
		err  string
	}{
		{
			name: "TeamMembers",
			call: func(cli *CLI) (any, error) { return cli.GetGroup("Team Members") },
			resp: &Group{
				ID:          "syn5ckyj3t5nff6wng2n2waeeq",
				Name:        "Team Members",
				Description: "All team members.",
				State:       GroupStateActive,
				CreatedAt:   time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, time.April, 20, 9, 41, 0, 0, time.UTC),
				Permissions: []GroupPermission{GroupPermissionCreateVaults},
				Type:        GroupTypeTeamMembers,
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
