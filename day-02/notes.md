# Kun 2: Funksiyalar (Functions)

## Funksiya nima?

Funksiya - ma'lum bir vazifani bajaradigan kod bloki. 
Funksiyalar kodni qayta ishlatishga, tuzilishga va oson tushunishga yordam beradi.

## Funksiya sintaksisi

```go
func functionName(param1 type, param2 type) returnType {
    // Kod
    return value
}
```

## Funksiya turlari

### 1. Oddiy funksiya (Simple Function)

Parametr va qaytarish qiymati bo'lmagan funksiya:

```go
func greet() {
    fmt.Println("Salom!")
}
```

### 2. Parametrli funksiya (Function with Parameters)

Funksiyaga ma'lumot uzatish:

```go
func greetPerson(name string) {
    fmt.Printf("Salom, %s!\n", name)
}
```

### 3. Qaytarish qiymatli funksiya (Function with Return Value)

Funksiyadan natija olish:

```go
func add(a int, b int) int {
    return a + b
}
```

### 4. Ko'p parametrli funksiya (Multiple Parameters)

Bir nechta parametr:

```go
func sum(x int, y int, z int) int {
    return x + y + z
}
```

### 5. Ko'p qaytarish qiymati (Multiple Return Values)

Go'da funksiya bir nechta qiymat qaytarishi mumkin:

```go
func divideAndRemainder(a int, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}
```

### 6. Nomlangan qaytarish qiymatlar (Named Return Values)

Qaytarish qiymatlariga nom berish:

```go
func calculate(a int, b int) (sum int, difference int) {
    sum = a + b
    difference = a - b
    return // avtomatik sum va difference qaytariladi
}
```

### 7. Variadic funksiya (Variadic Function)

Cheksiz parametr qabul qiladigan funksiya:

```go
func sum(numbers ...int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}
```

## Funksiya xususiyatlari

### 1. Funksiya e'lon qilish

Funksiyani `func` kalit so'zi bilan e'lon qilamiz:

```go
func functionName() {
    // kod
}
```

### 2. Parametrlar

Funksiyaga ma'lumot uzatish uchun parametrlar ishlatiladi:

```go
func greet(name string, age int) {
    fmt.Printf("Ism: %s, Yosh: %d\n", name, age)
}
```

### 3. Qaytarish qiymati (Return)

Funksiyadan natija olish:

```go
func square(x int) int {
    return x * x
}
```

### 4. Ko'p qaytarish qiymati

Go'da funksiya bir nechta qiymat qaytarishi mumkin:

```go
func divide(a, b int) (int, int, error) {
    if b == 0 {
        return 0, 0, errors.New("nolga bo'lish mumkin emas")
    }
    return a / b, a % b, nil
}
```

### 5. Blank identifier (_)

Qaytarish qiymatini e'tiborsiz qoldirish:

```go
quotient, _, _ := divide(10, 3) // faqat quotient kerak
```

## Funksiya turlari

### 1. Oddiy funksiya

```go
func greet() {
    fmt.Println("Salom!")
}
```

### 2. Parametrli funksiya

```go
func greetPerson(name string) {
    fmt.Printf("Salom, %s!\n", name)
}
```

### 3. Qaytarish qiymatli funksiya

```go
func add(a int, b int) int {
    return a + b
}
```

### 4. Ko'p qaytarish qiymati

```go
func divideAndRemainder(a int, b int) (int, int) {
    return a / b, a % b
}
```

### 5. Nomlangan qaytarish qiymatlar

```go
func calculate(a, b int) (sum int, difference int) {
    sum = a + b
    difference = a - b
    return
}
```

## Funksiya qoidalari

1. **Funksiya nomi** - katta harf bilan boshlansa, u public (ochiq)
2. **Funksiya nomi** - kichik harf bilan boshlansa, u private (yopiq)
3. **Parametrlar** - funksiya chaqirilganda uzatiladi
4. **Return** - funksiyadan qiymat qaytarish
5. **Multiple return** - bir nechta qiymat qaytarish mumkin

