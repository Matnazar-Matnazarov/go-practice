package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"time"
)

type healthResp struct {
	Status string `json:"status"`
	Time   string `json:"time"`
}

func startTestServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(healthResp{Status: "ok", Time: time.Now().Format(time.RFC3339)})
	})
	return httptest.NewServer(mux)
}
