package git

import (
	"os"
	"os/exec"
	"strconv"
)

// repoUserEmail sets the global git author email.
func ConfigAutocorrect(repo Repository) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"config",
		"--local",
		"help.autocorrect",
		repo.Autocorrect,
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	return cmd
}

// repoUserEmail sets the global git author email.
func ConfigUserEmail(repo Repository) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"config",
		"--local",
		"user.email",
		repo.Author.Email,
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	return cmd
}

// repoUserName sets the global git author name.
func ConfigUserName(repo Repository) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"config",
		"--local",
		"user.name",
		repo.Author.Name,
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	return cmd
}

// repoSSLVerify disables globally the git ssl verification.
func ConfigSSLVerify(repo Repository) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"config",
		"--local",
		"http.sslVerify",
		strconv.FormatBool(repo.SSLVerify),
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	return cmd
}
