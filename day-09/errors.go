package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// demonstrateErrorBasic - Error asoslari
func demonstrateErrorBasic() {
	fmt.Println("Error asoslari:")

	// Error yaratish
	err := errors.New("bu xato")
	fmt.Printf("  Xato: %v\n", err)

	// Errorf bilan
	err2 := fmt.Errorf("xato: %s", "ma'lumot topilmadi")
	fmt.Printf("  Xato: %v\n", err2)
}

// demonstrateErrorInGoroutine - Goroutine da xato
func demonstrateErrorInGoroutine() {
	fmt.Println("Goroutine da xato:")

	errChan := make(chan error, 1)

	go func() {
		// Xato simulatsiyasi
		errChan <- fmt.Errorf("goroutine xatosi")
	}()

	err := <-errChan
	if err != nil {
		fmt.Printf("  Xato qabul qilindi: %v\n", err)
	}
}

// demonstrateErrorMultipleGoroutines - Bir nechta goroutine da xato
func demonstrateErrorMultipleGoroutines() {
	fmt.Println("Bir nechta goroutine da xato:")

	errChan := make(chan error, 5)
	var wg sync.WaitGroup

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(id int) {
			defer wg.Done()
			if id%2 == 0 {
				errChan <- fmt.Errorf("goroutine %d xatosi", id)
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	fmt.Println("  Xatolar:")
	for err := range errChan {
		if err != nil {
			fmt.Printf("    %v\n", err)
		}
	}
}

// demonstrateErrorWithResult - Natija va xato bilan
func demonstrateErrorWithResult() {
	fmt.Println("Natija va xato bilan:")

	type Result struct {
		Value int
		Error error
	}

	results := make(chan Result, 5)
	var wg sync.WaitGroup

	wg.Add(5)
	for i := 1; i <= 5; i++ {
		go func(id int) {
			defer wg.Done()
			if id%2 == 0 {
				results <- Result{
					Value: 0,
					Error: fmt.Errorf("xato: %d", id),
				}
			} else {
				results <- Result{
					Value: id * 2,
					Error: nil,
				}
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("  Natijalar:")
	for result := range results {
		if result.Error != nil {
			fmt.Printf("    Xato: %v\n", result.Error)
		} else {
			fmt.Printf("    Muvaffaqiyatli: %d\n", result.Value)
		}
	}
}

// demonstrateErrorWrapping - Error wrapping
func demonstrateErrorWrapping() {
	fmt.Println("Error wrapping:")

	// Asosiy xato
	baseErr := errors.New("asosiy xato")

	// Xatoni o'rab olish
	wrappedErr := fmt.Errorf("qo'shimcha ma'lumot: %w", baseErr)
	fmt.Printf("  Wrapped error: %v\n", wrappedErr)

	// Asosiy xatoni olish
	if errors.Unwrap(wrappedErr) != nil {
		fmt.Printf("  Unwrapped: %v\n", errors.Unwrap(wrappedErr))
	}
}

// demonstrateErrorIs - errors.Is
func demonstrateErrorIs() {
	fmt.Println("errors.Is:")

	baseErr := errors.New("asosiy xato")
	wrappedErr := fmt.Errorf("wrapper: %w", baseErr)

	// errors.Is - xatoni tekshirish
	if errors.Is(wrappedErr, baseErr) {
		fmt.Println("  Xato topildi!")
	}

	// Yana bir misol
	targetErr := errors.New("maqsad xato")
	if !errors.Is(wrappedErr, targetErr) {
		fmt.Println("  Maqsad xato topilmadi")
	}
}

// CustomError - maxsus xato turi
type CustomError struct {
	Code    int
	Message string
}

// Error - error interface ni implement qilish
func (e *CustomError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

// demonstrateErrorAs - errors.As
func demonstrateErrorAs() {
	fmt.Println("errors.As:")

	err := &CustomError{Code: 404, Message: "topilmadi"}

	var customErr *CustomError
	if errors.As(err, &customErr) {
		fmt.Printf("  Custom error: Code=%d, Message=%s\n", customErr.Code, customErr.Message)
	}
}

// demonstrateErrorTimeout - Timeout bilan xato
func demonstrateErrorTimeout() {
	fmt.Println("Timeout bilan xato:")

	done := make(chan error, 1)

	go func() {
		time.Sleep(200 * time.Millisecond)
		done <- fmt.Errorf("vazifa xatosi")
	}()

	select {
	case err := <-done:
		fmt.Printf("  Xato: %v\n", err)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("  Timeout: Vazifa vaqtida yakunlanmadi")
	}
}

// demonstrateErrorContext - Context bilan xato
func demonstrateErrorContext() {
	fmt.Println("Context bilan xato:")

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	done := make(chan error, 1)

	go func() {
		time.Sleep(200 * time.Millisecond)
		done <- nil
	}()

	select {
	case err := <-done:
		if err != nil {
			fmt.Printf("  Xato: %v\n", err)
		} else {
			fmt.Println("  Muvaffaqiyatli")
		}
	case <-ctx.Done():
		fmt.Printf("  Context xato: %v\n", ctx.Err())
	}
}

// demonstrateErrorRecovery - Panic recovery
func demonstrateErrorRecovery() {
	fmt.Println("Panic recovery:")

	var wg sync.WaitGroup

	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(id int) {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("  Goroutine %d: Panic recovery: %v\n", id, r)
				}
			}()

			if id == 1 {
				panic(fmt.Sprintf("goroutine %d panic", id))
			}
			fmt.Printf("  Goroutine %d: Muvaffaqiyatli\n", id)
		}(i)
	}

	wg.Wait()
}

// demonstrateErrorPractical - Amaliy misol: API client
func demonstrateErrorPractical() {
	fmt.Println("Amaliy misol: API client (simulatsiya):")

	type APIResponse struct {
		Data  string
		Error error
	}

	results := make(chan APIResponse, 3)
	var wg sync.WaitGroup

	apis := []string{"api1", "api2", "api3"}

	wg.Add(len(apis))
	for i, api := range apis {
		go func(id int, endpoint string) {
			defer wg.Done()

			// Xato simulatsiyasi
			if id == 1 {
				results <- APIResponse{
					Data:  "",
					Error: fmt.Errorf("%s xatosi", endpoint),
				}
				return
			}

			results <- APIResponse{
				Data:  fmt.Sprintf("%s javob", endpoint),
				Error: nil,
			}
		}(i, api)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("  API javoblari:")
	for result := range results {
		if result.Error != nil {
			fmt.Printf("    Xato: %v\n", result.Error)
		} else {
			fmt.Printf("    Muvaffaqiyatli: %s\n", result.Data)
		}
	}
}
