package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Kun 21: HTTP Client va API Testing (net/http, httptest)")
	fmt.Println()

	fmt.Println("=== 1. HTTP Client asoslari ===")
	demonstrateClientBasics()
	fmt.Println()

	fmt.Println("=== 2. Custom Client va Timeout ===")
	demonstrateCustomClient()
	fmt.Println()

	fmt.Println("=== 3. Context bilan so'rov ===")
	demonstrateContextRequest()
	fmt.Println()

	fmt.Println("=== 4. httptest server bilan demo ===")
	demonstrateHTTPTest()
	fmt.Println()

	fmt.Println("=== Kun 21 yakunlandi! ===")
	fmt.Println("O'rganildi: http.Client, timeout, context, httptest")
}

func demonstrateClientBasics() {
	fmt.Println("  http.Get(url) — default client (timeout yo'q)")
	fmt.Println("  client.Do(req) — custom client bilan so'rov")
	fmt.Println("  resp.Body.Close() — har doim yopish")
}

func demonstrateCustomClient() {
	client := &http.Client{Timeout: 5 * time.Second}
	_ = client
	fmt.Println("  client := &http.Client{Timeout: 5 * time.Second}")
	fmt.Println("  Production'da timeout majburiy")
}

func demonstrateContextRequest() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_ = ctx
	fmt.Println("  NewRequestWithContext(ctx, method, url, body)")
	fmt.Println("  Context — timeout va cancellation uchun")
}

func demonstrateHTTPTest() {
	ts := startTestServer()
	defer ts.Close()
	client := ts.Client()
	resp, err := client.Get(ts.URL + "/api/health")
	if err != nil {
		fmt.Printf("  Xato: %v\n", err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("  Test server javob: status=%d\n", resp.StatusCode)
}
