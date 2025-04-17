package git

import (
	"fmt"
	"os"

	plugin_exec "github.com/thegeeklab/wp-plugin-go/v6/exec"
)

// RemoteRemove drops the defined remote from a git repo.
func (r *Repository) RemoteRemove() *plugin_exec.Cmd {
	args := []string{
		"remote",
		"rm",
		r.RemoteName,
	}

	cmd := plugin_exec.Command(gitBin, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = r.WorkDir

	return cmd
}

// RemoteAdd adds an additional remote to a git repo.
func (r *Repository) RemoteAdd() *plugin_exec.Cmd {
	args := []string{
		"remote",
		"add",
		r.RemoteName,
		r.RemoteURL,
	}

	cmd := plugin_exec.Command(gitBin, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = r.WorkDir

	return cmd
}

// RemotePush pushs the changes from the local head to a remote branch.
func (r *Repository) RemotePush() *plugin_exec.Cmd {
	args := []string{
		"push",
		r.RemoteName,
		fmt.Sprintf("HEAD:%s", r.Branch),
	}

	if r.ForcePush {
		args = append(args, "--force")
	}

	if r.PushFollowTags {
		args = append(args, "--follow-tags")
	}

	cmd := plugin_exec.Command(gitBin, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = r.WorkDir

	return cmd
}
