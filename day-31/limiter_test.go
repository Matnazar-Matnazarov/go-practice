package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"golang.org/x/time/rate"
)

func TestRateLimiter(t *testing.T) {
	// Bitta burst ruxsat beramiz. 
	limiter := NewIPRateLimiter(rate.Limit(10), 1)

	handler := RateLimitMiddleware(limiter, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "/", nil)
	req.RemoteAddr = "192.168.1.1:1234"

	// Birinchi so'rov (O'tishi kerak)
	w1 := httptest.NewRecorder()
	handler(w1, req)
	if w1.Code != http.StatusOK {
		t.Errorf("Kutilgan 200 OK, olingan %d", w1.Code)
	}

	// Ikkinchi so'rov (Burst to'ldi, Too Many Requests)
	w2 := httptest.NewRecorder()
	handler(w2, req)
	if w2.Code != http.StatusTooManyRequests {
		t.Errorf("Kutilgan 429 Too Many Requests, olingan %d", w2.Code)
	}
}
