package scan

import (
	"context"
	"fmt"

	"github.com/guardrailsio/guardrails-cli/internal/archiver"
	grclient "github.com/guardrailsio/guardrails-cli/internal/client/guardrails"
	"github.com/guardrailsio/guardrails-cli/internal/repository"
	"github.com/jedib0t/go-pretty/text"
)

// Handler contains scan command dependencies.
type Handler struct {
	Args       *Args
	Repository repository.Repository
	Archiver   archiver.Archiver
	GRClient   grclient.GuardRailsClient
}

// New instantiates new scan command handler.
func New(
	args *Args,
	repo repository.Repository,
	arc archiver.Archiver,
	grClient grclient.GuardRailsClient) *Handler {

	return &Handler{Args: args, Repository: repo, Archiver: arc, GRClient: grClient}
}

// Execute runs scan command.
func (h *Handler) Execute(ctx context.Context) error {
	fmt.Println(text.FgCyan.Sprintf("scanning %s ...\n", h.Args.Path))

	repoMetadata, err := h.Repository.GetMetadataFromRemoteURL()
	if err != nil {
		return err
	}

	fmt.Printf("Project name: %s\nGit provider: %s\n", repoMetadata.Name, repoMetadata.Provider)

	// get list of tracked files in git repository.
	filepaths, err := h.Repository.ListFiles()
	if err != nil {
		return err
	}

	// pass the list of the tracked files and compress it into zip file.
	projectZipName := fmt.Sprintf("%s.zip", repoMetadata.Name)
	projectZipBuf, err := h.Archiver.OutputZipToIOReader(repoMetadata.Path, filepaths)
	if err != nil {
		return err
	}

	createUploadURLReq := &grclient.CreateUploadURLReq{
		File: projectZipName,
	}
	createUploadURLResp, err := h.GRClient.CreateUploadURL(ctx, createUploadURLReq)
	if err != nil {
		return err
	}

	uploadProjectReq := &grclient.UploadProjectReq{
		UploadURL: createUploadURLResp.SignedURL,
		File:      projectZipBuf,
	}
	err = h.GRClient.UploadProject(ctx, uploadProjectReq)
	if err != nil {
		return err
	}

	triggerScanReq := &grclient.TriggerScanReq{
		Repository: repoMetadata.Name,
		SHA:        repoMetadata.CommitHash,
		Branch:     repoMetadata.Branch,
		FileName:   projectZipName,
	}
	_, err = h.GRClient.TriggerScan(ctx, triggerScanReq)
	if err != nil {
		return err
	}

	return nil
}
