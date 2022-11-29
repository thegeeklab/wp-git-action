package git

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

const netrcFile = `
machine %s
login %s
password %s
`

const configFile = `
Host *
StrictHostKeyChecking no
UserKnownHostsFile=/dev/null
`

// WriteKey writes the SSH private key.
func WriteSSHKey(privateKey string) error {
	home := "/root"

	if currentUser, err := user.Current(); err == nil {
		home = currentUser.HomeDir
	}

	sshpath := filepath.Join(
		home,
		".ssh")

	if err := os.MkdirAll(sshpath, 0o700); err != nil {
		return err
	}

	confpath := filepath.Join(
		sshpath,
		"config")

	if err := os.WriteFile(
		confpath,
		[]byte(configFile),
		0o700,
	); err != nil {
		return err
	}

	privpath := filepath.Join(
		sshpath,
		"id_rsa",
	)

	if err := os.WriteFile(
		privpath,
		[]byte(privateKey),
		0o600,
	); err != nil {
		return err
	}

	return nil
}

// WriteNetrc writes the netrc file.
func WriteNetrc(machine, login, password string) error {
	netrcContent := fmt.Sprintf(
		netrcFile,
		machine,
		login,
		password,
	)

	home := "/root"

	if currentUser, err := user.Current(); err == nil {
		home = currentUser.HomeDir
	}

	netpath := filepath.Join(
		home,
		".netrc",
	)

	return os.WriteFile(
		netpath,
		[]byte(netrcContent),
		0o600,
	)
}
