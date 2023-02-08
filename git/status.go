package git

import (
	"bytes"
	"os"

	"github.com/sirupsen/logrus"
	"golang.org/x/sys/execabs"
)

func Status(repo Repository) *execabs.Cmd {
	cmd := execabs.Command(
		gitBin,
		"status",
		"--porcelain",
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	return cmd
}

func IsDirty(repo Repository) bool {
	res := bytes.NewBufferString("")

	cmd := Status(repo)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr
	cmd.Stdout = res
	cmd.Stderr = res

	err := runCommand(cmd)
	if err != nil {
		return false
	}

	if res.Len() > 0 {
		logrus.Debug(res.String())

		return true
	}

	return false
}
