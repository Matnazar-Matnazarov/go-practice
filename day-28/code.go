package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	// Hub yaratish va ishga tushirish
	hub := NewHub()
	go hub.Run()

	// HTTP router
	mux := http.NewServeMux()

	// WebSocket endpoint
	mux.HandleFunc("/ws", hub.ServeHTTP)

	// HTTP API endpoints
	mux.HandleFunc("/broadcast", hub.BroadcastMessage)
	mux.HandleFunc("/clients", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"clients":` + strconv.Itoa(hub.GetClientCount()) + `}`))
	})

	// Static files (HTML frontend)
	mux.Handle("/", http.FileServer(http.Dir(".")))

	// Server yaratish
	addr := ":8080"
	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Server goroutine'da ishga tushirish
	go func() {
		log.Printf("🚀 WebSocket server listening on %s", addr)
		log.Printf("📡 Open http://localhost:8080 in your browser")
		log.Printf("💬 Connect: ws://localhost:8080/ws")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server: %v", err)
		}
	}()

	// Signal handling (graceful shutdown)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	log.Println("🛑 Signal received, starting graceful shutdown...")

	// Shutdown timeout
	shutdownTimeout := 10 * time.Second
	hub.broadcast <- []byte("👋 Server shutting down...")
	time.Sleep(100 * time.Millisecond)

	if err := server.Close(); err != nil {
		log.Printf("⚠️  Shutdown error: %v", err)
	}

	log.Printf("✅ Server stopped (timeout: %v)", shutdownTimeout)
}
