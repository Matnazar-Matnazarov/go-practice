# Kun 7: WaitGroup va Channel (Kengaytirilgan)

## 1. WaitGroup

### WaitGroup nima?

WaitGroup - goroutine larni kutish uchun ishlatiladigan sinxronizatsiya mexanizmi. Barcha goroutine lar yakunlanguncha kutadi.

**Asosiy operatsiyalar:**
- `Add(delta int)` - counter ni oshirish
- `Done()` - counter ni kamaytirish (Add(-1))
- `Wait()` - counter 0 bo'lguncha kutish

### WaitGroup asoslari

```go
var wg sync.WaitGroup

// 3 ta goroutine qo'shish
wg.Add(3)

go func() {
    defer wg.Done()
    // Vazifa
}()

go func() {
    defer wg.Done()
    // Vazifa
}()

go func() {
    defer wg.Done()
    // Vazifa
}()

// Barcha goroutine larni kutish
wg.Wait()
```

### WaitGroup qoidalari

1. **Add() ni goroutine dan oldin chaqirish**
   ```go
   wg.Add(1)  // ✅ To'g'ri
   go func() {
       defer wg.Done()
   }()
   ```

2. **Done() ni defer bilan ishlatish**
   ```go
   go func() {
       defer wg.Done()  // ✅ Xato bo'lsa ham chaqiriladi
       // Vazifa
   }()
   ```

3. **Wait() ni barcha Add() dan keyin chaqirish**
   ```go
   wg.Add(3)
   go func() { defer wg.Done() }()
   go func() { defer wg.Done() }()
   go func() { defer wg.Done() }()
   wg.Wait()  // ✅ Barcha goroutine lar yakunlanguncha kutadi
   ```

### WaitGroup bilan counter

```go
var wg sync.WaitGroup
var counter int
var mu sync.Mutex

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
fmt.Println(counter)  // 10
```

### WaitGroup bilan worker pool

```go
var wg sync.WaitGroup
jobs := make(chan int, 10)
results := make(chan int, 10)

// Job lar yuborish
go func() {
    for j := 1; j <= 10; j++ {
        jobs <- j
    }
    close(jobs)
}()

// Worker lar
workers := 3
wg.Add(workers)

for w := 1; w <= workers; w++ {
    go func(workerID int) {
        defer wg.Done()
        
        for job := range jobs {
            results <- job * 2
        }
    }(w)
}

// Worker larni kutish
go func() {
    wg.Wait()
    close(results)
}()
```

### WaitGroup bilan xato bilan ishlash

```go
var wg sync.WaitGroup
errors := make(chan error, 10)

wg.Add(10)

for i := 0; i < 10; i++ {
    go func(id int) {
        defer wg.Done()
        
        if id%2 == 0 {
            errors <- fmt.Errorf("xato: %d", id)
            return
        }
    }(i)
}

// Xatolarni yig'ish
go func() {
    wg.Wait()
    close(errors)
}()

// Xatolarni ko'rsatish
for err := range errors {
    if err != nil {
        fmt.Println(err)
    }
}
```

### WaitGroup bilan timeout

```go
var wg sync.WaitGroup
done := make(chan bool)

wg.Add(1)
go func() {
    defer wg.Done()
    time.Sleep(200 * time.Millisecond)
    done <- true
}()

select {
case <-done:
    wg.Wait()
    fmt.Println("Muvaffaqiyatli")
case <-time.After(100 * time.Millisecond):
    fmt.Println("Timeout")
}
```

### WaitGroup bilan ichki WaitGroup

```go
var outerWg sync.WaitGroup
var innerWg sync.WaitGroup

outerWg.Add(2)

// Birinchi goroutine
go func() {
    defer outerWg.Done()
    
    innerWg.Add(3)
    for i := 0; i < 3; i++ {
        go func() {
            defer innerWg.Done()
            // Vazifa
        }()
    }
    innerWg.Wait()
}()

outerWg.Wait()
```

### WaitGroup amaliy qo'llanilishi

1. **Parallel fayl o'qish**
2. **Parallel hisoblash**
3. **Worker pool**
4. **Batch processing**
5. **API request lar**

---

## 2. Channel (Kengaytirilgan)

### Channel turlari

#### 1. Unbuffered Channel

```go
ch := make(chan int)

// Blocking - receiver kutmoqda bo'lishi kerak
ch <- 10
value := <-ch
```

#### 2. Buffered Channel

```go
ch := make(chan int, 10)

// Non-blocking (buffer to'lmaguncha)
ch <- 1
ch <- 2
ch <- 3

// Olish
value := <-ch
```

### Channel yo'nalishi

#### Send-only Channel

```go
func sendOnly(ch chan<- int) {
    ch <- 10  // ✅ Yuborish mumkin
    // value := <-ch  // ❌ Xato
}
```

#### Receive-only Channel

```go
func receiveOnly(ch <-chan int) {
    value := <-ch  // ✅ Olish mumkin
    // ch <- 10  // ❌ Xato
}
```

### Channel yopish

