package git

import (
	"github.com/thegeeklab/wp-plugin-go/v2/types"
	"golang.org/x/sys/execabs"
)

// ForceAdd forces the addition of all dirty files.
func ForceAdd(repo Repository) *types.Cmd {
	cmd := execabs.Command(
		gitBin,
		"add",
		"--all",
		"--force",
	)
	cmd.Dir = repo.WorkDir

	return &types.Cmd{
		Cmd: cmd,
	}
}

// Add updates the index to match the working tree.
func Add(repo Repository) *types.Cmd {
	cmd := execabs.Command(
		gitBin,
		"add",
	)
	cmd.Dir = repo.WorkDir

	if repo.Add != "" {
		cmd.Args = append(cmd.Args, repo.Add)
	} else {
		cmd.Args = append(cmd.Args, "--all")
	}

	return &types.Cmd{
		Cmd: cmd,
	}
}

// TestCleanTree returns non-zero if diff between index and local repository.
func IsCleanTree(repo Repository) *types.Cmd {
	cmd := execabs.Command(
		gitBin,
		"diff-index",
		"--quiet",
		"HEAD",
		"--ignore-submodules",
	)
	cmd.Dir = repo.WorkDir

	return &types.Cmd{
		Cmd: cmd,
	}
}

// EmptyCommit simply create an empty commit.
func EmptyCommit(repo Repository) *types.Cmd {
	args := []string{
		"commit",
		"--allow-empty",
		"-m",
		repo.CommitMsg,
	}

	cmd := execabs.Command(gitBin, args...)
	cmd.Dir = repo.WorkDir

	if repo.NoVerify {
		cmd.Args = append(cmd.Args, "--no-verify")
	}

	return &types.Cmd{
		Cmd: cmd,
	}
}

func Commit(repo Repository) *types.Cmd {
	args := []string{
		"commit",
		"-m",
		repo.CommitMsg,
	}

	cmd := execabs.Command(gitBin, args...)
	cmd.Dir = repo.WorkDir

	if repo.NoVerify {
		cmd.Args = append(cmd.Args, "--no-verify")
	}

	return &types.Cmd{
		Cmd: cmd,
	}
}
