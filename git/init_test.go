package git

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name     string
		repo     Repository
		expected []string
	}{
		{
			name: "init repo",
			repo: Repository{
				WorkDir: "/path/to/repo",
				Branch:  "main",
			},
			expected: []string{gitBin, "init", "-b", "main"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := tt.repo.Init()
			require.Equal(t, tt.expected, cmd.Cmd.Args)
			require.Equal(t, tt.repo.WorkDir, cmd.Cmd.Dir)
		})
	}
}
