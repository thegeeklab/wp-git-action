package plugin

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

const (
	netrcFile = `machine %s
login %s
password %s
`
	configFile = `Host *
StrictHostKeyChecking no
UserKnownHostsFile=/dev/null
`
)

const (
	strictFilePerm = 0o600
	strictDirPerm  = 0o700
)

func GetUserHomeDir() string {
	home := "/root"

	if currentUser, err := user.Current(); err == nil {
		home = currentUser.HomeDir
	}

	return home
}

// WriteKey writes the SSH private key.
func WriteSSHKey(path, key string) error {
	sshPath := filepath.Join(path, ".ssh")
	confPath := filepath.Join(sshPath, "config")
	keyPath := filepath.Join(sshPath, "id_rsa")

	if err := os.MkdirAll(sshPath, strictDirPerm); err != nil {
		return fmt.Errorf("failed to create .ssh directory: %w", err)
	}

	if err := os.WriteFile(confPath, []byte(configFile), strictFilePerm); err != nil {
		return fmt.Errorf("failed to create .ssh/config file: %w", err)
	}

	if err := os.WriteFile(keyPath, []byte(key), strictFilePerm); err != nil {
		return fmt.Errorf("failed to create .ssh/id_rsa file: %w", err)
	}

	return nil
}

// WriteNetrc writes the netrc file.
func WriteNetrc(path, machine, login, password string) error {
	netrcPath := filepath.Join(path, ".netrc")
	netrcContent := fmt.Sprintf(netrcFile, machine, login, password)

	if err := os.WriteFile(netrcPath, []byte(netrcContent), strictFilePerm); err != nil {
		return fmt.Errorf("failed to create .netrc file: %w", err)
	}

	return nil
}
