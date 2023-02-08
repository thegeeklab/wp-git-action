package git

import (
	"os"

	"golang.org/x/sys/execabs"
)

// RemoteRemove drops the defined remote from a git repo.
func Init(repo Repository) *execabs.Cmd {
	cmd := execabs.Command(
		gitBin,
		"init",
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	return cmd
}
