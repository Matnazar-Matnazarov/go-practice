package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// 1. Basic Worker Pool

type Job struct {
	ID    int
	Value int
}

type Result struct {
	JobID  int
	Input  int
	Output int
}

func worker(id int, jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		// Sun'iy ish vaqti
		time.Sleep(time.Millisecond * time.Duration(100+rand.Intn(200)))
		output := job.Value * 2
		fmt.Printf("  [worker-%d] job=%d input=%d output=%d\n", id, job.ID, job.Value, output)
		results <- Result{JobID: job.ID, Input: job.Value, Output: output}
	}
}

func demoBasicWorkerPool() {
	rand.Seed(time.Now().UnixNano())

	jobs := make(chan Job)
	results := make(chan Result)

	workerCount := 3
	totalJobs := 10

	// Workerlarni ishga tushirish
	for w := 1; w <= workerCount; w++ {
		go worker(w, jobs, results)
	}

	// Ishlarni yuborish
	go func() {
		for i := 1; i <= totalJobs; i++ {
			jobs <- Job{ID: i, Value: i}
		}
		close(jobs)
	}()

	// Natijalarni o'qish
	for i := 0; i < totalJobs; i++ {
		res := <-results
		fmt.Printf("  [result] job=%d input=%d output=%d\n", res.JobID, res.Input, res.Output)
	}
}

// 2. Cancellable Worker Pool

func cancellableWorker(ctx context.Context, id int, jobs <-chan Job, results chan<- Result) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("  [worker-%d] stopping (context cancelled)\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("  [worker-%d] jobs channel closed\n", id)
				return
			}
			// Ishni bajarish
			output := job.Value * job.Value
			fmt.Printf("  [worker-%d] job=%d square(%d)=%d\n", id, job.ID, job.Value, output)
			results <- Result{JobID: job.ID, Input: job.Value, Output: output}
		}
	}
}

func demoCancellableWorkerPool() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	jobs := make(chan Job)
	results := make(chan Result)

	workerCount := 3
	totalJobs := 20

	for w := 1; w <= workerCount; w++ {
		go cancellableWorker(ctx, w, jobs, results)
	}

	go func() {
		for i := 1; i <= totalJobs; i++ {
			jobs <- Job{ID: i, Value: i}
		}
		close(jobs)
	}()

	// Natijalarni kontekst tugaguncha o'qish
	for {
		select {
		case <-ctx.Done():
			fmt.Println("  [main] context cancelled, stopping result collection")
			return
		case res := <-results:
			fmt.Printf("  [result] job=%d input=%d output=%d\n", res.JobID, res.Input, res.Output)
		}
	}
}


// 3. Pipeline


func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func filterEven(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			if n%2 == 0 {
				out <- n
			}
		}
	}()
	return out
}

func demoPipeline() {
	nums := gen(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	sq := square(nums)
	evens := filterEven(sq)

	for v := range evens {
		fmt.Printf("  pipeline output: %d\n", v)
	}
}


// 4. Rate Limiting (Ticker)


func demoRateLimiter() {
	requests := []int{1, 2, 3, 4, 5}
	limiter := time.NewTicker(300 * time.Millisecond)
	defer limiter.Stop()

	for _, req := range requests {
		<-limiter.C // Har iteratsiyada 300ms kutish
		fmt.Printf("  handling request %d at %s\n", req, time.Now().Format("15:04:05.000"))
	}
}

