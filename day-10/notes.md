# Kun 10: Pointers va Interfaces

## Kirish

Go dasturlash tilida **Pointers** va **Interfaces** fundamental tushunchalar. Pointers - manzil bilan ishlash, Interfaces - polimorfizm mexanizmi.

---

## 1. Pointers

### Pointer nima?

**Pointer** - bu o'zgaruvchining xotira manzilini saqlaydigan o'zgaruvchi.

**Pointer** - bu qiymat emas, balki qiymatning manzili.

### Pointer nima uchun kerak?

1. **Funksiyada o'zgartirish**
   - Funksiyaga qiymat uzatilsa, nusxa o'zgaradi
   - Pointer uzatilsa, asl o'zgaradi

2. **Katta ma'lumotlar bilan ishlash**
   - Katta struct yoki array ni nusxalash qimmat
   - Pointer orqali tezroq

3. **Memory samaradorligi**
   - Nusxa yaratish kerak emas
   - Xotira tejash

### Pointer vazifasi

Pointer quyidagi vazifalarni bajaradi:

1. **Manzil saqlash** - O'zgaruvchining xotira manzili
2. **O'zgartirish** - Funksiyada asl qiymatni o'zgartirish
3. **Samaradorlik** - Katta ma'lumotlar bilan ishlash
4. **Nil tekshirish** - Pointer nil bo'lishi mumkin

### Pointer afzalliklari

✅ **O'zgartirish** - Funksiyada asl qiymatni o'zgartirish
✅ **Samaradorlik** - Katta ma'lumotlar bilan tezroq
✅ **Memory** - Xotira tejash
✅ **Flexible** - Nil bo'lishi mumkin

### Pointer kamchiliklari

❌ **Murakkab** - Tushunish qiyin
❌ **Xavfli** - Nil pointer panic qiladi
❌ **Debug qiyin** - Manzil bilan ishlash qiyin

### Pointer operatsiyalari

#### & (Address operator)

```go
x := 42
p := &x  // x ning manzilini olish
```

#### * (Dereference operator)

```go
p := &x
*p = 100  // p manzilidagi qiymatni o'zgartirish
value := *p  // p manzilidagi qiymatni o'qish
```

#### new() funksiyasi

```go
p := new(int)  // Yangi int yaratadi va pointer qaytaradi
*p = 100
```

### Pointer turlari

- `*int` - int pointer
- `*string` - string pointer
- `*struct` - struct pointer
- `*[]int` - slice pointer (odatda kerak emas)

### Pointer nil

```go
var p *int
if p == nil {
    // Pointer nil
}
```

**Eslatma:** Nil pointer ni dereference qilish panic qiladi!

### Pointer receiver

#### Value receiver

```go
func (c Counter) Increment() {
    c.value++  // Nusxa o'zgaradi
}
```

#### Pointer receiver

```go
func (c *Counter) Increment() {
    c.value++  // Asl o'zgaradi
}
```

**Qoida:**
- O'zgartirish kerak bo'lsa - Pointer receiver
- Faqat o'qish kerak bo'lsa - Value receiver

### Pointer bilan ishlash

#### Funksiyaga uzatish

```go
func modify(x *int) {
    *x = 999
}

x := 10
modify(&x)  // x endi 999
```

#### Qaytarish

```go
func createInt(value int) *int {
    return &value
}
```

### Muhim qoidalar

1. **Nil tekshirish**
   ```go
   if p != nil {
       *p = 10
   }
   ```

2. **Pointer receiver**
   - O'zgartirish kerak bo'lsa - Pointer receiver
   - Faqat o'qish kerak bo'lsa - Value receiver

3. **Slice va Map**
   - Slice va Map allaqachon reference type
   - Pointer kerak emas (odatda)

4. **Array**
   - Array value type
   - Pointer kerak bo'lishi mumkin

---

## 2. Interfaces

### Interface nima?

**Interface** - bu metodlar to'plami. Agar type barcha metodlarni implement qilsa, u interface ni implement qiladi.

**Interface** - bu polimorfizm mexanizmi.

### Interface nima uchun kerak?

1. **Polimorfizm**
   - Turli type lar bir xil interface dan foydalanishi mumkin
   - Kod qayta ishlatilishi mumkin

2. **Abstraksiya**
   - Implementation ni yashirish
   - Faqat interface bilan ishlash

3. **Testing**
   - Mock qilish oson
   - Test qilish oson

### Interface vazifasi

Interface quyidagi vazifalarni bajaradi:

