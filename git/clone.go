package git

import (
	"fmt"
	"os"

	plugin_exec "github.com/thegeeklab/wp-plugin-go/v6/exec"
)

// FetchSource fetches the source from remote.
func (r *Repository) FetchSource() *plugin_exec.Cmd {
	args := []string{
		"fetch",
		"origin",
		fmt.Sprintf("+%s:", r.Branch),
	}

	cmd := plugin_exec.Command(gitBin, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = r.WorkDir

	return cmd
}

// CheckoutHead handles branch checkout.
func (r *Repository) CheckoutHead() *plugin_exec.Cmd {
	args := []string{
		"checkout",
		"-qf",
		r.Branch,
	}

	cmd := plugin_exec.Command(gitBin, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = r.WorkDir

	return cmd
}
