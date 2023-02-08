package git

import (
	"fmt"
	"os"

	"golang.org/x/sys/execabs"
)

// FetchSource fetches the source from remote.
func FetchSource(repo Repository) *execabs.Cmd {
	args := []string{
		"fetch",
		"origin",
		fmt.Sprintf("+%s:", repo.Branch),
	}

	cmd := execabs.Command(
		gitBin,
		args...,
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	return cmd
}

// CheckoutHead handles branch checkout.
func CheckoutHead(repo Repository) *execabs.Cmd {
	args := []string{
		"checkout",
		"-qf",
		repo.Branch,
	}

	cmd := execabs.Command(
		gitBin,
		args...,
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	return cmd
}
