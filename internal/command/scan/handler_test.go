package scan

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	mockarchiver "github.com/guardrailsio/guardrails-cli/internal/archiver/mock"
	grclient "github.com/guardrailsio/guardrails-cli/internal/client/guardrails"
	mockgrclient "github.com/guardrailsio/guardrails-cli/internal/client/guardrails/mock"
	"github.com/guardrailsio/guardrails-cli/internal/config"
	outputwriter "github.com/guardrailsio/guardrails-cli/internal/output"
	"github.com/guardrailsio/guardrails-cli/internal/repository"
	mockrepository "github.com/guardrailsio/guardrails-cli/internal/repository/mock"
	"github.com/guardrailsio/guardrails-cli/internal/tools/spinner"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScanCommandExecuteSuccess(t *testing.T) {
	ctx := context.Background()
	args := &Args{
		Token:  "test-token",
		Path:   "/var/www/guardrails-cli",
		Format: "pretty",
		Output: "",
		Quiet:  true,
	}
	cfg := config.New()
	outputWriter := outputwriter.New(args.Output)
	spinner := spinner.New()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mockrepository.NewMockRepository(mockCtrl)
	mockArc := mockarchiver.NewMockArchiver(mockCtrl)
	mockGrClient := mockgrclient.NewMockGuardRailsClient(mockCtrl)

	commitMessage := gofakeit.Sentence(5)
	sha := sha1.New()
	sha.Write([]byte(commitMessage))
	encrypted := sha.Sum(nil)
	commitHash := fmt.Sprintf("%x", encrypted)

	repoMetadata := &repository.Metadata{
		Path:       args.Path,
		Protocol:   "https",
		Provider:   "github",
		Name:       gofakeit.AppName(),
		Branch:     gofakeit.Word(),
		CommitHash: commitHash,
	}
	listOfFiles := []string{
		".gitignore",
		"LICENSE",
		"Makefile",
		"README.md",
		"bin/.gitignore",
		"cmd/common.go",
		"cmd/root.go",
		"cmd/scan.go",
		"etc/scripts/run-test.sh",
		"go.mod",
		"go.sum",
		"internal/archiver/archiver.go",
		"internal/archiver/mock/archiver.go",
		"internal/client/client.go",
		"internal/client/errors.go",
		"internal/client/guardrails/client.go",
		"internal/client/guardrails/message.go",
		"internal/client/guardrails/mock/client.go",
		"internal/command/scan/args.go",
		"internal/command/scan/handler.go",
		"internal/command/scan/message.go",
		"internal/config/config.go",
		"internal/formatter/csv/csv.go",
		"internal/formatter/json/json.go",
		"internal/formatter/pretty/pretty.go",
		"internal/formatter/sarif/sarif.go",
		"internal/outputter/outputter.go",
		"internal/repository/ignorelist.go",
		"internal/repository/mock/repository.go",
		"internal/repository/repository.go",
		"internal/tools/spinner/spinner.go",
		"main.go",
	}
	fileZipName := fmt.Sprintf("%s_%s.tar.gz", repoMetadata.Name, repoMetadata.CommitHash)
	fileZipByte := bytes.NewReader([]byte{})
	createUploadURLReq := &grclient.CreateUploadURLReq{
		File: fileZipName,
	}
	createUploadURLResp := &grclient.CreateUploadURLResp{
		SignedURL: gofakeit.URL(),
	}
	uploadProjectReq := &grclient.UploadProjectReq{
		UploadURL: createUploadURLResp.SignedURL,
		File:      fileZipByte,
	}
	triggerScanReq := &grclient.TriggerScanReq{
		Repository: repoMetadata.Name,
		SHA:        repoMetadata.CommitHash,
		Branch:     repoMetadata.Branch,
		FileName:   fileZipName,
	}
	triggerScanResp := &grclient.TriggerScanResp{
		ScanID:       uuid.New().String(),
		DashboardURL: gofakeit.URL(),
	}
	getScanDataReq := &grclient.GetScanDataReq{
		ScanID: triggerScanResp.ScanID,
	}

	file, err := os.Open("../../client/guardrails/mock/get_scan_data.json")
	require.NoError(t, err)
	defer file.Close()

	getScanDataResp := new(grclient.GetScanDataResp)
	err = json.NewDecoder(file).Decode(getScanDataResp)
	require.NoError(t, err)

	gomock.InOrder(
		mockRepo.EXPECT().GetMetadataFromRemoteURL().Return(repoMetadata, nil),
		mockRepo.EXPECT().ListFiles().Return(listOfFiles, nil),
		mockArc.EXPECT().OutputZipToIOReader(repoMetadata.Path, listOfFiles).Return(fileZipByte, nil),
		mockGrClient.EXPECT().CreateUploadURL(ctx, createUploadURLReq).Return(createUploadURLResp, nil),
		mockGrClient.EXPECT().UploadProject(ctx, uploadProjectReq).Return(nil),
		mockGrClient.EXPECT().TriggerScan(ctx, triggerScanReq).Return(triggerScanResp, nil),
		mockGrClient.EXPECT().GetScanData(ctx, getScanDataReq).Return(getScanDataResp, nil),
	)

	cmd := New(args, spinner, cfg, mockRepo, mockArc, outputWriter, mockGrClient)
	err = cmd.Execute(ctx)
	assert.Nil(t, err)
}
