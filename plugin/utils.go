package plugin

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/thegeeklab/drone-git-action/git"
	"golang.org/x/sys/execabs"
)

// helper function to simply wrap os execte command.
func execute(cmd *execabs.Cmd) error {
	logrus.Debug("+", strings.Join(cmd.Args, " "))

	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

func rsyncDirectories(pages Pages, repo git.Repository) *execabs.Cmd {
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

	cmd := execabs.Command(
		"rsync",
		args...,
	)
	cmd.Dir = pages.Directory

	return cmd
}
