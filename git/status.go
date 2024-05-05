package git

import (
	"github.com/rs/zerolog/log"
	"github.com/thegeeklab/wp-plugin-go/v2/types"
	"golang.org/x/sys/execabs"
)

// Status returns a command that runs `git status --porcelain` for the given repository.
func Status(repo Repository) *types.Cmd {
	cmd := execabs.Command(
		gitBin,
		"status",
		"--porcelain",
	)
	cmd.Dir = repo.WorkDir

	return &types.Cmd{
		Cmd: cmd,
	}
}

// IsDirty checks if the given repository has any uncommitted changes.
// It runs `git status --porcelain` and returns true if the output is non-empty,
// indicating that there are uncommitted changes in the repository.
// If there is an error running the git command, it returns false.
func IsDirty(repo Repository) bool {
	cmd := Status(repo)
	cmd.Dir = repo.WorkDir

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
