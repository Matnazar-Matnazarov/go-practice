package main

import (
	"fmt"
	"time"
)

// demonstrateChannelBasic - channel asoslari
func demonstrateChannelBasic() {
	fmt.Println("Channel asoslari:")

	// Unbuffered channel
	ch := make(chan string)

	// Goroutine dan ma'lumot yuborish
	go func() {
		ch <- "Salom, Channel!"
	}()

	// Ma'lumot olish
	message := <-ch
	fmt.Printf("  Qabul qilindi: %s\n", message)
}

// demonstrateChannelBuffered - buffered channel
func demonstrateChannelBuffered() {
	fmt.Println("Buffered channel:")

	// 3 ta elementli buffered channel
	ch := make(chan int, 3)

	// Ma'lumotlarni yuborish (block qilmaydi)
	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println("  3 ta element yuborildi")

	// Ma'lumotlarni olish
	fmt.Println("  Ma'lumotlarni olish:")
	for i := 0; i < 3; i++ {
		value := <-ch
		fmt.Printf("    %d\n", value)
	}
}

// demonstrateChannelDirection - channel yo'nalishi
func demonstrateChannelDirection() {
	fmt.Println("Channel yo'nalishi:")

	// Send-only channel
	sendCh := make(chan<- int)

	// Receive-only channel
	receiveCh := make(<-chan int)

	// Funksiya parametrlari
	go sendOnly(sendCh)
	go receiveOnly(receiveCh)

	fmt.Println("  Send-only va receive-only channel lar yaratildi")
	time.Sleep(50 * time.Millisecond)
}

// sendOnly - faqat yuborish
func sendOnly(ch chan<- int) {
	ch <- 10
}

// receiveOnly - faqat qabul qilish
func receiveOnly(ch <-chan int) {
	// value := <-ch // Bu ishlamaydi, ch nil
	_ = ch
}

// demonstrateChannelClose - channel yopish
func demonstrateChannelClose() {
	fmt.Println("Channel yopish:")

	ch := make(chan int, 5)

	// Ma'lumotlar yuborish
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch) // Channel ni yopish
	}()

	// Ma'lumotlarni olish (range bilan)
	fmt.Println("  Ma'lumotlarni olish:")
	for value := range ch {
		fmt.Printf("    %d\n", value)
	}

	// Yopilgan channel dan olish
	value, ok := <-ch
	fmt.Printf("  Yopilgan channel: value=%d, ok=%t\n", value, ok)
}

// demonstrateChannelSelect - select statement
func demonstrateChannelSelect() {
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
	}
}

// demonstrateChannelSelectMultiple - bir nechta select
func demonstrateChannelSelectMultiple() {
	fmt.Println("Select bilan bir nechta channel:")

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// Ma'lumotlar yuborish
	go func() {
		ch1 <- 1
		ch2 <- 2
		ch3 <- 3
	}()

	// Bir nechta select
	for i := 0; i < 3; i++ {
		select {
		case v1 := <-ch1:
			fmt.Printf("  Channel 1: %d\n", v1)
		case v2 := <-ch2:
			fmt.Printf("  Channel 2: %d\n", v2)
		case v3 := <-ch3:
			fmt.Printf("  Channel 3: %d\n", v3)
		}
	}
}

// demonstrateChannelTimeout - timeout bilan
func demonstrateChannelTimeout() {
	fmt.Println("Channel timeout:")

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

// demonstrateChannelNonBlocking - non-blocking
func demonstrateChannelNonBlocking() {
	fmt.Println("Non-blocking channel:")

	ch := make(chan int)

	// Non-blocking send
	select {
	case ch <- 10:
		fmt.Println("  Ma'lumot yuborildi")
	default:
		fmt.Println("  Channel to'ldi, yuborilmadi")
	}

	// Non-blocking receive
	select {
	case value := <-ch:
		fmt.Printf("  Ma'lumot qabul qilindi: %d\n", value)
	default:
		fmt.Println("  Ma'lumot yo'q")
	}
}

// demonstrateChannelPipeline - pipeline pattern
func demonstrateChannelPipeline() {
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
		fmt.Printf("    %d\n", result)
	}
}

