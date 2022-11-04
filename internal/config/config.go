package config

import "time"

// Config provides centralized configuration for the application.
type Config struct {
	HttpClient       *HttpClientConfig
	GuardRailsClient *GuardRailsClient
}

// New instantiates new config.
func New() *Config {
	return &Config{
		HttpClient: &HttpClientConfig{
			PollingInterval: 15 * time.Second,
			Timeout:         10 * time.Second,
			RetryTimeout:    30 * time.Minute,
		},
		GuardRailsClient: NewGuardRailsClientConfig(),
	}
}

// HttpClientConfig provides configuration for guardrails cli http client.
type HttpClientConfig struct {
	// PollingInterval define how long the client should wait before attempting to connect to API server.
	PollingInterval time.Duration
	// Timeout define maximum waiting time allowed for http connections before it considered timeout.
	Timeout time.Duration
	// RetryTimeout define maximum waiting time of doing backoff retry in case client get unexpected response from API server.
	RetryTimeout time.Duration
}
