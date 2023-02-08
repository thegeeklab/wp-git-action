package git

import (
	"os"

	"golang.org/x/sys/execabs"
)

// ForceAdd forces the addition of all dirty files.
func ForceAdd(repo Repository) *execabs.Cmd {
	cmd := execabs.Command(
		gitBin,
		"add",
		"--all",
		"--force",
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	return cmd
}

// Add updates the index to match the working tree.
func Add(repo Repository) *execabs.Cmd {
	cmd := execabs.Command(
		gitBin,
		"add",
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	if repo.Add != "" {
		cmd.Args = append(cmd.Args, repo.Add)
	} else {
		cmd.Args = append(cmd.Args, "--all")
	}

	return cmd
}

// TestCleanTree returns non-zero if diff between index and local repository.
func TestCleanTree(repo Repository) *execabs.Cmd {
	cmd := execabs.Command(
		gitBin,
		"diff-index",
		"--quiet",
		"HEAD",
		"--ignore-submodules",
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	return cmd
}

// EmptyCommit simply create an empty commit.
func EmptyCommit(repo Repository) *execabs.Cmd {
	args := []string{
		"commit",
		"--allow-empty",
		"-m",
		repo.CommitMsg,
	}

	cmd := execabs.Command(
		gitBin,
		args...,
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	if repo.NoVerify {
		cmd.Args = append(cmd.Args, "--no-verify")
	}

	return cmd
}

// ForceCommit commits every change while skipping CI.
func ForceCommit(repo Repository) *execabs.Cmd {
	args := []string{
		"commit",
		"-m",
		repo.CommitMsg,
	}

	cmd := execabs.Command(
		gitBin,
		args...,
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	if repo.NoVerify {
		cmd.Args = append(cmd.Args, "--no-verify")
	}

	return cmd
}
