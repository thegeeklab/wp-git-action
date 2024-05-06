package plugin

import (
	"fmt"
	"os"
	"path/filepath"
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
