package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type healthResponse struct {
	OK   bool   `json:"ok"`
	Time string `json:"time"`
}

func runHTTPServer(addr string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/health", healthHandler)
	mux.HandleFunc("/api/panic", panicHandler)

	h := withRecovery(withLogging(mux))

	srv := &http.Server{
		Addr:              addr,
		Handler:           h,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("server error: %v", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp := healthResponse{OK: true, Time: time.Now().Format(time.RFC3339)}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("encode error: %v", err)
	}
}

func panicHandler(w http.ResponseWriter, r *http.Request) {
	panic("intentional panic for recovery demo")
}

func withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("%s %s %v\n", r.Method, r.URL.Path, time.Since(start))
	})
}

func withRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if v := recover(); v != nil {
				log.Printf("panic recovered: %v", v)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
