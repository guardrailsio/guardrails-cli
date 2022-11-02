package scan

import (
	"context"
	"fmt"

	"github.com/briandowns/spinner"
	"github.com/cenkalti/backoff"
	"github.com/guardrailsio/guardrails-cli/internal/archiver"
	grclient "github.com/guardrailsio/guardrails-cli/internal/client/guardrails"
	"github.com/guardrailsio/guardrails-cli/internal/config"
	prettyFmt "github.com/guardrailsio/guardrails-cli/internal/format/pretty"
	outputwriter "github.com/guardrailsio/guardrails-cli/internal/output"
	"github.com/guardrailsio/guardrails-cli/internal/repository"
	"github.com/jedib0t/go-pretty/text"
)

var (
	ErrFailedToSaveOutput = func(err error) error {
		return fmt.Errorf("Couldn't save output, %s", err.Error())
	}
)

// Handler contains scan command dependencies.
type Handler struct {
	Args         *Args
	Spinner      *spinner.Spinner
	Config       *config.Config
	OutputWriter *outputwriter.OutputWriter
	Repository   repository.Repository
	Archiver     archiver.Archiver
	GRClient     grclient.GuardRailsClient
}

// New instantiates new scan command handler.
func New(
	args *Args,
	spinner *spinner.Spinner,
	config *config.Config,
	repo repository.Repository,
	arc archiver.Archiver,
	out *outputwriter.OutputWriter,
	grClient grclient.GuardRailsClient) *Handler {

	return &Handler{
		Args:         args,
		Spinner:      spinner,
		Config:       config,
		Repository:   repo,
		Archiver:     arc,
		OutputWriter: out,
		GRClient:     grClient,
	}
}

// Execute runs scan command.
func (h *Handler) Execute(ctx context.Context) error {
	w := h.OutputWriter.Writer

	h.displayScanningMessage()
	repoMetadata, err := h.Repository.GetMetadataFromRemoteURL()
	if err != nil {
		return err
	}
	h.stopLoadingMessage()

	if !h.Args.Quiet {
		fmt.Fprintf(w, "Project name: %s\nGit provider: %s\n", repoMetadata.Name, repoMetadata.Provider)
		if h.Args.Format == "" || h.Args.Format == FormatPretty {
			fmt.Fprintf(w, "Format: %s (default)\n", FormatPretty)
		} else {
			fmt.Fprintf(w, "Format: %s\n", h.Args.Format)
		}

		if h.Args.Output == "" {
			fmt.Fprintf(w, "Output: stdout\n")
		} else {
			fmt.Fprintf(w, "Output: %s\n", h.Args.Output)
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
	projectZipName := fmt.Sprintf("%s_%s.tar.gz", repoMetadata.Name, repoMetadata.CommitHash)
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

	if err := h.GetScanDataFormatter(getScanDataResp); err != nil {
		return err
	}

	if !h.Args.Quiet || h.Args.Format == FormatPretty {
		fmt.Fprintf(w, "\nView the detailed report in the dashboard\n%s\n", text.FgBlue.Sprint(getScanDataResp.Report))
	}

	if h.Args.Output != "" {
		if err := h.OutputWriter.SaveBufferToFile(); err != nil {
			return ErrFailedToSaveOutput(err)
		} else if !h.Args.Quiet {
			fmt.Printf("\n%s\n", prettyFmt.Success("Output saved"))
		}

		// make the buffer empty again.
		h.OutputWriter.Buffer.Reset()
	}

	return nil
}
