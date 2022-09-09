package client

import (
	"net/http"

	"github.com/guardrailsio/guardrails-cli/internal/config"
)

// New instantiates new Http client.
func New(cfg *config.HttpClientConfig) *http.Client {
	return &http.Client{Timeout: cfg.Timeout}
}
