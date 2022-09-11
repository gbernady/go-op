package op

import (
	"strings"
)

// Filter represents a filter flag passed to 1Password CLI some commands.
type Filter func() []string

// WithIncludeArchive expands list to include items in the Archive.
//
// Supported by the `ListItems` and `GetItem` APIs.
func WithIncludeArchive() Filter {
	return func() []string {
		return []string{"--include-archive"}
	}
}

// WithCategories limits results to only include items matching given categories.
//
// Supported by the `ListItems` API.
func WithCategories(categories ...Category) Filter {
	return func() []string {
		if len(categories) == 0 {
			return nil
		}
		var s []string
		for _, c := range categories {
			s = append(s, string(c))
		}
		return []string{"--categories", sanitize(strings.Join(s, ","))}
	}
}

// WithGroup limits results to only include resources matching given group.
//
// Supported by the `ListUsers` and `ListVaults` APIs.
// --group group   List users who belong to a group.
func WithGroup(name string) Filter {
	return func() []string {
		if name == "" {
			return nil
		}
		return []string{"--group", sanitize(name)}
	}
}

// WithFavorite limits results to only include favorite items.
//
// Supported by the `ListItems` API.
func WithFavorite() Filter {
	return func() []string {
		return []string{"--favorite"}
	}
}

// WithTags limits results to only include items matching given tags.
//
// Supported by the `ListItems` API.
func WithTags(tags ...string) Filter {
	return func() []string {
		if len(tags) == 0 {
			return nil
		}
		return []string{"--tags", sanitize(strings.Join(tags, ","))}
	}
}

// WithUser limits results to only include resources matching given user.
//
// Supported by the `ListGroups` and `ListVaults` APIs.
func WithUser(name string) Filter {
	return func() []string {
		if name == "" {
			return nil
		}
		return []string{"--user", sanitize(name)}
	}
}

// WithVault limits results to only include resources matching given vault.
//
// Supported by the `ListGroups`, `ListItems`, `GetItem` and `ListUsers` APIs.
func WithVault(name string) Filter {
	return func() []string {
		if name == "" {
			return nil
		}
		return []string{"--vault", sanitize(name)}
	}
}

func applyFilters(cmd []string, filters []Filter) []string {
	for _, f := range filters {
		cmd = append(cmd, f()...)
	}
	return cmd
}
