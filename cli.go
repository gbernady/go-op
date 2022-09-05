package op

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
)

// CLI represents the 1Password CLI.
type CLI struct {
	// Account specifies the account to execute the commands by account shorthand, sign-in address, account ID, or user ID.
	Account string

	// Config specifies the 1Password CLI configuration directory to use.
	Config string

	// Env specifies the environment in which 1Password CLI commands will be run.
	// Each entry is of the form "key=value".
	// If Env is nil (default), commands will use the current process's environment.
	Env []string
}

// Version returns 1Password CLI version.
func (c CLI) Version() (string, error) {
	b, err := c.execRaw([]string{"--version"}, nil)
	return string(b), err
}

func (c CLI) execRaw(cmd []string, args []string) ([]byte, error) {
	if c.Account != "" {
		cmd = append(cmd, "--account", c.Account)
	}
	cmd = append(cmd, args...)
	cc := exec.Command("op", cmd...)
	cc.Env = append(cc.Environ(), c.Env...)
	if errors.Is(cc.Err, exec.ErrDot) {
		cc.Err = nil
	}
	b, err := cc.Output()
	if err != nil {
		if ee := err.(*exec.ExitError); ee != nil {
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
