package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// demonstrateAtomicBasic - Atomic operatsiyalar asoslari
func demonstrateAtomicBasic() {
	fmt.Println("Atomic operatsiyalar asoslari:")

	var counter int64

	// Atomic increment
	atomic.AddInt64(&counter, 1)
	fmt.Printf("  Counter: %d\n", atomic.LoadInt64(&counter))

	// Atomic decrement
	atomic.AddInt64(&counter, -1)
	fmt.Printf("  Counter: %d\n", atomic.LoadInt64(&counter))

	// Atomic store
	atomic.StoreInt64(&counter, 100)
	fmt.Printf("  Counter: %d\n", atomic.LoadInt64(&counter))
}

// demonstrateAtomicIncrement - Atomic increment
func demonstrateAtomicIncrement() {
	fmt.Println("Atomic increment:")

	var counter int64
	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Printf("  Yakuniy counter: %d\n", atomic.LoadInt64(&counter))
}

// demonstrateAtomicCompareAndSwap - Compare and Swap
func demonstrateAtomicCompareAndSwap() {
	fmt.Println("Atomic Compare and Swap:")

	var value int64 = 10

	// Compare and Swap: value == 10 bo'lsa, 20 ga o'zgartir
	swapped := atomic.CompareAndSwapInt64(&value, 10, 20)
	fmt.Printf("  Swapped: %t, Value: %d\n", swapped, atomic.LoadInt64(&value))

	// Yana urinib ko'rish (value endi 20)
	swapped = atomic.CompareAndSwapInt64(&value, 10, 30)
	fmt.Printf("  Swapped: %t, Value: %d\n", swapped, atomic.LoadInt64(&value))
}

// demonstrateAtomicSwap - Swap operatsiyasi
func demonstrateAtomicSwap() {
	fmt.Println("Atomic Swap:")

	var value int64 = 100

	// Eski qiymatni olish va yangi qiymatni o'rnatish
	oldValue := atomic.SwapInt64(&value, 200)
	fmt.Printf("  Eski qiymat: %d\n", oldValue)
	fmt.Printf("  Yangi qiymat: %d\n", atomic.LoadInt64(&value))
}

// demonstrateAtomicLoadStore - Load va Store
func demonstrateAtomicLoadStore() {
	fmt.Println("Atomic Load va Store:")

	var value int64

	// Store
	atomic.StoreInt64(&value, 42)
	fmt.Printf("  Store: %d\n", atomic.LoadInt64(&value))

	// Load
	loaded := atomic.LoadInt64(&value)
	fmt.Printf("  Load: %d\n", loaded)
}

// demonstrateAtomicPointer - Pointer operatsiyalari
func demonstrateAtomicPointer() {
	fmt.Println("Atomic Pointer operatsiyalari:")

	var ptr atomic.Value
	var data = "Ma'lumot"

	// Store pointer
	ptr.Store(data)
	fmt.Printf("  Stored: %v\n", ptr.Load())

	// Load pointer
	loaded := ptr.Load()
	fmt.Printf("  Loaded: %v\n", loaded)
}

// demonstrateAtomicCounter - Counter misoli
func demonstrateAtomicCounter() {
	fmt.Println("Atomic Counter misoli:")

	type Counter struct {
		value int64
	}

	counter := &Counter{}

	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter.value, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Printf("  Yakuniy counter: %d\n", atomic.LoadInt64(&counter.value))
}

// demonstrateAtomicMutexComparison - Mutex bilan taqqoslash
func demonstrateAtomicMutexComparison() {
	fmt.Println("Atomic vs Mutex taqqoslash:")

	// Atomic
	var atomicCounter int64
	start := time.Now()
	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				atomic.AddInt64(&atomicCounter, 1)
			}
		}()
	}
	wg.Wait()
	atomicTime := time.Since(start)

	// Mutex
	var mutexCounter int64
	var mu sync.Mutex
	start = time.Now()

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				mu.Lock()
				mutexCounter++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	mutexTime := time.Since(start)

	fmt.Printf("  Atomic counter: %d, Vaqt: %v\n", atomic.LoadInt64(&atomicCounter), atomicTime)
	fmt.Printf("  Mutex counter: %d, Vaqt: %v\n", mutexCounter, mutexTime)
}

// demonstrateAtomicBool - Bool operatsiyalari
func demonstrateAtomicBool() {
	fmt.Println("Atomic Bool operatsiyalari:")

	var flag int32 // Go da atomic bool yo'q, int32 ishlatiladi

	// Set true (1)
	atomic.StoreInt32(&flag, 1)
	fmt.Printf("  Flag: %t\n", atomic.LoadInt32(&flag) == 1)

	// Set false (0)
	atomic.StoreInt32(&flag, 0)
	fmt.Printf("  Flag: %t\n", atomic.LoadInt32(&flag) == 1)
}

// demonstrateAtomicUint32 - Uint32 operatsiyalari
func demonstrateAtomicUint32() {
	fmt.Println("Atomic Uint32 operatsiyalari:")

	var value uint32

	atomic.StoreUint32(&value, 100)
	fmt.Printf("  Value: %d\n", atomic.LoadUint32(&value))

	atomic.AddUint32(&value, 50)
	fmt.Printf("  Value: %d\n", atomic.LoadUint32(&value))
}

// demonstrateAtomicPractical - Amaliy misol: Rate limiter
func demonstrateAtomicPractical() {
	fmt.Println("Amaliy misol: Rate limiter (simulatsiya):")

	type RateLimiter struct {
		count int64
		limit int64
	}

	limiter := &RateLimiter{limit: 100}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 15; j++ {
				current := atomic.AddInt64(&limiter.count, 1)
				if current <= limiter.limit {
					fmt.Printf("  Request %d-%d: Muvaffaqiyatli (count: %d)\n", id, j, current)
				} else {
					atomic.AddInt64(&limiter.count, -1)
					fmt.Printf("  Request %d-%d: Limit oshib ketdi\n", id, j)
				}
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("  Yakuniy count: %d\n", atomic.LoadInt64(&limiter.count))
}
