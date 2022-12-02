package git

import (
	"fmt"
	"os"
	"os/exec"
)

// RemoteRemove drops the defined remote from a git repo.
func RemoteRemove(repo Repository) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"remote",
		"rm",
		repo.RemoteName,
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	return cmd
}

// RemoteAdd adds an additional remote to a git repo.
func RemoteAdd(repo Repository) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"remote",
		"add",
		repo.RemoteName,
		repo.RemoteURL,
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	return cmd
}

// RemotePush pushs the changes from the local head to a remote branch.
func RemotePush(repo Repository) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"push",
		repo.RemoteName,
		fmt.Sprintf("HEAD:%s", repo.Branch),
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
