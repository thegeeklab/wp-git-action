package git

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/thegeeklab/wp-plugin-go/v2/types"
	"golang.org/x/sys/execabs"
)

// Status returns a command that runs `git status --porcelain` in the given repository's working directory.
func Status(repo Repository) *types.Cmd {
	cmd := execabs.Command(
		gitBin,
		"status",
		"--porcelain",
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	return &types.Cmd{
		Cmd: cmd,
	}
}

// IsDirty checks if the given repository has any uncommitted changes.
// It runs the `git status --porcelain` command and returns true if the output is non-empty,
// indicating that there are uncommitted changes in the repository.
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
