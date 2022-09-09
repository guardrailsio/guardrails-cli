package scan

import (
	"fmt"

	"github.com/guardrailsio/guardrails-cli/internal/repository"
	"github.com/jedib0t/go-pretty/text"
)

// Handler contains scan command dependencies.
type Handler struct {
	Args       *Args
	Repository repository.Repository
}

// New instantiates new scan command handler.
func New(args *Args, repo repository.Repository) *Handler {
	return &Handler{Args: args, Repository: repo}
}

// Execute runs scan command.
func (h *Handler) Execute() error {
	if err := h.Args.Validate(); err != nil {
		return err
	}

	fmt.Println(text.FgCyan.Sprintf("scanning %s ...\n", h.Args.Path))

	repoMetadata, err := h.Repository.GetMetadataFromRemoteURL()
	if err != nil {
		return err
	}

	fmt.Printf("Project name: %s\nGit provider: %s\n", repoMetadata.Name, repoMetadata.Provider)

	return nil
}
