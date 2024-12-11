package plugin

import (
	"fmt"

	"github.com/thegeeklab/wp-git-action/git"
	plugin_base "github.com/thegeeklab/wp-plugin-go/v4/plugin"
	"github.com/urfave/cli/v2"
)

//go:generate go run ../internal/doc/main.go -output=../docs/data/data-raw.yaml

// Plugin implements provide the plugin.
type Plugin struct {
	*plugin_base.Plugin
	Settings *Settings
}

// Settings for the Plugin.
type Settings struct {
	Action cli.StringSlice
	SSHKey string

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
	Exclude   cli.StringSlice
	Delete    bool
}

type GitAction string

func New(e plugin_base.ExecuteFunc, build ...string) *Plugin {
	p := &Plugin{
		Settings: &Settings{},
	}

	options := plugin_base.Options{
		Name:                "wp-git-action",
		Description:         "Perform git actions",
		Flags:               Flags(p.Settings, plugin_base.FlagsPluginCategory),
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
		&cli.StringSliceFlag{
			Name:        "action",
			Usage:       "git action to execute",
			EnvVars:     []string{"PLUGIN_ACTION"},
			Destination: &settings.Action,
			Required:    true,
			Category:    category,
		},
		&cli.StringFlag{
			Name:        "author-name",
			Usage:       "git author name",
			EnvVars:     []string{"PLUGIN_AUTHOR_NAME", "CI_COMMIT_AUTHOR"},
			Destination: &settings.Repo.Author.Name,
			Required:    true,
			Category:    category,
		},
		&cli.StringFlag{
			Name:        "author-email",
			Usage:       "git author email",
			EnvVars:     []string{"PLUGIN_AUTHOR_EMAIL", "CI_COMMIT_AUTHOR_EMAIL"},
			Destination: &settings.Repo.Author.Email,
			Required:    true,
			Category:    category,
		},

		&cli.StringFlag{
			Name:        "netrc.machine",
			Usage:       "netrc remote machine name",
			EnvVars:     []string{"PLUGIN_NETRC_MACHINE", "CI_NETRC_MACHINE"},
			Destination: &settings.Netrc.Machine,
			Value:       "github.com",
			Category:    category,
		},
		&cli.StringFlag{
			Name:        "netrc.username",
			Usage:       "netrc login user on the remote machine",
			EnvVars:     []string{"PLUGIN_NETRC_USERNAME", "CI_NETRC_USERNAME"},
			Destination: &settings.Netrc.Login,
			Value:       "token",
			Category:    category,
		},
		&cli.StringFlag{
			Name:        "netrc.password",
			Usage:       "netrc login password on the remote machine",
			EnvVars:     []string{"PLUGIN_NETRC_PASSWORD", "CI_NETRC_PASSWORD"},
			Destination: &settings.Netrc.Password,
			Category:    category,
		},
		&cli.StringFlag{
			Name:        "ssh-key",
			Usage:       "ssh private key for the remote repository",
			EnvVars:     []string{"PLUGIN_SSH_KEY"},
			Destination: &settings.SSHKey,
			Category:    category,
		},

		&cli.StringFlag{
			Name:        "remote-url",
			Usage:       "url of the remote repository",
			EnvVars:     []string{"PLUGIN_REMOTE_URL", "CI_REPO_CLONE_URL"},
			Destination: &settings.Repo.RemoteURL,
			Category:    category,
		},
		&cli.StringFlag{
			Name:        "branch",
			Usage:       "name of the git source branch",
			EnvVars:     []string{"PLUGIN_BRANCH"},
			Destination: &settings.Repo.Branch,
			Value:       "main",
			Category:    category,
		},
		&cli.StringFlag{
			Name:        "path",
			Usage:       "path to clone git repository",
			EnvVars:     []string{"PLUGIN_PATH"},
			Destination: &settings.Repo.WorkDir,
			Category:    category,
		},
		&cli.StringFlag{
			Name:        "commit-message",
			Usage:       "commit message",
			EnvVars:     []string{"PLUGIN_MESSAGE"},
			Destination: &settings.Repo.CommitMsg,
			Value:       "[skip ci] commit dirty state",
			Category:    category,
		},
		&cli.BoolFlag{
			Name:        "force-push",
			Usage:       "enable force push to remote repository",
			EnvVars:     []string{"PLUGIN_FORCE"},
			Destination: &settings.Repo.ForcePush,
			Value:       false,
			Category:    category,
		},
		&cli.BoolFlag{
			Name:        "followtags",
			Usage:       "follow tags for pushes to remote repository",
			EnvVars:     []string{"PLUGIN_FOLLOWTAGS"},
			Destination: &settings.Repo.PushFollowTags,
			Value:       false,
			Category:    category,
		},
		&cli.BoolFlag{
			Name:        "empty-commit",
			Usage:       "allow empty commits",
			EnvVars:     []string{"PLUGIN_EMPTY_COMMIT"},
			Destination: &settings.Repo.EmptyCommit,
			Value:       false,
			Category:    category,
		},
		&cli.BoolFlag{
			Name:        "no-verify",
			Usage:       "bypass the pre-commit and commit-msg hooks",
			EnvVars:     []string{"PLUGIN_NO_VERIFY"},
			Destination: &settings.Repo.NoVerify,
			Value:       false,
			Category:    category,
		},
		&cli.StringFlag{
			Name:        "pages.directory",
			Usage:       "source directory to be synchronized with the pages banch",
			EnvVars:     []string{"PLUGIN_PAGES_DIRECTORY"},
			Destination: &settings.Pages.Directory,
			Value:       "docs/",
			Category:    category,
		},
		&cli.StringSliceFlag{
			Name:        "pages.exclude",
			Usage:       "files or directories to exclude from the pages rsync command",
			EnvVars:     []string{"PLUGIN_PAGES_EXCLUDE"},
			Destination: &settings.Pages.Exclude,
			Category:    category,
		},
		&cli.BoolFlag{
			Name:        "pages.delete",
			Usage:       "add delete flag to pages rsync command",
			EnvVars:     []string{"PLUGIN_PAGES_DELETE"},
			Destination: &settings.Pages.Delete,
			Value:       true,
			Category:    category,
		},
	}
}
