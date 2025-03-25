package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			},
			want: []string{gitBin, "add", "--all"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := tt.repo.Add()
			assert.Equal(t, tt.want, cmd.Args)
			assert.Equal(t, tt.repo.WorkDir, cmd.Dir)
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
			cmd := tt.repo.IsCleanTree()
			assert.Equal(t, tt.want, cmd.Args)
			assert.Equal(t, tt.repo.WorkDir, cmd.Dir)
		})
	}
}

func TestCommit(t *testing.T) {
	tests := []struct {
		name string
		repo Repository
		want []string
	}{
		{
			name: "commit with message",
			repo: Repository{
				WorkDir:   "/path/to/repo",
				CommitMsg: "Initial commit",
			},
			want: []string{gitBin, "commit", "-m", "Initial commit"},
		},
		{
			name: "commit with empty commit",
			repo: Repository{
				WorkDir:     "/path/to/repo",
				CommitMsg:   "Empty commit",
				EmptyCommit: true,
			},
			want: []string{gitBin, "commit", "-m", "Empty commit", "--allow-empty"},
		},
		{
			name: "commit with no verify",
			repo: Repository{
				WorkDir:   "/path/to/repo",
				CommitMsg: "No verify commit",
				NoVerify:  true,
			},
			want: []string{gitBin, "commit", "-m", "No verify commit", "--no-verify"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := tt.repo.Commit()
			assert.Equal(t, tt.want, cmd.Args)
			assert.Equal(t, tt.repo.WorkDir, cmd.Dir)
		})
	}
}
