package guardrailsclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	httpClient "github.com/guardrailsio/guardrails-cli/internal/client"
	"github.com/guardrailsio/guardrails-cli/internal/config"
)

// GuardRailsClient variable that does static check to make sure that client struct implements GuardRailsClient interface.
var _ GuardRailsClient = (*client)(nil)

//go:generate mockgen -destination=mock/client.go -package=mockguardrailsclient . GuardRailsClient

// GuardRailsClient defines methods to interact with GuardRails API.
type GuardRailsClient interface {
	// CreateUploadURL call GuardRails API to create upload URL.
	CreateUploadURL(projectFilename string) (string, error)
	// UploadProject accepts url generated from CreateUploadURL method, compress the project, and upload it to designated url.
	UploadProject(uploadURL string, file io.Reader) error
	// TriggerScan call GuardRails API to trigger scan operation.
	TriggerScan() error
	// GetScanData call GuardRails API to get scan data from scan operation.
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

// CreateUploadURL implements guardrailsclient.GuardRailsClient interface.
func (c *client) CreateUploadURL(projectFilename string) (string, error) {
	url := "https://API.guardrails.io/v2/cli/trigger-zip-scan-upload-url"
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

// UploadProject implements guardrailsclient.GuardRailsClient interface.
func (c *client) UploadProject(uploadURL string, file io.Reader) error {
	fmt.Println("UploadURL: ", uploadURL)

	return nil
}

// TriggerScan implements guardrailsclient.GuardRailsClient interface.
func (c *client) TriggerScan() error {
	return nil
}

// GetScanData implements guardrailsclient.GuardRailsClient interface.
func (c *client) GetScanData() error {
	return nil
}
