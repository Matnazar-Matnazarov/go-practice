package main

import (
	"fmt"
)

// main funksiyasi - dastur boshlanish nuqtasi
func main() {
	fmt.Println("Kun 6: Queue (Navbat) va Goroutine (Parallel dasturlash)")
	fmt.Println()

	
	// 1. QUEUE (NAVBAT) MA'LUMOT TUZILMASI
	
	fmt.Println("=== 1. Queue (Navbat) ===")
	fmt.Println()

	// Queue asoslari
	fmt.Println("--- Queue asoslari ---")
	demonstrateQueueBasic()
	fmt.Println()

	// Queue operatsiyalari
	fmt.Println("--- Queue operatsiyalari ---")
	demonstrateQueueOperations()
	fmt.Println()

	// Front element
	fmt.Println("--- Front element ---")
	demonstrateQueueFront()
	fmt.Println()

	// String queue
	fmt.Println("--- String queue ---")
	demonstrateStringQueue()
	fmt.Println()

	// Generic queue
	fmt.Println("--- Generic queue ---")
	demonstrateGenericQueue()
	fmt.Println()

	// Circular queue
	fmt.Println("--- Circular queue (simulatsiya) ---")
	demonstrateQueueCircular()
	fmt.Println()

	// Priority queue
	fmt.Println("--- Priority queue (simulatsiya) ---")
	demonstrateQueuePriority()
	fmt.Println()

	// Amaliy misollar
	fmt.Println("--- Amaliy misollar ---")

	// Vazifa rejalashtirish
	fmt.Println("Vazifa rejalashtirish:")
	demonstrateQueueTaskScheduler()
	fmt.Println()

	// BFS simulatsiyasi
	fmt.Println("BFS simulatsiyasi:")
	demonstrateQueueBFS()
	fmt.Println()

	
	// 2. GOROUTINE (PARALLEL DASTURLASH)
	
	fmt.Println("=== 2. Goroutine (Parallel dasturlash) ===")
	fmt.Println()

	// Goroutine asoslari
	fmt.Println("--- Goroutine asoslari ---")
	demonstrateGoroutineBasic()
	fmt.Println()

	// Bir nechta goroutine
	fmt.Println("--- Bir nechta goroutine ---")
	demonstrateGoroutineMultiple()
	fmt.Println()

	// Kechikish bilan
	fmt.Println("--- Kechikish bilan goroutine ---")
	demonstrateGoroutineWithDelay()
	fmt.Println()

	// Ketma-ket goroutine
	fmt.Println("--- Ketma-ket goroutine (channel bilan) ---")
	demonstrateGoroutineSequential()
	fmt.Println()

	// Channel bilan goroutine
	fmt.Println("--- Channel bilan goroutine ---")
	demonstrateGoroutineWithChannel()
	fmt.Println()

	// Bir nechta xabar
	fmt.Println("--- Channel bilan bir nechta xabar ---")
	demonstrateGoroutineChannelMultiple()
	fmt.Println()

	// Worker pool
	fmt.Println("--- Worker pool (asosiy) ---")
	demonstrateGoroutineWorkerPool()
	fmt.Println()

	// Select statement
	fmt.Println("--- Select statement ---")
	demonstrateGoroutineSelect()
	fmt.Println()

	// Timeout
	fmt.Println("--- Timeout bilan goroutine ---")
	demonstrateGoroutineTimeout()
	fmt.Println()

	// Non-blocking channel
	fmt.Println("--- Non-blocking channel ---")
	demonstrateGoroutineNonBlocking()
	fmt.Println()

	// Race condition (asosiy)
	fmt.Println("--- Race condition (asosiy) ---")
	demonstrateGoroutineRaceCondition()
	fmt.Println()

	// Mutex (asosiy)
	fmt.Println("--- Mutex (asosiy) ---")
	demonstrateGoroutineWithMutex()
	fmt.Println()

	// Pipeline pattern
	fmt.Println("--- Pipeline pattern ---")
	demonstrateGoroutinePipeline()
	fmt.Println()

	// Fan-out pattern
	fmt.Println("--- Fan-out pattern (asosiy) ---")
	demonstrateGoroutineFanOut()
	fmt.Println()

	// Fan-in pattern
	fmt.Println("--- Fan-in pattern (asosiy) ---")
	demonstrateGoroutineFanIn()
	fmt.Println()

	// Context (asosiy)
	fmt.Println("--- Context (asosiy) ---")
	demonstrateGoroutineContext()
	fmt.Println()

	// Amaliy misollar
	fmt.Println("--- Amaliy misollar ---")

	// Parallel hisoblash
	fmt.Println("Parallel hisoblash:")
	demonstrateGoroutinePractical()
	fmt.Println()

	// Concurrent I/O
	fmt.Println("Concurrent I/O simulatsiyasi:")
	demonstrateGoroutineConcurrentIO()
	fmt.Println()

	
	// 3. YAKUNIY XULOSA
	
	fmt.Println("=== Kun 6 yakunlandi! ===")
	fmt.Println("O'rganildi:")
	fmt.Println("  ✓ Queue (navbat) - FIFO ma'lumot tuzilmasi")
	fmt.Println("  ✓ Queue operatsiyalari (enqueue, dequeue, front)")
	fmt.Println("  ✓ String va Generic queue")
	fmt.Println("  ✓ Circular va Priority queue (asosiy)")
	fmt.Println("  ✓ Goroutine - parallel dasturlash")
	fmt.Println("  ✓ Channel - goroutine lar orasida aloqa")
	fmt.Println("  ✓ Select statement - bir nechta channel")
	fmt.Println("  ✓ Worker pool, Pipeline, Fan-out, Fan-in pattern lar")
	fmt.Println("  ✓ Timeout va non-blocking operatsiyalar")
	fmt.Println("  ✓ Amaliy misollar va pattern lar")
}
