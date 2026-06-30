package main

import (
	"fmt"
	"log"
	"net/http"
	"golang.org/x/time/rate"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Salom! Muvaffaqiyatli so'rov."))
}

func main() {
	// Har soniyada 2 ta so'rov ruxsat etiladi (token bucket).
	// Bitta vaqtda maksimal 5 ta so'rov qabul qilinishi mumkin (burst).
	limiter := NewIPRateLimiter(rate.Limit(2), 5)

	http.HandleFunc("/", RateLimitMiddleware(limiter, helloHandler))

	fmt.Println("🚀 Rate Limiter server ishga tushdi: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
