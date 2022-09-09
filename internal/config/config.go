package config

import "time"

// New instantiates new config.
func New() *Config {
	return &Config{
		HttpClient: &HttpClientConfig{
			PollingInterval: 5 * time.Second,
			Timeout:         2 * time.Second,
			RetryTimeout:    30 * time.Minute,
		},
	}
}

// Config provides centralized configuration for the application.
type Config struct {
	HttpClient *HttpClientConfig
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
