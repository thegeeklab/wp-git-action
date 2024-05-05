package git

import (
	"github.com/thegeeklab/wp-plugin-go/v2/types"
	"golang.org/x/sys/execabs"
)

// Init creates a new Git repository in the specified directory.
func Init(repo Repository) *types.Cmd {
	cmd := execabs.Command(
		gitBin,
		"init",
	)
	cmd.Dir = repo.WorkDir

	return &types.Cmd{
		Cmd: cmd,
	}
}
