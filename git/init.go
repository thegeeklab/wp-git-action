package git

import (
	"github.com/thegeeklab/wp-plugin-go/v2/types"
	"golang.org/x/sys/execabs"
)

// Init creates a new Git repository in the specified directory.
func (r *Repository) Init() *types.Cmd {
	args := []string{
		"init",
		"-b",
		r.Branch,
	}

	cmd := execabs.Command(gitBin, args...)
	cmd.Dir = r.WorkDir

	return &types.Cmd{
		Cmd: cmd,
	}
}
