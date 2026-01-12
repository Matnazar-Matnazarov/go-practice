package main

import (
	"fmt"
	"sync"
	"time"
)

// demonstrateMutexBasic - Mutex asoslari
func demonstrateMutexBasic() {
	fmt.Println("Mutex asoslari:")

	var mu sync.Mutex
	var counter int

	// Mutex bilan lock/unlock
	mu.Lock()
	counter++
	fmt.Printf("  Counter: %d\n", counter)
	mu.Unlock()
}

// demonstrateMutexRaceCondition - Race condition muammosi
func demonstrateMutexRaceCondition() {
	fmt.Println("Race condition muammosi:")

	var counter int
	var wg sync.WaitGroup

	// Mutex siz (race condition)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			counter++ // Race condition!
		}()
	}

	wg.Wait()
	fmt.Printf("  Counter (race condition): %d (noto'g'ri bo'lishi mumkin)\n", counter)
}

// demonstrateMutexFixRaceCondition - Mutex bilan race condition ni hal qilish
func demonstrateMutexFixRaceCondition() {
	fmt.Println("Mutex bilan race condition ni hal qilish:")

	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Printf("  Counter (mutex bilan): %d (to'g'ri)\n", counter)
}

// demonstrateMutexMultipleOperations - Bir nechta operatsiya
func demonstrateMutexMultipleOperations() {
	fmt.Println("Mutex bilan bir nechta operatsiya:")

	var mu sync.Mutex
	var data map[string]int
	data = make(map[string]int)
	var wg sync.WaitGroup

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(id int) {
			defer wg.Done()

			mu.Lock()
			key := fmt.Sprintf("key%d", id)
			data[key] = id * 2
			fmt.Printf("  %s = %d\n", key, data[key])
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Printf("  Map o'lchami: %d\n", len(data))
}

// demonstrateRWMutexBasic - RWMutex asoslari
func demonstrateRWMutexBasic() {
	fmt.Println("RWMutex asoslari:")

	var rwmu sync.RWMutex
	var data int
	var wg sync.WaitGroup

	// Writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		rwmu.Lock()
		data = 100
		fmt.Printf("  Writer: data = %d\n", data)
		rwmu.Unlock()
	}()

	// Reader lar
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(id int) {
			defer wg.Done()
			rwmu.RLock()
			value := data
			fmt.Printf("  Reader %d: data = %d\n", id, value)
			rwmu.RUnlock()
		}(i)
	}

	wg.Wait()
}

// demonstrateRWMutexMultipleReaders - Bir nechta reader
func demonstrateRWMutexMultipleReaders() {
	fmt.Println("RWMutex bilan bir nechta reader:")

	var rwmu sync.RWMutex
	var counter int
	var wg sync.WaitGroup

	// Writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		rwmu.Lock()
		counter = 50
		fmt.Println("  Writer: Counter ni yangiladi")
		time.Sleep(50 * time.Millisecond)
		rwmu.Unlock()
	}()

	// Reader lar (parallel o'qish)
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(id int) {
			defer wg.Done()
			rwmu.RLock()
			value := counter
			fmt.Printf("  Reader %d: Counter = %d\n", id, value)
			time.Sleep(10 * time.Millisecond)
			rwmu.RUnlock()
		}(i)
	}

	wg.Wait()
}

// demonstrateMutexDeadlock - Deadlock muammosi
func demonstrateMutexDeadlock() {
	fmt.Println("Mutex deadlock muammosi (ehtiyotkorlik bilan):")

	var mu1, mu2 sync.Mutex
	var wg sync.WaitGroup

	// Birinchi goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		mu1.Lock()
		fmt.Println("  Goroutine 1: mu1 lock qildi")
		time.Sleep(50 * time.Millisecond)
		mu2.Lock()
		fmt.Println("  Goroutine 1: mu2 lock qildi")
		mu2.Unlock()
		mu1.Unlock()
	}()

	// Ikkinchi goroutine (teskari tartib - deadlock!)
	wg.Add(1)
	go func() {
		defer wg.Done()
		mu2.Lock()
		fmt.Println("  Goroutine 2: mu2 lock qildi")
		time.Sleep(50 * time.Millisecond)
		// mu1.Lock() // Bu deadlock ga olib keladi
		// Shuning uchun biz buni o'tkazib yuboramiz
		fmt.Println("  Goroutine 2: mu1 lock qilishdan oldin to'xtatildi (deadlock oldini olish)")
		mu2.Unlock()
	}()

	wg.Wait()
	fmt.Println("  Deadlock oldini olish: Har doim bir xil tartibda lock qilish")
}

