package guardrailsclient

// CreateUploadURLReq is CreateUploadURL http request body.
type CreateUploadURLReq struct {
	CLIToken string `json:"clitoken"`
	File     string `json:"file"`
}

// CreateUploadURLResp is CreateUploadURL http response body.
type CreateUploadURLResp struct {
	SignedURL string `json:"signedUrl"`
}

// TriggerScanReq is TriggerScan http request body.
type TriggerScanReq struct {
	CLIToken   string `json:"clitoken"`
	Repository string `json:"repository"`
	SHA        string `json:"sha"`
	Branch     string `json:"branch"`
	FileName   string `json:"fileName"`
}
