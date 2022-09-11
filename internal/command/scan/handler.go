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

	// get list of tracked files in git repository.
	filepaths, err := h.Repository.ListFiles()

	projectZipName := fmt.Sprintf("%s.zip", repoMetadata.Name)
	err = h.Archiver.OutputZipToFile(repoMetadata.Path, projectZipName, filepaths)
	if err != nil {
		return err
	}

	// // pass the list of the tracked files and compress it into zip file.
	// projectZipBuf, err := h.Archiver.OutputZipToIOReader(repoMetadata.Path, filepaths)
	// if err != nil {
	// 	return err
	// }

	// projectZipName := fmt.Sprintf("%s.zip", repoMetadata.Name)
	// uploadURL, err := h.GRClient.CreateUploadURL(projectZipName)
	// if err != nil {
	// 	return err
	// }

	// err = h.GRClient.UploadProject(uploadURL, projectZipBuf)
	// if err != nil {
	// 	return err
	// }

	return nil
}
