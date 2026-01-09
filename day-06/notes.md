# Kun 6: Queue (Navbat) va Goroutine (Parallel dasturlash)

## 1. Queue (Navbat) Ma'lumot Tuzilmasi

### Queue nima?

Queue (navbat) - FIFO (First In, First Out) prinsipiga asoslangan ma'lumot tuzilmasi. Birinchi qo'shilgan element birinchi olinadi.

```
Qo'shish: [1] → [1, 2] → [1, 2, 3]
Olish:    [1, 2, 3] → [2, 3] → [3] → []
```

### Queue operatsiyalari

#### 1. Enqueue (Qo'shish)
Elementni queue ning oxiriga qo'shadi.

```go
queue.Enqueue(10)
queue.Enqueue(20)
// Queue: [10, 20]
```

#### 2. Dequeue (Olib tashlash)
Elementni queue ning boshidan olib tashlaydi.

```go
item, ok := queue.Dequeue()
// item = 10, ok = true
// Queue: [20]
```

#### 3. Front (Ko'rish)
Birinchi elementni ko'radi, lekin olib tashlamaydi.

```go
front, ok := queue.Front()
// front = 10, ok = true
// Queue: [10, 20] (o'zgarmadi)
```

#### 4. IsEmpty (Bo'shligini tekshirish)
Queue bo'sh yoki yo'qligini tekshiradi.

```go
isEmpty := queue.IsEmpty()
// isEmpty = false
```

#### 5. Size (Uzunlik)
Queue dagi elementlar sonini qaytaradi.

```go
size := queue.Size()
// size = 2
```

### Queue turlari

#### 1. Oddiy Queue
Asosiy FIFO queue.

```go
type Queue struct {
    items []int
}
```

#### 2. String Queue
String elementlar uchun maxsus queue.

```go
type StringQueue struct {
    items []string
}
```

#### 3. Generic Queue
Har qanday tip uchun queue (interface{}).

```go
type GenericQueue struct {
    items []interface{}
}
```

#### 4. Circular Queue
Aylanma queue - oxiridan boshlanadi.

#### 5. Priority Queue
Prioritet bo'yicha tartiblangan queue.

### Queue amaliy qo'llanilishi

1. **Vazifa rejalashtirish** - CPU scheduler
2. **BFS (Breadth-First Search)** - Graph traversal
3. **Print queue** - Printer spooler
4. **Message queue** - Microservices
5. **Request queue** - Web server

### Misol: Vazifa rejalashtirish

```go
taskQueue := NewStringQueue()

// Vazifalarni qo'shish
taskQueue.Enqueue("Email tekshirish")
taskQueue.Enqueue("Kod yozish")
taskQueue.Enqueue("Test yozish")

// Vazifalarni bajarish
for !taskQueue.IsEmpty() {
    task, _ := taskQueue.Dequeue()
    fmt.Printf("Bajarilmoqda: %s\n", task)
}
```

### Misol: BFS simulatsiyasi

```go
queue := NewQueue()

// Graph node larni qo'shish
for _, node := range nodes {
    queue.Enqueue(node)
}

// Level bo'yicha qayta ishlash
for !queue.IsEmpty() {
    node, _ := queue.Dequeue()
    // Keyingi level node lar
    for _, next := range getNeighbors(node) {
        queue.Enqueue(next)
    }
}
```

---

## 2. Goroutine (Parallel dasturlash)

### Goroutine nima?

Goroutine - Go tilida parallel dasturlash uchun yengil thread. Har bir goroutine Go runtime tomonidan boshqariladi.

**Asosiy xususiyatlar:**
- Yengil (2KB stack)
- Ko'p goroutine lar (minglab)
- Automatic scheduling
- Channel orqali aloqa

### Goroutine yaratish

#### 1. Oddiy goroutine

```go
go func() {
    fmt.Println("Goroutine ishlayapti")
}()
```

#### 2. Funksiya goroutine

```go
go sayHello("Goroutine")
```

