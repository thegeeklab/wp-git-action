package git

const gitBin = "/usr/bin/git"

type Author struct {
	Name  string
	Email string
}

type Repository struct {
	RemoteName string
	RemoteURL  string
	Branch     string

	Add       string
	CommitMsg string

	Autocorrect           string
	NoVerify              bool
	InsecureSkipSSLVerify bool
	EmptyCommit           bool
	PushFollowTags        bool
	ForcePush             bool
	WorkDir               string
	InitExists            bool

	Author Author
}
