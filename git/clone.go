package git

import (
	"fmt"
	"os"
	"os/exec"
)

// FetchSource fetches the source from remote.
func FetchSource(repo Repository) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"fetch",
		"origin",
		fmt.Sprintf("+%s:", repo.Branch),
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	return cmd
}

// CheckoutHead handles branch checkout.
func CheckoutHead(repo Repository) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"checkout",
		"-qf",
		repo.Branch,
	)
	cmd.Dir = repo.WorkDir
	cmd.Stderr = os.Stderr

	return cmd
}
