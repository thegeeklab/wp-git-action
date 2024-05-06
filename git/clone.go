package git

import (
	"fmt"

	"github.com/thegeeklab/wp-plugin-go/v2/types"
	"golang.org/x/sys/execabs"
)

// FetchSource fetches the source from remote.
func (r Repository) FetchSource() *types.Cmd {
	args := []string{
		"fetch",
		"origin",
		fmt.Sprintf("+%s:", r.Branch),
	}

	cmd := execabs.Command(gitBin, args...)
	cmd.Dir = r.WorkDir

	return &types.Cmd{
		Cmd: cmd,
	}
}

// CheckoutHead handles branch checkout.
func (r Repository) CheckoutHead() *types.Cmd {
	args := []string{
		"checkout",
		"-qf",
		r.Branch,
	}

	cmd := execabs.Command(gitBin, args...)
	cmd.Dir = r.WorkDir

	return &types.Cmd{
		Cmd: cmd,
	}
}
