package repository

import (
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// Repository variable that does static check to make sure that repository struct implements Repository interface.
var _ Repository = (*repository)(nil)

var (
	ErrNotAValidGitRepo = func(path string) error {
		return fmt.Errorf("%s is not a valid git repository", path)
	}
	ErrGitRemoteURLNotFound = errors.New("repository doesn't have remote URLs")
)

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
	DirPath     string
	Protocol    string
	Provider    string
	UserAccount string
	RemoteURL   string
	RepoName    string
	Branch      string
	CommitHash  string
}

// New instantiates new repository.
func New(projectPath string) (Repository, error) {
	client, err := git.PlainOpen(projectPath)
	if err != nil {
		if errors.Is(err, git.ErrRepositoryNotExists) {
			return nil, ErrNotAValidGitRepo(projectPath)
		}

		return nil, err
	}

	return &repository{client: client, Metadata: &Metadata{DirPath: projectPath}}, nil
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
		return nil, ErrGitRemoteURLNotFound
	}

	// retrieves all possible information from git remote url.
	getMetadataFromRemoteURL(r.Metadata, remoteURLs[0])

	ref, err := r.client.Head()
	if err != nil {
		return nil, err
	}

	branchRe := regexp.MustCompile(`([^\/]+$)`)
	branch := branchRe.FindString(ref.Name().String())

	r.Metadata.RemoteURL = remoteURLs[0]
	r.Metadata.Branch = branch
	r.Metadata.CommitHash = ref.Hash().String()

	return r.Metadata, nil
}

func getMetadataFromRemoteURL(metadata *Metadata, gitRemoteURL string) {
	re := regexp.MustCompile(`(?P<Protocol>git@|http(s)?:\/\/)(.+@)*(?P<Provider>[\w\d\.-]+)(:[\d]+){0,1}(\/scm)?\/*(?P<Name>.*)`)
	matches := re.FindStringSubmatch(gitRemoteURL)

	protocolRe := regexp.MustCompile(`[^\w]`)
	protocol := protocolRe.ReplaceAllString(matches[re.SubexpIndex("Protocol")], "")

	providerRe := regexp.MustCompile(`^([\w\d-]+)([\.\w\d]*)`)
	provider := providerRe.FindStringSubmatch(matches[re.SubexpIndex("Provider")])

	// use url parser to extract username and repository name by extracting the last 2 parts of the url
	urlParts := strings.Split(gitRemoteURL, "/")

	userAccount := urlParts[len(urlParts)-2]
	// remove the "git@provider.com:" from user account part if exists
	userAccountParts := strings.Split(userAccount, ":")
	if len(userAccountParts) > 1 {
		userAccount = userAccountParts[1]
	}

	repoName := strings.TrimSuffix(urlParts[len(urlParts)-1], ".git")

	metadata.Protocol = protocol
	metadata.Provider = provider[1]
	metadata.UserAccount = userAccount
	metadata.RepoName = repoName
}

// ListFiles implements repository.Repository interface.
func (r *repository) ListFiles() ([]string, error) {
	filepaths := make([]string, 0)

	// Retrieves all tracked files inside git repository by get the HEAD commit and walk the git worktree.
	ref, err := r.client.Head()
	if err != nil {
		return nil, err
	}

	commit, err := r.client.CommitObject(ref.Hash())
	if err != nil {
		return nil, err
	}
	t, err := commit.Tree()
	if err != nil {
		return nil, err
	}
	treeWalker := object.NewTreeWalker(t, true, nil)

	for {
		name, _, err := treeWalker.Next()
		if err == io.EOF {
			break
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
