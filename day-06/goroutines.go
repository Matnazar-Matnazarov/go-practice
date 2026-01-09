package main

import (
	"fmt"
	"time"
)

// demonstrateGoroutineBasic - goroutine asoslari
func demonstrateGoroutineBasic() {
	fmt.Println("Goroutine asoslari:")

	// Oddiy goroutine
	go func() {
		fmt.Println("  Goroutine 1: Salom, Go!")
	}()

	// Funksiya goroutine
	go sayHello("Goroutine 2")

	// Kutilish (goroutine ishlash uchun vaqt)
	time.Sleep(100 * time.Millisecond)
}

// sayHello - goroutine funksiyasi
func sayHello(name string) {
	fmt.Printf("  %s: Salom, Go!\n", name)
}

// demonstrateGoroutineMultiple - bir nechta goroutine
func demonstrateGoroutineMultiple() {
	fmt.Println("Bir nechta goroutine:")

	for i := 1; i <= 5; i++ {
		go func(id int) {
			fmt.Printf("  Goroutine %d ishlayapti\n", id)
		}(i)
	}

	time.Sleep(200 * time.Millisecond)
}

// demonstrateGoroutineWithDelay - kechikish bilan goroutine
func demonstrateGoroutineWithDelay() {
	fmt.Println("Kechikish bilan goroutine:")

	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Printf("  Goroutine: %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	fmt.Println("  Asosiy dastur ishlayapti...")
	time.Sleep(500 * time.Millisecond)
}

// demonstrateGoroutineSequential - ketma-ket goroutine
func demonstrateGoroutineSequential() {
	fmt.Println("Ketma-ket goroutine (sync.WaitGroup simulatsiyasi):")

	done := make(chan bool, 3)

	go func() {
		fmt.Println("  Goroutine 1 bajarildi")
		done <- true
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		fmt.Println("  Goroutine 2 bajarildi")
		done <- true
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("  Goroutine 3 bajarildi")
		done <- true
	}()

	// Barcha goroutine larni kutish
	for i := 0; i < 3; i++ {
		<-done
	}
	fmt.Println("  Barcha goroutine lar yakunlandi")
}

// demonstrateGoroutineWithChannel - channel bilan goroutine
func demonstrateGoroutineWithChannel() {
	fmt.Println("Channel bilan goroutine:")

	// Channel yaratish
	ch := make(chan string)

	// Goroutine dan ma'lumot yuborish
	go func() {
		ch <- "Goroutine dan xabar"
	}()

	// Ma'lumot olish
	message := <-ch
	fmt.Printf("  Qabul qilindi: %s\n", message)
}

// demonstrateGoroutineChannelMultiple - bir nechta xabar
func demonstrateGoroutineChannelMultiple() {
	fmt.Println("Channel bilan bir nechta xabar:")

	ch := make(chan int, 5) // Buffered channel

	// Xabarlar yuborish
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			fmt.Printf("  Yuborildi: %d\n", i)
		}
		close(ch)
	}()

	// Xabarlarni olish
	fmt.Println("  Xabarlarni qabul qilish:")
	for value := range ch {
		fmt.Printf("  Qabul qilindi: %d\n", value)
	}
}

// demonstrateGoroutineWorkerPool - worker pool (asosiy)
func demonstrateGoroutineWorkerPool() {
	fmt.Println("Worker pool (asosiy):")

	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Worker goroutine
	worker := func(id int, jobs <-chan int, results chan<- int) {
		for job := range jobs {
			fmt.Printf("  Worker %d: Job %d bajarilmoqda\n", id, job)
			time.Sleep(50 * time.Millisecond)
			results <- job * 2
		}
	}

	// 3 ta worker yaratish
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Job lar yuborish
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Natijalarni olish
	for r := 1; r <= 5; r++ {
		result := <-results
		fmt.Printf("  Natija: %d\n", result)
	}
}

// demonstrateGoroutineSelect - select statement
func demonstrateGoroutineSelect() {
	fmt.Println("Select statement:")

	ch1 := make(chan string)
	ch2 := make(chan string)

	// Birinchi channel
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Birinchi channel"
	}()

	// Ikkinchi channel
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch2 <- "Ikkinchi channel"
	}()

	// Select - birinchi tayyor bo'lganni olish
	select {
	case msg1 := <-ch1:
		fmt.Printf("  %s\n", msg1)
	case msg2 := <-ch2:
		fmt.Printf("  %s\n", msg2)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("  Timeout")
	}
}

// demonstrateGoroutineTimeout - timeout bilan goroutine
func demonstrateGoroutineTimeout() {
	fmt.Println("Timeout bilan goroutine:")

	ch := make(chan string)

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch <- "Ma'lumot"
	}()

	select {
	case msg := <-ch:
		fmt.Printf("  Qabul qilindi: %s\n", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("  Timeout: Ma'lumot kelmadi")
	}
}

// demonstrateGoroutineNonBlocking - non-blocking channel
func demonstrateGoroutineNonBlocking() {
	fmt.Println("Non-blocking channel:")

	ch := make(chan string)

	// Non-blocking send
	select {
	case ch <- "Xabar":
		fmt.Println("  Xabar yuborildi")
	default:
		fmt.Println("  Channel to'ldi, xabar yuborilmadi")
	}

	// Non-blocking receive
	select {
	case msg := <-ch:
		fmt.Printf("  Xabar qabul qilindi: %s\n", msg)
	default:
		fmt.Println("  Xabar yo'q")
	}
}

