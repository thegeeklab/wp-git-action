package git

import (
	"strconv"

	"github.com/thegeeklab/wp-plugin-go/v2/types"
	"golang.org/x/sys/execabs"
)

// ConfigAutocorrect sets the local git autocorrect configuration for the given repository.
// The autocorrect setting determines how git handles minor typos in commands.
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

// ConfigUserEmail sets the global git author email.
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

// ConfigUserName configures the user.name git config setting for the given repository.
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

// ConfigSSLVerify configures the http.sslVerify git config setting for the given repository.
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
