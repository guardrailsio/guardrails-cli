package repository

import (
	"errors"
	"regexp"

	"github.com/go-git/go-git/v5"
)

// Repository variable that does static check to make sure that repository struct implements Repository interface.
var _ Repository = (*repository)(nil)

//go:generate mockgen -destination=mock/repository.go -package=mockrepository . Repository

// Repository defines methods to interact with git repository.
type Repository interface {
	GetMetadataFromRemoteURL() (*Metadata, error)
}

func New(projectPath string) (Repository, error) {
	client, err := git.PlainOpen(projectPath)
	if err != nil {
		return nil, err
	}

	return &repository{client: client}, nil
}

// Repository contains details of git repository.
type repository struct {
	client   *git.Repository
	Metadata *Metadata
}

// Metadata contains repository metadata.
type Metadata struct {
	Protocol string
	Provider string
	Name     string
}

// GetMetadataFromRemoteURL extracts information from git repository remote URL and parse it to repository.Repository.
// We only accepts 2 formats for now:
// - git@github.com:user/git-repo-name.git
// - https://github.com/user/git-repo-name.git
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

	return &Metadata{
		Protocol: protocol,
		Provider: provider[1],
		Name:     name[1],
	}, nil
}