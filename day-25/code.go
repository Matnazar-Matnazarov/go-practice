package main

import (
	"log/slog"
	"os"
)

func main() {
	// Environment-based logger: dev = text + debug, prod = JSON + info
	env := EnvDev
	if v := os.Getenv("ENV"); v == "prod" || v == "production" {
		env = EnvProd
	}
	level := ParseLevel(os.Getenv("LOG_LEVEL"))
	_ = NewLogger(LoggerConfig{Env: env, Level: level, Out: os.Stdout})

	slog.Info("Kun 25: Structured Logging (log/slog) demo boshlandi")

	// 1. Log levels
	slog.Debug("debug message — faqat LOG_LEVEL=debug da ko‘rinadi")
	slog.Info("info message", "port", 8080, "env", string(env))
	slog.Warn("warn message", "reason", "deprecated API")
	slog.Error("error message", "err", "connection refused")

	// 2. WithAttrs — request-scoped logger
	reqLog := WithRequestID(slog.Default(), "req-abc-123")
	reqLog.Info("handler started", "path", "/api/users")
	reqLog.Info("handler finished", "status", 200, "duration_ms", 12)

	// 3. WithGroup — nested attributes (JSON’da guruh)
	httpLog := WithGroup(slog.Default(), "http")
	httpLog.Info("request",
		"method", "GET",
		"path", "/api/health",
		"status", 200,
		"duration_ms", 5,
	)

	// 4. Structured error and context-like attributes
	slog.Info("order created",
		slog.Group("order",
			"id", "ord-001",
			"amount", 99.99,
			"currency", "USD",
		),
		slog.String("user_id", "user-42"),
	)

	slog.Info("Kun 25 demo yakunlandi", "topic", "slog")
}
