package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			cmd := tt.repo.ConfigAutocorrect()
			assert.Equal(t, tt.want, cmd.Args)
			assert.Equal(t, tt.repo.WorkDir, cmd.Dir)
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
			cmd := tt.repo.ConfigUserEmail()
			assert.Equal(t, tt.want, cmd.Args)
			assert.Equal(t, tt.repo.WorkDir, cmd.Dir)
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
			cmd := tt.repo.ConfigUserName()
			assert.Equal(t, tt.want, cmd.Args)
			assert.Equal(t, tt.repo.WorkDir, cmd.Dir)
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
			cmd := tt.repo.ConfigSSLVerify(tt.skipVerify)
			assert.Equal(t, tt.want, cmd.Args)
			assert.Equal(t, tt.repo.WorkDir, cmd.Dir)
		})
	}
}

func TestConfigSSHCommand(t *testing.T) {
	tests := []struct {
		name   string
		repo   Repository
		sshKey string
		want   []string
	}{
		{
			name:   "set SSH command with key",
			repo:   Repository{},
			sshKey: "/path/to/ssh/key",
			want:   []string{gitBin, "config", "--local", "core.sshCommand", "ssh -i /path/to/ssh/key"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := tt.repo.ConfigSSHCommand(tt.sshKey)
			assert.Equal(t, tt.want, cmd.Args)
		})
	}
}
