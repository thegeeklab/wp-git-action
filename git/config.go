package git

import (
	"os/exec"
	"strconv"
)

// SetUserEmail sets the global git author email.
func SetUserEmail(email string) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"config",
		"--local",
		"user.email",
		email)

	return cmd
}

// SetUserName sets the global git author name.
func SetUserName(author string) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"config",
		"--local",
		"user.name",
		author)

	return cmd
}

// SetSSLSkipVerify disables globally the git ssl verification.
func SetSSLVerify(sslVerify bool) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"config",
		"--local",
		"http.sslVerify",
		strconv.FormatBool(sslVerify))

	return cmd
}
