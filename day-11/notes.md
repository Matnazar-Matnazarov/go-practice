# Kun 11: Funksiyalar va Metodlar (Pointer, Interface)

## Kirish

Go dasturlash tilida **Funksiyalar** va **Metodlar** bilan ishlash - professional kod yozishning asosiy qismi. Pointer va Interface bilan birgalikda ishlatilganda, kuchli va moslashuvchan kod yozish mumkin.

---

## 1. Metodlar (Kengaytirilgan)

### Method Sets

**Method set** - bu type uchun mavjud bo'lgan barcha metodlar to'plami.

#### Value type method set

```go
type Person struct {
    Name string
}

func (p Person) GetName() string { return p.Name }
func (p *Person) SetName(name string) { p.Name = name }
```

Value type method set:
- Barcha value receiver metodlar
- Barcha pointer receiver metodlar (Go avtomatik qo'shadi)

#### Pointer type method set

```go
var p *Person
```

Pointer type method set:
- Barcha value receiver metodlar
- Barcha pointer receiver metodlar

**Qoida:** Go avtomatik konvertatsiya qiladi:
- `p.GetName()` = `(*p).GetName()`
- `p.SetName()` = `(&p).SetName()`

### Pointer Receiver Qoidalari

#### Qachon Pointer Receiver ishlatish kerak?

1. **O'zgartirish kerak bo'lsa**
   ```go
   func (c *Counter) Increment() {
       c.value++  // Asl o'zgaradi
   }
   ```

2. **Struct katta bo'lsa**
   ```go
   func (p *LargeStruct) Process() {
       // Nusxa yaratish qimmat
   }
   ```

3. **Consistency uchun**
   - Agar bitta metod pointer receiver bo'lsa, barchasi bo'lishi kerak

#### Qachon Value Receiver ishlatish kerak?

1. **Faqat o'qish kerak bo'lsa**
   ```go
   func (p Person) GetName() string {
       return p.Name
   }
   ```

2. **Struct kichik bo'lsa**
   - Kichik struct lar uchun value receiver yaxshi

3. **Immutability uchun**
   - O'zgartirishni oldini olish

### Method Expressions

**Method expression** - metodni funksiya sifatida ishlatish.

```go
p := Person{Name: "Ali"}

// Method expression
getNameFunc := Person.GetName
fmt.Println(getNameFunc(p))

setNameFunc := (*Person).SetName
setNameFunc(&p, "Aliyev")
```

**Method expression** - `Type.MethodName` formatida.

### Method Values

**Method value** - o'zgaruvchiga bog'langan metod.

```go
p := Person{Name: "Ali"}

// Method value
getName := p.GetName
fmt.Println(getName())  // "Ali"

// Method value - closure kabi ishlaydi
counter := Counter{value: 10}
increment := counter.IncrementPtr
increment()
```

**Method value** - closure kabi ishlaydi, tashqi o'zgaruvchilarga kirish mumkin.

### Method Chaining

**Method chaining** - metodlarni ketma-ket chaqirish.

```go
type Calculator struct {
    result float64
}

func (c *Calculator) Add(n float64) *Calculator {
    c.result += n
    return c  // Method chaining uchun
}

// Ishlatish
calc := &Calculator{result: 0}
result := calc.Add(10).Multiply(2).Subtract(5).GetResult()
```

**Qoida:** Method chaining uchun metod pointer qaytarishi kerak.

---

## 2. Funksiyalar (Pointer va Interface)

### Funksiya Turlari (Function Types)

**Function type** - funksiyani type sifatida e'lon qilish.

```go
// Funksiya turi
type Operation func(int, int) int

// Ishlatish
var add Operation = func(a, b int) int {
    return a + b
}

result := add(5, 3)
```

**Afzalliklari:**
- Kod qayta ishlatilishi mumkin
- Type safety
- Parametr sifatida uzatish oson

### Funksiya Qiymatlari (Function Values)

**Function value** - funksiyani o'zgaruvchiga saqlash.

```go
// Funksiyani o'zgaruvchiga saqlash
add := func(a, b int) int {
    return a + b
}

// Funksiyani boshqa o'zgaruvchiga nusxalash
addCopy := add

// Funksiyani slice ga saqlash
operations := []Operation{
    func(a, b int) int { return a + b },
    func(a, b int) int { return a - b },
}
```

**Function value** - birinchi darajali fuqarolik (first-class citizen).

### Closure bilan Pointer

**Closure** - tashqi o'zgaruvchilarga kirish imkoniyati.

```go
// Closure - tashqi o'zgaruvchilarga kirish
counter := 0
increment := func() {
    counter++
}

increment()
increment()
// counter = 2

// Closure bilan pointer
value := 10
modify := func() {
    value = 20  // Tashqi o'zgaruvchini o'zgartirish
}
```

**Eslatma:** Closure tashqi o'zgaruvchilarga reference orqali kirishadi.

### Higher-Order Funksiyalar bilan Interface

**Higher-order function** - funksiyani parametr yoki qaytarish sifatida ishlatish.

```go
// Interface bilan ishlash
type Processor interface {
    Process(int) int
}

// Funksiya interface ni implement qiladi
type FuncProcessor func(int) int

func (f FuncProcessor) Process(n int) int {
    return f(n)
}

// Higher-order funksiya
process := func(p Processor, n int) int {
    return p.Process(n)
}
```

**Afzalliklari:**
- Kod qayta ishlatilishi mumkin
- Testing oson
- Flexibility

### Callback Funksiyalar

**Callback** - funksiya boshqa funksiyaga uzatiladi.

```go
type Callback func(string)

process := func(name string, callback Callback) {
    fmt.Printf("Processing: %s\n", name)
    callback("Done")
}

// Ishlatish
process("Task 1", func(msg string) {
    fmt.Printf("Callback: %s\n", msg)
})
```

**Callback** - asinxron operatsiyalar uchun ishlatiladi.

### Funksiya Parametr sifatida

**Function as parameter** - funksiyani parametr sifatida uzatish.

```go
// Map funksiyasi
func mapFunc(slice []int, fn func(int) int) []int {
    result := make([]int, len(slice))
    for i, v := range slice {
        result[i] = fn(v)
    }
    return result
}

// Ishlatish
doubled := mapFunc(numbers, func(n int) int {
    return n * 2
})
```

**Map, Filter, Reduce** - funksional dasturlash patterns.

### Funksiya Qaytarish

**Function as return** - funksiyani qaytarish.

```go
// Funksiya qaytarish
createAdder := func(n int) func(int) int {
    return func(x int) int {
        return x + n
    }
}

add5 := createAdder(5)
add10 := createAdder(10)

fmt.Println(add5(10))   // 15
fmt.Println(add10(10))   // 20
```

**Function factory** - funksiya yaratish funksiyasi.

---

## 3. Interfaces (Kengaytirilgan)

### Interface Implementatsiyasi

**Interface implementation** - type interface metodlarini implement qiladi.

```go
type Reader interface {
    Read() string
}

type Writer interface {
    Write(string)
}

// File - ReadWriter interface ni implement qiladi
type File struct {
    content string
}

func (f *File) Read() string {
    return f.content
}

func (f *File) Write(s string) {
    f.content = s
}

// Ishlatish
var rw ReadWriter = &File{content: "Hello"}
```

**Qoida:** Barcha metodlar implement qilinishi kerak.

### Type Assertion bilan Metodlar

**Type assertion** - interface dan konkret type ga o'tkazish.

```go
var p Processor

p = &Adder{value: 10}

// Type assertion
if adder, ok := p.(*Adder); ok {
    fmt.Printf("Type: Adder, value: %d\n", adder.value)
}

// Type switch
switch v := p.(type) {
case *Adder:
    fmt.Printf("Adder: %d\n", v.value)
case *Multiplier:
    fmt.Printf("Multiplier: %d\n", v.factor)
}
```

**Type assertion** - interface dan konkret type ga o'tkazish.

### Interface Conversion

**Interface conversion** - bir interface dan boshqasiga o'tkazish.

```go
// File - ReadWriter
file := &File{content: "Hello"}
var rw ReadWriter = file

// ReadWriter dan Reader ga
var r Reader = rw
fmt.Println(r.Read())

// ReadWriter dan Writer ga
var w Writer = rw
w.Write("World")
```

**Interface conversion** - interface composition orqali.

### Interface Nil Tekshirish

**Interface nil check** - interface nil ekanligini tekshirish.

```go
var r Reader
fmt.Println(r == nil)  // true

// Nil pointer interface ga uzatilsa
var f *File
r = f
fmt.Println(r == nil)  // false! (nil pointer, lekin interface nil emas)

// To'g'ri nil tekshirish
if r != nil {
    if file, ok := r.(*File); ok && file != nil {
        // File is not nil
    }
}
```

**Eslatma:** Nil pointer interface ga uzatilsa, interface nil bo'lmaydi!

### Interface Composition bilan Metodlar

**Interface composition** - bir nechta interface ni birlashtirish.

```go
type ReadWriter interface {
    Reader
    Writer
}

type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}

// Ishlatish
var rwc ReadWriteCloser = &File{content: "Content"}
rwc.Read()
rwc.Write("New")
rwc.Close()
```

**Composition** - interface larni birlashtirish.

---

## 4. Design Patterns

### Builder Pattern

**Builder pattern** - murakkab obyektlarni qurish.

```go
type QueryBuilder struct {
    table  string
    where  string
    order  string
    limit  int
}

// Builder metodlari
qb := table("users")
where(qb, "age > 18")
orderBy(qb, "name")
limit(qb, 10)
query := build(qb)
```

**Builder pattern** - murakkab obyektlarni bosqichma-bosqich qurish.

### Strategy Pattern

**Strategy pattern** - algoritmlarni almashtirish.

```go
type SortStrategy interface {
    Sort([]int) []int
}

type Sorter struct {
    strategy SortStrategy
}

func (s *Sorter) SetStrategy(strategy SortStrategy) {
    s.strategy = strategy
}

// Ishlatish
sorter := &Sorter{}
sorter.SetStrategy(&BubbleSort{})
sorted := sorter.Sort(numbers)
```

**Strategy pattern** - runtime da algoritmni tanlash.

### Observer Pattern

**Observer pattern** - o'zgarishlarni kuzatish.

```go
type Observer interface {
    Update(string)
}

type Subject struct {
    observers []Observer
    state     string
}

func (s *Subject) Attach(observer Observer) {
    s.observers = append(s.observers, observer)
}

func (s *Subject) Notify() {
    for _, observer := range s.observers {
        observer.Update(s.state)
    }
}
```

**Observer pattern** - o'zgarishlarni barcha kuzatuvchilarga xabar berish.

---

## Xulosa

### Metodlar

**Nima:** Type ga bog'langan funksiyalar

**Vazifasi:**
- Method sets va pointer receiver
- Method expressions va method values
- Method chaining

**Afzalliklari:**
- Type safety
- Code organization
- Reusability

**Kamchiliklari:**
- Value vs Pointer receiver tanlash qiyin
- Method sets qoidalari murakkab

### Funksiyalar

**Nima:** Birinchi darajali fuqarolik (first-class citizen)

**Vazifasi:**
- Function types va function values
- Closure bilan pointer
- Higher-order functions
- Callback functions

**Afzalliklari:**
- Flexibility
- Code reusability
- Functional programming

**Kamchiliklari:**
- Performance overhead (kichik)
- Debug qiyin

### Interfaces

**Nima:** Metodlar to'plami (kontrakt)

**Vazifasi:**
- Interface implementation
- Type assertion
- Interface conversion
- Interface composition

**Afzalliklari:**
- Polymorphism
- Abstraction
- Testing
- Flexibility

**Kamchiliklari:**
- Performance overhead (kichik)
- Type assertion kerak

### Keyingi qadamlar

- Advanced design patterns
- Performance optimization
- Testing with interfaces
- Real-world examples
