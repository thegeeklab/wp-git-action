package plugin

import (
	"github.com/thegeeklab/wp-git-action/git"
	wp "github.com/thegeeklab/wp-plugin-go/v2/plugin"
	"github.com/urfave/cli/v2"
)

// Plugin implements provide the plugin.
type Plugin struct {
	*wp.Plugin
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

func New(options wp.Options, settings *Settings) *Plugin {
	p := &Plugin{}

	options.Execute = p.run

	p.Plugin = wp.New(options)
	p.Settings = settings

	return p
}
