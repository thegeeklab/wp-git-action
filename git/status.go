package git

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func Status(repo Repository) *exec.Cmd {
	cmd := exec.Command(
		"git",
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
		fmt.Print(res.String())
		return true
	}

	return false
}
