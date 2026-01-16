# Kun 12: Generics (Umumiy turlar)

## Kirish

Go dasturlash tilida **Generics** (Go 1.18+) - bu turli type lar bilan ishlash imkoniyatini beruvchi kuchli xususiyat. Generics orqali type-safe va qayta ishlatiladigan kod yozish mumkin.

---

## 1. Generics Asoslari

### Generics nima?

**Generics** - bu turli type lar bilan ishlash imkoniyatini beruvchi mexanizm.

**Generics** - bu type parametrlar orqali kod qayta ishlatilishi.

### Generics nima uchun kerak?

1. **Type Safety**
   - Compile time da type tekshirish
   - Runtime xatolarni oldini olish

2. **Code Reusability**
   - Bir xil kod turli type lar uchun
   - DRY (Don't Repeat Yourself) prinsipi

3. **Performance**
   - Type assertion overhead yo'q
   - Compile time da optimizatsiya

### Generics vazifasi

Generics quyidagi vazifalarni bajaradi:

1. **Type Safety** - Compile time da type tekshirish
2. **Code Reusability** - Kod qayta ishlatilishi
3. **Performance** - Type assertion overhead yo'q
4. **Flexibility** - Turli type lar bilan ishlash

### Generics afzalliklari

✅ **Type Safety** - Compile time da type tekshirish
✅ **Code Reusability** - Kod qayta ishlatilishi
✅ **Performance** - Type assertion overhead yo'q
✅ **Flexibility** - Turli type lar bilan ishlash

### Generics kamchiliklari

❌ **Murakkab** - Tushunish qiyin
❌ **Syntax** - Yangi syntax o'rganish kerak
❌ **Go 1.18+** - Eski versiyalarda ishlamaydi

---

## 2. Type Parameters

### Type Parameters E'lon Qilish

**Type parameters** - funksiya yoki type uchun type parametrlar.

```go
// Funksiya uchun type parameters
func Print[T any](value T) {
    fmt.Println(value)
}

// Type uchun type parameters
type Stack[T any] struct {
    items []T
}
```

**Syntax:** `[TypeName Constraint]`

### Type Constraints

**Type constraint** - type parametr uchun cheklovlar.

```go
// any - har qanday type
func Print[T any](value T) {
    fmt.Println(value)
}

// comparable - solishtirish mumkin bo'lgan type lar
func Find[T comparable](slice []T, value T) int {
    for i, v := range slice {
        if v == value {
            return i
        }
    }
    return -1
}
```

**Built-in constraints:**
- `any` - har qanday type
- `comparable` - solishtirish mumkin bo'lgan type lar

### Multiple Type Parameters

**Multiple type parameters** - bir nechta type parametrlar.

```go
func Swap[T, U any](a T, b U) (U, T) {
    return b, a
}

// Ishlatish
x, y := Swap(10, "hello")
```

**Qoida:** Har bir type parametr alohida constraint ga ega bo'lishi mumkin.

### Type Parameters bilan Struct

**Struct** da type parameters ishlatish.

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
    if len(s.items) == 0 {
        var zero T
        return zero
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item
}
```

**Qoida:** Struct metodlarida ham type parametr ishlatiladi.

---

## 3. Type Constraints (Kengaytirilgan)

### Interface Constraints

**Interface constraint** - interface orqali type cheklovlar.

```go
type Number interface {
    int | int8 | int16 | int32 | int64 |
    uint | uint8 | uint16 | uint32 | uint64 |
    float32 | float64
}

func Add[T Number](a, b T) T {
    return a + b
}
```

**Union types** - `|` orqali bir nechta type larni birlashtirish.

### Method Constraints

**Method constraint** - metodlar bo'lishi kerak bo'lgan type lar.

```go
type Stringer interface {
    String() string
}

func Print[T Stringer](value T) {
    fmt.Println(value.String())
}
```

**Qoida:** Interface constraint - metodlar to'plami.

### Custom Constraints

**Custom constraint** - o'z interface constraint yaratish.

```go
type Addable interface {
    int | float64 | string
}

func Sum[T Addable](values []T) T {
    var sum T
    for _, v := range values {
        sum += v
    }
    return sum
}
```

**Qoida:** Custom constraint - interface va union types.

### ~ Operator

**~ Operator** - underlying type lar uchun.

```go
type MyInt int

type Integer interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64
}

func Double[T Integer](value T) T {
    return value * 2
}
```

**~ Operator** - underlying type ni tekshiradi.

---

## 4. Generics bilan Funksiyalar

### Generic Funksiyalar

**Generic function** - type parametrlar bilan funksiya.

```go
func Map[T, U any](slice []T, fn func(T) U) []U {
    result := make([]U, len(slice))
    for i, v := range slice {
        result[i] = fn(v)
    }
    return result
}

