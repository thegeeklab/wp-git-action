package git

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		repo Repository
		want []string
	}{
		{
			name: "add all files",
			repo: Repository{
				WorkDir: "/path/to/repo",
				Add:     "",
			},
			want: []string{gitBin, "add", "--all"},
		},
		{
			name: "add specific file",
			repo: Repository{
				WorkDir: "/path/to/repo",
				Add:     "file.go",
			},
			want: []string{gitBin, "add", "file.go"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := Add(tt.repo)
			require.Equal(t, tt.want, cmd.Cmd.Args)
			require.Equal(t, tt.repo.WorkDir, cmd.Cmd.Dir)
		})
	}
}

func TestIsCleanTree(t *testing.T) {
	tests := []struct {
		name string
		repo Repository
		want []string
	}{
		{
			name: "clean working tree",
			repo: Repository{
				WorkDir: "/path/to/repo",
			},
			want: []string{gitBin, "diff-index", "--quiet", "HEAD", "--ignore-submodules"},
		},
		{
			name: "unclean working tree",
			repo: Repository{
				WorkDir: "/path/to/unclean/repo",
			},
			want: []string{gitBin, "diff-index", "--quiet", "HEAD", "--ignore-submodules"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := IsCleanTree(tt.repo)
			require.Equal(t, tt.want, cmd.Cmd.Args)
			require.Equal(t, tt.repo.WorkDir, cmd.Cmd.Dir)
		})
	}
}

func TestEmptyCommit(t *testing.T) {
	tests := []struct {
		name string
		repo Repository
		want []string
	}{
		{
			name: "empty commit with default options",
			repo: Repository{
				WorkDir:   "/path/to/repo",
				CommitMsg: "Empty commit",
			},
			want: []string{gitBin, "commit", "--allow-empty", "-m", "Empty commit"},
		},
		{
			name: "empty commit with no-verify option",
			repo: Repository{
				WorkDir:   "/path/to/repo",
				CommitMsg: "Empty commit",
				NoVerify:  true,
			},
			want: []string{gitBin, "commit", "--allow-empty", "-m", "Empty commit", "--no-verify"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := EmptyCommit(tt.repo)
			require.Equal(t, tt.want, cmd.Cmd.Args)
			require.Equal(t, tt.repo.WorkDir, cmd.Cmd.Dir)
		})
	}
}
