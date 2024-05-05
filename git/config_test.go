package git

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigAutocorrect(t *testing.T) {
	tests := []struct {
		name string
		repo Repository
		want []string
	}{
		{
			name: "enable autocorrect",
			repo: Repository{
				WorkDir:     "/path/to/repo",
				Autocorrect: "1",
			},
			want: []string{gitBin, "config", "--local", "help.autocorrect", "1"},
		},
		{
			name: "disable autocorrect",
			repo: Repository{
				WorkDir:     "/path/to/repo",
				Autocorrect: "0",
			},
			want: []string{gitBin, "config", "--local", "help.autocorrect", "0"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := ConfigAutocorrect(tt.repo)
			require.Equal(t, tt.want, cmd.Cmd.Args)
			require.Equal(t, tt.repo.WorkDir, cmd.Cmd.Dir)
		})
	}
}

func TestConfigUserEmail(t *testing.T) {
	tests := []struct {
		name string
		repo Repository
		want []string
	}{
		{
			name: "set user email",
			repo: Repository{
				WorkDir: "/path/to/repo",
				Author: Author{
					Email: "user@example.com",
				},
			},
			want: []string{gitBin, "config", "--local", "user.email", "user@example.com"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := ConfigUserEmail(tt.repo)
			require.Equal(t, tt.want, cmd.Cmd.Args)
			require.Equal(t, tt.repo.WorkDir, cmd.Cmd.Dir)
		})
	}
}

func TestConfigUserName(t *testing.T) {
	tests := []struct {
		name string
		repo Repository
		want []string
	}{
		{
			name: "set user name",
			repo: Repository{
				WorkDir: "/path/to/repo",
				Author: Author{
					Name: "John Doe",
				},
			},
			want: []string{gitBin, "config", "--local", "user.name", "John Doe"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := ConfigUserName(tt.repo)
			require.Equal(t, tt.want, cmd.Cmd.Args)
			require.Equal(t, tt.repo.WorkDir, cmd.Cmd.Dir)
		})
	}
}

func TestConfigSSLVerify(t *testing.T) {
	tests := []struct {
		name       string
		repo       Repository
		skipVerify bool
		want       []string
	}{
		{
			name:       "enable SSL verification",
			repo:       Repository{WorkDir: "/path/to/repo"},
			skipVerify: false,
			want:       []string{gitBin, "config", "--local", "http.sslVerify", "true"},
		},
		{
			name:       "disable SSL verification",
			repo:       Repository{WorkDir: "/path/to/repo"},
			skipVerify: true,
			want:       []string{gitBin, "config", "--local", "http.sslVerify", "false"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := ConfigSSLVerify(tt.repo, tt.skipVerify)
			require.Equal(t, tt.want, cmd.Cmd.Args)
			require.Equal(t, tt.repo.WorkDir, cmd.Cmd.Dir)
		})
	}
}