// Ishlatish
numbers := []int{1, 2, 3, 4, 5}
doubled := Map(numbers, func(n int) int {
    return n * 2
})
```

**Generic function** - turli type lar bilan ishlaydi.

### Generic Filter

**Generic filter** - shart bo'yicha filtrlash.

```go
func Filter[T any](slice []T, fn func(T) bool) []T {
    var result []T
    for _, v := range slice {
        if fn(v) {
            result = append(result, v)
        }
    }
    return result
}

// Ishlatish
numbers := []int{1, 2, 3, 4, 5}
evens := Filter(numbers, func(n int) bool {
    return n%2 == 0
})
```

**Generic filter** - shart bo'yicha filtrlash.

### Generic Reduce

**Generic reduce** - slice ni bitta qiymatga qisqartirish.

```go
func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
    result := initial
    for _, v := range slice {
        result = fn(result, v)
    }
    return result
}

// Ishlatish
numbers := []int{1, 2, 3, 4, 5}
sum := Reduce(numbers, 0, func(acc, n int) int {
    return acc + n
})
```

**Generic reduce** - slice ni bitta qiymatga qisqartirish.

### Generic Find

**Generic find** - slice dan qiymat topish.

```go
func Find[T comparable](slice []T, value T) (T, bool) {
    for _, v := range slice {
        if v == value {
            return v, true
        }
    }
    var zero T
    return zero, false
}
```

**Generic find** - comparable type lar uchun.

---

## 5. Generics bilan Data Structures

### Generic Stack

**Generic stack** - LIFO (Last In First Out) ma'lumot tuzilmasi.

```go
type Stack[T any] struct {
    items []T
}

func NewStack[T any]() *Stack[T] {
    return &Stack[T]{items: make([]T, 0)}
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, true
}
```

**Generic stack** - turli type lar uchun.

### Generic Queue

**Generic queue** - FIFO (First In First Out) ma'lumot tuzilmasi.

```go
type Queue[T any] struct {
    items []T
}

func NewQueue[T any]() *Queue[T] {
    return &Queue[T]{items: make([]T, 0)}
}

func (q *Queue[T]) Enqueue(item T) {
    q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
    if len(q.items) == 0 {
        var zero T
        return zero, false
    }
    item := q.items[0]
    q.items = q.items[1:]
    return item, true
}
```

**Generic queue** - turli type lar uchun.

### Generic Map

**Generic map** - key-value juftliklar.

```go
type KeyValue[K comparable, V any] struct {
    Key   K
    Value V
}

type GenericMap[K comparable, V any] struct {
    items map[K]V
}

func NewGenericMap[K comparable, V any]() *GenericMap[K, V] {
    return &GenericMap[K, V]{items: make(map[K]V)}
}

func (m *GenericMap[K, V]) Set(key K, value V) {
    m.items[key] = value
}

func (m *GenericMap[K, V]) Get(key K) (V, bool) {
    value, ok := m.items[key]
    return value, ok
}
```

**Generic map** - turli key va value type lar uchun.

---

## 6. Generics bilan Interface

### Generic Interface

**Generic interface** - type parametrlar bilan interface.

```go
type Container[T any] interface {
    Add(T)
    Get() (T, bool)
    Size() int
}

type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Add(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Get() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, true
}

func (s *Stack[T]) Size() int {
    return len(s.items)
}
```

**Generic interface** - turli type lar uchun interface.

### Type Assertion bilan Generics

**Type assertion** - generic type dan konkret type ga.

```go
func Process[T any](value T) {
    switch v := any(value).(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    default:
        fmt.Printf("Unknown: %v\n", v)
    }
}
```

**Type assertion** - generic type dan konkret type ga.

---

## 7. Muhim Qoidalar

### Type Inference

**Type inference** - type ni avtomatik aniqlash.

```go
// Type inference
Print(42)        // T = int
Print("hello")   // T = string

// Explicit type
Print[int](42)
```

**Qoida:** Go type ni avtomatik aniqlaydi.

### Zero Value

**Zero value** - type parametr uchun zero value.

```go
func GetZero[T any]() T {
    var zero T
    return zero
}
```

**Qoida:** `var zero T` - type uchun zero value.

### Pointer Types

**Pointer types** - generic type lar bilan pointer.

```go
func Modify[T any](value *T) {
    // value ni o'zgartirish
}
```

**Qoida:** Pointer types ham generic bo'lishi mumkin.

### Slice Types

**Slice types** - generic type lar bilan slice.

```go
func Process[T any](slice []T) {
    // slice ni qayta ishlash
}
```

**Qoida:** Slice types ham generic bo'lishi mumkin.

---

## Xulosa

### Generics

**Nima:** Type parametrlar orqali kod qayta ishlatilishi

**Vazifasi:**
- Type safety
- Code reusability
- Performance
- Flexibility

**Afzalliklari:**
- Type safety
- Code reusability
- Performance
- Flexibility

**Kamchiliklari:**
- Murakkab
- Syntax o'rganish kerak
- Go 1.18+ talab qiladi

### Keyingi qadamlar

- Advanced generics patterns
- Performance optimization
- Real-world examples
- Generics bilan testing
