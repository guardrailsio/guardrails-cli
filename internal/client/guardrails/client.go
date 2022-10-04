package guardrailsclient

import (
	"bytes"
	"context"
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
	// CreateUploadURL call GuardRails API to create upload URL.
	CreateUploadURL(ctx context.Context, req *CreateUploadURLReq) (*CreateUploadURLResp, error)
	// UploadProject accepts url generated from CreateUploadURL and upload it via presigned url.
	UploadProject(ctx context.Context, req *UploadProjectReq) error
	// TriggerScan call GuardRails API to trigger scan operation.
	TriggerScan(ctx context.Context, req *TriggerScanReq) (*TriggerScanResp, error)
	// GetScanData call GuardRails API to get scan data from scan operation.
	GetScanData(ctx context.Context, req *GetScanDataReq) (*GetScanDataResp, error)
}

type client struct {
	cfg        *config.HttpClientConfig
	httpclient *http.Client
	token      string
}

// New instantiates new GuardRailsClient.
func New(cfg *config.HttpClientConfig, token string) GuardRailsClient {
	return &client{cfg: cfg, httpclient: new(http.Client), token: token}
}

// CreateUploadURL implements guardrailsclient.GuardRailsClient interface.
func (c *client) CreateUploadURL(ctx context.Context, req *CreateUploadURLReq) (*CreateUploadURLResp, error) {
	url := "https://API.guardrails.io/v2/cli/trigger-zip-scan-upload-url"

	req.CLIToken = c.token
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, c.cfg.Timeout)
	defer cancel()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.httpclient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := parseHTTPRespStatusCode("CreateUploadURL", resp); err != nil {
		return nil, err
	}

	respBody := new(CreateUploadURLResp)
	if err := json.NewDecoder(resp.Body).Decode(respBody); err != nil {
		return nil, err
	}

	return respBody, nil
}

// UploadProject implements guardrailsclient.GuardRailsClient interface.
func (c *client) UploadProject(ctx context.Context, req *UploadProjectReq) error {
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPut, req.UploadURL, req.File)
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", "application/zip")

	resp, err := c.httpclient.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return httpClient.UnexpectedHTTPResponseFormatter("UploadProject", resp.StatusCode, resp.Body)
	}

	return nil
}

// TriggerScan implements guardrailsclient.GuardRailsClient interface.
func (c *client) TriggerScan(ctx context.Context, req *TriggerScanReq) (*TriggerScanResp, error) {
	url := "https://api.guardrails.io/v2/cli/trigger-zip-scan"

	req.CLIToken = c.token
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, c.cfg.Timeout)
	defer cancel()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.httpclient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := parseHTTPRespStatusCode("TriggerScan", resp); err != nil {
		return nil, err
	}

	respBody := new(TriggerScanResp)
	if err := json.NewDecoder(resp.Body).Decode(respBody); err != nil {
		return nil, err
	}

	return respBody, nil
}

// GetScanData implements guardrailsclient.GuardRailsClient interface.
func (c *client) GetScanData(ctx context.Context, req *GetScanDataReq) (*GetScanDataResp, error) {
	url := "https://api.guardrails.io/v2/cli/scan"

	ctx, cancel := context.WithTimeout(ctx, c.cfg.Timeout)
	defer cancel()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("clitoken", c.token)
	httpReq.Header.Set("idscan", req.ScanID)

	resp, err := c.httpclient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, httpClient.ErrNotFound
	}

	if err := parseHTTPRespStatusCode("GetScanData", resp); err != nil {
		return nil, err
	}

	respBody := new(GetScanDataResp)
	if err := json.NewDecoder(resp.Body).Decode(respBody); err != nil {
		return nil, err
	}

	return respBody, nil
}
