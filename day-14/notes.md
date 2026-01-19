# Kun 14: Packages va Modules (Paketlar va Modullar)

## Kirish

Go dasturlash tilida **Packages** va **Modules** - bu kodni tashkil qilish va qayta ishlatiladigan komponentlar yaratishning asosiy mexanizmi. Packages orqali kodni mantiqiy guruhlarga ajratish, Modules orqali esa loyihalarni boshqarish mumkin.

---

## 1. Packages Asoslari

### Package nima?

**Package** - bu bir xil maqsadga ega bo'lgan Go fayllarining to'plami.

**Package** - bu kodni mantiqiy guruhlarga ajratish mexanizmi.

### Package nima uchun kerak?

1. **Code Organization**
   - Kodni mantiqiy guruhlarga ajratish
   - Kodni tushunish va saqlash oson

2. **Code Reusability**
   - Kodni qayta ishlatish
   - Bir nechta loyihalarda ishlatish

3. **Encapsulation**
   - Kodni yashirish (exported/unexported)
   - API ni boshqarish

### Package vazifasi

Package quyidagi vazifalarni bajaradi:

1. **Code Organization** - Kodni tashkil qilish
2. **Namespace** - Nomlar maydonini ajratish
3. **Encapsulation** - Kodni yashirish
4. **Reusability** - Kodni qayta ishlatish

### Package afzalliklari

✅ **Code Organization** - Kodni tashkil qilish
✅ **Reusability** - Kodni qayta ishlatish
✅ **Encapsulation** - Kodni yashirish
✅ **Testing** - Test qilish oson

### Package kamchiliklari

❌ **Murakkab** - Katta loyihalarda tushunish qiyin
❌ **Dependencies** - Qaramliklar boshqarish kerak

---

## 2. Package E'lon Qilish

### Package Declaration

**Package declaration** - har bir Go faylining birinchi qatori.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

**Qoida:** Har bir fayl package bilan boshlanadi.

### Package Name Qoidalari

**Package name** - kichik harflar bilan yoziladi.

```go
package math      // ✅ To'g'ri
package Math      // ❌ Xato
package MATH      // ❌ Xato
package my_package // ✅ To'g'ri (lekin camelCase yaxshiroq)
package mypackage  // ✅ To'g'ri
```

**Qoida:** Package nomi kichik harflar bilan, camelCase yoki snake_case.

### main Package

**main package** - executable dastur uchun.

```go
package main

func main() {
    // Dastur kodlari
}
```

**Qoida:** `main` package va `main()` funksiyasi executable dastur uchun.

---

## 3. Import

### Import Statement

**Import** - boshqa package lardan foydalanish.

```go
package main

import "fmt"
import "math"

func main() {
    fmt.Println(math.Pi)
}
```

**Qoida:** Import - package dan foydalanish uchun.

### Multiple Imports

**Multiple imports** - bir nechta package larni import qilish.

```go
import (
    "fmt"
    "math"
    "strings"
)
```

**Qoida:** Multiple imports - qavs ichida.

### Import Aliases

**Import alias** - package ga boshqa nom berish.

```go
import (
    f "fmt"
    m "math"
)

func main() {
    f.Println(m.Pi)
}
```

**Qoida:** Alias - package nomini o'zgartirish.

### Dot Import

**Dot import** - package nomini yozmasdan ishlatish.

```go
import . "fmt"

func main() {
    Println("Hello") // fmt yozmasdan
}
```

**Qoida:** Dot import - package nomini tushirib qoldirish.

### Blank Import

**Blank import** - package ni import qilish, lekin ishlatmaslik.

```go
import _ "database/sql/driver"

// Package init() funksiyasi ishga tushadi
```

**Qoida:** Blank import - init() funksiyasini chaqirish uchun.

---

## 4. Exported va Unexported

### Exported Identifiers

**Exported** - katta harf bilan boshlanadi.

```go
package math

// Exported - boshqa package lardan ishlatish mumkin
func Add(a, b int) int {
    return a + b
}

var Pi = 3.14159
```

**Qoida:** Katta harf bilan boshlangan - exported.

### Unexported Identifiers

**Unexported** - kichik harf bilan boshlanadi.

```go
package math

// Unexported - faqat shu package ichida ishlatish mumkin
func add(a, b int) int {
    return a + b
}

var pi = 3.14159
```

**Qoida:** Kichik harf bilan boshlangan - unexported.