```go
ch := make(chan int, 5)

// Ma'lumotlar yuborish
go func() {
    for i := 1; i <= 5; i++ {
        ch <- i
    }
    close(ch)  // Channel ni yopish
}()

// Range bilan o'qish
for value := range ch {
    fmt.Println(value)
}

// Yopilgan channel dan olish
value, ok := <-ch
// ok = false (yopilgan)
```

### Select Statement

Bir nechta channel dan birinchi tayyor bo'lganini olish.

```go
select {
case msg1 := <-ch1:
    fmt.Println(msg1)
case msg2 := <-ch2:
    fmt.Println(msg2)
case <-time.After(1 * time.Second):
    fmt.Println("Timeout")
default:
    fmt.Println("Hech qaysi channel tayyor emas")
}
```

### Timeout

```go
select {
case msg := <-ch:
    fmt.Println(msg)
case <-time.After(100 * time.Millisecond):
    fmt.Println("Timeout")
}
```

### Non-blocking Channel

```go
// Non-blocking send
select {
case ch <- value:
    fmt.Println("Yuborildi")
default:
    fmt.Println("Channel to'ldi")
}

// Non-blocking receive
select {
case value := <-ch:
    fmt.Println(value)
default:
    fmt.Println("Ma'lumot yo'q")
}
```

### Ticker Channel

```go
ticker := time.NewTicker(1 * time.Second)
defer ticker.Stop()

done := make(chan bool)

go func() {
    time.Sleep(5 * time.Second)
    done <- true
}()

for {
    select {
    case <-done:
        return
    case t := <-ticker.C:
        fmt.Println("Tick at", t)
    }
}
```

### Timer Channel

```go
timer := time.NewTimer(2 * time.Second)

<-timer.C
fmt.Println("2 soniya o'tdi")
```

### time.After

```go
select {
case <-time.After(2 * time.Second):
    fmt.Println("2 soniya o'tdi")
}
```

### Pipeline Pattern

```go
// Stage 1: Input
numbers := make(chan int)
go func() {
    for i := 1; i <= 10; i++ {
        numbers <- i
    }
    close(numbers)
}()

// Stage 2: Processing
doubled := make(chan int)
go func() {
    for num := range numbers {
        doubled <- num * 2
    }
    close(doubled)
}()

// Stage 3: Output
for result := range doubled {
    fmt.Println(result)
}
```

### Fan-out Pattern

Bir channel ni bir nechta worker larga taqsimlash.

```go
input := make(chan int, 10)
output1 := make(chan int, 10)
output2 := make(chan int, 10)

// Fan-out
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
```

### Fan-in Pattern

Bir nechta channel ni bitta channel ga birlashtirish.

```go
merged := make(chan int)

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
```

### Worker Pool

```go
jobs := make(chan int, 10)
results := make(chan int, 10)

// Job lar yuborish
go func() {
    for j := 1; j <= 10; j++ {
        jobs <- j
    }
    close(jobs)
}()

// Worker lar
workers := 3
for w := 1; w <= workers; w++ {
    go func(workerID int) {
        for job := range jobs {
            results <- job * 2
        }
    }(w)
}

// Natijalarni olish
for i := 0; i < 10; i++ {
    result := <-results
    fmt.Println(result)
}
```

### Channel bilan xato bilan ishlash

```go
type Result struct {
    Value int
    Error error
}

results := make(chan Result, 10)

go func() {
    for i := 1; i <= 10; i++ {
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

for result := range results {
    if result.Error != nil {
        fmt.Printf("Xato: %v\n", result.Error)
    } else {
        fmt.Printf("Muvaffaqiyatli: %d\n", result.Value)
    }
}
```

### Context Pattern (Asosiy)

```go
done := make(chan bool)

go func() {
    time.Sleep(100 * time.Millisecond)
    close(done)
}()

for i := 1; i <= 3; i++ {
    go func(id int) {
        select {
        case <-done:
            fmt.Printf("Goroutine %d: To'xtatildi\n", id)
            return
        default:
            fmt.Printf("Goroutine %d: Ishlayapti\n", id)
        }
    }(i)
}
```

### Muhim qoidalar

1. **Channel ni yopish** - yuboruvchi tomonidan
2. **Range bilan o'qish** - yopilgan channel uchun
3. **Select default** - non-blocking operatsiyalar
4. **Timeout** - `time.After` yoki `time.NewTimer`
5. **Deadlock dan saqlanish** - channel lar to'liq bo'lmasligi kerak

### Amaliy qo'llanilishi

1. **Goroutine lar orasida aloqa**
2. **Worker pool**
3. **Pipeline processing**
4. **Event handling**
5. **Timeout va cancellation**

---

## Xulosa

**WaitGroup:**
- Goroutine larni kutish
- Add, Done, Wait operatsiyalari
- Worker pool, parallel processing
- Xato bilan ishlash

**Channel:**
- Buffered va unbuffered
- Send-only va receive-only
- Select, timeout, non-blocking
- Pipeline, Fan-out, Fan-in pattern lar
- Ticker va Timer

**Keyingi qadamlar:**
- Context package
- Mutex va RWMutex
- Atomic operations
- Advanced channel patterns
- Error handling in goroutines
