package git

import (
	"os"

	plugin_exec "github.com/thegeeklab/wp-plugin-go/v5/exec"
)

// Init creates a new Git repository in the specified directory.
func (r *Repository) Init() *plugin_exec.Cmd {
	args := []string{
		"init",
		"-b",
		r.Branch,
	}

	cmd := plugin_exec.Command(gitBin, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = r.WorkDir

	return cmd
}
