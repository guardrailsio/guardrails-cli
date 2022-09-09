package scan

import (
	"fmt"

	"github.com/guardrailsio/guardrails-cli/internal/archiver"
	guardrailsclient "github.com/guardrailsio/guardrails-cli/internal/client/guardrails"
	"github.com/guardrailsio/guardrails-cli/internal/repository"
	"github.com/jedib0t/go-pretty/text"
)

// Handler contains scan command dependencies.
type Handler struct {
	Args       *Args
	Repository repository.Repository
	Archiver   archiver.Archiver
	GRClient   guardrailsclient.GuardRailsClient
}

// New instantiates new scan command handler.
func New(
	args *Args,
	repo repository.Repository,
	arc archiver.Archiver,
	grclient guardrailsclient.GuardRailsClient) *Handler {

	return &Handler{Args: args, Repository: repo, Archiver: arc, GRClient: grclient}
}

// Execute runs scan command.
func (h *Handler) Execute() error {
	fmt.Println(text.FgCyan.Sprintf("scanning %s ...\n", h.Args.Path))

	repoMetadata, err := h.Repository.GetMetadataFromRemoteURL()
	if err != nil {
		return err
	}

	fmt.Printf("Project name: %s\nGit provider: %s\n", repoMetadata.Name, repoMetadata.Provider)

	// create and compress git project
	projectZipFilename := fmt.Sprintf("%s.zip", repoMetadata.Name)
	zipFile, err := h.Archiver.OutputZipToReader(repoMetadata.Path)
	if err != nil {
		return err
	}

	uploadURL, err := h.GRClient.CreateUploadURL(repoMetadata.Name)
	if err != nil {
		return err
	}

	return nil
}
