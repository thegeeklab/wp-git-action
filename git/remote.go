package git

import (
	"fmt"
	"os"

	"golang.org/x/sys/execabs"
)

// RemoteRemove drops the defined remote from a git repo.
func RemoteRemove(repo Repository) *execabs.Cmd {
	args := []string{
		"remote",
		"rm",
		repo.RemoteName,
	}

	cmd := execabs.Command(
		gitBin,
		args...,
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	return cmd
}

// RemoteAdd adds an additional remote to a git repo.
func RemoteAdd(repo Repository) *execabs.Cmd {
	args := []string{
		"remote",
		"add",
		repo.RemoteName,
		repo.RemoteURL,
	}

	cmd := execabs.Command(
		gitBin,
		args...,
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	return cmd
}

// RemotePush pushs the changes from the local head to a remote branch.
func RemotePush(repo Repository) *execabs.Cmd {
	args := []string{
		"push",
		repo.RemoteName,
		fmt.Sprintf("HEAD:%s", repo.Branch),
	}

	cmd := execabs.Command(
		gitBin,
		args...,
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	if repo.ForcePush {
		cmd.Args = append(
			cmd.Args,
			"--force",
		)
	}

	if repo.PushFollowTags {
		cmd.Args = append(
			cmd.Args,
			"--follow-tags")
	}

	return cmd
}
