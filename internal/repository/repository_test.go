package repository

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestParseGitRemoteURLToMetadataSuccess(t *testing.T) {
	userAccount := gofakeit.Username()
	repoName := fmt.Sprintf("%s-%s", gofakeit.LoremIpsumWord(), gofakeit.LoremIpsumWord())

	urls := []string{
		fmt.Sprintf("git@github.com:%s/%s.git", userAccount, repoName),
		fmt.Sprintf("http://github.com/%s/%s.git", userAccount, repoName),
		fmt.Sprintf("https://github.com/%s/%s.git", userAccount, repoName),
		fmt.Sprintf("http://username:password@bitbucket-server:7990/scm/%s/%s.git", userAccount, repoName),
		fmt.Sprintf("https://gitlab-on-premise.com/%s/%s.git", userAccount, repoName),
		fmt.Sprintf("http://username:password@gitlab-ee/scm/%s/%s.git", userAccount, repoName),
		fmt.Sprintf("https://username:password@sub-domain-1.sub-domain-2.host.xz/prefix1/prefix2/prefix3/%s/%s.git", userAccount, repoName),
	}

	metadata := new(Metadata)

	for _, url := range urls {
		getMetadataFromRemoteURL(metadata, url)
		assert.NotEmpty(t, metadata.Protocol)
		assert.NotEmpty(t, metadata.Provider)
		assert.Equal(t, metadata.UserAccount, userAccount)
		assert.Equal(t, metadata.RepoName, repoName)
	}
}
