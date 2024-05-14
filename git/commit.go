package git

import (
	"io"

	"github.com/thegeeklab/wp-plugin-go/v2/types"
	"golang.org/x/sys/execabs"
)

// Add updates the index to match the working tree.
func (r *Repository) Add() *types.Cmd {
	cmd := &types.Cmd{
		Cmd: execabs.Command(
			gitBin,
			"add",
			"--all",
		),
	}
	cmd.Dir = r.WorkDir

	return cmd
}

// IsCleanTree returns non-zero if diff between index and local repository.
func (r *Repository) IsCleanTree() *types.Cmd {
	args := []string{
		"diff-index",
		"--quiet",
		"HEAD",
		"--ignore-submodules",
	}

	cmd := &types.Cmd{
		Cmd: execabs.Command(gitBin, args...),
	}
	cmd.Dir = r.WorkDir
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	cmd.SetTrace(false)

	return cmd
}

// Commit creates a new commit with the specified commit message.
func (r *Repository) Commit() *types.Cmd {
	args := []string{
		"commit",
		"-m",
		r.CommitMsg,
	}

	if r.EmptyCommit {
		args = append(args, "--allow-empty")
	}

	if r.NoVerify {
		args = append(args, "--no-verify")
	}

	cmd := &types.Cmd{
		Cmd: execabs.Command(gitBin, args...),
	}
	cmd.Dir = r.WorkDir

	return cmd
}
