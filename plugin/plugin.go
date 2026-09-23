package plugin

import (
	"fmt"
	"slices"

	"github.com/thegeeklab/wp-git-action/git"
	plugin_base "github.com/thegeeklab/wp-plugin-go/v7/plugin"
	"github.com/urfave/cli/v3"
)

//go:generate go run ../hack/docs-gen/main.go -output=../docs/data/data.yaml

// Plugin implements provide the plugin.
type Plugin struct {
	*plugin_base.Plugin
	Settings *Settings
}

// Settings for the Plugin.
type Settings struct {
	Action []string
	SSHKey string

	CommitMessageFrom string

	Netrc Netrc
	Pages Pages
	Repo  git.Repository
}

type Netrc struct {
	Machine  string
	Login    string
	Password string
}

type Pages struct {
	Directory string
	Exclude   []string
	Delete    bool
}

type GitAction string

func New(e plugin_base.ExecuteFunc, build ...string) *Plugin {
	p := &Plugin{
		Settings: &Settings{},
	}

	options := plugin_base.Options{
		Name:        "wp-git-action",
		Description: "Perform git actions",
		Flags: slices.Concat(
			plugin_base.LoggingFlags(plugin_base.FlagsPluginCategory),
			plugin_base.NetworkFlags(plugin_base.FlagsPluginCategory),
			Flags(p.Settings, plugin_base.FlagsPluginCategory),
		),
		Execute:             p.run,
		HideWoodpeckerFlags: true,
	}

	if len(build) > 0 {
		options.Version = build[0]
	}

	if len(build) > 1 {
		options.VersionMetadata = fmt.Sprintf("date=%s", build[1])
	}

	if e != nil {
		options.Execute = e
	}

	p.Plugin = plugin_base.New(options)

	return p
}

