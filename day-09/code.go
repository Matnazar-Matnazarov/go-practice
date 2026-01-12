package main

import "fmt"

// main funksiyasi - dastur boshlanish nuqtasi
func main() {
	fmt.Println("Kun 9: Atomic Operations va Error Handling")
	fmt.Println()

	// 1. ATOMIC OPERATIONS

	fmt.Println("=== 1. Atomic Operations ===")
	fmt.Println()

	// Atomic asoslari
	fmt.Println("--- Atomic operatsiyalar asoslari ---")
	demonstrateAtomicBasic()
	fmt.Println()

	// Atomic increment
	fmt.Println("--- Atomic increment ---")
	demonstrateAtomicIncrement()
	fmt.Println()

	// Compare and Swap
	fmt.Println("--- Atomic Compare and Swap ---")
	demonstrateAtomicCompareAndSwap()
	fmt.Println()

	// Swap
	fmt.Println("--- Atomic Swap ---")
	demonstrateAtomicSwap()
	fmt.Println()

	// Load va Store
	fmt.Println("--- Atomic Load va Store ---")
	demonstrateAtomicLoadStore()
	fmt.Println()

	// Pointer operatsiyalari
	fmt.Println("--- Atomic Pointer operatsiyalari ---")
	demonstrateAtomicPointer()
	fmt.Println()

	// Counter misoli
	fmt.Println("--- Atomic Counter misoli ---")
	demonstrateAtomicCounter()
	fmt.Println()

	// Mutex bilan taqqoslash
	fmt.Println("--- Atomic vs Mutex taqqoslash ---")
	demonstrateAtomicMutexComparison()
	fmt.Println()

	// Bool operatsiyalari
	fmt.Println("--- Atomic Bool operatsiyalari ---")
	demonstrateAtomicBool()
	fmt.Println()

	// Uint32 operatsiyalari
	fmt.Println("--- Atomic Uint32 operatsiyalari ---")
	demonstrateAtomicUint32()
	fmt.Println()

	// Amaliy misol
	fmt.Println("--- Amaliy misol: Rate limiter ---")
	demonstrateAtomicPractical()
	fmt.Println()

	// 2. ERROR HANDLING

	fmt.Println("=== 2. Error Handling ===")
	fmt.Println()

	// Error asoslari
	fmt.Println("--- Error asoslari ---")
	demonstrateErrorBasic()
	fmt.Println()

	// Goroutine da xato
	fmt.Println("--- Goroutine da xato ---")
	demonstrateErrorInGoroutine()
	fmt.Println()

	// Bir nechta goroutine da xato
	fmt.Println("--- Bir nechta goroutine da xato ---")
	demonstrateErrorMultipleGoroutines()
	fmt.Println()

	// Natija va xato bilan
	fmt.Println("--- Natija va xato bilan ---")
	demonstrateErrorWithResult()
	fmt.Println()

	// Error wrapping
	fmt.Println("--- Error wrapping ---")
	demonstrateErrorWrapping()
	fmt.Println()

	// errors.Is
	fmt.Println("--- errors.Is ---")
	demonstrateErrorIs()
	fmt.Println()

	// errors.As
	fmt.Println("--- errors.As ---")
	demonstrateErrorAs()
	fmt.Println()

	// Timeout bilan xato
	fmt.Println("--- Timeout bilan xato ---")
	demonstrateErrorTimeout()
	fmt.Println()

	// Context bilan xato
	fmt.Println("--- Context bilan xato ---")
	demonstrateErrorContext()
	fmt.Println()

	// Panic recovery
	fmt.Println("--- Panic recovery ---")
	demonstrateErrorRecovery()
	fmt.Println()

	// Amaliy misol
	fmt.Println("--- Amaliy misol: API client ---")
	demonstrateErrorPractical()
	fmt.Println()

	// 3. YAKUNIY XULOSA

	fmt.Println("=== Kun 9 yakunlandi! ===")
	fmt.Println("O'rganildi:")
	fmt.Println("  ✓ Atomic operations - race condition siz operatsiyalar")
	fmt.Println("  ✓ Atomic operatsiyalar (Add, Load, Store, Swap, CompareAndSwap)")
	fmt.Println("  ✓ Atomic vs Mutex taqqoslash")
	fmt.Println("  ✓ Error handling - xato bilan ishlash")
	fmt.Println("  ✓ Error wrapping va unwrapping")
	fmt.Println("  ✓ errors.Is va errors.As")
	fmt.Println("  ✓ Goroutine larda xato bilan ishlash")
	fmt.Println("  ✓ Panic recovery")
	fmt.Println("  ✓ Context bilan xato bilan ishlash")
	fmt.Println("  ✓ Amaliy misollar: Rate limiter, API client")
}
