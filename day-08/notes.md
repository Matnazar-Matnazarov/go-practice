# Kun 8: Context va Mutex (Sinxronizatsiya)

## 1. Context

### Context nima?

Context - goroutine larni boshqarish, cancel qilish, timeout va deadline o'rnatish, qiymat tarqatish uchun ishlatiladigan mexanizma.

**Asosiy maqsadlar:**
- Goroutine larni cancel qilish
- Timeout va deadline o'rnatish
- Qiymat tarqatish (request ID, user ID, va hokazo)
- Context tarqalishi (propagation)

### Context turlari

#### 1. Background Context

```go
ctx := context.Background()
```

- Asosiy context
- Hech qachon cancel qilinmaydi
- Asosiy context sifatida ishlatiladi

#### 2. TODO Context

```go
ctx := context.TODO()
```

- Vaqtinchalik context
- Qayerda ishlatishni bilmasangiz
- Keyinchalik o'zgartiriladi

### Context operatsiyalari

#### WithValue - Qiymat qo'shish

```go
ctx := context.Background()
ctx = context.WithValue(ctx, "userID", "12345")
ctx = context.WithValue(ctx, "username", "john_doe")

// Qiymat olish
userID := ctx.Value("userID")
username := ctx.Value("username")
```

**Qoidalar:**
- Key va Value `interface{}` turida
- Key uchun alohida type yaratish tavsiya etiladi
- Context qiymatlarini o'zgartirib bo'lmaydi (immutable)

#### WithCancel - Cancel qilish

```go
ctx, cancel := context.WithCancel(context.Background())

go func() {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("To'xtatildi")
            return
        default:
            // Vazifa
        }
    }
}()

// Cancel qilish
cancel()
```

#### WithTimeout - Timeout

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

select {
case <-ctx.Done():
    fmt.Println("Timeout")
case result := <-done:
    fmt.Println(result)
}
```

#### WithDeadline - Deadline

```go
deadline := time.Now().Add(2 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()
```

### Context tarqalishi (Propagation)

#### Qiymat tarqalishi

```go
// Parent context
parentCtx := context.Background()
parentCtx = context.WithValue(parentCtx, "requestID", "req-123")

// Child context (qiymat saqlanadi)
childCtx, cancel := context.WithCancel(parentCtx)
requestID := childCtx.Value("requestID") // "req-123"
```

#### Timeout tarqalishi

```go
// Parent timeout
ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)

// Child timeout (parent dan qisqa)
childCtx, childCancel := context.WithTimeout(ctx, 500*time.Millisecond)
// Child timeout parent timeout dan uzun bo'lsa ham, parent timeout ishlaydi
```

### Context bilan ishlash

#### HTTP Request

```go
ctx := context.Background()
ctx = context.WithValue(ctx, "userID", "user-123")
ctx = context.WithValue(ctx, "requestID", "req-456")

ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
defer cancel()

// API chaqiruv
select {
case result := <-apiCall(ctx):
    fmt.Println(result)
case <-ctx.Done():
    fmt.Println("Request timeout")
}
```

#### Database Query

```go
ctx := context.Background()
ctx = context.WithValue(ctx, "db", "postgresql")
ctx = context.WithValue(ctx, "query", "SELECT * FROM users")

ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
defer cancel()

// Query
select {
case results := <-query(ctx):
    fmt.Println(results)
case <-ctx.Done():
    fmt.Println("Query timeout")
}
```

#### Worker Pool

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

jobs := make(chan int, 10)

// Worker lar
for w := 1; w <= 3; w++ {
    go func(workerID int) {
        for {
            select {
            case <-ctx.Done():
                return
            case job, ok := <-jobs:
                if !ok {
                    return
                }
                // Vazifa
            }
        }
    }(w)
}
```

### Context xato bilan ishlash

```go
type Result struct {
    Value int
    Error error
}

ctx, cancel := context.WithCancel(context.Background())
defer cancel()

results := make(chan Result, 3)

for i := 1; i <= 3; i++ {
    go func(id int) {
        select {
        case <-ctx.Done():
            results <- Result{Error: ctx.Err()}
            return
        default:
            if id == 2 {
                results <- Result{Error: fmt.Errorf("xato: %d", id)}
                return
            }
            results <- Result{Value: id * 2}
        }
    }(i)
}
```

### Muhim qoidalar

1. **Context ni funksiya parametri sifatida uzatish**
   ```go
   func processRequest(ctx context.Context) {
       // ...
   }
   ```

2. **Context ni o'zgartirib bo'lmaydi** (immutable)
   ```go
   ctx = context.WithValue(ctx, "key", "value") // Yangi context
   ```

3. **Cancel ni defer bilan chaqirish**
   ```go
   ctx, cancel := context.WithCancel(context.Background())
   defer cancel()
   ```

4. **Context.Done() ni tekshirish**
   ```go
   select {
   case <-ctx.Done():
       return ctx.Err()
   default:
       // Vazifa
   }
   ```

5. **Context ni nil sifatida uzatmaslik**
   ```go
   // ❌ Noto'g'ri
   func f(ctx context.Context) {
       if ctx == nil {
           ctx = context.Background()
       }
   }
   ```

---

## 2. Mutex (Sinxronizatsiya)

### Mutex nima?

Mutex (Mutual Exclusion) - bir vaqtning o'zida faqat bitta goroutine ga ruxsat beradigan lock mexanizmi.

**Asosiy maqsad:**
- Race condition ni oldini olish
- Shared resource ga xavfsiz kirish

### Race Condition

Race condition - bir nechta goroutine bir xil ma'lumotga bir vaqtda kirishga harakat qilganda yuzaga keladigan muammo.

