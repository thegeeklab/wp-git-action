package git

import (
	"strconv"

	"github.com/thegeeklab/wp-plugin-go/v3/types"
	"golang.org/x/sys/execabs"
)

// ConfigAutocorrect sets the local git autocorrect configuration for the given repository.
// The autocorrect setting determines how git handles minor typos in commands.
func (r *Repository) ConfigAutocorrect() *types.Cmd {
	args := []string{
		"config",
		"--local",
		"help.autocorrect",
		r.Autocorrect,
	}

	cmd := &types.Cmd{
		Cmd: execabs.Command(gitBin, args...),
	}
	cmd.Dir = r.WorkDir
	cmd.SetTrace(false)

	return cmd
}

// ConfigUserEmail sets the global git author email.
func (r *Repository) ConfigUserEmail() *types.Cmd {
	args := []string{
		"config",
		"--local",
		"user.email",
		r.Author.Email,
	}

	cmd := &types.Cmd{
		Cmd: execabs.Command(gitBin, args...),
	}
	cmd.Dir = r.WorkDir
	cmd.SetTrace(false)

	return cmd
}

// ConfigUserName configures the user.name git config setting for the given repository.
func (r *Repository) ConfigUserName() *types.Cmd {
	args := []string{
		"config",
		"--local",
		"user.name",
		r.Author.Name,
	}

	cmd := &types.Cmd{
		Cmd: execabs.Command(gitBin, args...),
	}
	cmd.Dir = r.WorkDir
	cmd.SetTrace(false)

	return cmd
}

// ConfigSSLVerify configures the http.sslVerify git config setting for the given repository.
func (r *Repository) ConfigSSLVerify(skipVerify bool) *types.Cmd {
	args := []string{
		"config",
		"--local",
		"http.sslVerify",
		strconv.FormatBool(!skipVerify),
	}

	cmd := &types.Cmd{
		Cmd: execabs.Command(gitBin, args...),
	}
	cmd.Dir = r.WorkDir
	cmd.SetTrace(false)

	return cmd
}

// ConfigSSHCommand sets custom SSH key.
func (r *Repository) ConfigSSHCommand(sshKey string) *types.Cmd {
	args := []string{
		"config",
		"--local",
		"core.sshCommand",
		"ssh -i " + sshKey,
	}

	cmd := &types.Cmd{
		Cmd: execabs.Command(gitBin, args...),
	}
	cmd.Dir = r.WorkDir
	cmd.SetTrace(false)

	return cmd
}
