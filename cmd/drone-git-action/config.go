package main

import (
	"github.com/thegeeklab/drone-git-action/plugin"
	"github.com/urfave/cli/v2"
)

// settingsFlags has the cli.Flags for the plugin.Settings.
func settingsFlags(settings *plugin.Settings, category string) []cli.Flag {
	return []cli.Flag{
		&cli.StringSliceFlag{
			Name:        "actions",
			Usage:       "git actions to to execute",
			EnvVars:     []string{"PLUGIN_ACTIONS"},
			Destination: &settings.Actions,
			Required:    true,
			Category:    category,
		},

		&cli.StringFlag{
			Name:        "commit.author.name",
			Usage:       "git author name",
			EnvVars:     []string{"PLUGIN_AUTHOR_NAME", "DRONE_COMMIT_AUTHOR"},
			Destination: &settings.Commit.Author.Name,
			Required:    true,
			Category:    category,
		},
		&cli.StringFlag{
			Name:        "commit.author.email",
			Usage:       "git author email",
			EnvVars:     []string{"PLUGIN_AUTHOR_EMAIL", "DRONE_COMMIT_AUTHOR_EMAIL"},
			Destination: &settings.Commit.Author.Email,
			Required:    true,
			Category:    category,
		},

		&cli.StringFlag{
			Name:        "netrc.machine",
			Usage:       "netrc remote machine name",
			EnvVars:     []string{"PLUGIN_NETRC_MACHINE", "DRONE_NETRC_MACHINE"},
			Destination: &settings.Netrc.Machine,
			Category:    category,
		},
		&cli.StringFlag{
			Name:        "netrc.username",
			Usage:       "netrc login user on the remote machine",
			EnvVars:     []string{"PLUGIN_NETRC_USERNAME", "DRONE_NETRC_USERNAME"},
			Destination: &settings.Netrc.Login,
			Category:    category,
		},
		&cli.StringFlag{
			Name:        "netrc.password",
			Usage:       "netrc login password on the remote machine",
			EnvVars:     []string{"PLUGIN_NETRC_PASSWORD", "DRONE_NETRC_PASSWORD"},
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
			Name:        "remote",
			Usage:       "url of the remote repository",
			EnvVars:     []string{"PLUGIN_REMOTE"},
			Destination: &settings.Remote,
			Category:    category,
		},
		&cli.StringFlag{
			Name:        "branch",
			Usage:       "name of the git branch",
			EnvVars:     []string{"PLUGIN_BRANCH"},
			Destination: &settings.Branch,
			Value:       "main",
			Category:    category,
		},

		&cli.StringFlag{
			Name:        "path",
			Usage:       "path to git repository",
			EnvVars:     []string{"PLUGIN_PATH"},
			Destination: &settings.Path,
			Category:    category,
		},

		&cli.StringFlag{
			Name:        "message",
			Usage:       "commit message",
			EnvVars:     []string{"PLUGIN_MESSAGE"},
			Destination: &settings.Message,
			Value:       "[skip ci] Commit dirty state",
			Category:    category,
		},

		&cli.BoolFlag{
			Name:        "force",
			Usage:       "enable force push to remote repository",
			EnvVars:     []string{"PLUGIN_FORCE"},
			Destination: &settings.Force,
			Category:    category,
		},
		&cli.BoolFlag{
			Name:        "followtags",
			Usage:       "follow tags for pushes to remote repository",
			EnvVars:     []string{"PLUGIN_FOLLOWTAGS"},
			Destination: &settings.FollowTags,
			Category:    category,
		},
		&cli.BoolFlag{
			Name:        "skip-verify",
			Usage:       "skip ssl verification of the remote machine",
			EnvVars:     []string{"PLUGIN_SKIP_VERIFY"},
			Destination: &settings.SkipVerify,
			Category:    category,
		},
		&cli.BoolFlag{
			Name:        "empty-commit",
			Usage:       "allow empty commits",
			EnvVars:     []string{"PLUGIN_EMPTY_COMMIT"},
			Destination: &settings.EmptyCommit,
			Category:    category,
		},
		&cli.BoolFlag{
			Name:        "no-verify",
			Usage:       "bypass the pre-commit and commit-msg hooks",
			EnvVars:     []string{"PLUGIN_NO_VERIFY"},
			Destination: &settings.NoVerify,
			Category:    category,
		},
	}
}
