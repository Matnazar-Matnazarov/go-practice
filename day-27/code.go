package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	addr := ":8080"
	server := NewServer(addr)

	go func() {
		log.Printf("server listening on %s (Ctrl+C or SIGTERM to graceful shutdown)", addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server: %v", err)
		}
	}()

	sigChan := NotifyShutdown()
	<-sigChan
	log.Println("signal received, starting graceful shutdown")

	shutdownTimeout := 10 * time.Second
	if err := ShutdownServer(server, shutdownTimeout); err != nil {
		log.Printf("shutdown error: %v", err)
	}
	log.Println("server stopped")
}
