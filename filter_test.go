package op

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithIncludeArchive(t *testing.T) {
	assert.Equal(t, []string{"--include-archive"}, WithIncludeArchive()())
}

func TestWithCategories(t *testing.T) {
	assert.Nil(t, WithCategories()())
	assert.Equal(t, []string{"--categories", `"API Credential"`}, WithCategories(CategoryAPICredential)())
	assert.Equal(t, []string{"--categories", `"API Credential,Login"`}, WithCategories(CategoryAPICredential, CategoryLogin)())
}

func TestWithGroup(t *testing.T) {
	assert.Nil(t, WithGroup("")())
	assert.Equal(t, []string{"--group", "foo"}, WithGroup("foo")())
}

func TestWithFavorite(t *testing.T) {
	assert.Equal(t, []string{"--favorite"}, WithFavorite()())
}

func TestWithTags(t *testing.T) {
	assert.Nil(t, WithTags()())
	assert.Equal(t, []string{"--tags", `"foo"`}, WithTags("foo")())
	assert.Equal(t, []string{"--tags", `"foo,bar baz"`}, WithTags("foo", "bar baz")())
}

func TestWithUser(t *testing.T) {
	assert.Nil(t, WithUser("")())
	assert.Equal(t, []string{"--user", "foo"}, WithUser("foo")())
}

func TestWithVault(t *testing.T) {
	assert.Nil(t, WithVault("")())
	assert.Equal(t, []string{"--vault", "foo"}, WithVault("foo")())
}
