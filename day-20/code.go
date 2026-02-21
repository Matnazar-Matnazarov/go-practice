package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Kun 20: HTTP Server va Web API (net/http)")
	fmt.Println()

	// 1. Handler va HandlerFunc
	fmt.Println("=== 1. Handler va HandlerFunc ===")
	demonstrateHandlerBasics()
	fmt.Println()

	// 2. Request ma'lumotlari
	fmt.Println("=== 2. Request (Method, URL, Header) ===")
	demonstrateRequestInfo()
	fmt.Println()

	// 3. Response (status, header, body)
	fmt.Println("=== 3. Response (WriteHeader, Header, Body) ===")
	demonstrateResponseBasics()
	fmt.Println()

	// 4. Demo HTTP server ishga tushirish
	fmt.Println("=== 4. Demo HTTP Server ===")
	fmt.Println("Server :8080 da 5 soniya ishlaydi. Browser da:")
	fmt.Println("  http://localhost:8080/")
	fmt.Println("  http://localhost:8080/greet/Dunyo")
	fmt.Println("  http://localhost:8080/api/health")
	fmt.Println("  http://localhost:8080/api/users")
	fmt.Println()

	go runHTTPServer(":8080")
	time.Sleep(5 * time.Second)

	fmt.Println("=== Kun 20 demo yakunlandi! ===")
	fmt.Println("O'rganildi:")
	fmt.Println("  ✓ http.Handler va HandlerFunc")
	fmt.Println("  ✓ ServeMux va routing")
	fmt.Println("  ✓ Request: Method, URL, Header, Body")
	fmt.Println("  ✓ Response: status, header, JSON")
	fmt.Println("  ✓ Middleware (logging)")
	fmt.Println("  ✓ http.Server va timeout'lar")
}

// demonstrateHandlerBasics — Handler va HandleFunc tushunchasi
func demonstrateHandlerBasics() {
	// HandlerFunc — oddiy funksiyani Handler ga aylantiradi
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HandlerFunc orqali javob\n"))
	})
	_ = handler // demo: faqat tip ko'rsatish
	fmt.Println("  HandlerFunc: func(ResponseWriter, *Request) -> Handler")
	fmt.Println("  ServeMux path bo'yicha handlerni tanlaydi")
}

// demonstrateRequestInfo — Request tarkibi
func demonstrateRequestInfo() {
	// Real so'rovsiz, faqat tushuncha
	fmt.Println("  r.Method     — GET, POST, ...")
	fmt.Println("  r.URL.Path   — yo'l")
	fmt.Println("  r.URL.Query()— query parametrlar")
	fmt.Println("  r.Header     — request header'lar")
	fmt.Println("  r.Body       — body (io.ReadAll / json.Decoder)")
}

// demonstrateResponseBasics — Response yozish
func demonstrateResponseBasics() {
	fmt.Println("  w.WriteHeader(status) — status kod (bir marta)")
	fmt.Println("  w.Header().Set(...)    — response header")
	fmt.Println("  w.Write([]byte)       — body")
	fmt.Println("  json.NewEncoder(w).Encode(v) — JSON javob")
}
