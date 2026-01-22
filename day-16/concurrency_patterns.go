package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func demoBoundedConcurrency() {
	// Semaphore orqali bir vaqtda ishlaydigan goroutine sonini cheklash.
	// Bu pattern real loyihada: API call'lar, DB query'lar, file I/O uchun juda foydali.
	const (
		jobs        = 10
		maxInFlight = 3
	)

	sem := make(chan struct{}, maxInFlight)
	var wg sync.WaitGroup

	start := time.Now()
	for i := 1; i <= jobs; i++ {
		wg.Add(1)
		sem <- struct{}{}
		jobID := i

		go func() {
			defer wg.Done()
			defer func() { <-sem }()

			// Simulyatsiya: ish vaqti
			time.Sleep(120 * time.Millisecond)
			fmt.Printf("  job=%d done\n", jobID)
		}()
	}

	wg.Wait()
	fmt.Printf("  done in %s (max in-flight=%d)\n", time.Since(start).Truncate(time.Millisecond), maxInFlight)
}

func demoContextCancellation() {
	// Ko‘p ish (fan-out) va xato bo‘lsa, qolganlarini bekor qilish.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	inputs := []int{1, 2, 3, 4, 5, 6}
	results := make(chan int)
	errCh := make(chan error, 1)

	var wg sync.WaitGroup
	for _, n := range inputs {
		wg.Add(1)
		n := n
		go func() {
			defer wg.Done()

			v, err := work(ctx, n)
			if err != nil {
				// Birinchi xato kelganda cancel qilamiz.
				select {
				case errCh <- err:
					cancel()
				default:
				}
				return
			}

			select {
			case results <- v:
			case <-ctx.Done():
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	count := 0
	for v := range results {
		count++
		fmt.Printf("  result=%d\n", v)
	}

	select {
	case err := <-errCh:
		fmt.Printf("  stopped early due to error: %v\n", err)
	default:
		fmt.Printf("  all done successfully (results=%d)\n", count)
	}
}

func work(ctx context.Context, n int) (int, error) {
	// Simulyatsiya: har bir ish 80ms.
	select {
	case <-time.After(80 * time.Millisecond):
		// Ataylab bitta input'da xato qaytaramiz.
		if n == 4 {
			return 0, errors.New("upstream failure (n=4)")
		}
		return n * n, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

func demoLeakPrevention() {
	// Goroutine leak ko‘pincha: o‘qilmay qolgan channel send, yoki hech qachon tugamaydigan loop.
	// Done channel bilan stop signal berish - eng sodda va ishonchli yo‘l.
	done := make(chan struct{})
	out := make(chan int)

	go func() {
		defer close(out)
		ticker := time.NewTicker(50 * time.Millisecond)
		defer ticker.Stop()

		i := 0
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				i++
				select {
				case out <- i:
				case <-done:
					return
				}
			}
		}
	}()

	for v := range out {
		fmt.Printf("  got=%d\n", v)
		if v >= 3 {
			close(done)
		}
	}
	fmt.Println("  producer stopped cleanly")
}

func demoGracefulShutdownConcept() {
	// Bu misol server yozmaydi (faqat pattern), lekin real loyihada:
	// - signal kelganda context cancel
	// - background worker'lar stop
	// - resource'lar (db, files) yopiladi
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	fmt.Println("  Press Ctrl+C to simulate shutdown (10s auto-timeout) ...")
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Printf("  shutdown triggered: %v\n", ctx.Err())
	}
}
