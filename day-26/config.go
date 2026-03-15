package main

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds application configuration loaded from environment (and optional file).
type Config struct {
	Env      string
	Port     int
	LogLevel string
	DBURL    string
	Debug    bool
}

// getEnv returns env value or defaultValue if unset or empty.
func getEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}

// LoadConfig loads Config from environment. Returns error if required vars are missing or invalid.
func LoadConfig() (*Config, error) {
	cfg := &Config{
		Env:      getEnv("ENV", "dev"),
		Port:     8080,
		LogLevel: getEnv("LOG_LEVEL", "info"),
		DBURL:    os.Getenv("DATABASE_URL"),
		Debug:    getEnv("DEBUG", "false") == "true" || getEnv("DEBUG", "") == "1",
	}

	if p := getEnv("PORT", ""); p != "" {
		port, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("PORT must be a number: %w", err)
		}
		cfg.Port = port
	}

	if err := validateConfig(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func validateConfig(cfg *Config) error {
	if cfg.Port < 1 || cfg.Port > 65535 {
		return fmt.Errorf("PORT must be 1-65535, got %d", cfg.Port)
	}
	// In real app you might require DATABASE_URL in prod:
	// if cfg.Env == "prod" && cfg.DBURL == "" { return fmt.Errorf("DATABASE_URL required in prod") }
	return nil
}
