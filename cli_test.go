package op

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/fatih/camelcase"
)

func TestVersion(t *testing.T) {
	// FIXME: implement
}

func mockOp(t *testing.T) string {
	tmp := t.TempDir()

	// copy `op` mock
	b, err := os.ReadFile(filepath.Join("testdata", "op"))
	if err != nil {
		t.Error(err)
	}
	if err := os.WriteFile(filepath.Join(tmp, "op"), b, 0744); err != nil {
		t.Error(err)
	}

	// copy test response
	tname := strings.ToLower(strings.Join(camelcase.Split(strings.ReplaceAll(t.Name(), "/", ""))[1:], "_"))
	b, err = os.ReadFile(filepath.Join("testdata", tname))
	if err != nil {
		t.Error(err)
	}
	if err := os.WriteFile(filepath.Join(tmp, "op_response"), b, 0644); err != nil {
		t.Error(err)
	}

	return filepath.Join(tmp, "op")
}
