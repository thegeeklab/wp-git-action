package git

import (
	"github.com/rs/zerolog/log"
	"github.com/thegeeklab/wp-plugin-go/v2/types"
	"golang.org/x/sys/execabs"
)

// Status returns a command that runs `git status --porcelain` for the given repository.
func (r Repository) Status() *types.Cmd {
	cmd := execabs.Command(
		gitBin,
		"status",
		"--porcelain",
	)
	cmd.Dir = r.WorkDir

	return &types.Cmd{
		Cmd: cmd,
	}
}

// IsDirty checks if the given repository has any uncommitted changes.
// It runs `git status --porcelain` and returns true if the output is non-empty,
// indicating that there are uncommitted changes in the repository.
// If there is an error running the git command, it returns false.
func (r Repository) IsDirty() bool {
	cmd := r.Status()
	cmd.Dir = r.WorkDir

	out, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}

	if len(out) > 0 {
		log.Debug().Msg(string(out))

		return true
	}

	return false
}
