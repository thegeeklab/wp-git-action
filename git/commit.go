package git

import (
	"github.com/thegeeklab/wp-plugin-go/v2/types"
	"golang.org/x/sys/execabs"
)

// Add updates the index to match the working tree.
func (r *Repository) Add() *types.Cmd {
	cmd := execabs.Command(
		gitBin,
		"add",
		"--all",
	)
	cmd.Dir = r.WorkDir

	return &types.Cmd{
		Cmd: cmd,
	}
}

// TestCleanTree returns non-zero if diff between index and local repository.
func (r *Repository) IsCleanTree() *types.Cmd {
	cmd := execabs.Command(
		gitBin,
		"diff-index",
		"--quiet",
		"HEAD",
		"--ignore-submodules",
	)
	cmd.Dir = r.WorkDir

	return &types.Cmd{
		Cmd: cmd,
	}
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

	cmd := execabs.Command(gitBin, args...)
	cmd.Dir = r.WorkDir

	return &types.Cmd{
		Cmd: cmd,
	}
}