// demonstrateGoroutineRaceCondition - race condition (asosiy)
func demonstrateGoroutineRaceCondition() {
	fmt.Println("Race condition (asosiy):")

	counter := 0

	// Bir nechta goroutine counter ni o'zgartirish
	for i := 0; i < 5; i++ {
		go func() {
			counter++
			fmt.Printf("  Counter: %d\n", counter)
		}()
	}

	time.Sleep(200 * time.Millisecond)
	fmt.Printf("  Yakuniy counter: %d\n", counter)
}

// demonstrateGoroutineWithMutex - mutex bilan (asosiy)
func demonstrateGoroutineWithMutex() {
	fmt.Println("Mutex bilan (asosiy):")
	fmt.Println("  Note: Mutex keyingi kunlarda batafsil o'rganiladi")
	fmt.Println("  Hozircha channel va select ishlatiladi")
}

// demonstrateGoroutinePipeline - pipeline pattern
func demonstrateGoroutinePipeline() {
	fmt.Println("Pipeline pattern:")

	// Stage 1: Sonlar yuborish
	numbers := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
		close(numbers)
	}()

	// Stage 2: Ikki ga ko'paytirish
	doubled := make(chan int)
	go func() {
		for num := range numbers {
			doubled <- num * 2
		}
		close(doubled)
	}()

	// Stage 3: Natijalarni chiqarish
	fmt.Println("  Pipeline natijalari:")
	for result := range doubled {
		fmt.Printf("  %d\n", result)
	}
}

// demonstrateGoroutineFanOut - fan-out pattern (asosiy)
func demonstrateGoroutineFanOut() {
	fmt.Println("Fan-out pattern (asosiy):")

	input := make(chan int, 5)
	output1 := make(chan int, 5)
	output2 := make(chan int, 5)

	// Input yuborish
	go func() {
		for i := 1; i <= 5; i++ {
			input <- i
		}
		close(input)
	}()

	// Fan-out: ikkita worker
	go func() {
		for num := range input {
			output1 <- num * 2
		}
		close(output1)
	}()

	go func() {
		for num := range input {
			output2 <- num * 3
		}
		close(output2)
	}()

	// Natijalarni olish
	fmt.Println("  Worker 1 natijalari:")
	for result := range output1 {
		fmt.Printf("    %d\n", result)
	}

	fmt.Println("  Worker 2 natijalari:")
	for result := range output2 {
		fmt.Printf("    %d\n", result)
	}
}

// demonstrateGoroutineFanIn - fan-in pattern (asosiy)
func demonstrateGoroutineFanIn() {
	fmt.Println("Fan-in pattern (asosiy):")

	ch1 := make(chan int, 3)
	ch2 := make(chan int, 3)
	merged := make(chan int, 6)

	// Birinchi channel
	go func() {
		for i := 1; i <= 3; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	// Ikkinchi channel
	go func() {
		for i := 4; i <= 6; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	// Fan-in: ikkita channel ni birlashtirish
	go func() {
		for {
			select {
			case val, ok := <-ch1:
				if ok {
					merged <- val
				} else {
					ch1 = nil
				}
			case val, ok := <-ch2:
				if ok {
					merged <- val
				} else {
					ch2 = nil
				}
			}

			if ch1 == nil && ch2 == nil {
				close(merged)
				break
			}
		}
	}()

	// Natijalarni olish
	fmt.Println("  Birlashtirilgan natijalar:")
	for result := range merged {
		fmt.Printf("    %d\n", result)
	}
}

// demonstrateGoroutineContext - context (asosiy)
func demonstrateGoroutineContext() {
	fmt.Println("Context (asosiy):")
	fmt.Println("  Note: Context keyingi kunlarda batafsil o'rganiladi")
	fmt.Println("  Hozircha timeout va cancel pattern larni ko'rdik")
}

// demonstrateGoroutinePractical - amaliy misol: parallel hisoblash
func demonstrateGoroutinePractical() {
	fmt.Println("Amaliy misol: Parallel hisoblash")

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	results := make(chan int, len(numbers))

	// Har bir sonni parallel hisoblash
	for _, num := range numbers {
		go func(n int) {
			square := n * n
			results <- square
		}(num)
	}

	// Natijalarni yig'ish
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += <-results
	}

	fmt.Printf("  Sonlar: %v\n", numbers)
	fmt.Printf("  Kvadratlar yig'indisi: %d\n", sum)
}

// demonstrateGoroutineConcurrentIO - amaliy misol: concurrent I/O
func demonstrateGoroutineConcurrentIO() {
	fmt.Println("Amaliy misol: Concurrent I/O simulatsiyasi")

	tasks := []string{"File 1", "File 2", "File 3"}
	results := make(chan string, len(tasks))

	// Har bir faylni parallel o'qish
	for _, task := range tasks {
		go func(filename string) {
			// I/O simulatsiyasi
			time.Sleep(100 * time.Millisecond)
			results <- fmt.Sprintf("%s o'qildi", filename)
		}(task)
	}

	// Natijalarni olish
	for i := 0; i < len(tasks); i++ {
		result := <-results
		fmt.Printf("  %s\n", result)
	}
}
