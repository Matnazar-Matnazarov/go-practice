package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// User — API da ishlatiladigan model
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// HealthResponse — /health uchun javob
type HealthResponse struct {
	Status string `json:"status"`
	Time   string `json:"time"`
}

// homeHandler — bosh sahifa
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Kun 20: HTTP Server va Web API\nPath: %s\nMethod: %s\n", r.URL.Path, r.Method)
}

// greetHandler — /greet/Name
func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/greet/"):]
	if name == "" {
		name = "Mehmon"
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Salom, %s!\n", name)
}

// healthHandler — /api/health
func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	resp := HealthResponse{Status: "ok", Time: time.Now().Format(time.RFC3339)}
	json.NewEncoder(w).Encode(resp)
}

// usersHandler — /api/users (GET — ro‘yxat, POST — yozuv qo‘shish demo)
func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		users := []User{
			{ID: 1, Name: "Ali"},
			{ID: 2, Name: "Vali"},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(users); err != nil {
			log.Printf("Encode error: %v", err)
		}
	case http.MethodPost:
		var u User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		// Demo: faqat echo qilamiz
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(u)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// loggingMiddleware — so‘rovlarni log qiladi
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

// runHTTPServer — demo HTTP server (bloklaydi)
func runHTTPServer(addr string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/greet/", greetHandler)
	mux.HandleFunc("/api/health", healthHandler)
	mux.HandleFunc("/api/users", usersHandler)

	wrapped := loggingMiddleware(mux)
	server := &http.Server{
		Addr:         addr,
		Handler:      wrapped,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("Server %s da ishga tushdi. To'xtatish: Ctrl+C\n", addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
