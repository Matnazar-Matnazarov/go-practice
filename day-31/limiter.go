package main

import (
	"net"
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

// IPRateLimiter IP manzillari bo'yicha limitlarni saqlaydi
type IPRateLimiter struct {
	ips map[string]*rate.Limiter
	mu  *sync.RWMutex
	r   rate.Limit
	b   int
}

// NewIPRateLimiter yangi limiter yaratadi
func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	i := &IPRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}
	return i
}

// AddIP yangi IP manzilini qo'shadi yoki eskisini qaytaradi
func (i *IPRateLimiter) AddIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter, exists := i.ips[ip]
	if !exists {
		limiter = rate.NewLimiter(i.r, i.b)
		i.ips[ip] = limiter
	}

	return limiter
}

// GetLimiter IP uchun limiterni qaytaradi
func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.RLock()
	limiter, exists := i.ips[ip]
	if !exists {
		i.mu.RUnlock()
		return i.AddIP(ip)
	}
	i.mu.RUnlock()
	return limiter
}

// RateLimitMiddleware HTTP so'rovlarni cheklash uchun middleware
func RateLimitMiddleware(limiter *IPRateLimiter, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Mijozning IP manzilini ajratib olish
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			// Agar RemoteAddr noto'g'ri bo'lsa, IP sifatida RemoteAddr ishlatiladi
			ip = r.RemoteAddr
		}

		l := limiter.GetLimiter(ip)
		if !l.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	}
}
