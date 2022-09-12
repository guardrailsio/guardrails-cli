package guardrailsclient

import "io"

// CreateUploadURLReq is CreateUploadURL http request body.
type CreateUploadURLReq struct {
	CLIToken string `json:"clitoken"`
	File     string `json:"file"`
}

// CreateUploadURLResp is CreateUploadURL http response body.
type CreateUploadURLResp struct {
	SignedURL string `json:"signedUrl"`
}

// UploadProjectReq is UploadProject http request body.
type UploadProjectReq struct {
	UploadURL string
	File      io.Reader
}

// TriggerScanReq is TriggerScan http request body.
type TriggerScanReq struct {
	CLIToken   string `json:"clitoken"`
	Repository string `json:"repository"`
	SHA        string `json:"sha"`
	Branch     string `json:"branch"`
	FileName   string `json:"fileName"`
}

// TriggerScanResp is TriggerScan http response body.
type TriggerScanResp struct {
	ScanID       string `json:"idScan"`
	DashboardURL string `json:"dashboardUrl"`
}