// demonstrateChannelFanOut - fan-out pattern
func demonstrateChannelFanOut() {
	fmt.Println("Fan-out pattern:")

	input := make(chan int, 10)
	output1 := make(chan int, 10)
	output2 := make(chan int, 10)

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

// demonstrateChannelFanIn - fan-in pattern
func demonstrateChannelFanIn() {
	fmt.Println("Fan-in pattern:")

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

// demonstrateChannelWorkerPool - worker pool
func demonstrateChannelWorkerPool() {
	fmt.Println("Worker pool:")

	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// Job lar yuborish
	go func() {
		for j := 1; j <= 10; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	// 3 ta worker
	workers := 3
	for w := 1; w <= workers; w++ {
		go func(workerID int) {
			for job := range jobs {
				fmt.Printf("  Worker %d: Job %d bajarilmoqda\n", workerID, job)
				time.Sleep(50 * time.Millisecond)
				results <- job * 2
			}
		}(w)
	}

	// Natijalarni olish
	go func() {
		for i := 0; i < 10; i++ {
			result := <-results
			fmt.Printf("  Natija: %d\n", result)
		}
	}()

	time.Sleep(1 * time.Second)
}

// demonstrateChannelTicker - ticker channel
func demonstrateChannelTicker() {
	fmt.Println("Ticker channel:")

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	done := make(chan bool)

	go func() {
		time.Sleep(500 * time.Millisecond)
		done <- true
	}()

	fmt.Println("  Ticker natijalari:")
	for {
		select {
		case <-done:
			fmt.Println("  Ticker to'xtatildi")
			return
		case t := <-ticker.C:
			fmt.Printf("    Tick at %v\n", t.Format("15:04:05.000"))
		}
	}
}

// demonstrateChannelTimer - timer channel
func demonstrateChannelTimer() {
	fmt.Println("Timer channel:")

	timer := time.NewTimer(200 * time.Millisecond)

	fmt.Println("  Timer kutmoqda...")
	<-timer.C
	fmt.Println("  Timer yakunlandi!")
}

// demonstrateChannelAfter - time.After
func demonstrateChannelAfter() {
	fmt.Println("time.After channel:")

	fmt.Println("  2 soniya kutmoqda...")
	<-time.After(2 * time.Second)
	fmt.Println("  2 soniya o'tdi!")
}

// demonstrateChannelContext - context pattern (asosiy)
func demonstrateChannelContext() {
	fmt.Println("Context pattern (asosiy):")

	done := make(chan bool)

	go func() {
		time.Sleep(100 * time.Millisecond)
		close(done)
	}()

	// Bir nechta goroutine
	for i := 1; i <= 3; i++ {
		go func(id int) {
			select {
			case <-done:
				fmt.Printf("  Goroutine %d: To'xtatildi\n", id)
				return
			default:
				fmt.Printf("  Goroutine %d: Ishlayapti\n", id)
			}
		}(i)
	}

	time.Sleep(200 * time.Millisecond)
}

// demonstrateChannelErrorHandling - xato bilan ishlash
func demonstrateChannelErrorHandling() {
	fmt.Println("Channel bilan xato bilan ishlash:")

	type Result struct {
		Value int
		Error error
	}

	results := make(chan Result, 5)

	// Vazifalarni bajarish
	go func() {
		for i := 1; i <= 5; i++ {
			var err error
			if i%2 == 0 {
				err = fmt.Errorf("xato: %d", i)
			}

			results <- Result{
				Value: i,
				Error: err,
			}
		}
		close(results)
	}()

	// Natijalarni ko'rsatish
	fmt.Println("  Natijalar:")
	for result := range results {
		if result.Error != nil {
			fmt.Printf("    Xato: %v\n", result.Error)
		} else {
			fmt.Printf("    Muvaffaqiyatli: %d\n", result.Value)
		}
	}
}

// demonstrateChannelPractical - amaliy misol: parallel hisoblash
func demonstrateChannelPractical() {
	fmt.Println("Amaliy misol: Parallel hisoblash:")

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