#### 3. Bir nechta goroutine

```go
for i := 1; i <= 5; i++ {
    go func(id int) {
        fmt.Printf("Goroutine %d\n", id)
    }(i)
}
```

### Channel (Kanal)

Channel - goroutine lar orasida aloqa qilish uchun.

#### 1. Channel yaratish

```go
ch := make(chan int)        // Unbuffered
ch := make(chan int, 10)    // Buffered (10 element)
```

#### 2. Ma'lumot yuborish va olish

```go
// Yuborish
ch <- 10

// Olish
value := <-ch
```

#### 3. Channel yopish

```go
close(ch)
```

#### 4. Range bilan o'qish

```go
for value := range ch {
    fmt.Println(value)
}
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
select {
case ch <- value:
    fmt.Println("Yuborildi")
default:
    fmt.Println("Channel to'ldi")
}
```

### Goroutine Pattern lar

#### 1. Worker Pool

Bir nechta worker goroutine lar job larni parallel bajaradi.

```go
jobs := make(chan int, 10)
results := make(chan int, 10)

// Worker lar
for w := 1; w <= 3; w++ {
    go worker(w, jobs, results)
}

// Job lar yuborish
for j := 1; j <= 10; j++ {
    jobs <- j
}
close(jobs)
```

#### 2. Pipeline

Ma'lumotni bir nechta stage orqali o'tkazish.

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

#### 3. Fan-out

Bir channel ni bir nechta worker larga taqsimlash.

```go
input := make(chan int)
output1 := make(chan int)
output2 := make(chan int)

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

#### 4. Fan-in

Bir nechta channel ni bitta channel ga birlashtirish.

```go
merged := make(chan int)

go func() {
    for {
        select {
        case val := <-ch1:
            merged <- val
        case val := <-ch2:
            merged <- val
        }
    }
}()
```

### Race Condition

Bir nechta goroutine bir xil ma'lumotni o'zgartirganda yuzaga keladi.

```go
counter := 0

for i := 0; i < 10; i++ {
    go func() {
        counter++  // Race condition!
    }()
}
```

**Yechim:** Mutex yoki channel ishlatish (keyingi kunlarda).

### Amaliy misollar

#### 1. Parallel hisoblash

```go
numbers := []int{1, 2, 3, 4, 5}
results := make(chan int, len(numbers))

for _, num := range numbers {
    go func(n int) {
        results <- n * n
    }(num)
}

sum := 0
for i := 0; i < len(numbers); i++ {
    sum += <-results
}
```

#### 2. Concurrent I/O

```go
files := []string{"file1.txt", "file2.txt", "file3.txt"}
results := make(chan string, len(files))

for _, file := range files {
    go func(filename string) {
        // I/O operatsiyasi
        results <- readFile(filename)
    }(file)
}

for i := 0; i < len(files); i++ {
    fmt.Println(<-results)
}
```

### Muhim qoidalar

1. **Goroutine ni kutish** - `time.Sleep` yoki channel
2. **Channel ni yopish** - `close(ch)` yuboruvchi tomonidan
3. **Deadlock dan saqlanish** - channel lar to'liq bo'lmasligi kerak
4. **Race condition** - mutex yoki channel bilan hal qilish
5. **Resource leak** - goroutine larni to'g'ri yopish

### Keyingi kunlar

- Mutex va RWMutex
- WaitGroup
- Context
- Atomic operations
- Channel patterns (advanced)

---

## Xulosa

**Queue:**
- FIFO ma'lumot tuzilmasi
- Enqueue, Dequeue, Front operatsiyalari
- Amaliy qo'llanilishi: task scheduling, BFS

**Goroutine:**
- Parallel dasturlash
- Channel orqali aloqa
- Select, Timeout, Non-blocking
- Pattern lar: Worker Pool, Pipeline, Fan-out, Fan-in

**Keyingi qadamlar:**
- Mutex va synchronization
- Context va cancellation
- Advanced channel patterns
- Error handling in goroutines
