# Kun 8: Context va Mutex (Sinxronizatsiya)

## Kirish

Go dasturlash tilida parallel dasturlash (concurrency) juda muhim. Goroutine lar orasida to'g'ri boshqaruv va sinxronizatsiya uchun **Context** va **Mutex** mexanizmlari ishlatiladi.

---

## 1. Context

### Context nima?

**Context** - goroutine larni boshqarish, cancel qilish, timeout va deadline o'rnatish, qiymat tarqatish uchun ishlatiladigan mexanizma.

**Context** - bu goroutine lar orasida ma'lumot va signal tarqatish uchun standart usul.

### Context nima uchun kerak?

1. **Goroutine larni to'xtatish (Cancellation)**
   - Uzoq davom etadigan operatsiyalarni to'xtatish
   - Foydalanuvchi request ni bekor qilganda
   - Timeout bo'lganda

2. **Timeout va Deadline**
   - Operatsiyalar uchun maksimal vaqt belgilash
   - Uzoq davom etadigan operatsiyalarni cheklash
   - Resource leak larni oldini olish

3. **Qiymat tarqatish (Value Propagation)**
   - Request ID, User ID kabi ma'lumotlarni tarqatish
   - Logging va tracing uchun
   - Authentication va authorization ma'lumotlari

4. **Context tarqalishi (Propagation)**
   - Parent context dan child context ga ma'lumot o'tkazish
   - HTTP request dan database query ga ma'lumot o'tkazish

### Context vazifasi

Context quyidagi vazifalarni bajaradi:

1. **Goroutine larni boshqarish** - cancel, timeout, deadline
2. **Ma'lumot tarqatish** - request ID, user ID va boshqalar
3. **Resource management** - goroutine larni to'g'ri yopish
4. **Error propagation** - xatolarni tarqatish

### Context afzalliklari

✅ **Standart mexanizma** - Go standart kutubxonasida
✅ **Type-safe** - Compile time da tekshiriladi
✅ **Immutable** - O'zgartirib bo'lmaydi (xavfsiz)
✅ **Propagation** - Avtomatik tarqaladi
✅ **Cancellation** - Oson cancel qilish
✅ **Timeout** - Oson timeout o'rnatish

### Context kamchiliklari

❌ **Overhead** - Kichik overhead (juda kichik)
❌ **Immutable** - Har safar yangi context yaratiladi
❌ **Value type** - Key uchun alohida type yaratish tavsiya etiladi

### Context turlari

#### 1. Background Context

```go
ctx := context.Background()
```

- **Nima:** Asosiy context, hech qachon cancel qilinmaydi
- **Qachon ishlatiladi:** Dastur boshlanishida, asosiy context sifatida
- **Xususiyati:** Hech qachon timeout yoki cancel bo'lmaydi

#### 2. TODO Context

```go
ctx := context.TODO()
```

- **Nima:** Vaqtinchalik context
- **Qachon ishlatiladi:** Qayerda ishlatishni bilmasangiz, keyinchalik o'zgartiriladi
- **Xususiyati:** Background ga o'xshaydi, lekin maqsad aniq emas

### Context operatsiyalari

#### WithValue - Qiymat qo'shish

**Nima:** Context ga qiymat qo'shish

**Qachon ishlatiladi:**
- Request ID, User ID kabi ma'lumotlarni saqlash
- Logging va tracing uchun
- Authentication ma'lumotlari

**Qoidalar:**
- Key va Value `interface{}` turida
- Key uchun alohida type yaratish tavsiya etiladi (type safety)
- Context qiymatlarini o'zgartirib bo'lmaydi (immutable)

**Misol:**
```go
type contextKey string

const userIDKey contextKey = "userID"

ctx := context.Background()
ctx = context.WithValue(ctx, userIDKey, "12345")

// Qiymat olish
userID := ctx.Value(userIDKey)
```

#### WithCancel - Cancel qilish

**Nima:** Context ni cancel qilish imkoniyatini beradi

**Qachon ishlatiladi:**
- Goroutine larni to'xtatish kerak bo'lganda
- Foydalanuvchi request ni bekor qilganda
- Shartli to'xtatish kerak bo'lganda

**Qoidalar:**
- `cancel()` funksiyasini chaqirish kerak
- `defer cancel()` bilan ishlatish tavsiya etiladi
- Cancel qilingan context `ctx.Done()` channel orqali signal beradi

