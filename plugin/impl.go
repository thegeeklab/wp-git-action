package plugin

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/thegeeklab/wp-plugin-go/v2/file"
	"github.com/thegeeklab/wp-plugin-go/v2/types"
	"github.com/thegeeklab/wp-plugin-go/v2/util"
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

const (
	ActionClone  Action = "clone"
	ActionCommit Action = "commit"
	ActionPush   Action = "push"
	ActionPages  Action = "pages"
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

	if p.Settings.Repo.WorkDir == "" {
		p.Settings.Repo.WorkDir, err = os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get working directory: %w", err)
		}
	}

	for _, actionStr := range p.Settings.Action.Value() {
		action := Action(actionStr)
		switch action {
		case ActionClone:
			continue
		case ActionCommit:
			continue
		case ActionPush:
			if p.Settings.SSHKey == "" && p.Settings.Netrc.Password == "" {
				return ErrAuthSourceNotSet
			}
		case ActionPages:
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
			return fmt.Errorf("%w: %s", ErrActionUnknown, actionStr)
		}
	}

	return nil
}

// Execute provides the implementation of the plugin.
func (p *Plugin) Execute() error {
	var err error

	homeDir := util.GetUserHomeDir()
	batchCmd := make([]*types.Cmd, 0)
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
			return fmt.Errorf("failed to unset git env vars '%s': %w", env, err)
		}
	}

	if err := os.Setenv("GIT_TERMINAL_PROMPT", "0"); err != nil {
		return fmt.Errorf("failed to git env var': %w", err)
	}

	// Write SSH key and netrc file.
	if p.Settings.SSHKey != "" {
		batchCmd = append(batchCmd, p.Settings.Repo.ConfigSSHCommand(p.Settings.SSHKey))
	}

	netrc := p.Settings.Netrc
	if err := WriteNetrc(homeDir, netrc.Machine, netrc.Login, netrc.Password); err != nil {
		return err
	}

	// Handle repo initialization.
	if err := os.MkdirAll(p.Settings.Repo.WorkDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create working directory: %w", err)
	}

	p.Settings.Repo.IsEmpty, err = file.IsDirEmpty(p.Settings.Repo.WorkDir)
	if err != nil {
		return fmt.Errorf("failed to check working directory: %w", err)
	}

	isDir, err := file.IsDir(filepath.Join(p.Settings.Repo.WorkDir, ".git"))
	if err != nil {
		return fmt.Errorf("failed to check working directory: %w", err)
	}

	if !isDir {
		batchCmd = append(batchCmd, p.Settings.Repo.Init())
	}

	// Handle repo configuration.
	batchCmd = append(batchCmd, p.Settings.Repo.ConfigAutocorrect())
	batchCmd = append(batchCmd, p.Settings.Repo.ConfigUserName())
	batchCmd = append(batchCmd, p.Settings.Repo.ConfigUserEmail())
	batchCmd = append(batchCmd, p.Settings.Repo.ConfigSSLVerify(p.Network.InsecureSkipVerify))

	for _, actionStr := range p.Settings.Action.Value() {
		action := Action(actionStr)
		switch action {
		case ActionClone:
			cmds, err := p.handleClone()
			if err != nil {
				return err
			}

			batchCmd = append(batchCmd, cmds...)
		case ActionCommit:
			batchCmd = append(batchCmd, p.handleCommit()...)
		case ActionPush:
			batchCmd = append(batchCmd, p.handlePush()...)
		case ActionPages:
			cmds, err := p.handlePages()
			if err != nil {
				return err
			}

			batchCmd = append(batchCmd, cmds...)
		}
	}

	for _, cmd := range batchCmd {
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}

// handleClone clones the remote repository into the configured working directory.
// If the working directory is not empty, it returns an error.
func (p *Plugin) handleClone() ([]*types.Cmd, error) {
	var cmds []*types.Cmd

	if !p.Settings.Repo.IsEmpty {
		return cmds, fmt.Errorf("%w: %s exists and not empty", ErrGitCloneDestintionNotValid, p.Settings.Repo.WorkDir)
	}

	if p.Settings.Repo.RemoteURL != "" {
		cmds = append(cmds, p.Settings.Repo.RemoteAdd())
	}

	cmds = append(cmds, p.Settings.Repo.FetchSource())
	cmds = append(cmds, p.Settings.Repo.CheckoutHead())

	return cmds, nil
}

// HandleCommit commits changes locally.
func (p *Plugin) handleCommit() []*types.Cmd {
	var cmds []*types.Cmd

	cmds = append(cmds, p.Settings.Repo.Add())

	if err := p.Settings.Repo.IsCleanTree().Run(); err != nil || p.Settings.Repo.EmptyCommit {
		cmds = append(cmds, p.Settings.Repo.Commit())
	}

	return cmds
}

// HandlePush pushs changes to remote.
func (p *Plugin) handlePush() []*types.Cmd {
	return []*types.Cmd{p.Settings.Repo.RemotePush()}
}

// HandlePages syncs, commits and pushes the changes from the pages directory to the pages branch.
func (p *Plugin) handlePages() ([]*types.Cmd, error) {
	var cmds []*types.Cmd

	defer os.RemoveAll(p.Settings.Repo.WorkDir)

	ccmd, err := p.handleClone()
	if err != nil {
		return cmds, err
	}

	cmds = append(cmds, ccmd...)
	cmds = append(cmds,
		SyncDirectories(
			p.Settings.Pages.Exclude.Value(),
			p.Settings.Pages.Delete,
			p.Settings.Pages.Directory,
			p.Settings.Repo.WorkDir,
		),
	)

	cmds = append(cmds, p.handleCommit()...)
	cmds = append(cmds, p.handlePush()...)

	return cmds, nil
}
