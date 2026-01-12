package main

import "fmt"

// main funksiyasi - dastur boshlanish nuqtasi
func main() {
	fmt.Println("Kun 8: Context va Mutex (Sinxronizatsiya)")
	fmt.Println()

	// 1. CONTEXT

	fmt.Println("=== 1. Context ===")
	fmt.Println()

	// Context asoslari
	fmt.Println("--- Context asoslari ---")
	demonstrateContextBasic()
	fmt.Println()

	// Context bilan qiymat
	fmt.Println("--- Context bilan qiymat ---")
	demonstrateContextWithValue()
	fmt.Println()

	// Context bilan cancel
	fmt.Println("--- Context bilan cancel ---")
	demonstrateContextWithCancel()
	fmt.Println()

	// Context bilan timeout
	fmt.Println("--- Context bilan timeout ---")
	demonstrateContextWithTimeout()
	fmt.Println()

	// Context bilan deadline
	fmt.Println("--- Context bilan deadline ---")
	demonstrateContextWithDeadline()
	fmt.Println()

	// Ichki context
	fmt.Println("--- Ichki context ---")
	demonstrateContextNested()
	fmt.Println()

	// Bir nechta goroutine
	fmt.Println("--- Context bilan bir nechta goroutine ---")
	demonstrateContextMultipleGoroutines()
	fmt.Println()

	// Timeout tarqalishi
	fmt.Println("--- Context timeout tarqalishi ---")
	demonstrateContextTimeoutPropagation()
	fmt.Println()

	// Qiymat tarqalishi
	fmt.Println("--- Context qiymat tarqalishi ---")
	demonstrateContextValuePropagation()
	fmt.Println()

	// HTTP request simulatsiyasi
	fmt.Println("--- HTTP request simulatsiyasi ---")
	demonstrateContextHTTPRequest()
	fmt.Println()

	// Database query simulatsiyasi
	fmt.Println("--- Database query simulatsiyasi ---")
	demonstrateContextDatabaseQuery()
	fmt.Println()

	// Worker pool bilan
	fmt.Println("--- Context bilan worker pool ---")
	demonstrateContextWorkerPool()
	fmt.Println()

	// Xato bilan ishlash
	fmt.Println("--- Context bilan xato bilan ishlash ---")
	demonstrateContextErrorHandling()
	fmt.Println()

	// Amaliy misol
	fmt.Println("--- Amaliy misol: API client ---")
	demonstrateContextPractical()
	fmt.Println()

	// 2. MUTEX

	fmt.Println("=== 2. Mutex (Sinxronizatsiya) ===")
	fmt.Println()

	// Mutex asoslari
	fmt.Println("--- Mutex asoslari ---")
	demonstrateMutexBasic()
	fmt.Println()

	// Race condition muammosi
	fmt.Println("--- Race condition muammosi ---")
	demonstrateMutexRaceCondition()
	fmt.Println()

	// Mutex bilan race condition ni hal qilish
	fmt.Println("--- Mutex bilan race condition ni hal qilish ---")
	demonstrateMutexFixRaceCondition()
	fmt.Println()

	// Bir nechta operatsiya
	fmt.Println("--- Mutex bilan bir nechta operatsiya ---")
	demonstrateMutexMultipleOperations()
	fmt.Println()

	// RWMutex asoslari
	fmt.Println("--- RWMutex asoslari ---")
	demonstrateRWMutexBasic()
	fmt.Println()

	// Bir nechta reader
	fmt.Println("--- RWMutex bilan bir nechta reader ---")
	demonstrateRWMutexMultipleReaders()
	fmt.Println()

	// Deadlock muammosi
	fmt.Println("--- Mutex deadlock muammosi ---")
	demonstrateMutexDeadlock()
	fmt.Println()

	// Xavfsiz counter
	fmt.Println("--- Mutex bilan xavfsiz counter ---")
	demonstrateMutexSafeCounter()
	fmt.Println()

	// Cache simulatsiyasi
	fmt.Println("--- RWMutex bilan cache simulatsiyasi ---")
	demonstrateRWMutexCache()
	fmt.Println()

	// Bank hisob simulatsiyasi
	fmt.Println("--- Mutex bilan bank hisob simulatsiyasi ---")
	demonstrateMutexBankAccount()
	fmt.Println()

	// sync.Once
	fmt.Println("--- sync.Once ---")
	demonstrateMutexOnce()
	fmt.Println()

	// Condition variable
	fmt.Println("--- Condition variable ---")
	demonstrateMutexCondition()
	fmt.Println()

	// Amaliy misol
	fmt.Println("--- Amaliy misol: Counter service ---")
	demonstrateMutexPractical()
	fmt.Println()

	// 3. YAKUNIY XULOSA

	fmt.Println("=== Kun 8 yakunlandi! ===")
	fmt.Println("O'rganildi:")
	fmt.Println("  ✓ Context - goroutine larni boshqarish")
	fmt.Println("  ✓ Context operatsiyalari (WithValue, WithCancel, WithTimeout, WithDeadline)")
	fmt.Println("  ✓ Context qiymat va timeout tarqalishi")
	fmt.Println("  ✓ Context bilan HTTP request va Database query")
	fmt.Println("  ✓ Mutex - race condition ni hal qilish")
	fmt.Println("  ✓ RWMutex - parallel o'qish va yozish")
	fmt.Println("  ✓ sync.Once - bir marta bajarish")
	fmt.Println("  ✓ Condition variable - shartli kutish")
	fmt.Println("  ✓ Amaliy misollar: Cache, Bank account, Counter service")
}
