# Kun 18: Testing (Test Qilish)

## Kirish

Go dasturlash tilida **Testing** - bu kodning to'g'ri ishlashini tekshirishning asosiy mexanizmi. Go'da testing built-in package bilan qo'llab-quvvatlanadi va professional dasturlashning ajralmas qismi.

---

## 1. Testing Asoslari

### Testing nima?

**Testing** - bu kodning to'g'ri ishlashini tekshirish jarayoni.

**Testing** - bu xatolarni erta aniqlash va kod sifatini ta'minlash.

### Testing nima uchun kerak?

1. **Xatolarni aniqlash**
   - Kod xatolarini erta topish
   - Regression xatolarini oldini olish

2. **Kod sifatini ta'minlash**
   - Kodning to'g'ri ishlashini kafolatlash
   - Refactoring xavfsizligi

3. **Hujjatlashtirish**
   - Kodning qanday ishlashini ko'rsatish
   - Misollar va use case'lar

### Testing vazifasi

Testing quyidagi vazifalarni bajaradi:

1. **Unit Testing** - Alohida funksiyalarni test qilish
2. **Integration Testing** - Komponentlarni birgalikda test qilish
3. **Benchmark Testing** - Performance test qilish
4. **Example Testing** - Misollar orqali test qilish

### Testing afzalliklari

✅ **Xatolarni aniqlash** - erta xatolarni topish
✅ **Kod sifatini ta'minlash** - To'g'ri ishlashni kafolatlash
✅ **Refactoring xavfsizligi** - Kodni xavfsiz o'zgartirish
✅ **Hujjatlashtirish** - Kodning qanday ishlashini ko'rsatish

---

## 2. Test Fayllari

### Test Fayl Qoidalari

**Test fayl** - `_test.go` bilan tugaydi.

```go
// math_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
    // Test kodlari
}
```

**Qoida:** Test fayl nomi `*_test.go` formatida bo'lishi kerak.

### Test Funksiyalari

**Test funksiyasi** - `Test` bilan boshlanadi.

```go
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
}
```

**Qoida:** Test funksiyasi `Test` bilan boshlanadi va `*testing.T` qabul qiladi.

### Test Ishlatish

**go test** - testlarni ishga tushirish.

```bash
go test                    # Barcha testlarni ishga tushirish
go test -v                 # Batafsil natija
go test -run TestAdd        # Faqat TestAdd ni ishga tushirish
go test ./...              # Barcha package lardagi testlar
```

**Qoida:** `go test` - testlarni ishga tushirish.

---

## 3. Table-Driven Tests

### Table-Driven Test Pattern

**Table-driven test** - bir nechta test case'larni jadvalda saqlash.

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a        int
        b        int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -2, -3, -5},
        {"zero", 0, 0, 0},
        {"mixed", -2, 3, 1},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
            }
        })
    }
}
```

**Qoida:** Table-driven test - ko'p test case'larni qisqa va tushunarli yozish.

### t.Run() - Subtest

**t.Run()** - subtest yaratish.

```go
t.Run("test name", func(t *testing.T) {
    // Test kodlari
})
```

**Qoida:** t.Run() - subtest yaratish va natijani alohida ko'rsatish.

---

## 4. Test Helpers

### Test Helper Funksiyalar

**Test helper** - test kodini qisqartirish.

```go
func assertEqual(t *testing.T, got, want int) {
    t.Helper()
    if got != want {
        t.Errorf("got %d; want %d", got, want)
    }
}

