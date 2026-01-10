package main

import (
	"fmt"
	"sync"
	"time"
)

// demonstrateWaitGroupBasic - WaitGroup asoslari
func demonstrateWaitGroupBasic() {
	fmt.Println("WaitGroup asoslari:")

	var wg sync.WaitGroup

	// 3 ta goroutine qo'shish
	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println("  Goroutine 1 bajarildi")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("  Goroutine 2 bajarildi")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("  Goroutine 3 bajarildi")
	}()

	// Barcha goroutine larni kutish
	wg.Wait()
	fmt.Println("  Barcha goroutine lar yakunlandi")
}

// demonstrateWaitGroupMultiple - bir nechta goroutine
func demonstrateWaitGroupMultiple() {
	fmt.Println("WaitGroup bilan bir nechta goroutine:")

	var wg sync.WaitGroup
	count := 5

	wg.Add(count)

	for i := 1; i <= count; i++ {
		go func(id int) {
			defer wg.Done()
			fmt.Printf("  Goroutine %d ishlayapti\n", id)
			time.Sleep(50 * time.Millisecond)
			fmt.Printf("  Goroutine %d yakunlandi\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("  Barcha goroutine lar yakunlandi")
}

// demonstrateWaitGroupWithDelay - kechikish bilan
func demonstrateWaitGroupWithDelay() {
	fmt.Println("WaitGroup bilan kechikish:")

	var wg sync.WaitGroup
	delays := []int{100, 200, 300}

	wg.Add(len(delays))

	for i, delay := range delays {
		go func(id int, d int) {
			defer wg.Done()
			fmt.Printf("  Goroutine %d: %dms kutmoqda\n", id, d)
			time.Sleep(time.Duration(d) * time.Millisecond)
			fmt.Printf("  Goroutine %d: Yakunlandi\n", id)
		}(i+1, delay)
	}

	fmt.Println("  Barcha goroutine larni kutmoqdamiz...")
	wg.Wait()
	fmt.Println("  Barcha goroutine lar yakunlandi")
}

// demonstrateWaitGroupCounter - counter bilan
func demonstrateWaitGroupCounter() {
	fmt.Println("WaitGroup bilan counter:")

	var wg sync.WaitGroup
	var counter int
	var mu sync.Mutex

	goroutines := 10
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(id int) {
			defer wg.Done()

			// Counter ni oshirish (mutex bilan)
			mu.Lock()
			counter++
			current := counter
			mu.Unlock()

			fmt.Printf("  Goroutine %d: Counter = %d\n", id, current)
		}(i)
	}

	wg.Wait()
	fmt.Printf("  Yakuniy counter: %d\n", counter)
}

// demonstrateWaitGroupWorkerPool - worker pool bilan
func demonstrateWaitGroupWorkerPool() {
	fmt.Println("WaitGroup bilan worker pool:")

	var wg sync.WaitGroup
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Job lar yuborish
	go func() {
		for j := 1; j <= 5; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	// 3 ta worker
	workers := 3
	wg.Add(workers)

	for w := 1; w <= workers; w++ {
		go func(workerID int) {
			defer wg.Done()

			for job := range jobs {
				fmt.Printf("  Worker %d: Job %d bajarilmoqda\n", workerID, job)
				time.Sleep(50 * time.Millisecond)
				results <- job * 2
				fmt.Printf("  Worker %d: Job %d yakunlandi\n", workerID, job)
			}
		}(w)
	}

	// Worker larni kutish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Natijalarni olish
	fmt.Println("  Natijalar:")
	for result := range results {
		fmt.Printf("    Natija: %d\n", result)
	}
}

// demonstrateWaitGroupErrorHandling - xato bilan ishlash
func demonstrateWaitGroupErrorHandling() {
	fmt.Println("WaitGroup bilan xato bilan ishlash:")

	var wg sync.WaitGroup
	errors := make(chan error, 5)

	tasks := []int{1, 2, 3, 4, 5}
	wg.Add(len(tasks))

	for _, task := range tasks {
		go func(t int) {
			defer wg.Done()

			// Xato simulatsiyasi (juft sonlar uchun)
			if t%2 == 0 {
				errors <- fmt.Errorf("task %d xato", t)
				return
			}

			fmt.Printf("  Task %d muvaffaqiyatli\n", t)
		}(task)
	}

	// Xatolarni yig'ish
	go func() {
		wg.Wait()
		close(errors)
	}()

	// Xatolarni ko'rsatish
	fmt.Println("  Xatolar:")
	for err := range errors {
		if err != nil {
			fmt.Printf("    %v\n", err)
		}
	}
}

// demonstrateWaitGroupParallelProcessing - parallel qayta ishlash
func demonstrateWaitGroupParallelProcessing() {
	fmt.Println("WaitGroup bilan parallel qayta ishlash:")

	var wg sync.WaitGroup
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	results := make(chan int, len(numbers))

	wg.Add(len(numbers))

	// Har bir sonni parallel qayta ishlash
	for _, num := range numbers {
		go func(n int) {
			defer wg.Done()

			// Hisoblash (simulatsiya)
			time.Sleep(10 * time.Millisecond)
			square := n * n
			results <- square
		}(num)
	}

	// Natijalarni yig'ish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Natijalarni ko'rsatish
	fmt.Println("  Natijalar:")
	sum := 0
	for result := range results {
		sum += result
		fmt.Printf("    %d\n", result)
	}
	fmt.Printf("  Yig'indi: %d\n", sum)
}

// demonstrateWaitGroupSequentialTasks - ketma-ket vazifalar
func demonstrateWaitGroupSequentialTasks() {
	fmt.Println("WaitGroup bilan ketma-ket vazifalar:")

	var wg sync.WaitGroup

	// Stage 1
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("  Stage 1: Ma'lumotlarni yuklash")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("  Stage 1: Yakunlandi")
	}()

	wg.Wait()

	// Stage 2
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("  Stage 2: Ma'lumotlarni qayta ishlash")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("  Stage 2: Yakunlandi")
	}()

	wg.Wait()

	// Stage 3
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("  Stage 3: Natijalarni saqlash")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("  Stage 3: Yakunlandi")
	}()

	wg.Wait()
	fmt.Println("  Barcha stage lar yakunlandi")
}

