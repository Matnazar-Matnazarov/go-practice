package main

import (
	"io"
	"log/slog"
	"os"
)

// Env is the runtime environment (dev, prod).
type Env string

const (
	EnvDev  Env = "dev"
	EnvProd Env = "prod"
)

// LoggerConfig configures the global logger.
type LoggerConfig struct {
	Env   Env
	Level slog.Level
	Out   io.Writer
}

// NewLogger creates a slog.Logger with JSON handler for prod and text for dev.
// Sets it as default so slog.Info/Error etc. use it.
func NewLogger(cfg LoggerConfig) *slog.Logger {
	if cfg.Out == nil {
		cfg.Out = os.Stderr
	}
	opts := &slog.HandlerOptions{
		Level: cfg.Level,
		AddSource: cfg.Env == EnvDev,
	}
	var h slog.Handler
	if cfg.Env == EnvProd {
		h = slog.NewJSONHandler(cfg.Out, opts)
	} else {
		h = slog.NewTextHandler(cfg.Out, opts)
	}
	logger := slog.New(h)
	slog.SetDefault(logger)
	return logger
}

// ParseLevel converts string to slog.Level (e.g. "debug" -> LevelDebug).
func ParseLevel(s string) slog.Level {
	switch s {
	case "debug", "DEBUG":
		return slog.LevelDebug
	case "info", "INFO":
		return slog.LevelInfo
	case "warn", "WARN":
		return slog.LevelWarn
	case "error", "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

// WithRequestID returns a logger that adds "request_id" to every log record.
// Simulates request-scoped logging used in HTTP middleware.
func WithRequestID(logger *slog.Logger, requestID string) *slog.Logger {
	if logger == nil {
		logger = slog.Default()
	}
	return logger.With(slog.String("request_id", requestID))
}

// WithGroup returns a logger that nests all attributes under the given group name.
func WithGroup(logger *slog.Logger, group string) *slog.Logger {
	if logger == nil {
		logger = slog.Default()
	}
	return logger.WithGroup(group)
}
