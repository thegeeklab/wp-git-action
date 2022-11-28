package plugin

import (
	"fmt"
	"os"

	"github.com/thegeeklab/drone-git-action/git"
	"github.com/urfave/cli/v2"
)

type Netrc struct {
	Machine  string
	Login    string
	Password string
}

type Commit struct {
	Author Author
}

type Author struct {
	Name  string
	Email string
}

// Settings for the Plugin.
type Settings struct {
	Actions     cli.StringSlice
	SSHKey      string
	Remote      string
	Branch      string
	Path        string
	Message     string
	Force       bool
	FollowTags  bool
	SkipVerify  bool
	EmptyCommit bool
	NoVerify    bool

	Netrc  Netrc
	Commit Commit
	Author Author
}

// Validate handles the settings validation of the plugin.
func (p *Plugin) Validate() error {
	if p.settings.SSHKey == "" && p.settings.Netrc.Password == "" {
		return fmt.Errorf("either SSH key or netrc password are required")
	}

	return nil
}

// Execute provides the implementation of the plugin.
func (p *Plugin) Execute() error {
	if p.settings.Path != "" {
		if err := os.MkdirAll(p.settings.Path, os.ModePerm); err != nil {
			return err
		}

		if err := os.Chdir(p.settings.Path); err != nil {
			return err
		}
	}

	if err := git.GlobalName(p.settings.Commit.Author.Name).Run(); err != nil {
		return err
	}

	if err := git.GlobalUser(p.settings.Commit.Author.Email).Run(); err != nil {
		return err
	}

	if p.settings.SkipVerify {
		if err := git.SkipVerify().Run(); err != nil {
			return err
		}
	}

	if p.settings.SSHKey != "" {
		if err := git.WriteSSHKey(p.settings.SSHKey); err != nil {
			return err
		}
	}

	if err := git.WriteNetrc(p.settings.Netrc.Machine, p.settings.Netrc.Login, p.settings.Netrc.Password); err != nil {
		return err
	}

	for _, action := range p.settings.Actions.Value() {
		switch action {
		case "clone":
			if err := p.handleClone(); err != nil {
				return err
			}
		case "commit":
			if err := p.handleCommit(); err != nil {
				return err
			}
		case "push":
			if err := p.handlePush(); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unknown action %s", action)
		}
	}

	return nil
}
