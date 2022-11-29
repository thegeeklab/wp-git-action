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
	Actions           cli.StringSlice
	SSHKey            string
	Remote            string
	Branch            string
	Path              string
	Message           string
	Force             bool
	FollowTags        bool
	InsecureSSLVerify bool
	EmptyCommit       bool
	NoVerify          bool

	Netrc  Netrc
	Commit Commit
	Author Author
}

// Validate handles the settings validation of the plugin.
func (p *Plugin) Validate() error {
	for _, action := range p.settings.Actions.Value() {
		switch action {
		case "clone":
			continue
		case "commit":
			continue
		case "push":
			if p.settings.SSHKey == "" && p.settings.Netrc.Password == "" {
				return fmt.Errorf("either SSH key or netrc password are required")
			}
		default:
			return fmt.Errorf("unknown action %s", action)
		}
	}

	return nil
}

// Execute provides the implementation of the plugin.
func (p *Plugin) Execute() error {
	for _, env := range []string{"GIT_AUTHOR_NAME", "GIT_AUTHOR_EMAIL"} {
		if err := os.Unsetenv(env); err != nil {
			return err
		}
	}
	if err := os.Setenv("GIT_TERMINAL_PROMPT", "0"); err != nil {
		return err
	}

	if p.settings.Path != "" {
		if err := p.initRepo(); err != nil {
			return err
		}
	}

	if err := git.SetUserName(p.settings.Commit.Author.Name).Run(); err != nil {
		return err
	}
	if err := git.SetUserEmail(p.settings.Commit.Author.Email).Run(); err != nil {
		return err
	}
	if err := git.SetSSLVerify(p.settings.InsecureSSLVerify).Run(); err != nil {
		return err
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
		}
	}

	return nil
}
