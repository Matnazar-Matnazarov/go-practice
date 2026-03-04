package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Kun 22: HTTP Middleware va Error Handling (net/http)")
	fmt.Println()

	fmt.Println("Demo server :8080 da 5 soniya ishlaydi.")
	fmt.Println("  http://localhost:8080/api/health")
	fmt.Println("  http://localhost:8080/api/panic")
	fmt.Println()

	go runHTTPServer(":8080")
	time.Sleep(5 * time.Second)

	fmt.Println("=== Kun 22 demo yakunlandi! ===")
}
