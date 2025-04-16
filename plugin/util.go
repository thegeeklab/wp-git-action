package plugin

import (
	"fmt"
	"os"
	"path/filepath"

	plugin_exec "github.com/thegeeklab/wp-plugin-go/v5/exec"
)

const (
	netrcFile = `machine %s
login %s
password %s
`
)

const strictFilePerm = 0o600

// WriteNetrc writes the netrc file.
func WriteNetrc(path, machine, login, password string) error {
	netrcPath := filepath.Join(path, ".netrc")
	netrcContent := fmt.Sprintf(netrcFile, machine, login, password)

	if err := os.WriteFile(netrcPath, []byte(netrcContent), strictFilePerm); err != nil {
		return fmt.Errorf("failed to create .netrc file: %w", err)
	}

	return nil
}

// ExecBatch executes a batch of commands. If any command in the batch fails, the function will return the error.
func ExecBatch(batchCmd []*plugin_exec.Cmd) error {
	for _, cmd := range batchCmd {
		if cmd == nil {
			continue
		}

		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}
