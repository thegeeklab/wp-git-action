package git

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/thegeeklab/wp-plugin-go/v2/trace"
	"golang.org/x/sys/execabs"
)

const (
	netrcFile = `
machine %s
login %s
password %s
`
	configFile = `
Host *
StrictHostKeyChecking no
UserKnownHostsFile=/dev/null
`
)

const (
	strictFilePerm = 0o600
	strictDirPerm  = 0o600
)

// WriteKey writes the SSH private key.
func WriteSSHKey(privateKey string) error {
	home := "/root"

	if currentUser, err := user.Current(); err == nil {
		home = currentUser.HomeDir
	}

	sshpath := filepath.Join(home, ".ssh")

	if err := os.MkdirAll(sshpath, strictDirPerm); err != nil {
		return err
	}

	confpath := filepath.Join(sshpath, "config")

	if err := os.WriteFile(
		confpath,
		[]byte(configFile),
		strictFilePerm,
	); err != nil {
		return err
	}

	privpath := filepath.Join(sshpath, "id_rsa")

	return os.WriteFile(
		privpath,
		[]byte(privateKey),
		strictFilePerm,
	)
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
		strictFilePerm,
	)
}

func runCommand(cmd *execabs.Cmd) error {
	if cmd.Stdout == nil {
		cmd.Stdout = os.Stdout
	}

	if cmd.Stderr == nil {
		cmd.Stderr = os.Stderr
	}

	trace.Cmd(cmd)

	return cmd.Run()
}