**Misol:**
```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

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

**Nima:** Context ga timeout o'rnatadi

**Qachon ishlatiladi:**
- HTTP request lar uchun
- Database query lar uchun
- Uzoq davom etadigan operatsiyalar uchun

**Qoidalar:**
- Vaqt o'tgach avtomatik cancel bo'ladi
- `defer cancel()` bilan ishlatish tavsiya etiladi
- Timeout bo'lganda `ctx.Err()` `context.DeadlineExceeded` qaytaradi

**Misol:**
```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

select {
case <-ctx.Done():
    fmt.Println("Timeout:", ctx.Err())
case result := <-done:
    fmt.Println(result)
}
```

#### WithDeadline - Deadline

**Nima:** Context ga aniq vaqt (deadline) o'rnatadi

**Qachon ishlatiladi:**
- Aniq vaqtda to'xtatish kerak bo'lganda
- Schedule qilingan vazifalar uchun

**Qoidalar:**
- `time.Time` turida deadline beriladi
- Deadline o'tgach avtomatik cancel bo'ladi
- Timeout dan farqi - aniq vaqt belgilanadi

**Misol:**
```go
deadline := time.Now().Add(2 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()
```

### Context tarqalishi (Propagation)

#### Qiymat tarqalishi

Parent context dan child context ga qiymatlar avtomatik tarqaladi.

```go
// Parent context
parentCtx := context.Background()
parentCtx = context.WithValue(parentCtx, "requestID", "req-123")

// Child context (qiymat saqlanadi)
childCtx, cancel := context.WithCancel(parentCtx)
requestID := childCtx.Value("requestID") // "req-123"
```

#### Timeout tarqalishi

Parent context timeout child context ga ham ta'sir qiladi.

```go
// Parent timeout (200ms)
ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)

// Child timeout (500ms) - lekin parent timeout ishlaydi
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

### Context muhim qoidalari

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

6. **Context ni struct field sifatida saqlamaslik**
   - Context ni har safar parametr sifatida uzatish kerak

---

## 2. Mutex (Sinxronizatsiya)

### Mutex nima?

**Mutex** (Mutual Exclusion) - bir vaqtning o'zida faqat bitta goroutine ga ruxsat beradigan lock mexanizmi.

**Mutex** - bu shared resource ga xavfsiz kirish uchun mexanizma.

### Mutex nima uchun kerak?

1. **Race Condition ni oldini olish**
   - Bir nechta goroutine bir xil ma'lumotga kirishganda
   - Data corruption ni oldini olish
   - To'g'ri natijalar olish

2. **Shared Resource ga xavfsiz kirish**
   - Map, slice, counter kabi ma'lumotlarga
   - Bir vaqtda faqat bitta goroutine o'zgartirishi mumkin
   - Data consistency ni ta'minlash

3. **Critical Section ni himoya qilish**
   - Muhim kod qismlarini himoya qilish
   - Bir vaqtda faqat bitta goroutine bajarishi mumkin

### Mutex vazifasi

Mutex quyidagi vazifalarni bajaradi:

1. **Lock** - Resource ga kirishni bloklash
2. **Unlock** - Resource ni ochish
3. **Synchronization** - Goroutine lar orasida sinxronizatsiya
4. **Data protection** - Shared data ni himoya qilish

### Mutex afzalliklari

✅ **Oddiy** - Oson ishlatish
✅ **Tez** - Past overhead
✅ **To'g'ri** - Race condition ni hal qiladi
✅ **Flexible** - Har qanday data type bilan ishlaydi

### Mutex kamchiliklari

❌ **Deadlock** - Noto'g'ri ishlatilsa deadlock bo'lishi mumkin
❌ **Performance** - Lock kutish vaqt oladi
❌ **Complexity** - Murakkab dasturlarda qiyin boshqarish

### Race Condition

**Race Condition** - bir nechta goroutine bir xil ma'lumotga bir vaqtda kirishga harakat qilganda yuzaga keladigan muammo.

**Muammo:**
- Data corruption
- Noto'g'ri natijalar
- Dastur xato ishlashi

**Misol:**
```go
// Race condition misoli
var counter int

// 10 ta goroutine
for i := 0; i < 10; i++ {
    go func() {
        counter++ // Race condition!
        // counter++ uchta operatsiya:
        // 1. counter ni o'qish
        // 2. 1 qo'shish
        // 3. counter ga yozish
        // Bir nechta goroutine bir vaqtda bajarilsa, ba'zi qo'shishlar yo'qoladi
    }()
}
// Counter noto'g'ri bo'lishi mumkin (masalan, 7 yoki 8 bo'lishi mumkin, 10 emas)
```

