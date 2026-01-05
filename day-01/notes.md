# Kun 1: Go dasturlash tiliga kirish

## Go nima?

Go (yoki Golang) - Google tomonidan 2009-yilda yaratilgan ochiq manbali dasturlash tili. 
U tez, sodda va zamonaviy dasturlash tilidir.

## Asosiy xususiyatlar

- **Sodda sintaksis** - o'qish va yozish oson
- **Tez kompilyatsiya** - kod tez ishga tushadi
- **Kuchli tipizatsiya** - xatolarni oldini oladi
- **Garbage collection** - xotirani avtomatik boshqaradi
- **Concurrency** - parallel dasturlar yozish oson

## Go dasturining struktura

Har bir Go dasturi `package` bilan boshlanadi. 
Asosiy dastur uchun `package main` ishlatiladi.

```go
package main

import "fmt"

func main() {
    fmt.Println("Salom, Go!")
}
```

## Asosiy tushunchalar

### 1. Package (Paket)
- Kodning modullari
- `main` - dastur boshlanish nuqtasi
- Boshqa paketlar `import` orqali chaqiriladi

### 2. Import (Import)
- Tashqi paketlarni chaqirish
- `fmt` - formatlash va chiqarish uchun
- `os` - operatsion tizim bilan ishlash uchun

### 3. Function (Funksiya)
- `func` kalit so'zi bilan yoziladi
- `main()` - dastur boshlanish funksiyasi
- Har bir funksiya nomi bilan boshlanadi

### 4. O'zgaruvchilar (Variables)

Go'da 3 usulda o'zgaruvchi e'lon qilish mumkin:

```go
// 1. var bilan (tip ko'rsatiladi)
var name string = "Go"

// 2. var bilan (tip avtomatik aniqlanadi)
var age = 25

// 3. Qisqa sintaksis (faqat funksiya ichida)
city := "Toshkent"
```

### 5. Asosiy ma'lumot turlari (Data Types)

- `string` - matnlar ("salom")
- `int` - butun sonlar (42)
- `float64` - o'nlik sonlar (3.14)
- `bool` - mantiqiy qiymatlar (true/false)

### 6. Print funksiyalari

- `fmt.Print()` - oddiy chiqarish
- `fmt.Println()` - chiqarish va yangi qator
- `fmt.Printf()` - formatlangan chiqarish

### 7. String operatsiyalari (Matn bilan ishlash)

Go'da string bilan ishlash uchun asosiy operatsiyalar:

```go
// String birlashtirish
text1 := "Salom"
text2 := "Dunyo"
combined := text1 + " " + text2

// fmt.Sprintf() - formatlangan matn yaratish
name := "Matnazar"
age := 25
formatted := fmt.Sprintf("Ism: %s, Yosh: %d", name, age)
```

**Asosiy string operatsiyalar:**
- `+` - matnlarni birlashtirish
- `fmt.Sprintf()` - formatlangan matn yaratish
- String uzunligi va boshqa operatsiyalar keyingi kunda o'rganiladi

## Bugun o'rganilganlar

✅ Go dasturining asosiy struktura
✅ Package va import tushunchalari
✅ O'zgaruvchilar va ma'lumot turlari
✅ Print funksiyalari
✅ Konstanta va matematik amallar
✅ String bilan asosiy ishlash (birlashtirish, formatlash)
✅ Sodda dastur yozish

## Keyingi kun

Kun 2: Funksiyalar, parametrlar va qaytarish qiymatlari
