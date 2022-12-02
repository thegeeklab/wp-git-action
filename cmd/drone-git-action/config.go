package main

import (
	"github.com/thegeeklab/drone-git-action/plugin"
	"github.com/urfave/cli/v2"
)

// settingsFlags has the cli.Flags for the plugin.Settings.
func settingsFlags(settings *plugin.Settings, category string) []cli.Flag {
	return []cli.Flag{
		&cli.StringSliceFlag{
			Name:        "action",
			Usage:       "git action to to execute",
			EnvVars:     []string{"PLUGIN_ACTION"},
			Destination: &settings.Action,
			Required:    true,
			Category:    category,
		},
		&cli.StringFlag{
			Name:        "author-name",
			Usage:       "git author name",
			EnvVars:     []string{"PLUGIN_AUTHOR_NAME", "DRONE_COMMIT_AUTHOR"},
			Destination: &settings.Repo.Author.Name,
			Required:    true,
			Category:    category,
		},
		&cli.StringFlag{
			Name:        "author-email",
			Usage:       "git author email",
			EnvVars:     []string{"PLUGIN_AUTHOR_EMAIL", "DRONE_COMMIT_AUTHOR_EMAIL"},
			Destination: &settings.Repo.Author.Email,
			Required:    true,
			Category:    category,
		},

		&cli.StringFlag{
			Name:        "netrc.machine",
			Usage:       "netrc remote machine name",
			EnvVars:     []string{"PLUGIN_NETRC_MACHINE", "DRONE_NETRC_MACHINE"},
			Destination: &settings.Netrc.Machine,
			Value:       "github.com",
			Category:    category,
		},
		&cli.StringFlag{
			Name:        "netrc.username",
			Usage:       "netrc login user on the remote machine",
			EnvVars:     []string{"PLUGIN_NETRC_USERNAME", "DRONE_NETRC_USERNAME"},
			Destination: &settings.Netrc.Login,
			Value:       "token",
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
			Name:        "remote-url",
			Usage:       "url of the remote repository",
			EnvVars:     []string{"PLUGIN_REMOTE_URL", "DRONE_REMOTE_URL"},
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
			Name:        "insecure-ssl-verify",
			Usage:       "set SSL verification of the remote machine",
			EnvVars:     []string{"PLUGIN_INSECURE_SSL_VERIFY"},
			Destination: &settings.Repo.InsecureSSLVerify,
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
			Usage:       "source directory for pages sync",
			EnvVars:     []string{"PLUGIN_PAGES_DIRECTORY"},
			Destination: &settings.Pages.Directory,
			Value:       "docs/",
			Category:    category,
		},
		&cli.StringSliceFlag{
			Name:        "pages.exclude",
			Usage:       "exclude flag added to pages rsnyc command",
			EnvVars:     []string{"PLUGIN_PAGES_EXCLUDE"},
			Destination: &settings.Pages.Exclude,
			Category:    category,
		},
		&cli.BoolFlag{
			Name:        "pages.delete",
			Usage:       "delete flag added to pages rsync command",
			EnvVars:     []string{"PLUGIN_PAGES_DELETE"},
			Destination: &settings.Pages.Delete,
			Value:       true,
			Category:    category,
		},
	}
}