### Encapsulation

**Encapsulation** - kodni yashirish.

```go
package counter

// Exported
type Counter struct {
    value int // Unexported field
}

// Exported method
func (c *Counter) Increment() {
    c.value++
}

// Unexported method
func (c *Counter) increment() {
    c.value++
}
```

**Qoida:** Encapsulation - API ni boshqarish.

---

## 5. init() Funksiyasi

### init() Funksiyasi

**init()** - package yuklanganda avtomatik chaqiriladi.

```go
package main

import "fmt"

func init() {
    fmt.Println("init() funksiyasi chaqirildi")
}

func main() {
    fmt.Println("main() funksiyasi chaqirildi")
}
```

**Qoida:** init() - package yuklanganda avtomatik chaqiriladi.

### Multiple init() Funksiyalari

**Multiple init()** - bir nechta init() funksiyalari.

```go
package main

func init() {
    fmt.Println("init() 1")
}

func init() {
    fmt.Println("init() 2")
}

func main() {
    // init() funksiyalari ketma-ket chaqiriladi
}
```

**Qoida:** Multiple init() - ketma-ket chaqiriladi.

### init() Order

**init() order** - chaqirilish tartibi.

1. Import edilgan package larning init()
2. Package o'zining init()
3. main() funksiyasi

**Qoida:** init() - import dan keyin, main() dan oldin.

---

## 6. Modules

### Module nima?

**Module** - Go loyihasi va uning qaramliklari.

**Module** - go.mod fayli bilan belgilanadi.

### Module E'lon Qilish

**go mod init** - yangi module yaratish.

```bash
go mod init example.com/myproject
```

**Qoida:** go mod init - yangi module yaratish.

### go.mod Fayli

**go.mod** - module ma'lumotlari.

```go
module example.com/myproject

go 1.18

require (
    github.com/example/package v1.2.3
)
```

**Qoida:** go.mod - module va qaramliklar.

### Module Path

**Module path** - module manzili.

```go
module github.com/user/project
```

**Qoida:** Module path - module identifikatori.

---

## 7. Dependencies

### Dependency Management

**Dependencies** - boshqa package lardan foydalanish.

```bash
go get github.com/example/package
```

**Qoida:** go get - package ni o'rnatish.

### go.sum Fayli

**go.sum** - qaramliklar checksum.

**Qoida:** go.sum - xavfsizlik uchun.

### Version Management

**Version** - package versiyasi.

```bash
go get github.com/example/package@v1.2.3
go get github.com/example/package@latest
```

**Qoida:** Version - package versiyasini belgilash.

---

## 8. Internal Packages

### Internal Package

**Internal** - faqat ota package dan ishlatish mumkin.

```go
// internal/helper/helper.go
package helper

func Helper() {
    // Faqat ota package dan ishlatish mumkin
}
```

**Qoida:** internal/ - faqat ota package dan ishlatish.

---

## 9. Vendor Directory

### Vendor

**vendor/** - qaramliklarni lokal saqlash.

```bash
go mod vendor
```

**Qoida:** vendor/ - qaramliklarni lokal saqlash.

---

## 10. Muhim Qoidalar

### Package Naming

**Package naming** - package nomini tanlash.

- Kichik harflar
- Qisqa va tushunarli
- Underscore yoki camelCase

### Import Organization

**Import organization** - import larni tartibga solish.

1. Standard library
2. Third-party packages
3. Local packages

### Best Practices

1. **Package nomi** - kichik harflar, qisqa
2. **Exported** - faqat kerakli narsalar
3. **init()** - kam ishlatish
4. **Dependencies** - minimal qaramliklar

---

## Xulosa

### Packages

**Nima:** Bir xil maqsadga ega bo'lgan Go fayllari to'plami

**Vazifasi:**
- Code organization
- Namespace
- Encapsulation
- Reusability

**Afzalliklari:**
- Code organization
- Reusability
- Encapsulation
- Testing

**Kamchiliklari:**
- Murakkab
- Dependencies boshqarish

### Modules

**Nima:** Go loyihasi va uning qaramliklari

**Vazifasi:**
- Dependency management
- Version control
- Project organization

**Afzalliklari:**
- Dependency management
- Version control
- Reproducible builds

### Keyingi qadamlar

- Advanced package patterns
- Module design best practices
- Dependency management strategies
- Real-world examples
