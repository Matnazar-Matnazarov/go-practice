package main

import "fmt"

// main funksiyasi - dastur boshlanish nuqtasi
func main() {
	fmt.Println("Kun 7: WaitGroup va Channel (Kengaytirilgan)")
	fmt.Println()

	
	// 1. WAITGROUP
	
	fmt.Println("=== 1. WaitGroup ===")
	fmt.Println()

	// WaitGroup asoslari
	fmt.Println("--- WaitGroup asoslari ---")
	demonstrateWaitGroupBasic()
	fmt.Println()

	// Bir nechta goroutine
	fmt.Println("--- Bir nechta goroutine ---")
	demonstrateWaitGroupMultiple()
	fmt.Println()

	// Kechikish bilan
	fmt.Println("--- Kechikish bilan WaitGroup ---")
	demonstrateWaitGroupWithDelay()
	fmt.Println()

	// Counter bilan
	fmt.Println("--- Counter bilan WaitGroup ---")
	demonstrateWaitGroupCounter()
	fmt.Println()

	// Worker pool
	fmt.Println("--- Worker pool bilan WaitGroup ---")
	demonstrateWaitGroupWorkerPool()
	fmt.Println()

	// Xato bilan ishlash
	fmt.Println("--- Xato bilan ishlash ---")
	demonstrateWaitGroupErrorHandling()
	fmt.Println()

	// Parallel qayta ishlash
	fmt.Println("--- Parallel qayta ishlash ---")
	demonstrateWaitGroupParallelProcessing()
	fmt.Println()

	// Ketma-ket vazifalar
	fmt.Println("--- Ketma-ket vazifalar ---")
	demonstrateWaitGroupSequentialTasks()
	fmt.Println()

	// Timeout bilan
	fmt.Println("--- Timeout bilan WaitGroup ---")
	demonstrateWaitGroupTimeout()
	fmt.Println()

	// Ichki WaitGroup
	fmt.Println("--- Ichki WaitGroup ---")
	demonstrateWaitGroupNested()
	fmt.Println()

	// Amaliy misol
	fmt.Println("--- Amaliy misol: Fayllarni parallel o'qish ---")
	demonstrateWaitGroupPractical()
	fmt.Println()

	
	// 2. CHANNEL (KENGAYTIRILGAN)
	
	fmt.Println("=== 2. Channel (Kengaytirilgan) ===")
	fmt.Println()

	// Channel asoslari
	fmt.Println("--- Channel asoslari ---")
	demonstrateChannelBasic()
	fmt.Println()

	// Buffered channel
	fmt.Println("--- Buffered channel ---")
	demonstrateChannelBuffered()
	fmt.Println()

	// Channel yo'nalishi
	fmt.Println("--- Channel yo'nalishi ---")
	demonstrateChannelDirection()
	fmt.Println()

	// Channel yopish
	fmt.Println("--- Channel yopish ---")
	demonstrateChannelClose()
	fmt.Println()

	// Select statement
	fmt.Println("--- Select statement ---")
	demonstrateChannelSelect()
	fmt.Println()

	// Bir nechta select
	fmt.Println("--- Bir nechta select ---")
	demonstrateChannelSelectMultiple()
	fmt.Println()

	// Timeout
	fmt.Println("--- Channel timeout ---")
	demonstrateChannelTimeout()
	fmt.Println()

	// Non-blocking
	fmt.Println("--- Non-blocking channel ---")
	demonstrateChannelNonBlocking()
	fmt.Println()

	// Pipeline pattern
	fmt.Println("--- Pipeline pattern ---")
	demonstrateChannelPipeline()
	fmt.Println()

	// Fan-out pattern
	fmt.Println("--- Fan-out pattern ---")
	demonstrateChannelFanOut()
	fmt.Println()

	// Fan-in pattern
	fmt.Println("--- Fan-in pattern ---")
	demonstrateChannelFanIn()
	fmt.Println()

	// Worker pool
	fmt.Println("--- Worker pool ---")
	demonstrateChannelWorkerPool()
	fmt.Println()

	// Ticker channel
	fmt.Println("--- Ticker channel ---")
	demonstrateChannelTicker()
	fmt.Println()

	// Timer channel
	fmt.Println("--- Timer channel ---")
	demonstrateChannelTimer()
	fmt.Println()

	// time.After
	fmt.Println("--- time.After channel ---")
	demonstrateChannelAfter()
	fmt.Println()

	// Context pattern
	fmt.Println("--- Context pattern (asosiy) ---")
	demonstrateChannelContext()
	fmt.Println()

	// Xato bilan ishlash
	fmt.Println("--- Xato bilan ishlash ---")
	demonstrateChannelErrorHandling()
	fmt.Println()

	// Amaliy misol
	fmt.Println("--- Amaliy misol: Parallel hisoblash ---")
	demonstrateChannelPractical()
	fmt.Println()

	
	// 3. YAKUNIY XULOSA
	
	fmt.Println("=== Kun 7 yakunlandi! ===")
	fmt.Println("O'rganildi:")
	fmt.Println("  ✓ WaitGroup - goroutine larni kutish")
	fmt.Println("  ✓ WaitGroup operatsiyalari (Add, Done, Wait)")
	fmt.Println("  ✓ WaitGroup bilan worker pool, parallel processing")
	fmt.Println("  ✓ WaitGroup bilan xato bilan ishlash")
	fmt.Println("  ✓ Channel - buffered va unbuffered")
	fmt.Println("  ✓ Channel yo'nalishi (send-only, receive-only)")
	fmt.Println("  ✓ Channel yopish va range bilan o'qish")
	fmt.Println("  ✓ Select statement - bir nechta channel")
	fmt.Println("  ✓ Timeout va non-blocking operatsiyalar")
	fmt.Println("  ✓ Ticker va Timer channel lar")
	fmt.Println("  ✓ Pipeline, Fan-out, Fan-in pattern lar")
	fmt.Println("  ✓ Worker pool va amaliy misollar")
}
