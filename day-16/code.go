package main

import "fmt"

func main() {
	fmt.Println("Kun 16: Advanced Concurrency Patterns")
	fmt.Println()

	fmt.Println("=== 1. Bounded concurrency (semaphore) ===")
	fmt.Println()
	demoBoundedConcurrency()
	fmt.Println()

	fmt.Println("=== 2. Context cancellation (fan-out work) ===")
	fmt.Println()
	demoContextCancellation()
	fmt.Println()

	fmt.Println("=== 3. Goroutine leak prevention (done channel) ===")
	fmt.Println()
	demoLeakPrevention()
	fmt.Println()

	fmt.Println("=== 4. Graceful shutdown (signal + context) ===")
	fmt.Println()
	demoGracefulShutdownConcept()
	fmt.Println()

	fmt.Println("=== Kun 16 yakunlandi! ===")
}