// demonstrateMutexSafeCounter - Xavfsiz counter
func demonstrateMutexSafeCounter() {
	fmt.Println("Mutex bilan xavfsiz counter:")

	type SafeCounter struct {
		mu    sync.Mutex
		value int
	}

	counter := &SafeCounter{}

	var wg sync.WaitGroup
	wg.Add(10)

	// Increment
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			counter.mu.Lock()
			counter.value++
			counter.mu.Unlock()
		}()
	}

	// Decrement
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			counter.mu.Lock()
			counter.value--
			counter.mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Printf("  Yakuniy qiymat: %d\n", counter.value)
}

// demonstrateRWMutexCache - Cache simulatsiyasi
func demonstrateRWMutexCache() {
	fmt.Println("RWMutex bilan cache simulatsiyasi:")

	type Cache struct {
		mu    sync.RWMutex
		items map[string]string
	}

	cache := &Cache{
		items: make(map[string]string),
	}

	var wg sync.WaitGroup

	// Writer lar
	wg.Add(2)
	go func() {
		defer wg.Done()
		cache.mu.Lock()
		cache.items["key1"] = "value1"
		fmt.Println("  Writer 1: key1 qo'shildi")
		cache.mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		cache.mu.Lock()
		cache.items["key2"] = "value2"
		fmt.Println("  Writer 2: key2 qo'shildi")
		cache.mu.Unlock()
	}()

	// Reader lar
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(id int) {
			defer wg.Done()
			cache.mu.RLock()
			value1 := cache.items["key1"]
			value2 := cache.items["key2"]
			fmt.Printf("  Reader %d: key1=%s, key2=%s\n", id, value1, value2)
			cache.mu.RUnlock()
		}(i)
	}

	wg.Wait()
}

// demonstrateMutexBankAccount - Bank hisob simulatsiyasi
func demonstrateMutexBankAccount() {
	fmt.Println("Mutex bilan bank hisob simulatsiyasi:")

	type BankAccount struct {
		mu      sync.Mutex
		balance int
	}

	account := &BankAccount{balance: 1000}

	var wg sync.WaitGroup

	// Deposit
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(amount int) {
			defer wg.Done()
			account.mu.Lock()
			account.balance += amount
			fmt.Printf("  Deposit: +%d, Balance: %d\n", amount, account.balance)
			account.mu.Unlock()
		}(100)
	}

	// Withdraw
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func(amount int) {
			defer wg.Done()
			account.mu.Lock()
			if account.balance >= amount {
				account.balance -= amount
				fmt.Printf("  Withdraw: -%d, Balance: %d\n", amount, account.balance)
			} else {
				fmt.Printf("  Withdraw: -%d, Noto'g'ri (balance yetarli emas)\n", amount)
			}
			account.mu.Unlock()
		}(150)
	}

	wg.Wait()
	fmt.Printf("  Yakuniy balance: %d\n", account.balance)
}

// demonstrateMutexOnce - sync.Once
func demonstrateMutexOnce() {
	fmt.Println("sync.Once (bir marta bajarish):")

	var once sync.Once
	var wg sync.WaitGroup

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(id int) {
			defer wg.Done()
			once.Do(func() {
				fmt.Printf("  Goroutine %d: Birinchi marta bajarildi\n", id)
			})
			fmt.Printf("  Goroutine %d: Bajarildi\n", id)
		}(i)
	}

	wg.Wait()
}

// demonstrateMutexCondition - Condition variable (asosiy)
func demonstrateMutexCondition() {
	fmt.Println("Condition variable (asosiy):")

	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	var ready bool
	var wg sync.WaitGroup

	// Waiter
	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		for !ready {
			fmt.Println("  Waiter: Kutmoqda...")
			cond.Wait()
		}
		fmt.Println("  Waiter: Tayyor!")
		mu.Unlock()
	}()

	// Signaler
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(200 * time.Millisecond)
		mu.Lock()
		ready = true
		fmt.Println("  Signaler: Tayyor bo'ldi")
		cond.Signal()
		mu.Unlock()
	}()

	wg.Wait()
}

// demonstrateMutexPractical - Amaliy misol: Counter service
func demonstrateMutexPractical() {
	fmt.Println("Amaliy misol: Counter service:")

	type CounterService struct {
		mu      sync.RWMutex
		counters map[string]int
	}

	service := &CounterService{
		counters: make(map[string]int),
	}

	var wg sync.WaitGroup

	// Increment
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(key string) {
			defer wg.Done()
			service.mu.Lock()
			service.counters[key]++
			fmt.Printf("  Increment: %s = %d\n", key, service.counters[key])
			service.mu.Unlock()
		}(fmt.Sprintf("counter%d", i%3))
	}

	// Read
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(key string) {
			defer wg.Done()
			service.mu.RLock()
			value := service.counters[key]
			fmt.Printf("  Read: %s = %d\n", key, value)
			service.mu.RUnlock()
		}(fmt.Sprintf("counter%d", i))
	}

	wg.Wait()
}
