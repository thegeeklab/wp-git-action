package git

import (
	"os"

	plugin_exec "github.com/thegeeklab/wp-plugin-go/v5/exec"
)

// Add updates the index to match the working tree.
func (r *Repository) Add() *plugin_exec.Cmd {
	cmd := plugin_exec.Command(gitBin, "add", "--all")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = r.WorkDir

	return cmd
}

// IsCleanTree returns non-zero if diff between index and local repository.
func (r *Repository) IsCleanTree() *plugin_exec.Cmd {
	args := []string{
		"diff-index",
		"--quiet",
		"HEAD",
		"--ignore-submodules",
	}

	cmd := plugin_exec.Command(gitBin, args...)
	cmd.Dir = r.WorkDir
	cmd.Trace = false

	return cmd
}

// Commit creates a new commit with the specified commit message.
func (r *Repository) Commit() *plugin_exec.Cmd {
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

	cmd := plugin_exec.Command(gitBin, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = r.WorkDir

	return cmd
}
