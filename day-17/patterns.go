package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// 1) PIPELINE

func demoPipeline() {
	ctx, cancel := context.WithTimeout(context.Background(), 700*time.Millisecond)
	defer cancel()

	in := gen(ctx, 1, 2, 3, 4, 5, 6)
	stage1 := mapStage(ctx, in, func(n int) (int, error) {
		time.Sleep(60 * time.Millisecond)
		return n * n, nil
	})
	stage2 := filterStage(ctx, stage1, func(n int) bool {
		return n%2 == 0
	})

	for v := range stage2 {
		fmt.Printf("  out=%d\n", v)
	}
	fmt.Printf("  pipeline done: %v\n", ctx.Err())
}

func gen(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func mapStage(ctx context.Context, in <-chan int, fn func(int) (int, error)) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case n, ok := <-in:
				if !ok {
					return
				}
				v, err := fn(n)
				if err != nil {
					return
				}
				select {
				case out <- v:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func filterStage(ctx context.Context, in <-chan int, keep func(int) bool) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case n, ok := <-in:
				if !ok {
					return
				}
				if !keep(n) {
					continue
				}
				select {
				case out <- n:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}


// 2) WORKER POOL (fan-out/fan-in)


func demoWorkerPool() {
	const (
		workers = 3
		jobs    = 10
	)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	jobCh := make(chan int)
	out := make(chan string)

	var wg sync.WaitGroup
	for w := 1; w <= workers; w++ {
		wg.Add(1)
		workerID := w
		go func() {
			defer wg.Done()
			for {
				select {
				case j, ok := <-jobCh:
					if !ok {
						return
					}
					// Simulyatsiya: ish vaqti
					time.Sleep(80 * time.Millisecond)
					msg := fmt.Sprintf("worker=%d processed job=%d", workerID, j)
					select {
					case out <- msg:
					case <-ctx.Done():
						return
					}
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	go func() {
		defer close(jobCh)
		for j := 1; j <= jobs; j++ {
			select {
			case jobCh <- j:
			case <-ctx.Done():
				return
			}
		}
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	count := 0
	for msg := range out {
		count++
		fmt.Printf("  %s\n", msg)
	}
	fmt.Printf("  done (results=%d, err=%v)\n", count, ctx.Err())
}


// 3) RATE LIMITING (token bucket)


func demoRateLimiting() {
	// Token bucket: har 200ms 1 token, maksimum 3 token "burst".
	const (
		perToken = 200 * time.Millisecond
		burst    = 3
	)

	tokens := make(chan struct{}, burst)
	for i := 0; i < burst; i++ {
		tokens <- struct{}{}
	}

	ticker := time.NewTicker(perToken)
	defer ticker.Stop()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for range ticker.C {
			select {
			case tokens <- struct{}{}:
			default:
				// bucket full
			}
		}
	}()

	start := time.Now()
	for i := 1; i <= 10; i++ {
		<-tokens
		fmt.Printf("  request=%d allowed at %s\n", i, time.Since(start).Truncate(time.Millisecond))
	}

	// tozalash
	_ = done
}


// 4) ERRGROUP-STYLE (cancel on first error)


type Group struct {
	ctx    context.Context
	cancel context.CancelFunc

	wg  sync.WaitGroup
	mu  sync.Mutex
	err error
}

func WithContext(parent context.Context) (*Group, context.Context) {
	ctx, cancel := context.WithCancel(parent)
	return &Group{ctx: ctx, cancel: cancel}, ctx
}

func (g *Group) Go(fn func(context.Context) error) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		if err := fn(g.ctx); err != nil {
			g.mu.Lock()
			if g.err == nil {
				g.err = err
				g.cancel()
			}
			g.mu.Unlock()
		}
	}()
}

func (g *Group) Wait() error {
	g.wg.Wait()
	g.cancel()
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.err
}

func demoErrGroup() {
	g, ctx := WithContext(context.Background())

	inputs := []int{1, 2, 3, 4, 5}
	results := make(chan int)

	// producer
	g.Go(func(ctx context.Context) error {
		defer close(results)
		for _, n := range inputs {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case results <- n:
			}
		}
		return nil
	})

	// consumers (2 ta ishchi)
	out := make(chan int)
	var fanInWG sync.WaitGroup
	for w := 0; w < 2; w++ {
		fanInWG.Add(1)
		workerID := w + 1
		g.Go(func(ctx context.Context) error {
			defer fanInWG.Done()
			for {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case n, ok := <-results:
					if !ok {
						return nil
					}
					time.Sleep(70 * time.Millisecond)
					if n == 4 {
						return fmt.Errorf("worker=%d failed on n=4", workerID)
					}
					select {
					case out <- n * 10:
					case <-ctx.Done():
						return ctx.Err()
					}
				}
			}
		})
	}

	go func() {
		fanInWG.Wait()
		close(out)
	}()

	sum := 0
	for v := range out {
		sum += v
		fmt.Printf("  got=%d\n", v)
	}

	err := g.Wait()
	if err == nil {
		fmt.Printf("  all done, sum=%d\n", sum)
		return
	}
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		fmt.Printf("  stopped: %v (sum=%d)\n", err, sum)
		return
	}
	fmt.Printf("  stopped early due to error: %v (sum=%d)\n", err, sum)
	_ = ctx
}

