package git

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFetchSource(t *testing.T) {
	tests := []struct {
		name string
		repo Repository
		want []string
	}{
		{
			name: "fetch from origin with branch",
			repo: Repository{
				WorkDir: "/path/to/repo",
				Branch:  "main",
			},
			want: []string{gitBin, "fetch", "origin", "+main:"},
		},
		{
			name: "fetch from origin with different branch",
			repo: Repository{
				WorkDir: "/path/to/repo",
				Branch:  "develop",
			},
			want: []string{gitBin, "fetch", "origin", "+develop:"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := tt.repo.FetchSource()
			require.Equal(t, tt.want, cmd.Cmd.Args)
			require.Equal(t, tt.repo.WorkDir, cmd.Cmd.Dir)
		})
	}
}

func TestCheckoutHead(t *testing.T) {
	tests := []struct {
		name string
		repo Repository
		want []string
	}{
		{
			name: "checkout head with branch",
			repo: Repository{
				WorkDir: "/path/to/repo",
				Branch:  "main",
			},
			want: []string{gitBin, "checkout", "-qf", "main"},
		},
		{
			name: "checkout head with different branch",
			repo: Repository{
				WorkDir: "/path/to/repo",
				Branch:  "develop",
			},
			want: []string{gitBin, "checkout", "-qf", "develop"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := tt.repo.CheckoutHead()
			require.Equal(t, tt.want, cmd.Cmd.Args)
			require.Equal(t, tt.repo.WorkDir, cmd.Cmd.Dir)
		})
	}
}
