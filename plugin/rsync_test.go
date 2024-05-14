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
		debug   bool
		src     string
		dest    string
		want    []string
	}{
		{
			name:    "exclude .git and other patterns",
			exclude: []string{"*.log", "temp/"},
			del:     false,
			debug:   false,
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
			debug:   false,
			src:     "/path/to/src",
			dest:    "/path/to/dest",
			want:    []string{"rsync", "-r", "--exclude", ".git", "--delete", ".", "/path/to/dest"},
		},
		{
			name:    "debug output enabled",
			exclude: []string{},
			del:     false,
			debug:   true,
			src:     "/path/to/src",
			dest:    "/path/to/dest",
			want:    []string{"rsync", "-r", "--exclude", ".git", "--stats", ".", "/path/to/dest"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := SyncDirectories(tt.exclude, tt.del, tt.src, tt.dest, tt.debug)
			assert.Equal(t, tt.want, cmd.Cmd.Args)
			assert.Equal(t, tt.src, cmd.Cmd.Dir)
		})
	}
}
