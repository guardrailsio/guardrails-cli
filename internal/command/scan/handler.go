package scan

import (
	"context"
	"fmt"

	"github.com/briandowns/spinner"
	"github.com/cenkalti/backoff"
	"github.com/guardrailsio/guardrails-cli/internal/archiver"
	grclient "github.com/guardrailsio/guardrails-cli/internal/client/guardrails"
	"github.com/guardrailsio/guardrails-cli/internal/config"
	prettyFmt "github.com/guardrailsio/guardrails-cli/internal/formatter/pretty"
	"github.com/guardrailsio/guardrails-cli/internal/repository"
	"github.com/jedib0t/go-pretty/text"
)

// Handler contains scan command dependencies.
type Handler struct {
	Args       *Args
	Spinner    *spinner.Spinner
	Config     *config.Config
	Repository repository.Repository
	Archiver   archiver.Archiver
	GRClient   grclient.GuardRailsClient
}

// New instantiates new scan command handler.
func New(
	args *Args,
	spinner *spinner.Spinner,
	config *config.Config,
	repo repository.Repository,
	arc archiver.Archiver,
	grClient grclient.GuardRailsClient) *Handler {

	return &Handler{Args: args, Spinner: spinner, Config: config, Repository: repo, Archiver: arc, GRClient: grClient}
}

// Execute runs scan command.
func (h *Handler) Execute(ctx context.Context) error {
	h.displayScanningMessage()
	repoMetadata, err := h.Repository.GetMetadataFromRemoteURL()
	if err != nil {
		return err
	}
	h.stopLoadingMessage()

	if !h.Args.Quiet {
		fmt.Printf("Project name: %s\nGit provider: %s\n", repoMetadata.Name, repoMetadata.Provider)
		if h.Args.Format == "" || h.Args.Format == "pretty" {
			fmt.Printf("Format: pretty (default)\n")
		} else {
			fmt.Printf("Format: %s\n", h.Args.Format)
		}

		if h.Args.Output == "" {
			fmt.Printf("Output: none\n")
		} else {
			fmt.Printf("Output: %s\n", h.Args.Output)
		}
	}

	// get list of tracked files in git repository.
	filepaths, err := h.Repository.ListFiles()
	if err != nil {
		return err
	}

	// pass the list of the tracked files and compress it into zip file.
	h.displayCompressingMessage(repoMetadata.Name)
	projectZipBuf, err := h.Archiver.OutputZipToIOReader(repoMetadata.Path, filepaths)
	if err != nil {
		return err
	}
	h.stopLoadingMessage()

	// create presigned url for uploading the compressed file
	projectZipName := fmt.Sprintf("%s.zip", repoMetadata.Name)
	createUploadURLReq := &grclient.CreateUploadURLReq{
		File: projectZipName,
	}
	createUploadURLResp, err := h.GRClient.CreateUploadURL(ctx, createUploadURLReq)
	if err != nil {
		return err
	}

	// upload the compressed project files
	h.displayUploadingMessage(projectZipName)
	uploadProjectReq := &grclient.UploadProjectReq{
		UploadURL: createUploadURLResp.SignedURL,
		File:      projectZipBuf,
	}
	err = h.GRClient.UploadProject(ctx, uploadProjectReq)
	if err != nil {
		return err
	}
	h.stopLoadingMessage()

	// call GuardRails trigger scan API
	triggerScanReq := &grclient.TriggerScanReq{
		Repository: repoMetadata.Name,
		SHA:        repoMetadata.CommitHash,
		Branch:     repoMetadata.Branch,
		FileName:   projectZipName,
	}
	triggerScanResp, err := h.GRClient.TriggerScan(ctx, triggerScanReq)
	if err != nil {
		return err
	}

	h.displayRetrievingScanResultMessage(repoMetadata.Name)
	bo := backoff.NewConstantBackOff(h.Config.HttpClient.PollingInterval)
	backoffCtx, backoffCtxCancel := context.WithTimeout(ctx, h.Config.HttpClient.RetryTimeout)
	defer backoffCtxCancel()

	bc := backoff.WithContext(bo, backoffCtx)

	getScanDataReq := &grclient.GetScanDataReq{
		ScanID: triggerScanResp.ScanID,
	}

	var getScanDataResp *grclient.GetScanDataResp
	err = backoff.Retry(func() error {
		if getScanDataResp, err = h.GRClient.GetScanData(ctx, getScanDataReq); err != nil {
			return err
		}

		return nil
	}, bc)
	if err != nil {
		return err
	}
	h.stopLoadingMessage()

	if !h.Args.Quiet {
		switch h.Args.Format {
		default:
			if getScanDataResp.TotalVulnerabilities == 0 && getScanDataResp.NewVulnerabilities == 0 {
				fmt.Printf("\n%s", prettyFmt.Success("No issues detected, well done!"))
			}
		}

		fmt.Printf("\n\nView the detailed report in the dashboard\n%s", text.FgBlue.Sprint(triggerScanResp.DashboardURL))
	}

	return nil
}
