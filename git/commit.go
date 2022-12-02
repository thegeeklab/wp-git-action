package git

import (
	"os"
	"os/exec"
)

// ForceAdd forces the addition of all dirty files.
func ForceAdd(repo Repository) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"add",
		"--all",
		"--force",
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	return cmd
}

// Add updates the index to match the working tree.
func Add(repo Repository) *exec.Cmd {
	cmd := exec.Command(
		"git",
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

// TestCleanTree returns non-zero if diff between index and local repository
func TestCleanTree(repo Repository) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"diff-index",
		"--quiet",
		"HEAD",
		"--ignore-submodules",
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	return cmd
}

// EmptyCommit simply create an empty commit
func EmptyCommit(repo Repository) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"commit",
		"--allow-empty",
		"-m",
		repo.CommitMsg,
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	if repo.NoVerify {
		cmd.Args = append(cmd.Args, "--no-verify")
	}

	return cmd
}

// ForceCommit commits every change while skipping CI.
func ForceCommit(repo Repository) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"commit",
		"-m",
		repo.CommitMsg,
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	if repo.NoVerify {
		cmd.Args = append(cmd.Args, "--no-verify")
	}

	return cmd
}