// Flags returns a slice of CLI flags for the plugin.
func Flags(settings *Settings, category string) []cli.Flag {
	return []cli.Flag{
		// Git action to execute.
		//
		// Supported actions: `clone | commit | push | pages`. Specified actions are executed in the specified order
		//
		// - **clone:** Clones the repository in `remote_url` and checks out the `branch` to `path`.
		// - **commit:** Adds a commit to the default repository or the repository in `remote_url`.
		// - **push:** Pushes all commits to the default repository or the repository set in `remote_url`.
		// - **pages:** The `pages` action is a special action that cannot be combined with other actions. It is intended for
		//   use for GitHub pages. It synchronizes the contents of `pages_directory` with the target `branch` using `rsync`
		//   and pushes the changes automatically.
		&cli.StringSliceFlag{
			Name:        "action",
			Usage:       "git action to execute",
			Sources:     cli.EnvVars("PLUGIN_ACTION"),
			Destination: &settings.Action,
			Required:    true,
			Category:    category,
		},
		// Git author name.
		&cli.StringFlag{
			Name:        "author-name",
			Usage:       "git author name",
			Sources:     cli.EnvVars("PLUGIN_AUTHOR_NAME", "CI_COMMIT_AUTHOR"),
			Destination: &settings.Repo.Author.Name,
			Required:    true,
			Category:    category,
		},
		// Git author email.
		&cli.StringFlag{
			Name:        "author-email",
			Usage:       "git author email",
			Sources:     cli.EnvVars("PLUGIN_AUTHOR_EMAIL", "CI_COMMIT_AUTHOR_EMAIL"),
			Destination: &settings.Repo.Author.Email,
			Required:    true,
			Category:    category,
		},

		// Netrc remote machine name.
		&cli.StringFlag{
			Name:        "netrc.machine",
			Usage:       "netrc remote machine name",
			Sources:     cli.EnvVars("PLUGIN_NETRC_MACHINE", "CI_NETRC_MACHINE"),
			Destination: &settings.Netrc.Machine,
			Value:       "github.com",
			Category:    category,
		},
		// Netrc login user on the remote machine.
		&cli.StringFlag{
			Name:        "netrc.username",
			Usage:       "netrc login user on the remote machine",
			Sources:     cli.EnvVars("PLUGIN_NETRC_USERNAME", "CI_NETRC_USERNAME"),
			Destination: &settings.Netrc.Login,
			Value:       "token",
			Category:    category,
		},
		// Netrc login password on the remote machine.
		&cli.StringFlag{
			Name:        "netrc.password",
			Usage:       "netrc login password on the remote machine",
			Sources:     cli.EnvVars("PLUGIN_NETRC_PASSWORD", "CI_NETRC_PASSWORD"),
			Destination: &settings.Netrc.Password,
			Category:    category,
		},
		// Ssh private key for the remote repository.
		&cli.StringFlag{
			Name:        "ssh-key",
			Usage:       "ssh private key for the remote repository",
			Sources:     cli.EnvVars("PLUGIN_SSH_KEY"),
			Destination: &settings.SSHKey,
			Category:    category,
		},

		// Url of the remote repository.
		&cli.StringFlag{
			Name:        "remote-url",
			Usage:       "url of the remote repository",
			Sources:     cli.EnvVars("PLUGIN_REMOTE_URL", "CI_REPO_CLONE_URL"),
			Destination: &settings.Repo.RemoteURL,
			Category:    category,
		},
		// Name of the git source branch.
		&cli.StringFlag{
			Name:        "branch",
			Usage:       "name of the git source branch",
			Sources:     cli.EnvVars("PLUGIN_BRANCH"),
			Destination: &settings.Repo.Branch,
			Value:       "main",
			Category:    category,
		},
		// Path to clone git repository.
		&cli.StringFlag{
			Name:        "path",
			Usage:       "path to clone git repository",
			Sources:     cli.EnvVars("PLUGIN_PATH"),
			Destination: &settings.Repo.WorkDir,
			Category:    category,
		},
		// Delete the working directory after the git action.
		&cli.BoolFlag{
			Name:        "cleanup",
			Usage:       "delete the working directory after the git action",
			Sources:     cli.EnvVars("PLUGIN_CLEANUP"),
			Destination: &settings.Repo.Cleanup,
			Value:       true,
			Category:    category,
		},
		// Commit message.
		&cli.StringFlag{
			Name:        "commit-message",
			Usage:       "commit message",
			Sources:     cli.EnvVars("PLUGIN_MESSAGE"),
			Destination: &settings.Repo.CommitMsg,
			Value:       "[skip ci] commit dirty state",
			Category:    category,
		},
		// Read the commit message from the named environment variable.
		//
		// Useful for adopting the message of the triggering CI commit (e.g. `CI_COMMIT_MESSAGE` on Woodpecker)
		// without running into YAML substitution and escaping problems for values containing `:` or `"`.
		//
		// When set and the named variable resolves to a non-empty value, that value overrides `message`. If
		// the variable is unset or empty, `message` is used as a fallback (either an explicit value or the
		// static default).
		&cli.StringFlag{
			Name:        "commit-message-from",
			Usage:       "name of an environment variable to read the commit message from",
			Sources:     cli.EnvVars("PLUGIN_COMMIT_MESSAGE_FROM"),
			Destination: &settings.CommitMessageFrom,
			Category:    category,
		},
		// Enable force push to remote repository.
		&cli.BoolFlag{
			Name:        "force-push",
			Usage:       "enable force push to remote repository",
			Sources:     cli.EnvVars("PLUGIN_FORCE"),
			Destination: &settings.Repo.ForcePush,
			Value:       false,
			Category:    category,
		},
		// Follow tags for pushes to remote repository.
		//
		// Push all the `refs` that would be pushed without this option, and also push annotated tags
		// in `refs/tags` that are missing from the remote.
		&cli.BoolFlag{
			Name:        "followtags",
			Usage:       "follow tags for pushes to remote repository",
			Sources:     cli.EnvVars("PLUGIN_FOLLOWTAGS"),
			Destination: &settings.Repo.PushFollowTags,
			Value:       false,
			Category:    category,
		},
		// Allow empty commits.
		//
		// Usually recording a commit that has the exact same tree as its sole parent commit is a mistake,
		// and those commits are not allowed by default.
		&cli.BoolFlag{
			Name:        "empty-commit",
			Usage:       "allow empty commits",
			Sources:     cli.EnvVars("PLUGIN_EMPTY_COMMIT"),
			Destination: &settings.Repo.EmptyCommit,
			Value:       false,
			Category:    category,
		},
		// Bypass the pre-commit and commit-msg hooks.
		&cli.BoolFlag{
			Name:        "no-verify",
			Usage:       "bypass the pre-commit and commit-msg hooks",
			Sources:     cli.EnvVars("PLUGIN_NO_VERIFY"),
			Destination: &settings.Repo.NoVerify,
			Value:       false,
			Category:    category,
		},
		// Source directory to be synchronized with the pages branch.
		&cli.StringFlag{
			Name:        "pages.directory",
			Usage:       "source directory to be synchronized with the pages banch",
			Sources:     cli.EnvVars("PLUGIN_PAGES_DIRECTORY"),
			Destination: &settings.Pages.Directory,
			Value:       "docs/",
			Category:    category,
		},
		// Files or directories to exclude from the pages rsync command.
		&cli.StringSliceFlag{
			Name:        "pages.exclude",
			Usage:       "files or directories to exclude from the pages rsync command",
			Sources:     cli.EnvVars("PLUGIN_PAGES_EXCLUDE"),
			Destination: &settings.Pages.Exclude,
			Category:    category,
		},
		// Add delete flag to pages rsync command.
		//
		// When set to `true`, the `--delete` flag is added to the rsync command to remove files
		// from the branch that do not exist in the `pages_directory` either.
		&cli.BoolFlag{
			Name:        "pages.delete",
			Usage:       "add delete flag to pages rsync command",
			Sources:     cli.EnvVars("PLUGIN_PAGES_DELETE"),
			Destination: &settings.Pages.Delete,
			Value:       true,
			Category:    category,
		},
	}
}
