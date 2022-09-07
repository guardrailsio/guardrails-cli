package scan

import (
	"errors"
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/guardrailsio/guardrails-cli/internal/project"
	"github.com/jedib0t/go-pretty/text"
)

// Handler contains scan command dependencies.
type Handler struct {
	Args    *Args
	GitRepo *git.Repository
}

// New instantiates new scan command handler.
func New(args *Args, gitRepo *git.Repository) *Handler {
	return &Handler{Args: args, GitRepo: gitRepo}
}

// Execute runs scan command.
func (h *Handler) Execute() error {
	if err := h.Args.Validate(); err != nil {
		return err
	}

	fmt.Println(text.FgCyan.Sprintf("scanning %s ...\n", h.Args.Path))

	cfg, err := h.GitRepo.Config()
	if err != nil {
		return err
	}

	// TODO: currently we only take first remote URL from origin. It could be expanded later since git can have multiple remote urls.
	remoteURLs := cfg.Remotes["origin"].URLs
	if len(remoteURLs) == 0 {
		return errors.New("repository doesn't have remote URLs")
	}

	repo := project.GetProjectFromRemoteURL(remoteURLs[0])

	fmt.Printf("Project name: %s\nGit provider: %s\n", repo.Name, repo.Provider)

	return nil
}
