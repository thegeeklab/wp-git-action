package plugin

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	plugin_exec "github.com/thegeeklab/wp-plugin-go/v6/exec"
	plugin_file "github.com/thegeeklab/wp-plugin-go/v6/file"
	plugin_util "github.com/thegeeklab/wp-plugin-go/v6/util"
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
	GitActionClone  GitAction = "clone"
	GitActionCommit GitAction = "commit"
	GitActionPush   GitAction = "push"
	GitActionPages  GitAction = "pages"
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

	for _, actionStr := range p.Settings.Action {
		action := GitAction(actionStr)
		switch action {
		case GitActionClone:
			continue
		case GitActionCommit:
			continue
		case GitActionPush:
			if p.Settings.SSHKey == "" && p.Settings.Netrc.Password == "" {
				return ErrAuthSourceNotSet
			}
		case GitActionPages:
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

			if len(p.Settings.Action) > 1 {
				return ErrPagesActionNotExclusive
			}
		default:
			return fmt.Errorf("%w: %s", ErrActionUnknown, actionStr)
		}
	}

	return nil
}

// Execute provides the implementation of the plugin.
//
//nolint:gocognit
func (p *Plugin) Execute() error {
	var err error

	homeDir := plugin_util.GetUserHomeDir()
	batchCmd := make([]*plugin_exec.Cmd, 0)
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

	if p.Settings.Repo.Cleanup {
		defer os.RemoveAll(p.Settings.Repo.WorkDir)
	}

	p.Settings.Repo.IsEmpty, err = plugin_file.IsDirEmpty(p.Settings.Repo.WorkDir)
	if err != nil {
		return fmt.Errorf("failed to check working directory: %w", err)
	}

	isDir, err := plugin_file.IsDir(filepath.Join(p.Settings.Repo.WorkDir, ".git"))
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

	if err := ExecBatch(batchCmd); err != nil {
		return err
	}

	for _, actionStr := range p.Settings.Action {
		action := GitAction(actionStr)
		switch action {
		case GitActionClone:
			log.Debug().Msg("Compose action cmd: clone")

			if err := p.handleClone(); err != nil {
				return err
			}
		case GitActionCommit:
			log.Debug().Msg("Compose action cmd: commit")

			if err := p.handleCommit(); err != nil {
				return err
			}
		case GitActionPush:
			log.Debug().Msg("Compose action cmd: push")

			if err := p.handlePush(); err != nil {
				return err
			}
		case GitActionPages:
			log.Debug().Msg("Compose action cmd: pages")

			if err := p.handleClone(); err != nil {
				return err
			}

			if err := p.handlePages(); err != nil {
				return err
			}

			if err := p.handleCommit(); err != nil {
				return err
			}

			if err := p.handlePush(); err != nil {
				return err
			}
		}
	}

	return nil
}

// handleClone clones the remote repository into the configured working directory.
// If the working directory is not empty, it returns an error.
func (p *Plugin) handleClone() error {
	var batchCmd []*plugin_exec.Cmd

	if !p.Settings.Repo.IsEmpty {
		return fmt.Errorf("%w: %s exists and not empty", ErrGitCloneDestintionNotValid, p.Settings.Repo.WorkDir)
	}

	if p.Settings.Repo.RemoteURL != "" {
		batchCmd = append(batchCmd, p.Settings.Repo.RemoteAdd())
	}

	batchCmd = append(batchCmd, p.Settings.Repo.FetchSource())
	batchCmd = append(batchCmd, p.Settings.Repo.CheckoutHead())

	return ExecBatch(batchCmd)
}

// HandleCommit commits changes locally.
func (p *Plugin) handleCommit() error {
	if err := p.Settings.Repo.Add().Run(); err != nil {
		return err
	}

	if err := p.Settings.Repo.IsCleanTree().Run(); err == nil {
		if !p.Settings.Repo.EmptyCommit {
			log.Debug().Msg("Commit skipped: no changes")

			return nil
		}
	}

	return p.Settings.Repo.Commit().Run()
}

// HandlePush pushs changes to remote.
func (p *Plugin) handlePush() error {
	return p.Settings.Repo.RemotePush().Run()
}

// HandlePages syncs, commits and pushes the changes from the pages directory to the pages branch.
func (p *Plugin) handlePages() error {
	log.Debug().
		Str("src", p.Settings.Pages.Directory).
		Str("dest", p.Settings.Repo.WorkDir).
		Msg("handlePages")

	return SyncDirectories(
		p.Settings.Pages.Exclude,
		p.Settings.Pages.Delete,
		p.Settings.Pages.Directory,
		p.Settings.Repo.WorkDir,
		(zerolog.GlobalLevel() == zerolog.DebugLevel),
	).Run()
}
