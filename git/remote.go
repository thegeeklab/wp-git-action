package git

import (
	"fmt"

	"github.com/thegeeklab/wp-plugin-go/v2/types"
	"golang.org/x/sys/execabs"
)

// RemoteRemove drops the defined remote from a git repo.
func RemoteRemove(repo Repository) *types.Cmd {
	args := []string{
		"remote",
		"rm",
		repo.RemoteName,
	}

	cmd := execabs.Command(gitBin, args...)
	cmd.Dir = repo.WorkDir

	return &types.Cmd{
		Cmd: cmd,
	}
}

// RemoteAdd adds an additional remote to a git repo.
func RemoteAdd(repo Repository) *types.Cmd {
	args := []string{
		"remote",
		"add",
		repo.RemoteName,
		repo.RemoteURL,
	}

	cmd := execabs.Command(gitBin, args...)
	cmd.Dir = repo.WorkDir

	return &types.Cmd{
		Cmd: cmd,
	}
}

// RemotePush pushs the changes from the local head to a remote branch.
func RemotePush(repo Repository) *types.Cmd {
	args := []string{
		"push",
		repo.RemoteName,
		fmt.Sprintf("HEAD:%s", repo.Branch),
	}

	cmd := execabs.Command(gitBin, args...)
	cmd.Dir = repo.WorkDir

	if repo.ForcePush {
		cmd.Args = append(cmd.Args, "--force")
	}

	if repo.PushFollowTags {
		cmd.Args = append(cmd.Args, "--follow-tags")
	}

	return &types.Cmd{
		Cmd: cmd,
	}
}
