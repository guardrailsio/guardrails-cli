package guardrailsclient

import (
	"net/http"

	httpClient "github.com/guardrailsio/guardrails-cli/internal/client"
	"github.com/guardrailsio/guardrails-cli/internal/config"
)

// GuardRailsClient variable that does static check to make sure that client struct implements GuardRailsClient interface.
var _ GuardRailsClient = (*client)(nil)

//go:generate mockgen -destination=mock/client.go -package=mockguardrailsclient . GuardRailsClient

// GuardRailsClient defines methods to interact with GuardRails API.
type GuardRailsClient interface {
	CreateUploadURL() error
	UploadProject() error
	TriggerScan() error
	GetScanData() error
}

type client struct {
	cfg    *config.HttpClientConfig
	client *http.Client
	token  string
}

// New instantiates new GuardRailsClient.
func New(cfg *config.HttpClientConfig, token string) GuardRailsClient {
	c := httpClient.New(cfg)

	return &client{cfg: cfg, client: c, token: token}
}

// CreateUploadURL call GuardRails api to create upload URL.
func (c *client) CreateUploadURL() error {
	return nil
}

// UploadProject accepts url generated from CreateUploadURL method, compress the project, and upload it to designated url.
func (c *client) UploadProject() error {
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
