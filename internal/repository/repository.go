package repository

import (
	"errors"
	"fmt"
	"io"
	"regexp"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// Repository variable that does static check to make sure that repository struct implements Repository interface.
var _ Repository = (*repository)(nil)

//go:generate mockgen -destination=mock/repository.go -package=mockrepository . Repository

// Repository defines methods to interact with git repository.
type Repository interface {
	// GetMetadataFromRemoteURL extracts information from git repository remote URL and parse it to repository.Repository.
	// It only accepts 2 formats for now: "git@github.com:user/git-repo-name.git" and "https://github.com/user/git-repo-name.git".
	GetMetadataFromRemoteURL() (*Metadata, error)
	// ListFiles walking through git worktree and list of all files in repository whether is tracked or untracked.
	// This function respect rules defined in every .gitignore and also will exclude all files with type in ignoreList.
	ListFiles() ([]string, error)
}

// Repository contains details of git repository.
type repository struct {
	client   *git.Repository
	Metadata *Metadata
}

// Metadata contains repository metadata.
type Metadata struct {
	Path     string
	Protocol string
	Provider string
	Name     string
}

// New instantiates new repository.
func New(projectPath string) (Repository, error) {
	client, err := git.PlainOpen(projectPath)
	if err != nil {
		if errors.Is(err, git.ErrRepositoryNotExists) {
			return nil, fmt.Errorf("%s is not a valid git repository", projectPath)
		}

		return nil, err
	}

	return &repository{client: client, Metadata: &Metadata{Path: projectPath}}, nil
}

// GetMetadataFromRemoteURL implements repository.Repository interface.
func (r *repository) GetMetadataFromRemoteURL() (*Metadata, error) {
	cfg, err := r.client.Config()
	if err != nil {
		return nil, err
	}

	// TODO: currently we only take first remote URL from origin. It could be expanded later since git can have multiple remote urls.
	remoteURLs := cfg.Remotes["origin"].URLs
	if len(remoteURLs) == 0 {
		return nil, errors.New("repository doesn't have remote URLs")
	}

	remoteURL := remoteURLs[0]

	re := regexp.MustCompile(`(?P<Protocol>git@|http(s)?:\/\/)(.+@)*(?P<Provider>[\w\d\.]+)(:[\d]+){0,1}\/*(?P<Name>.*)`)
	matches := re.FindStringSubmatch(remoteURL)

	protocolRe := regexp.MustCompile(`[^\w]`)
	protocol := protocolRe.ReplaceAllString(matches[re.SubexpIndex("Protocol")], "")

	providerRe := regexp.MustCompile(`^(.*?)\.`)
	provider := providerRe.FindStringSubmatch(matches[re.SubexpIndex("Provider")])

	nameRe := regexp.MustCompile(`\/(.*)\.git$`)
	name := nameRe.FindStringSubmatch(matches[re.SubexpIndex("Name")])

	r.Metadata.Protocol = protocol
	r.Metadata.Provider = provider[1]
	r.Metadata.Name = name[1]

	return r.Metadata, nil
}

// ListFiles implements repository.Repository interface.
func (r *repository) ListFiles() ([]string, error) {
	filepaths := make([]string, 0)
	deletedFilepaths := make([]string, 0)

	// Unfortunately, the tree walker method only appends tracked file (file that already commited to git) which will leave untracked files left behind.
	// Thus we need to retrieve the untracked files via "git status". Git status has 2 status area : worktree and staging.
	// The difference lies whether the untracked files is already added to staging via "git add" or not. In short, we need to include untracked files in worktree area and
	// added files in staging area. Another status that affects the ListFiles operations is "deleted" because uncommited deleted file still list in worktree but in fact
	// has been deleted physically. Other statuses like modified or renamed won't be included.

	workTree, err := r.client.Worktree()
	if err != nil {
		return nil, err
	}
	workTreeStatus, err := workTree.Status()
	if err != nil {
		return nil, err
	}

	for name, status := range workTreeStatus {
		if status.Worktree == git.Untracked || status.Staging == git.Added {
			filepaths = append(filepaths, name)
		}
		if status.Worktree == git.Deleted || status.Staging == git.Deleted {
			deletedFilepaths = append(deletedFilepaths, name)
		}
	}

	// Retrieves all tracked files inside git repository by get the HEAD commit and walk the git worktree.
	ref, err := r.client.Head()
	if err != nil {
		return nil, err
	}

	commit, err := r.client.CommitObject(ref.Hash())
	t, err := commit.Tree()
	treeWalker := object.NewTreeWalker(t, true, nil)

TREEWALKER:
	for {
		name, _, err := treeWalker.Next()
		if err == io.EOF {
			break
		}

		// TODO: we could improve this to use slices.Contains that introduced in go 1.18 but for compatibility, let's do it in the old golang way
		for _, d := range deletedFilepaths {
			if d == name {
				continue TREEWALKER
			}
		}

		isIgnored := isFileTypeIgnored(name)
		if isIgnored {
			continue
		}

		filepaths = append(filepaths, name)
	}
	defer treeWalker.Close()

	return filepaths, nil
}
