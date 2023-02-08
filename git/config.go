package git

import (
	"os"
	"strconv"

	"golang.org/x/sys/execabs"
)

// repoUserEmail sets the global git author email.
func ConfigAutocorrect(repo Repository) *execabs.Cmd {
	args := []string{
		"config",
		"--local",
		"help.autocorrect",
		repo.Autocorrect,
	}

	cmd := execabs.Command(
		gitBin,
		args...,
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	return cmd
}

// repoUserEmail sets the global git author email.
func ConfigUserEmail(repo Repository) *execabs.Cmd {
	args := []string{
		"config",
		"--local",
		"user.email",
		repo.Author.Email,
	}

	cmd := execabs.Command(
		gitBin,
		args...,
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	return cmd
}

// repoUserName sets the global git author name.
func ConfigUserName(repo Repository) *execabs.Cmd {
	args := []string{
		"config",
		"--local",
		"user.name",
		repo.Author.Name,
	}

	cmd := execabs.Command(
		gitBin,
		args...,
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	return cmd
}

// repoSSLVerify disables globally the git ssl verification.
func ConfigSSLVerify(repo Repository) *execabs.Cmd {
	args := []string{
		"config",
		"--local",
		"http.sslVerify",
		strconv.FormatBool(repo.SSLVerify),
	}

	cmd := execabs.Command(
		gitBin,
		args...,
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	return cmd
}
