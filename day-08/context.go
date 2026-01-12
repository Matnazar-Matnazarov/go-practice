package main

import (
	"context"
	"fmt"
	"time"
)

// demonstrateContextBasic - Context asoslari
func demonstrateContextBasic() {
	fmt.Println("Context asoslari:")

	// Background context
	ctx := context.Background()
	fmt.Printf("  Background context: %v\n", ctx)

	// TODO context
	todoCtx := context.TODO()
	fmt.Printf("  TODO context: %v\n", todoCtx)

	// Context qiymatini olish
	fmt.Printf("  Context nil: %t\n", ctx == nil)
}

// demonstrateContextWithValue - Context bilan qiymat
func demonstrateContextWithValue() {
	fmt.Println("Context bilan qiymat:")

	ctx := context.Background()

	// Qiymat qo'shish
	ctx = context.WithValue(ctx, "userID", "12345")
	ctx = context.WithValue(ctx, "username", "john_doe")

	// Qiymat olish
	userID := ctx.Value("userID")
	username := ctx.Value("username")

	fmt.Printf("  User ID: %v\n", userID)
	fmt.Printf("  Username: %v\n", username)
}

// demonstrateContextWithCancel - Context bilan cancel
func demonstrateContextWithCancel() {
	fmt.Println("Context bilan cancel:")

	ctx, cancel := context.WithCancel(context.Background())

	// Goroutine
	go func() {
		for i := 0; i < 5; i++ {
			select {
			case <-ctx.Done():
				fmt.Printf("  Goroutine to'xtatildi: %v\n", ctx.Err())
				return
			default:
				fmt.Printf("  Goroutine ishlayapti: %d\n", i)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	time.Sleep(300 * time.Millisecond)
	cancel() // Cancel qilish
	time.Sleep(100 * time.Millisecond)
}

// demonstrateContextWithTimeout - Context bilan timeout
func demonstrateContextWithTimeout() {
	fmt.Println("Context bilan timeout:")

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	// Goroutine
	done := make(chan bool)
	go func() {
		time.Sleep(300 * time.Millisecond) // Timeout dan uzun
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("  Vazifa yakunlandi")
	case <-ctx.Done():
		fmt.Printf("  Timeout: %v\n", ctx.Err())
	}
}

// demonstrateContextWithDeadline - Context bilan deadline
func demonstrateContextWithDeadline() {
	fmt.Println("Context bilan deadline:")

	deadline := time.Now().Add(200 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	// Goroutine
	done := make(chan bool)
	go func() {
		time.Sleep(300 * time.Millisecond)
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("  Vazifa yakunlandi")
	case <-ctx.Done():
		fmt.Printf("  Deadline o'tdi: %v\n", ctx.Err())
	}
}

// demonstrateContextNested - Ichki context
func demonstrateContextNested() {
	fmt.Println("Ichki context:")

	parentCtx := context.Background()
	parentCtx = context.WithValue(parentCtx, "parent", "value1")

	// Child context
	childCtx, cancel := context.WithCancel(parentCtx)
	childCtx = context.WithValue(childCtx, "child", "value2")

	// Qiymatlarni olish
	fmt.Printf("  Parent value: %v\n", childCtx.Value("parent"))
	fmt.Printf("  Child value: %v\n", childCtx.Value("child"))

	// Cancel
	cancel()
	fmt.Printf("  Child context canceled: %v\n", childCtx.Err())
}

// demonstrateContextMultipleGoroutines - Bir nechta goroutine
func demonstrateContextMultipleGoroutines() {
	fmt.Println("Context bilan bir nechta goroutine:")

	ctx, cancel := context.WithCancel(context.Background())

	// 3 ta goroutine
	for i := 1; i <= 3; i++ {
		go func(id int) {
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("  Goroutine %d: To'xtatildi\n", id)
					return
				default:
					fmt.Printf("  Goroutine %d: Ishlayapti\n", id)
					time.Sleep(100 * time.Millisecond)
				}
			}
		}(i)
	}

	time.Sleep(300 * time.Millisecond)
	cancel()
	time.Sleep(100 * time.Millisecond)
}

// demonstrateContextTimeoutPropagation - Timeout tarqalishi
func demonstrateContextTimeoutPropagation() {
	fmt.Println("Context timeout tarqalishi:")

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	// Birinchi goroutine
	go func() {
		select {
		case <-ctx.Done():
			fmt.Printf("  Goroutine 1: %v\n", ctx.Err())
		}
	}()

	// Ikkinchi goroutine (child context)
	childCtx, childCancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer childCancel()

	go func() {
		select {
		case <-childCtx.Done():
			fmt.Printf("  Goroutine 2: %v\n", childCtx.Err())
		}
	}()

	time.Sleep(300 * time.Millisecond)
}

// demonstrateContextValuePropagation - Qiymat tarqalishi
func demonstrateContextValuePropagation() {
	fmt.Println("Context qiymat tarqalishi:")

	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", "req-12345")

	// Funksiya chaqirish
	processRequest(ctx)
}

// processRequest - Request ni qayta ishlash
func processRequest(ctx context.Context) {
	requestID := ctx.Value("requestID")
	fmt.Printf("  Request ID: %v\n", requestID)

	// Child context
	childCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Child context dan qiymat olish
	fmt.Printf("  Child context Request ID: %v\n", childCtx.Value("requestID"))
}

// demonstrateContextHTTPRequest - HTTP request simulatsiyasi
func demonstrateContextHTTPRequest() {
	fmt.Println("HTTP request simulatsiyasi:")

	ctx := context.Background()
	ctx = context.WithValue(ctx, "userID", "user-123")
	ctx = context.WithValue(ctx, "requestID", "req-456")

	// Request timeout
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// API chaqiruv simulatsiyasi
	done := make(chan string)
	go func() {
		time.Sleep(500 * time.Millisecond)
		done <- "API javob"
	}()

	select {
	case result := <-done:
		fmt.Printf("  %s\n", result)
		fmt.Printf("  User ID: %v\n", ctx.Value("userID"))
		fmt.Printf("  Request ID: %v\n", ctx.Value("requestID"))
	case <-ctx.Done():
		fmt.Printf("  Request timeout: %v\n", ctx.Err())
	}
}

// demonstrateContextDatabaseQuery - Database query simulatsiyasi
func demonstrateContextDatabaseQuery() {
	fmt.Println("Database query simulatsiyasi:")

	ctx := context.Background()
	ctx = context.WithValue(ctx, "db", "postgresql")
	ctx = context.WithValue(ctx, "query", "SELECT * FROM users")

	// Query timeout
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	// Query simulatsiyasi
	done := make(chan []string)
	go func() {
		time.Sleep(200 * time.Millisecond)
		done <- []string{"user1", "user2", "user3"}
	}()

	select {
	case results := <-done:
		fmt.Printf("  Natijalar: %v\n", results)
		fmt.Printf("  DB: %v\n", ctx.Value("db"))
		fmt.Printf("  Query: %v\n", ctx.Value("query"))
	case <-ctx.Done():
		fmt.Printf("  Query timeout: %v\n", ctx.Err())
	}
}

// demonstrateContextWorkerPool - Worker pool bilan
func demonstrateContextWorkerPool() {
	fmt.Println("Context bilan worker pool:")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// Job lar yuborish
	go func() {
		for j := 1; j <= 5; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	// Worker lar
	workers := 3
	for w := 1; w <= workers; w++ {
		go func(workerID int) {
			for {
				select {
				case <-ctx.Done():
					return
				case job, ok := <-jobs:
					if !ok {
						return
					}
					fmt.Printf("  Worker %d: Job %d\n", workerID, job)
					results <- job * 2
				}
			}
		}(w)
	}

	// Natijalarni olish
	for i := 0; i < 5; i++ {
		result := <-results
		fmt.Printf("  Natija: %d\n", result)
	}
}

// demonstrateContextErrorHandling - Xato bilan ishlash
func demonstrateContextErrorHandling() {
	fmt.Println("Context bilan xato bilan ishlash:")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Xato bilan vazifa
	type Result struct {
		Value int
		Error error
	}

	results := make(chan Result, 3)

	for i := 1; i <= 3; i++ {
		go func(id int) {
			select {
			case <-ctx.Done():
				results <- Result{Error: ctx.Err()}
				return
			default:
				if id == 2 {
					results <- Result{Error: fmt.Errorf("xato: %d", id)}
					return
				}
				results <- Result{Value: id * 2}
			}
		}(i)
	}

	// Natijalarni ko'rsatish
	for i := 0; i < 3; i++ {
		result := <-results
		if result.Error != nil {
			fmt.Printf("  Xato: %v\n", result.Error)
		} else {
			fmt.Printf("  Natija: %d\n", result.Value)
		}
	}
}

// demonstrateContextPractical - Amaliy misol: API client
func demonstrateContextPractical() {
	fmt.Println("Amaliy misol: API client simulatsiyasi:")

	ctx := context.Background()
	ctx = context.WithValue(ctx, "apiKey", "secret-key-123")
	ctx = context.WithValue(ctx, "endpoint", "/api/users")

	// Request timeout
	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()

	// API chaqiruv
	done := make(chan map[string]interface{})
	go func() {
		time.Sleep(300 * time.Millisecond)
		done <- map[string]interface{}{
			"status": 200,
			"data":   []string{"user1", "user2"},
		}
	}()

	select {
	case response := <-done:
		fmt.Printf("  Status: %v\n", response["status"])
		fmt.Printf("  Data: %v\n", response["data"])
		fmt.Printf("  API Key: %v\n", ctx.Value("apiKey"))
	case <-ctx.Done():
		fmt.Printf("  Request failed: %v\n", ctx.Err())
	}
}
