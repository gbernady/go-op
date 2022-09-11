package op

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

// CLI represents the 1Password CLI.
type CLI struct {
	// Account specifies the account to execute the commands by account shorthand, sign-in address, account ID, or user ID.
	Account string

	// Config specifies the 1Password CLI configuration directory to use.
	Config string

	// Path specifies an absolute path to the 1Password CLI executable.
	// If not set, exec.LookPath() will be utilized to attempt to find the `op` executable on $PATH.
	Path string
}

// Version returns 1Password CLI version.
func (c CLI) Version() (string, error) {
	b, err := c.execRaw([]string{"--version"}, nil)
	return strings.TrimSpace(string(b)), err
}

func (c CLI) execRaw(cmd []string, args []string) ([]byte, error) {
	if c.Account != "" {
		cmd = append(cmd, "--account", c.Account)
	}
	if c.Config != "" {
		cmd = append(cmd, "--config", c.Config)
	}
	cmd = append(cmd, args...)

	path := c.Path
	if path == "" {
		p, err := exec.LookPath("op")
		if err != nil && !errors.Is(err, exec.ErrDot) {
			return nil, err
		}
		path = p
	}

	op := &exec.Cmd{
		Path: path,
		Args: append([]string{path}, cmd...),
	}
	b, err := op.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("%s: %s", ee, ee.Stderr)
		}
		return nil, err
	}
	return b, err
}

func (c CLI) execJSON(cmd []string, args []string, v any) error {
	cmd = append(cmd, "--format", "json", "--iso-timestamps")
	b, err := c.execRaw(cmd, args)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &v)
	return err
}

func sanitize(s string) string {
	return fmt.Sprintf("%q", s)
}
