package plugin

import (
	"os"

	plugin_exec "github.com/thegeeklab/wp-plugin-go/v5/exec"
)

func SyncDirectories(exclude []string, del bool, src, dest string, debug bool) *plugin_exec.Cmd {
	args := []string{
		"-r",
		"--exclude",
		".git",
	}

	for _, item := range exclude {
		args = append(
			args,
			"--exclude",
			item,
		)
	}

	if del {
		args = append(
			args,
			"--delete",
		)
	}

	if debug {
		args = append(
			args,
			"--stats",
		)
	}

	args = append(
		args,
		".",
		dest,
	)

	cmd := plugin_exec.Command("rsync", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = src

	return cmd
}
