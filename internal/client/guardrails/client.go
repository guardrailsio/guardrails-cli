package guardrailsclient

import (
	"bytes"
	"encoding/json"
	"net/http"

	httpClient "github.com/guardrailsio/guardrails-cli/internal/client"
	"github.com/guardrailsio/guardrails-cli/internal/config"
)

// GuardRailsClient variable that does static check to make sure that client struct implements GuardRailsClient interface.
var _ GuardRailsClient = (*client)(nil)

//go:generate mockgen -destination=mock/client.go -package=mockguardrailsclient . GuardRailsClient

// GuardRailsClient defines methods to interact with GuardRails API.
type GuardRailsClient interface {
	CreateUploadURL(projectFilename string) (string, error)
	UploadProject(uploadURL string) error
	TriggerScan() error
	GetScanData() error
}

type client struct {
	cfg        *config.HttpClientConfig
	httpclient *http.Client
	token      string
}

// New instantiates new GuardRailsClient.
func New(cfg *config.HttpClientConfig, token string) GuardRailsClient {
	c := httpClient.New(cfg)

	return &client{cfg: cfg, httpclient: c, token: token}
}

// CreateUploadURL call GuardRails api to create upload URL.
func (c *client) CreateUploadURL(projectFilename string) (string, error) {
	url := "https://api.guardrails.io/v2/cli/trigger-zip-scan-upload-url"
	contentType := "application/json"
	req := &CreateUploadURLReq{
		CLIToken: c.token,
		File:     projectFilename,
	}
	reqBody, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	resp, err := c.httpclient.Post(url, contentType, bytes.NewReader(reqBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody := new(CreateUploadURLResp)
	if err := json.NewDecoder(resp.Body).Decode(respBody); err != nil {
		return "", err
	}

	return respBody.SignedURL, nil
}

// UploadProject accepts url generated from CreateUploadURL method, compress the project, and upload it to designated url.
func (c *client) UploadProject(uploadURL string) error {
	return nil
}

// TriggerScan call GuardRails api to trigger scan operation.
func (c *client) TriggerScan() error {
	return nil
}

// GetScanData call GuardRails api to get scan data from scan operation.
func (c *client) GetScanData() error {
	return nil
}
