package git

import (
	"fmt"

	"github.com/thegeeklab/wp-plugin-go/v3/types"
	"golang.org/x/sys/execabs"
)

// RemoteRemove drops the defined remote from a git repo.
func (r *Repository) RemoteRemove() *types.Cmd {
	args := []string{
		"remote",
		"rm",
		r.RemoteName,
	}

	cmd := &types.Cmd{
		Cmd: execabs.Command(gitBin, args...),
	}
	cmd.Dir = r.WorkDir

	return cmd
}

// RemoteAdd adds an additional remote to a git repo.
func (r *Repository) RemoteAdd() *types.Cmd {
	args := []string{
		"remote",
		"add",
		r.RemoteName,
		r.RemoteURL,
	}

	cmd := &types.Cmd{
		Cmd: execabs.Command(gitBin, args...),
	}
	cmd.Dir = r.WorkDir

	return cmd
}

// RemotePush pushs the changes from the local head to a remote branch.
func (r *Repository) RemotePush() *types.Cmd {
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

	cmd := &types.Cmd{
		Cmd: execabs.Command(gitBin, args...),
	}
	cmd.Dir = r.WorkDir

	return cmd
}