### Mutex bilan hal qilish

```go
var mu sync.Mutex
var counter int

for i := 0; i < 10; i++ {
    go func() {
        mu.Lock()      // Lock qilish
        counter++      // Xavfsiz operatsiya
        mu.Unlock()    // Unlock qilish
    }()
}
// Counter har doim to'g'ri bo'ladi (10)
```

### Mutex operatsiyalari

#### Lock va Unlock

**Lock** - Resource ga kirishni bloklash
**Unlock** - Resource ni ochish

```go
var mu sync.Mutex

mu.Lock()
// Xavfsiz kod - faqat bitta goroutine bajaradi
mu.Unlock()
```

**Qoidalar:**
- Lock dan keyin har doim Unlock qilish kerak
- Defer bilan ishlatish tavsiya etiladi
- Lock qilingan mutex ni qayta lock qilish mumkin emas (deadlock)

#### Defer bilan ishlatish

```go
mu.Lock()
defer mu.Unlock() // Xato bo'lsa ham unlock qilinadi
// Xavfsiz kod
```

### RWMutex (Read-Write Mutex)

**RWMutex** - parallel o'qishga ruxsat beradi, lekin yozish uchun to'liq lock.

**Afzalliklari:**
- Bir nechta reader lar parallel o'qishi mumkin
- Writer lar uchun to'liq lock
- O'qish ko'p, yozish kam bo'lsa samarali

**Qachon ishlatiladi:**
- Cache, database kabi ma'lumotlar uchun
- Ko'p o'qish, kam yozish bo'lsa
- Performance muhim bo'lsa

**Misol:**
```go
var rwmu sync.RWMutex
var data int

// Reader lar (parallel o'qish)
rwmu.RLock()
value := data
rwmu.RUnlock()

// Writer (to'liq lock)
rwmu.Lock()
data = 100
rwmu.Unlock()
```

### Deadlock

**Deadlock** - ikkita yoki undan ko'p mutex lar bir-birini kutganda yuzaga keladi.

**Muammo:**
- Dastur to'xtab qoladi
- Goroutine lar kutib qoladi
- Resource lar qulflangan qoladi

**Misol:**
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
- Lock qilish tartibini rejalashtirish

### sync.Once

**sync.Once** - funksiyani faqat bir marta bajarish.

**Qachon ishlatiladi:**
- Initialization kodlari uchun
- Singleton pattern uchun
- Bir marta bajarilishi kerak bo'lgan kodlar uchun

**Misol:**
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

**Condition Variable** - shartli kutish uchun.

**Qachon ishlatiladi:**
- Goroutine lar orasida signal berish
- Shart bajarilguncha kutish
- Producer-Consumer pattern uchun

**Misol:**
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

### Mutex muhim qoidalari

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
   - Lock qilish tartibini rejalashtirish

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

6. **Lock qilish vaqtini minimallashtirish**
   - Faqat kerakli kod qismini lock qilish
   - I/O operatsiyalarni lock dan tashqarida bajarish

### Mutex vs Channel

**Mutex qachon:**
- Shared state ni himoya qilish
- Oddiy counter, map, slice
- Performance muhim
- Oddiy sinxronizatsiya

**Channel qachon:**
- Goroutine lar orasida aloqa
- Data transfer
- Pipeline, worker pool
- Message passing

**Qoida:**
- **"Don't communicate by sharing memory; share memory by communicating"**
- Lekin ba'zi hollarda Mutex yaxshiroq

---

## Xulosa

### Context

**Nima:** Goroutine larni boshqarish mexanizmi

**Vazifasi:**
- Cancel, timeout, deadline
- Qiymat tarqatish
- HTTP request, database query

**Afzalliklari:**
- Standart mexanizma
- Type-safe
- Immutable
- Propagation

**Kamchiliklari:**
- Kichik overhead
- Har safar yangi context

### Mutex

**Nima:** Shared resource ga xavfsiz kirish mexanizmi

**Vazifasi:**
- Race condition ni hal qilish
- Shared resource ga xavfsiz kirish
- RWMutex - parallel o'qish
- sync.Once, Condition variable

**Afzalliklari:**
- Oddiy va tez
- To'g'ri natijalar
- Flexible

**Kamchiliklari:**
- Deadlock xavfi
- Performance overhead
- Complexity

### Keyingi qadamlar

- Atomic operations
- Channel patterns (advanced)
- Error handling in goroutines
- Testing concurrent code
- Performance optimization