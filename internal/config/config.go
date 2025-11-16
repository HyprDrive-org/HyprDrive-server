package config

import (
	"fmt"
	"os"
)

// Config contains runtime settings for the HyprDrive server.
type Config struct {
	HTTPPort    string
	Environment string
}

// Load reads configuration from environment variables and falls back to defaults.
func Load() (*Config, error) {
	cfg := &Config{
		HTTPPort:    getEnv("HYPRDRIVE_HTTP_PORT", "8080"),
		Environment: getEnv("HYPRDRIVE_ENV", "development"),
	}

	if cfg.HTTPPort == "" {
		return nil, fmt.Errorf("http port cannot be empty")
	}

	return cfg, nil
}

func getEnv(key, def string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return def
}
