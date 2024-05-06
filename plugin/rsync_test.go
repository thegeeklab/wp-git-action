package plugin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSyncDirectories(t *testing.T) {
	tests := []struct {
		name    string
		exclude []string
		del     bool
		src     string
		dest    string
		want    []string
	}{
		{
			name:    "exclude .git and other patterns",
			exclude: []string{"*.log", "temp/"},
			del:     false,
			src:     "/path/to/src",
			dest:    "/path/to/dest",
			want: []string{
				"rsync", "-r", "--exclude", ".git", "--exclude", "*.log",
				"--exclude", "temp/", ".", "/path/to/dest",
			},
		},
		{
			name:    "delete enabled",
			exclude: []string{},
			del:     true,
			src:     "/path/to/src",
			dest:    "/path/to/dest",
			want:    []string{"rsync", "-r", "--exclude", ".git", "--delete", ".", "/path/to/dest"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := SyncDirectories(tt.exclude, tt.del, tt.src, tt.dest)
			assert.Equal(t, tt.want, cmd.Cmd.Args)
			assert.Equal(t, tt.src, cmd.Cmd.Dir)
		})
	}
}
