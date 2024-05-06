package git

type Author struct {
	Name  string
	Email string
}

type Repository struct {
	RemoteName string
	RemoteURL  string
	Branch     string

	CommitMsg string

	Autocorrect    string
	NoVerify       bool
	EmptyCommit    bool
	PushFollowTags bool
	ForcePush      bool
	WorkDir        string
	IsEmpty        bool

	Author Author
}

const gitBin = "/usr/bin/git"
