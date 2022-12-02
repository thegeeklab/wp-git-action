package plugin

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/thegeeklab/drone-git-action/git"
)

// helper function to simply wrap os execte command.
func execute(cmd *exec.Cmd) error {
	fmt.Println("+", strings.Join(cmd.Args, " "))

	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

func rsyncDirectories(pages Pages, repo git.Repository) *exec.Cmd {
	args := []string{
		"-r",
		"--exclude",
		".git",
	}

	for _, item := range pages.Exclude.Value() {
		args = append(
			args,
			"--exclude",
			item,
		)
	}

	if pages.Delete {
		args = append(
			args,
			"--delete",
		)
	}

	args = append(
		args,
		".",
		repo.WorkDir,
	)

	cmd := exec.Command(
		"rsync",
		args...,
	)
	cmd.Dir = pages.Directory

	return cmd
}
