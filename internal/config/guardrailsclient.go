package config

import (
	"errors"
	"os"
)

var (
	ErrMissingGuardRailsAPIHost = errors.New("missing mandatory GUARDRAILS_API_HOST environment variables.")
)

// GuardRailsClient contains configuration for guardrails client.
type GuardRailsClient struct {
	APIHost string
}

func NewGuardRailsClientConfig() *GuardRailsClient {
	apiHost := os.Getenv("GUARDRAILS_API_HOST")
	if apiHost == "" {
		apiHost = "https://api.guardrails.io"
	}

	return &GuardRailsClient{APIHost: apiHost}
}
