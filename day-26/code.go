package main

import (
	"fmt"
	"os"
)

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "config: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Kun 26: Configuration va Environment (12-factor)")
	fmt.Println()
	fmt.Println("=== Config (sensitive fields masked) ===")
	fmt.Printf("  ENV          = %s\n", cfg.Env)
	fmt.Printf("  PORT         = %d\n", cfg.Port)
	fmt.Printf("  LOG_LEVEL    = %s\n", cfg.LogLevel)
	fmt.Printf("  DEBUG        = %v\n", cfg.Debug)
	fmt.Printf("  DATABASE_URL = %s\n", maskSecret(cfg.DBURL))
	fmt.Println()
	fmt.Println("Config yuklandi va validate qilindi. Dastur ishga tayyor.")
}

func maskSecret(s string) string {
	if s == "" {
		return "(not set)"
	}
	if len(s) <= 8 {
		return "***"
	}
	return s[:4] + "..." + s[len(s)-2:]
}
