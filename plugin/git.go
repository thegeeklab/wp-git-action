package plugin

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/thegeeklab/drone-git-action/git"
)

// InitRepo initializes the repository.
func (p Plugin) initRepo() error {
	path := filepath.Join(p.settings.Path, ".git")
	if err := os.MkdirAll(p.settings.Path, os.ModePerm); err != nil {
		return err
	}

	if err := os.Chdir(p.settings.Path); err != nil {
		return err
	}

	if isDirEmpty(path) {
		return execute(exec.Command(
			"git",
			"init",
		))
	}

	return nil
}

// AddRemote adds a remote to repository.
func (p Plugin) addRemote() error {
	if p.settings.Remote != "" {
		if err := execute(git.RemoteAdd("origin", p.settings.Remote)); err != nil {
			return err
		}
	}

	return nil
}

// FetchSource fetches the source from remote.
func (p Plugin) fetchSource() error {
	return execute(exec.Command(
		"git",
		"fetch",
		"origin",
		fmt.Sprintf("+%s:", p.settings.Branch),
	))
}

// CheckoutHead handles branch checkout.
func (p Plugin) checkoutHead() error {
	return execute(exec.Command(
		"git",
		"checkout",
		"-qf",
		p.settings.Branch,
	))
}

// HandleClone clones remote.
func (p Plugin) handleClone() error {
	if err := p.addRemote(); err != nil {
		return err
	}

	if err := p.fetchSource(); err != nil {
		return err
	}

	if err := p.checkoutHead(); err != nil {
		return err
	}

	return nil
}

// HandleCommit commits changes locally.
func (p Plugin) handleCommit() error {
	if err := execute(git.Add()); err != nil {
		return err
	}

	if err := execute(git.TestCleanTree()); err != nil {
		if err := execute(git.ForceCommit(p.settings.Message, p.settings.NoVerify)); err != nil {
			return err
		}
	} else {
		if p.settings.EmptyCommit {
			if err := execute(git.EmptyCommit(p.settings.Message, p.settings.NoVerify)); err != nil {
				return err
			}
		}
	}

	return nil
}

// HandlePush pushs changes to remote.
func (p Plugin) handlePush() error {
	return execute(git.RemotePushNamedBranch(
		"origin",
		p.settings.Branch,
		p.settings.Branch,
		p.settings.Force,
		p.settings.FollowTags,
	))
}
