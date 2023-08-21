package plugin

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/thegeeklab/wp-git-action/git"
)

var (
	ErrAuthSourceNotSet           = errors.New("either SSH key or netrc password is required")
	ErrPagesDirectoryNotExist     = errors.New("pages directory must exist")
	ErrPagesDirectoryNotValid     = errors.New("pages directory not valid")
	ErrPagesSourceNotSet          = errors.New("pages source directory must be set")
	ErrPagesActionNotExclusive    = errors.New("pages action is mutual exclusive")
	ErrActionUnknown              = errors.New("action not found")
	ErrGitCloneDestintionNotValid = errors.New("destination not valid")
)

//nolint:revive
func (p *Plugin) run(ctx context.Context) error {
	if err := p.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	if err := p.Execute(); err != nil {
		return fmt.Errorf("execution failed: %w", err)
	}

	return nil
}

// Validate handles the settings validation of the plugin.
func (p *Plugin) Validate() error {
	var err error

	p.Settings.Repo.Autocorrect = "never"
	p.Settings.Repo.RemoteName = "origin"
	p.Settings.Repo.Add = ""

	if p.Settings.Repo.WorkDir == "" {
		p.Settings.Repo.WorkDir, err = os.Getwd()
	}

	if err != nil {
		return err
	}

	for _, action := range p.Settings.Action.Value() {
		switch action {
		case "clone":
			continue
		case "commit":
			continue
		case "push":
			if p.Settings.SSHKey == "" && p.Settings.Netrc.Password == "" {
				return ErrAuthSourceNotSet
			}
		case "pages":
			p.Settings.Pages.Directory = filepath.Join(p.Settings.Repo.WorkDir, p.Settings.Pages.Directory)
			p.Settings.Repo.WorkDir = filepath.Join(p.Settings.Repo.WorkDir, ".tmp")

			if _, err := os.Stat(p.Settings.Pages.Directory); os.IsNotExist(err) {
				return fmt.Errorf("%w: '%s' not found", ErrPagesDirectoryNotExist, p.Settings.Pages.Directory)
			}

			if info, _ := os.Stat(p.Settings.Pages.Directory); !info.IsDir() {
				return fmt.Errorf("%w: '%s' not a directory", ErrPagesDirectoryNotValid, p.Settings.Pages.Directory)
			}

			if p.Settings.SSHKey == "" && p.Settings.Netrc.Password == "" {
				return ErrAuthSourceNotSet
			}

			if p.Settings.Pages.Directory == "" {
				return ErrPagesSourceNotSet
			}

			if len(p.Settings.Action.Value()) > 1 {
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

	if err := git.ConfigAutocorrect(p.Settings.Repo).Run(); err != nil {
		return err
	}

	if err := git.ConfigUserName(p.Settings.Repo).Run(); err != nil {
		return err
	}

	if err := git.ConfigUserEmail(p.Settings.Repo).Run(); err != nil {
		return err
	}

	if err := git.ConfigSSLVerify(p.Settings.Repo).Run(); err != nil {
		return err
	}

	if p.Settings.SSHKey != "" {
		if err := git.WriteSSHKey(p.Settings.SSHKey); err != nil {
			return err
		}
	}

	if err := git.WriteNetrc(p.Settings.Netrc.Machine, p.Settings.Netrc.Login, p.Settings.Netrc.Password); err != nil {
		return err
	}

	for _, action := range p.Settings.Action.Value() {
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
	path := filepath.Join(p.Settings.Repo.WorkDir, ".git")

	if err := os.MkdirAll(p.Settings.Repo.WorkDir, os.ModePerm); err != nil {
		return err
	}

	if _, err := os.Stat(path); !os.IsNotExist(err) {
		p.Settings.Repo.InitExists = true

		return nil
	}

	return execute(git.Init(p.Settings.Repo))
}

// HandleClone clones remote.
func (p *Plugin) handleClone() error {
	if p.Settings.Repo.InitExists {
		return fmt.Errorf("%w: %s exists and not empty", ErrGitCloneDestintionNotValid, p.Settings.Repo.WorkDir)
	}

	if p.Settings.Repo.RemoteURL != "" {
		if err := execute(git.RemoteAdd(p.Settings.Repo)); err != nil {
			return err
		}
	}

	if err := execute(git.FetchSource(p.Settings.Repo)); err != nil {
		return err
	}

	return execute(git.CheckoutHead(p.Settings.Repo))
}

// HandleCommit commits changes locally.
func (p *Plugin) handleCommit() error {
	if err := execute(git.Add(p.Settings.Repo)); err != nil {
		return err
	}

	if err := execute(git.TestCleanTree(p.Settings.Repo)); err != nil {
		if err := execute(git.ForceCommit(p.Settings.Repo)); err != nil {
			return err
		}
	}

	if p.Settings.Repo.EmptyCommit {
		if err := execute(git.EmptyCommit(p.Settings.Repo)); err != nil {
			return err
		}
	}

	return nil
}

// HandlePush pushs changes to remote.
func (p *Plugin) handlePush() error {
	return execute(git.RemotePush(p.Settings.Repo))
}

// HandlePages syncs, commits and pushes the changes from the pages directory to the pages branch.
func (p *Plugin) handlePages() error {
	defer os.RemoveAll(p.Settings.Repo.WorkDir)

	if err := p.handleClone(); err != nil {
		return err
	}

	if err := execute(
		rsyncDirectories(p.Settings.Pages, p.Settings.Repo),
	); err != nil {
		return err
	}

	if err := p.handleCommit(); err != nil {
		return err
	}

	return p.handlePush()
}
