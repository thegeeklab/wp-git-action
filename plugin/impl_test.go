package plugin

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thegeeklab/wp-git-action/git"
)

func TestValidate_CommitMessageFrom(t *testing.T) {
	const ciVar = "TEST_CI_COMMIT_MESSAGE"
	const staticDefault = "[skip ci] commit dirty state"

	tests := []struct {
		name              string
		commitMessageFrom string
		envValue          string
		initialMsg        string
		wantMsg           string
	}{
		{
			name:              "from-flag unset leaves default",
			commitMessageFrom: "",
			initialMsg:        staticDefault,
			wantMsg:           staticDefault,
		},
		{
			name:              "from-flag set but env empty falls back to message",
			commitMessageFrom: ciVar,
			envValue:          "",
			initialMsg:        "scheduled sync",
			wantMsg:           "scheduled sync",
		},
		{
			name:              "from-flag set and env populated overrides default",
			commitMessageFrom: ciVar,
			envValue:          "feat(api): add motion photo facet",
			initialMsg:        staticDefault,
			wantMsg:           "feat(api): add motion photo facet",
		},
		{
			name:              "from-flag set and env populated overrides explicit message",
			commitMessageFrom: ciVar,
			envValue:          "from CI",
			initialMsg:        "explicit message",
			wantMsg:           "from CI",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(ciVar, tt.envValue)

			p := &Plugin{Settings: &Settings{
				Action:            []string{"commit"},
				CommitMessageFrom: tt.commitMessageFrom,
				Repo:              git.Repository{CommitMsg: tt.initialMsg, WorkDir: t.TempDir()},
			}}

			assert.NoError(t, p.Validate())
			assert.Equal(t, tt.wantMsg, p.Settings.Repo.CommitMsg)
		})
	}
}
