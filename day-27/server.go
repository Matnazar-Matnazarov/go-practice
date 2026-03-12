package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

// NewServer returns an HTTP server configured for the graceful shutdown demo.
// Call server.ListenAndServe() in a goroutine, then ShutdownServer(server, timeout) on signal.
func NewServer(addr string) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Kun 27: Graceful Shutdown demo\nPath: %s\n", r.URL.Path)
	})
	mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})
	mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
		w.Write([]byte("slow response"))
	})
	return &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}

// ShutdownServer gracefully shuts down the HTTP server with a timeout.
func ShutdownServer(server *http.Server, timeout time.Duration) error {
	if server == nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	log.Printf("shutdown: draining connections (timeout %v)", timeout)
	err := server.Shutdown(ctx)
	if err != nil && err != context.DeadlineExceeded {
		return err
	}
	if err == context.DeadlineExceeded {
		log.Printf("shutdown: timeout; some connections may have been closed forcibly")
	}
	return nil
}