```go
// Race condition misoli
var counter int

// 10 ta goroutine
for i := 0; i < 10; i++ {
    go func() {
        counter++ // Race condition!
    }()
}
// Counter noto'g'ri bo'lishi mumkin
```

### Mutex bilan hal qilish

```go
var mu sync.Mutex
var counter int

for i := 0; i < 10; i++ {
    go func() {
        mu.Lock()
        counter++
        mu.Unlock()
    }()
}
// Counter to'g'ri bo'ladi
```

### Mutex operatsiyalari

#### Lock va Unlock

```go
var mu sync.Mutex

mu.Lock()
// Xavfsiz kod
mu.Unlock()
```

**Qoidalar:**
- Lock dan keyin har doim Unlock qilish kerak
- Defer bilan ishlatish tavsiya etiladi
- Lock qilingan mutex ni qayta lock qilish mumkin emas (deadlock)

#### Defer bilan ishlatish

```go
mu.Lock()
defer mu.Unlock()
// Xavfsiz kod
```

### RWMutex (Read-Write Mutex)

RWMutex - parallel o'qishga ruxsat beradi, lekin yozish uchun to'liq lock.

**Afzalliklari:**
- Bir nechta reader lar parallel o'qishi mumkin
- Writer lar uchun to'liq lock

```go
var rwmu sync.RWMutex
var data int

// Reader lar (parallel)
rwmu.RLock()
value := data
rwmu.RUnlock()

// Writer (to'liq lock)
rwmu.Lock()
data = 100
rwmu.Unlock()
```

### Deadlock

Deadlock - ikkita yoki undan ko'p mutex lar bir-birini kutganda yuzaga keladi.

```go
// Deadlock misoli
var mu1, mu2 sync.Mutex

// Goroutine 1
go func() {
    mu1.Lock()
    mu2.Lock() // mu2 kutmoqda
    mu2.Unlock()
    mu1.Unlock()
}()

// Goroutine 2
go func() {
    mu2.Lock()
    mu1.Lock() // mu1 kutmoqda (deadlock!)
    mu1.Unlock()
    mu2.Unlock()
}()
```

**Oldini olish:**
- Har doim bir xil tartibda lock qilish
- Timeout qo'shish
- Mutex ni qisqa vaqtga saqlash

### sync.Once

sync.Once - funksiyani faqat bir marta bajarish.

```go
var once sync.Once

for i := 0; i < 10; i++ {
    go func() {
        once.Do(func() {
            fmt.Println("Birinchi marta bajarildi")
        })
    }()
}
// "Birinchi marta bajarildi" faqat bir marta chiqadi
```

### Condition Variable

Condition variable - shartli kutish uchun.

```go
var mu sync.Mutex
cond := sync.NewCond(&mu)
var ready bool

// Waiter
go func() {
    mu.Lock()
    for !ready {
        cond.Wait() // Kutmoqda
    }
    mu.Unlock()
}()

// Signaler
go func() {
    mu.Lock()
    ready = true
    cond.Signal() // Yoki cond.Broadcast()
    mu.Unlock()
}()
```

### Amaliy misollar

#### Xavfsiz Counter

```go
type SafeCounter struct {
    mu    sync.Mutex
    value int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}

func (c *SafeCounter) Get() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.value
}
```

#### Cache

```go
type Cache struct {
    mu    sync.RWMutex
    items map[string]string
}

func (c *Cache) Get(key string) string {
    c.mu.RLock()
    defer c.mu.RUnlock()
    return c.items[key]
}

func (c *Cache) Set(key, value string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.items[key] = value
}
```

#### Bank Account

```go
type BankAccount struct {
    mu      sync.Mutex
    balance int
}

func (a *BankAccount) Deposit(amount int) {
    a.mu.Lock()
    defer a.mu.Unlock()
    a.balance += amount
}

func (a *BankAccount) Withdraw(amount int) error {
    a.mu.Lock()
    defer a.mu.Unlock()
    if a.balance < amount {
        return fmt.Errorf("balance yetarli emas")
    }
    a.balance -= amount
    return nil
}
```

### Muhim qoidalar

1. **Lock dan keyin Unlock qilish**
   ```go
   mu.Lock()
   defer mu.Unlock() // ✅ Tavsiya etiladi
   ```

2. **Qisqa vaqtga lock qilish**
   ```go
   mu.Lock()
   // Qisqa kod
   mu.Unlock()
   ```

3. **Deadlock dan saqlanish**
   - Har doim bir xil tartibda lock qilish
   - Timeout qo'shish

4. **RWMutex ni to'g'ri ishlatish**
   - Ko'p o'qish, kam yozish bo'lsa RWMutex
   - Ko'p yozish bo'lsa oddiy Mutex

5. **Mutex ni struct ichida saqlash**
   ```go
   type SafeCounter struct {
       mu    sync.Mutex
       value int
   }
   ```

### Mutex vs Channel

**Mutex qachon:**
- Shared state ni himoya qilish
- Oddiy counter, map, slice
- Performance muhim

**Channel qachon:**
- Goroutine lar orasida aloqa
- Data transfer
- Pipeline, worker pool

---

## Xulosa

**Context:**
- Goroutine larni boshqarish
- Cancel, timeout, deadline
- Qiymat tarqatish
- HTTP request, database query

**Mutex:**
- Race condition ni hal qilish
- Shared resource ga xavfsiz kirish
- RWMutex - parallel o'qish
- sync.Once, Condition variable

**Keyingi qadamlar:**
- Atomic operations
- Channel patterns (advanced)
- Error handling in goroutines
- Testing concurrent code
