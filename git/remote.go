package git

import (
	"fmt"

	"github.com/thegeeklab/wp-plugin-go/v2/types"
	"golang.org/x/sys/execabs"
)

// RemoteRemove drops the defined remote from a git repo.
func (r *Repository) RemoteRemove() *types.Cmd {
	args := []string{
		"remote",
		"rm",
		r.RemoteName,
	}

	cmd := execabs.Command(gitBin, args...)
	cmd.Dir = r.WorkDir

	return &types.Cmd{
		Cmd: cmd,
	}
}

// RemoteAdd adds an additional remote to a git repo.
func (r *Repository) RemoteAdd() *types.Cmd {
	args := []string{
		"remote",
		"add",
		r.RemoteName,
		r.RemoteURL,
	}

	cmd := execabs.Command(gitBin, args...)
	cmd.Dir = r.WorkDir

	return &types.Cmd{
		Cmd: cmd,
	}
}

// RemotePush pushs the changes from the local head to a remote branch.
func (r *Repository) RemotePush() *types.Cmd {
	args := []string{
		"push",
		r.RemoteName,
		fmt.Sprintf("HEAD:%s", r.Branch),
	}

	cmd := execabs.Command(gitBin, args...)
	cmd.Dir = r.WorkDir

	if r.ForcePush {
		cmd.Args = append(cmd.Args, "--force")
	}

	if r.PushFollowTags {
		cmd.Args = append(cmd.Args, "--follow-tags")
	}

	return &types.Cmd{
		Cmd: cmd,
	}
}
