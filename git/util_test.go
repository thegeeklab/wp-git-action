package git

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriteSSHKey(t *testing.T) {
	tests := []struct {
		name       string
		privateKey string
		dir        string
		wantErr    bool
	}{
		{
			name:       "valid private key",
			privateKey: "valid_private_key",
			dir:        t.TempDir(),
			wantErr:    false,
		},
		{
			name:       "empty private key",
			privateKey: "",
			dir:        t.TempDir(),
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WriteSSHKey(tt.dir, tt.privateKey)
			if tt.wantErr {
				require.Error(t, err)

				return
			}

			require.NoError(t, err)

			privateKeyPath := filepath.Join(tt.dir, ".ssh", "id_rsa")
			_, err = os.Stat(privateKeyPath)
			require.NoError(t, err)

			configPath := filepath.Join(tt.dir, ".ssh", "config")
			_, err = os.Stat(configPath)
			require.NoError(t, err)
		})
	}
}

func TestWriteNetrc(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		machine  string
		login    string
		password string
		wantErr  bool
	}{
		{
			name:     "valid input",
			path:     t.TempDir(),
			machine:  "example.com",
			login:    "user",
			password: "pass",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WriteNetrc(tt.path, tt.machine, tt.login, tt.password)
			if tt.wantErr {
				require.Error(t, err)

				return
			}

			require.NoError(t, err)

			netrcPath := filepath.Join(tt.path, ".netrc")
			_, err = os.Stat(netrcPath)
			require.NoError(t, err)

			content, err := os.ReadFile(netrcPath)
			require.NoError(t, err)

			expected := fmt.Sprintf("machine %s\nlogin %s\npassword %s\n", tt.machine, tt.login, tt.password)
			require.Equal(t, expected, string(content))
		})
	}
}
