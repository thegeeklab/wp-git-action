package git

import (
	"strconv"

	plugin_exec "github.com/thegeeklab/wp-plugin-go/v3/exec"
)

// ConfigAutocorrect sets the local git autocorrect configuration for the given repository.
// The autocorrect setting determines how git handles minor typos in commands.
func (r *Repository) ConfigAutocorrect() *plugin_exec.Cmd {
	args := []string{
		"config",
		"--local",
		"help.autocorrect",
		r.Autocorrect,
	}

	cmd := plugin_exec.Command(gitBin, args...)
	cmd.Dir = r.WorkDir
	cmd.Trace = false

	return cmd
}

// ConfigUserEmail sets the global git author email.
func (r *Repository) ConfigUserEmail() *plugin_exec.Cmd {
	args := []string{
		"config",
		"--local",
		"user.email",
		r.Author.Email,
	}

	cmd := plugin_exec.Command(gitBin, args...)
	cmd.Dir = r.WorkDir
	cmd.Trace = false

	return cmd
}

// ConfigUserName configures the user.name git config setting for the given repository.
func (r *Repository) ConfigUserName() *plugin_exec.Cmd {
	args := []string{
		"config",
		"--local",
		"user.name",
		r.Author.Name,
	}

	cmd := plugin_exec.Command(gitBin, args...)
	cmd.Dir = r.WorkDir
	cmd.Trace = false

	return cmd
}

// ConfigSSLVerify configures the http.sslVerify git config setting for the given repository.
func (r *Repository) ConfigSSLVerify(skipVerify bool) *plugin_exec.Cmd {
	args := []string{
		"config",
		"--local",
		"http.sslVerify",
		strconv.FormatBool(!skipVerify),
	}

	cmd := plugin_exec.Command(gitBin, args...)
	cmd.Dir = r.WorkDir
	cmd.Trace = false

	return cmd
}

// ConfigSSHCommand sets custom SSH key.
func (r *Repository) ConfigSSHCommand(sshKey string) *plugin_exec.Cmd {
	args := []string{
		"config",
		"--local",
		"core.sshCommand",
		"ssh -i " + sshKey,
	}

	cmd := plugin_exec.Command(gitBin, args...)
	cmd.Dir = r.WorkDir
	cmd.Trace = false

	return cmd
}
