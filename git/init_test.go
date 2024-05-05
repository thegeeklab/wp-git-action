package git

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	repo := Repository{
		WorkDir: "/path/to/repo",
	}

	cmd := Init(repo)
	require.Equal(t, []string{gitBin, "init"}, cmd.Cmd.Args)
	require.Equal(t, repo.WorkDir, cmd.Cmd.Dir)

	// Test with an empty work directory
	repo.WorkDir = ""
	cmd = Init(repo)
	require.Equal(t, []string{gitBin, "init"}, cmd.Cmd.Args)
	require.Empty(t, cmd.Cmd.Dir)
}
