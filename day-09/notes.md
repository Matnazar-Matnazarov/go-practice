# Kun 9: Atomic Operations va Error Handling

## Kirish

Go dasturlash tilida parallel dasturlashda **Atomic operations** va **Error handling** juda muhim. Atomic operations race condition siz operatsiyalar uchun, Error handling esa xatolarni to'g'ri boshqarish uchun ishlatiladi.

---

## 1. Atomic Operations

### Atomic nima?

**Atomic operations** - bu bir vaqtda faqat bitta goroutine tomonidan bajariladigan operatsiyalar. Bu operatsiyalar CPU darajasida atomik (bo'linmas) bo'ladi.

**Atomic** - bu mutex dan tezroq va oddiy counter, flag kabi operatsiyalar uchun samarali.

### Atomic nima uchun kerak?

1. **Race Condition ni oldini olish**
   - Mutex dan tezroq
   - Oddiy operatsiyalar uchun (increment, decrement)
   - Performance muhim bo'lsa

2. **Lock-free operatsiyalar**
   - Mutex lock qilish kerak emas
   - Past overhead
   - Tez ishlash

3. **Counter va Flag operatsiyalari**
   - Oddiy counter increment/decrement
   - Boolean flag
   - Uint32, Int64 operatsiyalari

### Atomic vazifasi

Atomic quyidagi vazifalarni bajaradi:

1. **Atomik operatsiyalar** - CPU darajasida atomik
2. **Race condition siz** - Mutex siz xavfsiz
3. **Tez ishlash** - Mutex dan tezroq
4. **Lock-free** - Lock qilish kerak emas

### Atomic afzalliklari

✅ **Tez** - Mutex dan tezroq
✅ **Oddiy** - Oson ishlatish
✅ **Lock-free** - Lock qilish kerak emas
✅ **Performance** - Past overhead

### Atomic kamchiliklari

❌ **Cheklangan** - Faqat oddiy operatsiyalar
❌ **Type cheklangan** - Int32, Int64, Uint32, Uint64, Pointer
❌ **Murakkab operatsiyalar uchun emas** - Faqat oddiy operatsiyalar

### Atomic operatsiyalari

#### AddInt64 - Qo'shish

```go
var counter int64
atomic.AddInt64(&counter, 1)  // Increment
atomic.AddInt64(&counter, -1) // Decrement
```

#### LoadInt64 - O'qish

```go
value := atomic.LoadInt64(&counter)
```

#### StoreInt64 - Yozish

```go
atomic.StoreInt64(&counter, 100)
```

#### SwapInt64 - Almashtirish

```go
oldValue := atomic.SwapInt64(&counter, 200)
// Eski qiymatni qaytaradi va yangi qiymatni o'rnatadi
```

#### CompareAndSwapInt64 - Taqqoslash va almashtirish

```go
swapped := atomic.CompareAndSwapInt64(&value, 10, 20)
// value == 10 bo'lsa, 20 ga o'zgartir va true qaytar
// Aks holda false qaytar
```

### Atomic turlari

- `int32`, `int64` - Integer operatsiyalar
- `uint32`, `uint64` - Unsigned integer operatsiyalar
- `uintptr` - Pointer operatsiyalar
- `Value` - Har qanday type uchun

### Atomic vs Mutex

**Atomic qachon:**
- Oddiy counter, flag
- Performance muhim
- Lock-free kerak

**Mutex qachon:**
- Murakkab operatsiyalar
- Bir nechta o'zgaruvchi
- Critical section

**Taqqoslash:**
- Atomic tezroq (2-3 marta)
- Mutex ko'proq imkoniyatlar

### Amaliy misollar

#### Counter

```go
type Counter struct {
    value int64
}

func (c *Counter) Increment() {
    atomic.AddInt64(&c.value, 1)
}

func (c *Counter) Get() int64 {
    return atomic.LoadInt64(&c.value)
}
```

#### Rate Limiter

```go
type RateLimiter struct {
    count int64
    limit int64
}

func (r *RateLimiter) Allow() bool {
    current := atomic.AddInt64(&r.count, 1)
    if current > r.limit {
        atomic.AddInt64(&r.count, -1)
        return false
    }
    return true
}
```

---

## 2. Error Handling

### Error nima?

**Error** - dasturda yuzaga keladigan noto'g'ri holat yoki muammo.

Go da error - bu `error` interface turi.

### Error nima uchun kerak?

1. **Xatolarni boshqarish**
   - Dastur xato ishlaganda
   - Ma'lumot topilmasa
   - Network xatosi

2. **Xatolarni tarqatish**
   - Funksiyadan funksiyaga
   - Goroutine dan goroutine ga
   - Layer dan layer ga

3. **Xatolarni log qilish**
   - Debug qilish uchun
   - Monitoring uchun
   - Troubleshooting uchun

### Error vazifasi

Error quyidagi vazifalarni bajaradi:

1. **Xatolarni ko'rsatish** - Nima noto'g'ri ketganini
2. **Xatolarni tarqatish** - Funksiyalar orasida
3. **Xatolarni boshqarish** - Xatolarga javob berish
4. **Xatolarni log qilish** - Debug va monitoring

### Error afzalliklari

✅ **Oddiy** - Oson ishlatish
✅ **Type-safe** - Compile time da tekshiriladi
✅ **Explicit** - Aniq ko'rinadi
✅ **Flexible** - Har qanday xato

### Error kamchiliklari

❌ **Ko'p kod** - Har safar tekshirish kerak
❌ **Verbose** - Ko'p kod yozish kerak
❌ **Error handling** - Har safar boshqarish kerak

### Error yaratish

#### errors.New

```go
err := errors.New("xato xabari")
```

#### fmt.Errorf

```go
err := fmt.Errorf("xato: %s", "ma'lumot topilmadi")
```

#### Custom Error

```go
type CustomError struct {
    Code    int
    Message string
}

func (e *CustomError) Error() string {
    return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}
```

### Error tekshirish

#### if err != nil

```go
result, err := someFunction()
if err != nil {
    return err
}
```

#### errors.Is

```go
if errors.Is(err, targetErr) {
    // Xato topildi
}
```

#### errors.As

```go
var customErr *CustomError
if errors.As(err, &customErr) {
    // Custom error topildi
    fmt.Println(customErr.Code)
}
```

### Error wrapping

#### fmt.Errorf bilan

```go
baseErr := errors.New("asosiy xato")
wrappedErr := fmt.Errorf("qo'shimcha: %w", baseErr)
```

#### errors.Unwrap

```go
unwrapped := errors.Unwrap(wrappedErr)
```

### Goroutine larda Error

#### Channel orqali

```go
errChan := make(chan error, 1)

go func() {
    errChan <- someFunction()
}()

err := <-errChan
if err != nil {
    // Xato bilan ishlash
}
```

#### Result struct bilan

```go
type Result struct {
    Value int
    Error error
}

results := make(chan Result, 1)

go func() {
    value, err := someFunction()
    results <- Result{Value: value, Error: err}
}()

result := <-results
if result.Error != nil {
    // Xato bilan ishlash
}
```

### Panic va Recovery

#### Panic

```go
panic("kritik xato")
```

#### Recovery

```go
defer func() {
    if r := recover(); r != nil {
        fmt.Printf("Panic recovery: %v\n", r)
    }
}()

// Kod
```

### Context bilan Error

```go
ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
defer cancel()

select {
case err := <-done:
    if err != nil {
        return err
    }
case <-ctx.Done():
    return ctx.Err()
}
```

### Muhim qoidalar

1. **Error ni e'tiborsiz qoldirmaslik**
   ```go
   // ❌ Noto'g'ri
   someFunction()
   
   // ✅ To'g'ri
   err := someFunction()
   if err != nil {
       return err
   }
   ```

2. **Error ni qaytarish**
   ```go
   func f() error {
       if err != nil {
           return err
       }
       return nil
   }
   ```

3. **Error wrapping**
   ```go
   return fmt.Errorf("qo'shimcha: %w", err)
   ```

4. **errors.Is va errors.As ishlatish**
   ```go
   if errors.Is(err, targetErr) {
       // ...
   }
   ```

5. **Panic ni kam ishlatish**
   - Faqat kritik xatolarda
   - Recovery bilan ishlatish

---

## Xulosa

### Atomic Operations

**Nima:** Race condition siz operatsiyalar

**Vazifasi:**
- Atomik operatsiyalar
- Lock-free
- Tez ishlash

**Afzalliklari:**
- Tez
- Oddiy
- Lock-free

**Kamchiliklari:**
- Cheklangan operatsiyalar
- Type cheklangan

### Error Handling

**Nima:** Xatolarni boshqarish mexanizmi

**Vazifasi:**
- Xatolarni ko'rsatish
- Xatolarni tarqatish
- Xatolarni boshqarish

**Afzalliklari:**
- Oddiy
- Type-safe
- Explicit

**Kamchiliklari:**
- Ko'p kod
- Verbose

### Keyingi qadamlar

- Advanced error handling patterns
- Error logging va monitoring
- Testing error cases
- Performance optimization
