package plugin

import (
	"os/user"
)

func getUserHomeDir() string {
	home := "/root"

	if currentUser, err := user.Current(); err == nil {
		home = currentUser.HomeDir
	}

	return home
}
