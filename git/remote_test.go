package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoteRemove(t *testing.T) {
	tests := []struct {
		name string
		repo Repository
		want []string
	}{
		{
			name: "remove remote",
			repo: Repository{
				WorkDir:    "/path/to/repo",
				RemoteName: "origin",
			},
			want: []string{gitBin, "remote", "rm", "origin"},
		},
		{
			name: "remove custom remote name",
			repo: Repository{
				WorkDir:    "/path/to/repo",
				RemoteName: "upstream",
			},
			want: []string{gitBin, "remote", "rm", "upstream"},
		},
		{
			name: "remove remote with empty work dir",
			repo: Repository{
				RemoteName: "origin",
			},
			want: []string{gitBin, "remote", "rm", "origin"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := tt.repo.RemoteRemove()
			assert.Equal(t, tt.want, cmd.Cmd.Args)
			assert.Equal(t, tt.repo.WorkDir, cmd.Cmd.Dir)
		})
	}
}

func TestRemoteAdd(t *testing.T) {
	tests := []struct {
		name string
		repo Repository
		want []string
	}{
		{
			name: "add remote with valid inputs",
			repo: Repository{
				WorkDir:    "/path/to/repo",
				RemoteName: "origin",
				RemoteURL:  "https://example.com/repo.git",
			},
			want: []string{gitBin, "remote", "add", "origin", "https://example.com/repo.git"},
		},
		{
			name: "add remote with empty work dir",
			repo: Repository{
				RemoteName: "origin",
				RemoteURL:  "https://example.com/repo.git",
			},
			want: []string{gitBin, "remote", "add", "origin", "https://example.com/repo.git"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := tt.repo.RemoteAdd()
			assert.Equal(t, tt.want, cmd.Cmd.Args)
			assert.Equal(t, tt.repo.WorkDir, cmd.Cmd.Dir)
		})
	}
}

func TestRemotePush(t *testing.T) {
	tests := []struct {
		name       string
		repo       Repository
		want       []string
		forcePush  bool
		followTags bool
	}{
		{
			name: "push with default options",
			repo: Repository{
				WorkDir:    "/path/to/repo",
				RemoteName: "origin",
				Branch:     "main",
			},
			want: []string{gitBin, "push", "origin", "HEAD:main"},
		},
		{
			name: "push with force option",
			repo: Repository{
				WorkDir:    "/path/to/repo",
				RemoteName: "origin",
				Branch:     "main",
				ForcePush:  true,
			},
			want:      []string{gitBin, "push", "origin", "HEAD:main", "--force"},
			forcePush: true,
		},
		{
			name: "push with follow tags option",
			repo: Repository{
				WorkDir:        "/path/to/repo",
				RemoteName:     "origin",
				Branch:         "main",
				PushFollowTags: true,
			},
			want:       []string{gitBin, "push", "origin", "HEAD:main", "--follow-tags"},
			followTags: true,
		},
		{
			name: "push with empty work dir",
			repo: Repository{
				RemoteName: "origin",
				Branch:     "main",
			},
			want: []string{gitBin, "push", "origin", "HEAD:main"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := tt.repo.RemotePush()
			assert.Equal(t, tt.want, cmd.Cmd.Args)
			assert.Equal(t, tt.repo.WorkDir, cmd.Cmd.Dir)
		})
	}
}
