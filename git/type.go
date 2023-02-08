package git

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

	Autocorrect       string
	NoVerify          bool
	InsecureSSLVerify bool
	EmptyCommit       bool
	PushFollowTags    bool
	ForcePush         bool
	SSLVerify         bool
	WorkDir           string
	InitExists        bool

	Author Author
}

const gitBin = "/usr/bin/git"
