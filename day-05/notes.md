# Kun 5: Struct, Funksiyalar va For Loop (Kengaytirilgan)

## Struct (Struktura)

### Struct nima?

Struct - turli tipdagi ma'lumotlarni birlashtiruvchi ma'lumot tuzilmasi. 
OOP da class ga o'xshaydi, lekin metodlar alohida.

### Struct e'lon qilish

```go
type Person struct {
    Name string
    Age  int
    City string
}
```

### Struct yaratish

```go
// Usul 1: Field nomlari bilan
person1 := Person{
    Name: "Matnazar",
    Age:  25,
    City: "Toshkent",
}

// Usul 2: Qisqa sintaksis
person2 := Person{"Ali", 30, "Samarqand"}

// Usul 3: Field nomlarsiz
person3 := Person{
    "Vali",
    28,
    "Buxoro",
}
```

### Struct fieldlarga kirish

```go
person := Person{"Matnazar", 25, "Toshkent"}
fmt.Println(person.Name)  // Matnazar
fmt.Println(person.Age)   // 25
fmt.Println(person.City)  // Toshkent

// Fieldlarni o'zgartirish
person.Age = 26
```

### Struct pointer

```go
person := &Person{"Matnazar", 25, "Toshkent"}
fmt.Println(person.Name)  // Avtomatik dereference
```

### Struct embedding (Inheritance)

Go'da klassik inheritance yo'q, lekin embedding orqali xuddi shunday ishlaydi:

```go
type Person struct {
    Name string
    Age  int
}

type Student struct {
    Person      // Embedding
    StudentID int
    Grade     string
}

student := Student{
    Person: Person{"Matnazar", 25},
    StudentID: 12345,
    Grade: "A",
}
```

### Nested struct

Struct ichida struct:

```go
type Address struct {
    Street string
    City   string
}

type User struct {
    Person  Person
    Address Address
}
```

### Struct metodlar

Struct ga metodlar qo'shish:

```go
// Value receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Pointer receiver
func (p *Person) UpdateAge(newAge int) {
    p.Age = newAge
}
```

**Value vs Pointer receiver:**
- **Value receiver** - struct nusxasida ishlaydi (o'zgarmaydi)
- **Pointer receiver** - asl struct da ishlaydi (o'zgaradi)

### Struct array, slice, map

```go
// Array
people := [3]Person{...}

// Slice
students := []Student{...}

// Map
employees := map[int]Employee{...}
```

## Funksiyalar (Kengaytirilgan)

### Funksiya kompozitsiyasi

Funksiyalarni bir-biriga uzatish:

```go
result := multiplyByTwo(addOne(sum(numbers)))
```

### Higher-order funksiyalar

Funksiyani parametr sifatida uzatish:

```go
func applyFunction(numbers []int, fn func(int) int) []int {
    result := make([]int, len(numbers))
    for i, num := range numbers {
        result[i] = fn(num)
    }
    return result
}
```

### Closure funksiyalar

Ichki funksiya tashqi o'zgaruvchilarga kirish huquqiga ega:

```go
func createCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

### Rekursiv funksiyalar

O'zini o'zi chaqiradigan funksiyalar:

```go
func factorialRecursive(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorialRecursive(n-1)
}
```

### Defer funksiyalar

Funksiya tugagunga qadar kechiktiriladi:

```go
defer fmt.Println("Bu oxirida chiqadi")
fmt.Println("Bu birinchi chiqadi")
```

**Defer xususiyatlari:**
- LIFO (Last In First Out) - teskari tartibda ishlaydi
- Resource tozalash uchun qulay
- Error handling da foydali

### Anonim funksiyalar

Nomisiz funksiyalar:

```go
func() {
    fmt.Println("Anonim funksiya")
}()

// O'zgaruvchiga berish
add := func(a, b int) int {
    return a + b
}
```

### Variadic funksiyalar

Cheksiz parametr qabul qilish:

```go
func variadicSum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Slice ni uzatish
sum := variadicSum(slice...)
```

### Error handling

Funksiyadan error qaytarish:

```go
func divideSafe(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("nolga bo'lish mumkin emas")
    }
    return a / b, nil
}

// Ishlatish
result, err := divideSafe(10, 2)
if err != nil {
    fmt.Println("Xato:", err)
}
```

## For Loop (Kengaytirilgan)

### Loop optimizatsiyasi

```go
// Oddiy loop
for i := 0; i < len(numbers); i++ {
    // kod
}

// Range loop (optimallashtirilgan)
for _, num := range numbers {
    // kod
}
```

### Qidirish algoritmlari

**Linear search:**
```go
func linearSearch(arr []int, target int) int {
    for i, value := range arr {
        if value == target {
            return i
        }
    }
    return -1
}
```

**Binary search:**
```go
func binarySearch(arr []int, target int) int {
    left, right := 0, len(arr)-1
    for left <= right {
        mid := (left + right) / 2
        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}
```

### Sort algoritmlari

**Bubble sort:**
```go
func bubbleSort(arr []int) []int {
    result := make([]int, len(arr))
    copy(result, arr)
    
    for i := 0; i < len(result)-1; i++ {
        for j := 0; j < len(result)-i-1; j++ {
            if result[j] > result[j+1] {
                result[j], result[j+1] = result[j+1], result[j]
            }
        }
    }
    return result
}
```

### Loop transformatsiya

**Map, Filter, Reduce pattern:**

```go
// Map - transformatsiya
doubled := make([]int, len(numbers))
for i, num := range numbers {
    doubled[i] = num * 2
}

// Filter - filtrlash
evens := []int{}
for _, num := range numbers {
    if num%2 == 0 {
        evens = append(evens, num)
    }
}

// Reduce - yig'indi
sum := 0
for _, num := range numbers {
    sum += num
}
```

### Loop performance

- **Pre-allocate slice** - capacity belgilash
- **Range loop** - index loop dan tezroq
- **Break early** - keraksiz iteratsiyalarni oldini olish

## Bugun o'rganilganlar

✅ Struct (struktura) - e'lon qilish, fieldlar
✅ Struct embedding va nested struct
✅ Struct metodlar (value va pointer receiver)
✅ Struct array, slice, map
✅ Funksiya kompozitsiyasi va higher-order funksiyalar
✅ Closure va rekursiv funksiyalar
✅ Defer funksiyalar va error handling
✅ For loop optimizatsiyasi va kengaytirilgan misollar
✅ Qidirish va sort algoritmlari
✅ Loop transformatsiya (map, filter, reduce)

## Keyingi kun

Kun 6: Pointer, Interface va Error handling
