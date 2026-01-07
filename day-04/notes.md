# Kun 4: Map (Xarita) va For Loop (Tsikllar)

## Map (Xarita) Ma'lumot Tuzilmasi

### Map nima?

Map - kalit-qiymat juftligi saqlaydigan ma'lumot tuzilmasi. 
Hash table asosida ishlaydi va tez qidirish imkonini beradi.

### Map e'lon qilish

```go
// Usul 1: make() bilan
var ages map[string]int
ages = make(map[string]int)

// Usul 2: Literal sintaksis
ages := map[string]int{
    "Matnazar": 25,
    "Ali":      30,
}

// Usul 3: Qisqa sintaksis
ages := make(map[string]int)
ages["Matnazar"] = 25
```

### Map operatsiyalari

- **Element qo'shish**: `map[key] = value`
- **Element o'qish**: `value := map[key]`
- **Element o'chirish**: `delete(map, key)`
- **Uzunlik**: `len(map)`
- **Element borligini tekshirish**: `value, exists := map[key]`

### Map xususiyatlari

1. **Reference type** - nusxalashda reference olinadi
2. **Dinamik** - elementlar qo'shish va o'chirish mumkin
3. **Tez qidirish** - O(1) vaqt murakkabligi
4. **Unordered** - elementlar tartibsiz saqlanadi
5. **Nil map** - nil map ga yozib bo'lmaydi

### Map iteratsiyasi

```go
// Kalit va qiymat
for key, value := range map {
    // kod
}

// Faqat kalit
for key := range map {
    // kod
}

// Faqat qiymat
for _, value := range map {
    // kod
}
```

### Map turlari

- **String to int**: `map[string]int`
- **Int to string**: `map[int]string`
- **String to slice**: `map[string][]int`
- **Nested map**: `map[string]map[string]int`

### Map amaliy misollar

- So'zlar sonini hisoblash
- Eng katta/kichik qiymat topish
- Map larni birlashtirish
- Map ni filtrlash
- Kalitlar va qiymatlarni olish
- Map ni teskari qilish

## For Loop (Tsikllar)

### For loop turlari

#### 1. Oddiy for loop

```go
for initialization; condition; increment {
    // kod
}
```

**Misol:**
```go
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}
```

#### 2. Shartli for loop

```go
for condition {
    // kod
}
```

**Misol:**
```go
i := 1
for i <= 5 {
    fmt.Println(i)
    i++
}
```

#### 3. Cheksiz for loop

```go
for {
    // kod
    if condition {
        break
    }
}
```

**Misol:**
```go
count := 0
for {
    count++
    if count > 5 {
        break
    }
    fmt.Println(count)
}
```

### For range loop

For range - array, slice, string, map bilan ishlash uchun qulay.

#### Array bilan

```go
arr := [5]int{1, 2, 3, 4, 5}

// Index va value
for index, value := range arr {
    fmt.Printf("Index %d: %d\n", index, value)
}

// Faqat value
for _, value := range arr {
    fmt.Println(value)
}

// Faqat index
for index := range arr {
    fmt.Println(index)
}
```

#### Slice bilan

```go
slice := []int{1, 2, 3, 4, 5}
for index, value := range slice {
    fmt.Printf("Index %d: %d\n", index, value)
}
```

#### String bilan

```go
text := "Go"

// Byte (index va byte)
for i := 0; i < len(text); i++ {
    fmt.Printf("%c", text[i])
}

// Rune (index va rune)
for index, char := range text {
    fmt.Printf("Index %d: %c\n", index, char)
}
```

#### Map bilan

```go
ages := map[string]int{
    "Matnazar": 25,
    "Ali":      30,
}

// Key va value
for name, age := range ages {
    fmt.Printf("%s: %d\n", name, age)
}

// Faqat key
for name := range ages {
    fmt.Println(name)
}
```

### Ichki for looplar (Nested Loops)

```go
for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
        fmt.Printf("%d x %d = %d\n", i, j, i*j)
    }
}
```

### Break va Continue

#### Break

Loop dan chiqish:

```go
for i := 1; i <= 10; i++ {
    if i > 5 {
        break
    }
    fmt.Println(i)
}
```

#### Continue

Keyingi iteratsiyaga o'tish:

```go
for i := 1; i <= 10; i++ {
    if i%2 == 0 {
        continue
    }
    fmt.Println(i)
}
```

### Label bilan Break va Continue

Ichki looplardan tashqaridagi loop ni boshqarish:

```go
OuterLoop:
for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
        if i == 2 && j == 2 {
            break OuterLoop
        }
        fmt.Printf("i=%d, j=%d\n", i, j)
    }
}
```

### For loop amaliy misollar

- Array/Slice yig'indisi
- Eng katta/kichik son topish
- Juft sonlar sonini hisoblash
- Array/Slice ni teskari qilish
- Faktorial hisoblash
- Fibonacci ketma-ketligi
- Pattern chiqarish
- Qidirish algoritmlari
- Takrorlangan elementlarni olib tashlash

## Map va For Loop birgalikda

Map bilan for range loop ishlatish:

```go
ages := map[string]int{
    "Matnazar": 25,
    "Ali":      30,
}

for name, age := range ages {
    fmt.Printf("%s: %d yosh\n", name, age)
}
```

## Bugun o'rganilganlar

✅ Map (xarita) - e'lon qilish, operatsiyalar
✅ Map iteratsiyasi va element tekshirish
✅ Turli tipdagi maplar (nested, string to slice)
✅ Map nusxalash, tozalash, filtrlash
✅ For loop - oddiy, shartli, cheksiz
✅ For range - array, slice, string, map
✅ Ichki for looplar (nested loops)
✅ Break, continue, label
✅ Amaliy algoritmlar va misollar

## Keyingi kun

Kun 5: Struct (struktura), metodlar va interface
