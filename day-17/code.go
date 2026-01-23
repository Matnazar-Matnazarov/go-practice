package main

import "fmt"

func main() {
	fmt.Println("Kun 17: Concurrency Patterns II (Pipeline, Worker Pool, Rate Limit, ErrGroup)")
	fmt.Println()

	fmt.Println("=== 1. Pipeline (generator -> stages -> sink) ===")
	fmt.Println()
	demoPipeline()
	fmt.Println()

	fmt.Println("=== 2. Fan-out / Fan-in (worker pool) ===")
	fmt.Println()
	demoWorkerPool()
	fmt.Println()

	fmt.Println("=== 3. Rate limiting (token bucket) ===")
	fmt.Println()
	demoRateLimiting()
	fmt.Println()

	fmt.Println("=== 4. ErrGroup-style cancel on first error ===")
	fmt.Println()
	demoErrGroup()
	fmt.Println()

	fmt.Println("=== Kun 17 yakunlandi! ===")
}