1. **Polimorfizm** - Turli type lar bir xil interface
2. **Abstraksiya** - Implementation ni yashirish
3. **Kontrakt** - Metodlar to'plami
4. **Flexibility** - Kod qayta ishlatilishi mumkin

### Interface afzalliklari

✅ **Polimorfizm** - Turli type lar bir xil interface
✅ **Abstraksiya** - Implementation ni yashirish
✅ **Testing** - Mock qilish oson
✅ **Flexibility** - Kod qayta ishlatilishi mumkin

### Interface kamchiliklari

❌ **Performance** - Kichik overhead
❌ **Murakkab** - Tushunish qiyin
❌ **Type assertion** - Type ni tekshirish kerak

### Interface e'lon qilish

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

### Interface implementatsiyasi

```go
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}
```

**Eslatma:** Go da interface ni "implement qilish" aniq yozilmaydi. Agar barcha metodlar mavjud bo'lsa, avtomatik implement qilinadi.

### Empty interface

```go
var i interface{}
i = 42
i = "hello"
i = true
```

**Empty interface** - har qanday type ni qabul qiladi.

### Type assertion

#### Type assertion (safe)

```go
s, ok := i.(string)
if ok {
    // s string
}
```

#### Type assertion (unsafe)

```go
s := i.(string)  // Panic qilishi mumkin!
```

#### Type switch

```go
switch v := i.(type) {
case string:
    fmt.Println("String:", v)
case int:
    fmt.Println("Int:", v)
default:
    fmt.Println("Unknown:", v)
}
```

### Interface composition

```go
type Reader interface {
    Read() string
}

type Writer interface {
    Write(string)
}

type ReadWriter interface {
    Reader
    Writer
}
```

**Composition** - bir nechta interface ni birlashtirish.

### Polymorphism

```go
type Animal interface {
    Sound() string
}

animals := []Animal{Dog{}, Cat{}}
for _, animal := range animals {
    fmt.Println(animal.Sound())
}
```

**Polymorphism** - turli type lar bir xil interface dan foydalanishi.

### Nil interface

```go
var i interface{}
fmt.Println(i == nil)  // true

var s *string
i = s
fmt.Println(i == nil)  // false! (nil pointer, lekin interface nil emas)
```

**Eslatma:** Nil pointer interface ga uzatilsa, interface nil bo'lmaydi!

### Stringer interface

```go
type Stringer interface {
    String() string
}

type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%s (%d yosh)", p.Name, p.Age)
}

p := Person{Name: "Ali", Age: 25}
fmt.Println(p)  // String() avtomatik chaqiriladi
```

**Stringer** - `fmt` package da avtomatik ishlatiladi.

### Error interface

```go
type error interface {
    Error() string
}

type CustomError struct {
    Code    int
    Message string
}

func (e *CustomError) Error() string {
    return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}
```

**Error** - Go da xatolarni ko'rsatish uchun ishlatiladi.

### Interface embedding

```go
type ReadWriter interface {
    Reader
    Writer
    Close()
}
```

**Embedding** - interface ichida boshqa interface larni qo'shish.

### Muhim qoidalar

1. **Interface ni implement qilish**
   - Barcha metodlar mavjud bo'lishi kerak
   - Metodlar to'g'ri signature ga ega bo'lishi kerak

2. **Type assertion**
   - Safe version ishlatish (ok tekshirish)
   - Unsafe version panic qilishi mumkin

3. **Nil interface**
   - Nil pointer interface ga uzatilsa, interface nil bo'lmaydi
   - Interface ni nil tekshirish murakkab

4. **Pointer receiver**
   - Interface metodlari pointer receiver bo'lishi mumkin
   - Value receiver ham ishlaydi

5. **Empty interface**
   - `interface{}` har qanday type ni qabul qiladi
   - Type assertion kerak

---

## Xulosa

### Pointers

**Nima:** O'zgaruvchining xotira manzili

**Vazifasi:**
- Manzil saqlash
- O'zgartirish
- Samaradorlik

**Afzalliklari:**
- O'zgartirish
- Samaradorlik
- Memory tejash

**Kamchiliklari:**
- Murakkab
- Xavfli (nil pointer)
- Debug qiyin

### Interfaces

**Nima:** Metodlar to'plami (kontrakt)

**Vazifasi:**
- Polimorfizm
- Abstraksiya
- Kontrakt

**Afzalliklari:**
- Polimorfizm
- Abstraksiya
- Testing
- Flexibility

**Kamchiliklari:**
- Performance overhead
- Murakkab
- Type assertion kerak

### Keyingi qadamlar

- Advanced interface patterns
- Type assertions va type switches
- Interface composition patterns
- Performance optimization
