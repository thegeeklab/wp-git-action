package git

import (
	"fmt"

	"github.com/thegeeklab/wp-plugin-go/v2/types"
	"golang.org/x/sys/execabs"
)

// FetchSource fetches the source from remote.
func FetchSource(repo Repository) *types.Cmd {
	args := []string{
		"fetch",
		"origin",
		fmt.Sprintf("+%s:", repo.Branch),
	}

	cmd := execabs.Command(gitBin, args...)
	cmd.Dir = repo.WorkDir

	return &types.Cmd{
		Cmd: cmd,
	}
}

// CheckoutHead handles branch checkout.
func CheckoutHead(repo Repository) *types.Cmd {
	args := []string{
		"checkout",
		"-qf",
		repo.Branch,
	}

	cmd := execabs.Command(gitBin, args...)
	cmd.Dir = repo.WorkDir

	return &types.Cmd{
		Cmd: cmd,
	}
}