func TestAdd(t *testing.T) {
    assertEqual(t, Add(2, 3), 5)
}
```

**Qoida:** `t.Helper()` - helper funksiyani stack trace dan olib tashlash.

### Setup va Teardown

**Setup** - testdan oldin tayyorlash.

**Teardown** - testdan keyin tozalash.

```go
func TestWithSetup(t *testing.T) {
    // Setup
    db := setupDatabase(t)
    defer teardownDatabase(t, db)

    // Test
    result := db.Query("SELECT * FROM users")
    // ...
}
```

**Qoida:** Setup va teardown - test muhitini tayyorlash va tozalash.

---

## 5. Error Testing

### Error Test Qilish

**Error test** - xato holatlarini test qilish.

```go
func TestDivide(t *testing.T) {
    _, err := Divide(10, 0)
    if err == nil {
        t.Error("expected error for division by zero")
    }
}
```

**Qoida:** Error test - xato holatlarini tekshirish.

### errors.Is() va errors.As()

**errors.Is()** - xato turini tekshirish.

**errors.As()** - xato type'ini tekshirish.

```go
func TestDivide(t *testing.T) {
    _, err := Divide(10, 0)
    if !errors.Is(err, ErrDivisionByZero) {
        t.Errorf("expected ErrDivisionByZero, got %v", err)
    }
}
```

**Qoida:** errors.Is() va errors.As() - xato turini tekshirish.

---

## 6. Benchmark Testing

### Benchmark Funksiyalari

**Benchmark funksiyasi** - `Benchmark` bilan boshlanadi.

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}
```

**Qoida:** Benchmark funksiyasi `Benchmark` bilan boshlanadi va `*testing.B` qabul qiladi.

### Benchmark Ishlatish

**go test -bench** - benchmark testlarni ishga tushirish.

```bash
go test -bench=.              # Barcha benchmarklar
go test -bench=BenchmarkAdd   # Faqat BenchmarkAdd
go test -bench=. -benchmem    # Memory allocation ko'rsatish
```

**Qoida:** `go test -bench` - benchmark testlarni ishga tushirish.

### b.N - Iteratsiyalar

**b.N** - benchmark iteratsiyalari soni.

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}
```

**Qoida:** b.N - Go avtomatik aniqlaydi, qancha iteratsiya kerak.

---

## 7. Example Testing

### Example Funksiyalari

**Example funksiyasi** - `Example` bilan boshlanadi.

```go
func ExampleAdd() {
    result := Add(2, 3)
    fmt.Println(result)
    // Output: 5
}
```

**Qoida:** Example funksiyasi `Example` bilan boshlanadi va `// Output:` comment bilan natija ko'rsatiladi.

### Example Ishlatish

**go test -run Example** - example testlarni ishga tushirish.

```bash
go test -run Example
```

**Qoida:** Example testlar - hujjatlashtirish va misollar.

---

## 8. Test Coverage

### Coverage Test Qilish

**Coverage** - kodning qancha qismi test qilingan.

```bash
go test -cover              # Coverage foizini ko'rsatish
go test -coverprofile=coverage.out  # Coverage faylini yaratish
go tool cover -html=coverage.out    # HTML report yaratish
```

**Qoida:** Coverage - kodning qancha qismi test qilingan.

---

## 9. Integration Testing

### Integration Test

**Integration test** - bir nechta komponentlarni birgalikda test qilish.

```go
func TestIntegration(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping integration test")
    }
    // Integration test kodlari
}
```

**Qoida:** `testing.Short()` - qisqa testlarni o'tkazib yuborish.

---

## 10. Muhim Qoidalar

### Test Best Practices

1. **Test nomi** - tushunarli va tavsiflovchi
2. **Table-driven tests** - ko'p test case'lar uchun
3. **Test isolation** - har bir test mustaqil
4. **Test coverage** - muhim kodlar uchun yuqori coverage
5. **Error testing** - xato holatlarini test qilish

### Test Naming

**Test nomi** - tushunarli va tavsiflovchi.

```go
func TestAdd_PositiveNumbers(t *testing.T) { }
func TestAdd_NegativeNumbers(t *testing.T) { }
func TestAdd_Zero(t *testing.T) { }
```

**Qoida:** Test nomi - funksiya nomi va holatni ko'rsatishi kerak.

---

## Xulosa

### Testing

**Nima:** Kodning to'g'ri ishlashini tekshirish

**Vazifasi:**
- Unit testing
- Integration testing
- Benchmark testing
- Example testing

**Afzalliklari:**
- Xatolarni aniqlash
- Kod sifatini ta'minlash
- Refactoring xavfsizligi
- Hujjatlashtirish

### Keyingi qadamlar

- Advanced testing patterns
- Mock testing
- Test fixtures
- Test organization
- Real-world examples
