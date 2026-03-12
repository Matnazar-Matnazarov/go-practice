package main

import (
	"os"
	"os/signal"
	"syscall"
)

// NotifyShutdown registers SIGINT and SIGTERM and returns a channel that
// receives when either signal is sent. Caller should block on <-sigChan.
// Buffer size 1 avoids blocking the signal handler.
func NotifyShutdown() <-chan os.Signal {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	return sigChan
}