// demonstrateWaitGroupTimeout - timeout bilan (channel bilan)
func demonstrateWaitGroupTimeout() {
	fmt.Println("WaitGroup bilan timeout (channel bilan):")

	var wg sync.WaitGroup
	done := make(chan bool)

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(200 * time.Millisecond)
		fmt.Println("  Goroutine yakunlandi")
		done <- true
	}()

	// Timeout bilan kutish
	select {
	case <-done:
		wg.Wait()
		fmt.Println("  Muvaffaqiyatli yakunlandi")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("  Timeout: Goroutine hali yakunlanmadi")
	}
}

// demonstrateWaitGroupNested - ichki WaitGroup
func demonstrateWaitGroupNested() {
	fmt.Println("WaitGroup bilan ichki WaitGroup:")

	var outerWg sync.WaitGroup
	var innerWg sync.WaitGroup

	outerWg.Add(2)

	// Birinchi goroutine
	go func() {
		defer outerWg.Done()

		innerWg.Add(3)
		for i := 1; i <= 3; i++ {
			go func(id int) {
				defer innerWg.Done()
				fmt.Printf("  Inner goroutine %d-1\n", id)
			}(i)
		}
		innerWg.Wait()
		fmt.Println("  Outer goroutine 1 yakunlandi")
	}()

	// Ikkinchi goroutine
	go func() {
		defer outerWg.Done()

		innerWg.Add(2)
		for i := 1; i <= 2; i++ {
			go func(id int) {
				defer innerWg.Done()
				fmt.Printf("  Inner goroutine %d-2\n", id)
			}(i)
		}
		innerWg.Wait()
		fmt.Println("  Outer goroutine 2 yakunlandi")
	}()

	outerWg.Wait()
	fmt.Println("  Barcha goroutine lar yakunlandi")
}

// demonstrateWaitGroupPractical - amaliy misol: fayllarni parallel o'qish
func demonstrateWaitGroupPractical() {
	fmt.Println("Amaliy misol: Fayllarni parallel o'qish (simulatsiya):")

	var wg sync.WaitGroup
	files := []string{"file1.txt", "file2.txt", "file3.txt", "file4.txt", "file5.txt"}
	results := make(chan string, len(files))

	wg.Add(len(files))

	for _, filename := range files {
		go func(file string) {
			defer wg.Done()

			// Fayl o'qish simulatsiyasi
			time.Sleep(50 * time.Millisecond)
			results <- fmt.Sprintf("%s o'qildi (%d bytes)", file, len(file)*10)
		}(filename)
	}

	// Natijalarni yig'ish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Natijalarni ko'rsatish
	fmt.Println("  Natijalar:")
	for result := range results {
		fmt.Printf("    %s\n", result)
	}
}
