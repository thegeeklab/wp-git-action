package plugin

import (
	"github.com/thegeeklab/wp-plugin-go/v2/types"

	"golang.org/x/sys/execabs"
)

func SyncDirectories(exclude []string, del bool, src, dest string) *types.Cmd {
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

	args = append(
		args,
		".",
		dest,
	)

	cmd := &types.Cmd{
		Cmd: execabs.Command("rsync", args...),
	}
	cmd.Dir = src

	return cmd
}
