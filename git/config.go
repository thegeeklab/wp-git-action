package git

import (
	"strconv"

	"github.com/thegeeklab/wp-plugin-go/v2/types"
	"golang.org/x/sys/execabs"
)

// repoUserEmail sets the global git author email.
func ConfigAutocorrect(repo Repository) *types.Cmd {
	args := []string{
		"config",
		"--local",
		"help.autocorrect",
		repo.Autocorrect,
	}

	cmd := execabs.Command(gitBin, args...)
	cmd.Dir = repo.WorkDir

	return &types.Cmd{
		Cmd: cmd,
	}
}

// repoUserEmail sets the global git author email.
func ConfigUserEmail(repo Repository) *types.Cmd {
	args := []string{
		"config",
		"--local",
		"user.email",
		repo.Author.Email,
	}

	cmd := execabs.Command(gitBin, args...)
	cmd.Dir = repo.WorkDir

	return &types.Cmd{
		Cmd: cmd,
	}
}

// repoUserName sets the global git author name.
func ConfigUserName(repo Repository) *types.Cmd {
	args := []string{
		"config",
		"--local",
		"user.name",
		repo.Author.Name,
	}

	cmd := execabs.Command(gitBin, args...)
	cmd.Dir = repo.WorkDir

	return &types.Cmd{
		Cmd: cmd,
	}
}

// ConfigSSLVerify disables globally the git ssl verification.
func ConfigSSLVerify(repo Repository, skipVerify bool) *types.Cmd {
	args := []string{
		"config",
		"--local",
		"http.sslVerify",
		strconv.FormatBool(!skipVerify),
	}

	cmd := execabs.Command(gitBin, args...)
	cmd.Dir = repo.WorkDir

	return &types.Cmd{
		Cmd: cmd,
	}
}
