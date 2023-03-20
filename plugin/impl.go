package plugin

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/thegeeklab/drone-git-action/git"
	"github.com/urfave/cli/v2"
)

type Netrc struct {
	Machine  string
	Login    string
	Password string
}

type Pages struct {
	Directory string
	Exclude   cli.StringSlice
	Delete    bool
}

// Settings for the Plugin.
type Settings struct {
	Action cli.StringSlice
	SSHKey string

	Netrc Netrc
	Pages Pages
	Repo  git.Repository
}

var (
	ErrAuthSourceNotSet           = errors.New("either SSH key or netrc password is required")
	ErrPagesDirectoryNotExist     = errors.New("pages directory must exist")
	ErrPagesDirectoryNotValid     = errors.New("pages directory not valid")
	ErrPagesSourceNotSet          = errors.New("pages source directory must be set")
	ErrPagesActionNotExclusive    = errors.New("pages action is mutual exclusive")
	ErrActionUnknown              = errors.New("action not found")
	ErrGitCloneDestintionNotValid = errors.New("destination not valid")
)

// Validate handles the settings validation of the plugin.
func (p *Plugin) Validate() error {
	var err error

	p.settings.Repo.Autocorrect = "never"
	p.settings.Repo.RemoteName = "origin"
	p.settings.Repo.Add = ""

	if p.settings.Repo.WorkDir == "" {
		p.settings.Repo.WorkDir, err = os.Getwd()
	}

	if err != nil {
		return err
	}

	for _, action := range p.settings.Action.Value() {
		switch action {
		case "clone":
			continue
		case "commit":
			continue
		case "push":
			if p.settings.SSHKey == "" && p.settings.Netrc.Password == "" {
				return ErrAuthSourceNotSet
			}
		case "pages":
			p.settings.Pages.Directory = filepath.Join(p.settings.Repo.WorkDir, p.settings.Pages.Directory)
			p.settings.Repo.WorkDir = filepath.Join(p.settings.Repo.WorkDir, ".tmp")

			if _, err := os.Stat(p.settings.Pages.Directory); os.IsNotExist(err) {
				return fmt.Errorf("%w: '%s' not found", ErrPagesDirectoryNotExist, p.settings.Pages.Directory)
			}

			if info, _ := os.Stat(p.settings.Pages.Directory); !info.IsDir() {
				return fmt.Errorf("%w: '%s' not a directory", ErrPagesDirectoryNotValid, p.settings.Pages.Directory)
			}

			if p.settings.SSHKey == "" && p.settings.Netrc.Password == "" {
				return ErrAuthSourceNotSet
			}

			if p.settings.Pages.Directory == "" {
				return ErrPagesSourceNotSet
			}

			if len(p.settings.Action.Value()) > 1 {
				return ErrPagesActionNotExclusive
			}
		default:
			return fmt.Errorf("%w: %s", ErrActionUnknown, action)
		}
	}

	return nil
}

// Execute provides the implementation of the plugin.
func (p *Plugin) Execute() error {
	gitEnv := []string{
		"GIT_AUTHOR_NAME",
		"GIT_AUTHOR_EMAIL",
		"GIT_AUTHOR_DATE",
		"GIT_COMMITTER_NAME",
		"GIT_COMMITTER_EMAIL",
		"GIT_COMMITTER_DATE",
	}

	for _, env := range gitEnv {
		if err := os.Unsetenv(env); err != nil {
			return err
		}
	}

	if err := os.Setenv("GIT_TERMINAL_PROMPT", "0"); err != nil {
		return err
	}

	if err := p.handleInit(); err != nil {
		return err
	}

	if err := git.ConfigAutocorrect(p.settings.Repo).Run(); err != nil {
		return err
	}

	if err := git.ConfigUserName(p.settings.Repo).Run(); err != nil {
		return err
	}

	if err := git.ConfigUserEmail(p.settings.Repo).Run(); err != nil {
		return err
	}

	if err := git.ConfigSSLVerify(p.settings.Repo).Run(); err != nil {
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

	for _, action := range p.settings.Action.Value() {
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
		case "pages":
			if err := p.handlePages(); err != nil {
				return err
			}
		}
	}

	return nil
}

// handleInit initializes the repository.
func (p *Plugin) handleInit() error {
	path := filepath.Join(p.settings.Repo.WorkDir, ".git")

	if err := os.MkdirAll(p.settings.Repo.WorkDir, os.ModePerm); err != nil {
		return err
	}

	if _, err := os.Stat(path); !os.IsNotExist(err) {
		p.settings.Repo.InitExists = true

		return nil
	}

	return execute(git.Init(p.settings.Repo))
}

// HandleClone clones remote.
func (p *Plugin) handleClone() error {
	if p.settings.Repo.InitExists {
		return fmt.Errorf("%w: %s exists and not empty", ErrGitCloneDestintionNotValid, p.settings.Repo.WorkDir)
	}

	if p.settings.Repo.RemoteURL != "" {
		if err := execute(git.RemoteAdd(p.settings.Repo)); err != nil {
			return err
		}
	}

	if err := execute(git.FetchSource(p.settings.Repo)); err != nil {
		return err
	}

	return execute(git.CheckoutHead(p.settings.Repo))
}

// HandleCommit commits changes locally.
func (p *Plugin) handleCommit() error {
	if err := execute(git.Add(p.settings.Repo)); err != nil {
		return err
	}

	if err := execute(git.TestCleanTree(p.settings.Repo)); err != nil {
		if err := execute(git.ForceCommit(p.settings.Repo)); err != nil {
			return err
		}
	}

	if p.settings.Repo.EmptyCommit {
		if err := execute(git.EmptyCommit(p.settings.Repo)); err != nil {
			return err
		}
	}

	return nil
}

// HandlePush pushs changes to remote.
func (p *Plugin) handlePush() error {
	return execute(git.RemotePush(p.settings.Repo))
}

// HandlePages syncs, commits and pushes the changes from the pages directory to the pages branch.
func (p *Plugin) handlePages() error {
	defer os.RemoveAll(p.settings.Repo.WorkDir)

	if err := p.handleClone(); err != nil {
		return err
	}

	if err := execute(
		rsyncDirectories(p.settings.Pages, p.settings.Repo),
	); err != nil {
		return err
	}

	if err := p.handleCommit(); err != nil {
		return err
	}

	return p.handlePush()
}
