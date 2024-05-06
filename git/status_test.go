package git

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStatus(t *testing.T) {
	tests := []struct {
		name string
		repo Repository
		want []string
	}{
		{
			name: "with work dir",
			repo: Repository{
				WorkDir: "/path/to/repo",
			},
			want: []string{gitBin, "status", "--porcelain"},
		},
		{
			name: "without work dir",
			repo: Repository{},
			want: []string{gitBin, "status", "--porcelain"},
		},
		{
			name: "with custom stderr",
			repo: Repository{
				WorkDir: "/path/to/repo",
			},
			want: []string{gitBin, "status", "--porcelain"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := tt.repo.Status()
			require.Equal(t, tt.want, cmd.Cmd.Args)
			require.Equal(t, tt.repo.WorkDir, cmd.Cmd.Dir)
		})
	}
}

func TestIsDirty(t *testing.T) {
	tests := []struct {
		name string
		repo Repository
		want bool
	}{
		{
			name: "dirty repo",
			repo: Repository{
				WorkDir: t.TempDir(),
				Branch:  "main",
			},
			want: true,
		},
		{
			name: "clean repo",
			repo: Repository{
				WorkDir: t.TempDir(),
				Branch:  "main",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.repo.Init().Run(); err != nil {
				require.NoError(t, err)
			}

			if tt.want {
				_, err := os.Create(filepath.Join(tt.repo.WorkDir, "dummy"))
				require.NoError(t, err)
			}

			isDirty := tt.repo.IsDirty()
			require.Equal(t, tt.want, isDirty)
		})
	}
}