## Funksiya chaqirish

```go
// Parametrsiz
greet()

// Parametrli
greetPerson("Matnazar")

// Qaytarish qiymatli
result := add(5, 3)

// Ko'p qaytarish qiymati
quotient, remainder := divideAndRemainder(10, 3)
```

## String operatsiyalari (Matn bilan ishlash)

Go'da string bilan ishlash uchun `strings` paketi ishlatiladi.

### 1. Asosiy string operatsiyalari

```go
import "strings"

text := "Go dasturlash tili"

// len() - string uzunligi
length := len(text)

// strings.Contains() - matn ichida qidirish
contains := strings.Contains(text, "Go")

// strings.HasPrefix() - boshlanishini tekshirish
hasPrefix := strings.HasPrefix(text, "Go")

// strings.HasSuffix() - oxirini tekshirish
hasSuffix := strings.HasSuffix(text, "tili")
```

### 2. String o'zgartirish

```go
// strings.ToUpper() - katta harflarga o'tkazish
upper := strings.ToUpper("go")

// strings.ToLower() - kichik harflarga o'tkazish
lower := strings.ToLower("GO")

// strings.TrimSpace() - bo'sh joylarni olib tashlash
trimmed := strings.TrimSpace("  Go  ")

// strings.Trim() - belgilarni olib tashlash
trimmed2 := strings.Trim("!!!Go!!!", "!")
```

### 3. String bo'lish va birlashtirish

```go
// strings.Split() - matnni bo'lish
words := strings.Split("Go Python Java", " ")

// strings.Join() - matnlarni birlashtirish
joined := strings.Join([]string{"Go", "Python"}, " | ")

// strings.Fields() - bo'sh joylar bo'yicha bo'lish
fields := strings.Fields("Go Python Java")
```

### 4. String almashtirish

```go
// strings.Replace() - matnni almashtirish (bitta o'rin)
replaced := strings.Replace("Go Go", "Go", "Golang", 1)

// strings.ReplaceAll() - barcha o'rinlarni almashtirish
replacedAll := strings.ReplaceAll("Go Go Go", "Go", "Golang")
```

### 5. String qidirish va solishtirish

```go
// strings.Index() - belgi pozitsiyasini topish
index := strings.Index("Go dasturlash", "dasturlash")

// strings.Count() - belgi sonini hisoblash
count := strings.Count("Go Go", "o")

// strings.Compare() - matnlarni solishtirish
// -1: birinchi kichik, 0: teng, 1: birinchi katta
result := strings.Compare("Go", "Python")

// strings.EqualFold() - katta-kichik harfni e'tiborsiz solishtirish
equal := strings.EqualFold("Go", "GO")
```

### 6. Boshqa foydali string funksiyalari

```go
// strings.Repeat() - matnni takrorlash
repeated := strings.Repeat("Go ", 3)

// strings.Title() - har bir so'zning birinchi harfini katta qilish
title := strings.Title("go dasturlash tili")
```

### 7. Amaliy misollar

```go
// Email tekshirish
func validateEmail(email string) bool {
    return strings.Contains(email, "@") && strings.Contains(email, ".")
}

// Matnni formatlash
func formatText(text string) string {
    trimmed := strings.TrimSpace(text)
    return strings.Title(trimmed)
}
```

## Bugun o'rganilganlar

✅ Funksiya sintaksisi va tuzilishi
✅ Parametrlar va qaytarish qiymatlari
✅ Ko'p qaytarish qiymatlari
✅ Nomlangan qaytarish qiymatlar
✅ Variadic funksiyalar
✅ Funksiya chaqirish va ishlatish
✅ String operatsiyalari (len, Contains, Split, Join, Replace, va boshqalar)
✅ strings paketi bilan ishlash

## Keyingi kun

Kun 3: Shartli operatorlar (if, else, switch) va tsikllar (for, while)
