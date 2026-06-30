<div align="center">
    
# 🚀 Go 30-Day Journey


![Go Version](https://img.shields.io/badge/Go-1.18%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Progress](https://img.shields.io/badge/Progress-43.3%25-00ADD8?style=for-the-badge)
![License](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)
![Status](https://img.shields.io/badge/Status-Active-success?style=for-the-badge)

**Go dasturlash tilini 0 dan o'rganish: 30 kunlik mashqlar va loyihalar**

*Learning Go from scratch in 30 days. Daily notes, exercises, and mini-projects.*

[![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)](https://golang.org/)
[![GitHub](https://img.shields.io/badge/GitHub-Repository-black?style=flat&logo=github)](https://github.com)

</div>

---

## 📋 Table of Contents (Mundarija)

- [📁 Struktura](#-struktura-structure)
- [🚀 Dasturni ishga tushirish](#-dasturni-ishga-tushirish-how-to-run)
- [📚 O'rganish rejasi](#-organish-rejasi-learning-plan)
- [📊 Progress](#-progress-jarayon)
- [🛠️ Texnologiyalar](#️-texnologiyalar-va-xususiyatlar-technologies--features)
- [🌿 Branch strategiyasi](#-branch-strategiyasi-branch-strategy)
- [📝 Eslatmalar](#-eslatmalar-notes)
- [🎯 Maqsad](#-maqsad-goal)

## 📁 Struktura (Structure)

<details>
<summary><b>📂 Loyiha strukturasini ko'rish (Click to expand)</b></summary>

```

```
go-practice/
├── README.md
├── day-01/          → Kun 1: Asosiy tushunchalar
│   ├── notes.md     → O'rganilgan nazariya va qoidalar
│   └── code.go      → Amaliy mashqlar va misollar
├── day-02/          → Kun 2: Funksiyalar va parametrlar
│   ├── notes.md     → O'rganilgan nazariya va qoidalar
│   ├── code.go      → Asosiy dastur (main funksiya)
│   └── functions.go → Barcha funksiyalar (alohida fayl)
├── day-03/          → Kun 3: Shartli operatorlar, Array va Slice
│   ├── notes.md     → O'rganilgan nazariya va qoidalar
│   ├── code.go      → Asosiy dastur (main funksiya)
│   ├── conditions.go → Shartli operatorlar (if, switch)
│   └── arrays_slices.go → Array va Slice operatsiyalari
├── day-04/          → Kun 4: Map (Xarita) va For Loop
│   ├── notes.md     → O'rganilgan nazariya va qoidalar
│   ├── code.go      → Asosiy dastur (main funksiya)
│   ├── maps.go      → Map operatsiyalari
│   └── loops.go     → For loop turlari
├── day-05/          → Kun 5: Struct, Funksiyalar va For Loop (Kengaytirilgan)
│   ├── notes.md     → O'rganilgan nazariya va qoidalar
│   ├── code.go      → Asosiy dastur (main funksiya)
│   ├── structs.go   → Struct operatsiyalari va metodlar
│   ├── functions_advanced.go → Kengaytirilgan funksiyalar
│   └── loops_advanced.go → Kengaytirilgan for loop
├── day-06/          → Kun 6: Queue (Navbat) va Goroutine (Parallel dasturlash)
│   ├── notes.md     → O'rganilgan nazariya va qoidalar
│   ├── code.go      → Asosiy dastur (main funksiya)
│   ├── queue.go     → Queue ma'lumot tuzilmasi
│   ├── goroutines.go → Goroutine va Channel operatsiyalari
│   └── run.sh       → Ishga tushirish skripti
├── day-07/          → Kun 7: WaitGroup va Channel (Kengaytirilgan)
│   ├── notes.md     → O'rganilgan nazariya va qoidalar
│   ├── code.go      → Asosiy dastur (main funksiya)
│   ├── waitgroup.go → WaitGroup operatsiyalari
│   ├── channels.go  → Channel (kengaytirilgan) operatsiyalari
│   └── run.sh       → Ishga tushirish skripti
├── day-08/          → Kun 8: Context va Mutex (Sinxronizatsiya)
│   ├── notes.md     → O'rganilgan nazariya va qoidalar
│   ├── code.go      → Asosiy dastur (main funksiya)
│   ├── context.go   → Context operatsiyalari
│   ├── mutex.go     → Mutex va RWMutex operatsiyalari
│   └── run.sh       → Ishga tushirish skripti
├── day-09/          → Kun 9: Atomic Operations va Error Handling
│   ├── notes.md     → O'rganilgan nazariya va qoidalar
│   ├── code.go      → Asosiy dastur (main funksiya)
│   ├── atomic.go    → Atomic operatsiyalar
│   ├── errors.go    → Error handling
│   └── run.sh       → Ishga tushirish skripti
├── day-10/          → Kun 10: Pointers va Interfaces
│   ├── notes.md     → O'rganilgan nazariya va qoidalar
│   ├── code.go      → Asosiy dastur (main funksiya)
│   ├── pointers.go  → Pointer operatsiyalari
│   ├── interfaces.go → Interface operatsiyalari
│   └── run.sh       → Ishga tushirish skripti
├── day-11/          → Kun 11: Funksiyalar va Metodlar (Pointer, Interface)
│   ├── notes.md     → O'rganilgan nazariya va qoidalar
│   ├── code.go      → Asosiy dastur (main funksiya)
│   ├── functions_advanced.go → Kengaytirilgan funksiyalar
│   ├── interfaces_advanced.go → Kengaytirilgan interface operatsiyalari
│   ├── methods.go   → Metodlar va method sets
│   └── run.sh       → Ishga tushirish skripti
├── day-12/          → Kun 12: Generics (Umumiy turlar)
│   ├── notes.md     → O'rganilgan nazariya va qoidalar
│   ├── code.go      → Asosiy dastur (main funksiya)
│   ├── generics.go  → Generics misollari va funksiyalar
│   └── run.sh       → Ishga tushirish skripti
├── day-13/          → Kun 13: Reflection (Refleksiya)
│   ├── notes.md     → O'rganilgan nazariya va qoidalar
│   ├── code.go      → Asosiy dastur (main funksiya)
│   ├── reflection.go → Reflection misollari va funksiyalar
│   └── run.sh       → Ishga tushirish skripti
├── ...
├── day-30/          → Kun 30: Yakuniy loyiha
│   ├── notes.md
│   └── code.go
└── exercises/       → Kichik loyihalar
    └── small-projects/
```

</details>

## 🚀 Dasturni ishga tushirish (How to run)

```bash
# Kun 1 kodini ishga tushirish
cd day-01
go run code.go

# Yoki kompilyatsiya qilib ishga tushirish
go build code.go
./code

# Kun 2 kodini ishga tushirish (bir nechta fayl)
cd day-02
go run *.go

# Yoki to'g'ridan-to'g'ri
go run day-02/*.go

# Kompilyatsiya qilib ishga tushirish
cd day-02
go build
./day-02

# Kun 3 kodini ishga tushirish
cd day-03
go run *.go

# Yoki to'g'ridan-to'g'ri
go run day-03/*.go

# Kun 4 kodini ishga tushirish
cd day-04
go run *.go

# Yoki to'g'ridan-to'g'ri
go run day-04/*.go

# Kun 5 kodini ishga tushirish
cd day-05
go run *.go

# Yoki to'g'ridan-to'g'ri
go run day-05/*.go

# Kun 6 kodini ishga tushirish
cd day-06
go run *.go

# Yoki to'g'ridan-to'g'ri
go run day-06/*.go

# Yoki run.sh orqali
cd day-06
./run.sh

# Kun 7 kodini ishga tushirish
cd day-07
go run *.go

# Yoki to'g'ridan-to'g'ri
go run day-07/*.go

# Yoki run.sh orqali
cd day-07
./run.sh

# Kun 8 kodini ishga tushirish
cd day-08
go run *.go

# Yoki to'g'ridan-to'g'ri
go run day-08/*.go

# Yoki run.sh orqali
cd day-08
./run.sh

# Kun 9 kodini ishga tushirish
cd day-09
go run *.go

# Yoki to'g'ridan-to'g'ri
go run day-09/*.go

# Yoki run.sh orqali
cd day-09
./run.sh

# Kun 10 kodini ishga tushirish
cd day-10
go run *.go

# Yoki to'g'ridan-to'g'ri
go run day-10/*.go

# Yoki run.sh orqali
cd day-10
./run.sh

# Kun 11 kodini ishga tushirish
cd day-11
go run *.go

# Yoki to'g'ridan-to'g'ri
go run day-11/*.go

# Yoki run.sh orqali
cd day-11
./run.sh

# Kun 12 kodini ishga tushirish
cd day-12
go run *.go

# Yoki to'g'ridan-to'g'ri
go run day-12/*.go

# Yoki run.sh orqali
cd day-12
./run.sh

# Kun 13 kodini ishga tushirish
cd day-13
go run *.go

# Yoki to'g'ridan-to'g'ri
go run day-13/*.go

# Yoki run.sh orqali
cd day-13
./run.sh
```

## 📚 O'rganish rejasi (Learning Plan)

### 📅 30-Kunlik Reja

Bu loyiha 30 kunlik strukturalashtirilgan o'rganish rejasini taklif qiladi. Har bir kun yangi mavzular va amaliy mashqlarni o'z ichiga oladi.

- **Kun 1-5**: Asosiy tushunchalar (o'zgaruvchilar, funksiyalar, ma'lumot turlari)
  - ✅ **Kun 1**: Package, import, o'zgaruvchilar, ma'lumot turlari, print funksiyalari
  - ✅ **Kun 2**: Funksiyalar, parametrlar, qaytarish qiymatlari, variadic funksiyalar
  - ✅ **Kun 3**: Shartli operatorlar (if, switch), Array va Slice
  - ✅ **Kun 4**: Map (xarita), For Loop (tsikllar)
  - ✅ **Kun 5**: Struct (struktura), Funksiyalar (kengaytirilgan), For Loop (kengaytirilgan)
- **Kun 6-10**: Murakkab ma'lumot tuzilmalari va parallel dasturlash
  - ✅ **Kun 6**: Queue (navbat), Goroutine va Channel (parallel dasturlash boshlang'ich)
  - ✅ **Kun 7**: WaitGroup va Channel (kengaytirilgan)
  - ✅ **Kun 8**: Context va Mutex (sinxronizatsiya)
  - ✅ **Kun 9**: Atomic Operations va Error Handling
  - ✅ **Kun 10**: Pointers va Interfaces
- **Kun 11-15**: Funksiyalar va metodlar (pointer, interface, generics, reflection)
  - ✅ **Kun 11**: Funksiyalar va Metodlar (Pointer, Interface) - Method sets, method expressions, method values, method chaining, higher-order functions, design patterns
  - ✅ **Kun 12**: Generics (Umumiy turlar) - Type parameters, type constraints, generic functions, generic data structures, interface constraints
  - ✅ **Kun 13**: Reflection (Refleksiya) - Type inspection, value manipulation, struct fields, method invocation, slice/map reflection
  - ✅ **Kun 14**: Packages va Modules (Paketlar va Modullar)
  - ✅ **Kun 15**: Packages & Modules (Advanced)
- **Kun 16-20**: Concurrency (goroutine, channel) - Advanced concurrency patterns
  - ✅ **Kun 16**: Advanced Concurrency Patterns (bounded concurrency, cancellation, leak prevention, graceful shutdown)
- **Kun 21-25**: File I/O, error handling, testing - File operations, JSON/XML, testing frameworks
- **Kun 26-30**: Loyihalar va amaliyot - Real-world projects and applications

## 🌿 Branch strategiyasi (Branch Strategy)

| Branch | Maqsad | Status |
|--------|--------|--------|
| `main` | ✅ Stable progress / final working code | Production-ready |
| `dev` | 🔄 Daily exercises / experiments | Development |
| `day-XX` | 📝 Har bir kun uchun alohida branch | Feature branches |

### 🔀 Branch ishlatish

```bash
# Yangi kun uchun branch yaratish
git checkout -b day-14

# O'zgarishlarni commit qilish
git add .
git commit -m "feat: Add Day 14 - [Mavzu]"

# Remote ga push qilish
git push -u origin day-14
```

## 📝 Eslatmalar (Notes)

### 📚 Fayl tuzilishi

Har bir kun papkasida quyidagi fayllar mavjud:

- 📄 **`notes.md`** - Nazariya, qoidalar va tushuntirishlar
- 💻 **`code.go`** - Asosiy dastur (main funksiya)
- 🔧 **`*.go`** - Qo'shimcha Go fayllar (mavzuga qarab)
- 🚀 **`run.sh`** - Ishga tushirish skripti (ba'zi kunlar uchun)

### 🌍 Til

- Barcha kodlar **o'zbek tilida** izohlar bilan yozilgan
- Hujjatlar **o'zbek va ingliz** tillarida
- Kod misollari **o'zbek tilida** izohlar bilan

### 💡 Maslahatlar

- ✅ Har bir kunda `notes.md` faylini to'liq o'qing
- ✅ Kodlarni o'zgartirib, tushunishni mustahkamlang
- ✅ Har bir mavzuni mustahkamlang keyingi kunlarga o'tishdan oldin
- ✅ Xatolarni o'qib, tushunishga harakat qiling

### 📚 Fayl tuzilishi

Har bir kun papkasida quyidagi fayllar mavjud:

- 📄 **`notes.md`** - Nazariya, qoidalar va tushuntirishlar
- 💻 **`code.go`** - Asosiy dastur (main funksiya)
- 🔧 **`*.go`** - Qo'shimcha Go fayllar (mavzuga qarab)
- 🚀 **`run.sh`** - Ishga tushirish skripti (ba'zi kunlar uchun)

### 🌍 Til

- Barcha kodlar **o'zbek tilida** izohlar bilan yozilgan
- Hujjatlar **o'zbek va ingliz** tillarida
- Kod misollari **o'zbek tilida** izohlar bilan

### 💡 Maslahatlar

- ✅ Har bir kunda `notes.md` faylini to'liq o'qing
- ✅ Kodlarni o'zgartirib, tushunishni mustahkamlang
- ✅ Har bir mavzuni mustahkamlang keyingi kunlarga o'tishdan oldin
- ✅ Xatolarni o'qib, tushunishga harakat qiling

## 🎯 Maqsad (Goal)

<div align="center">

### 🎓 30 kun ichida Go dasturlash tilini professional darajada o'zlashtirish

**Asosiy maqsadlar:**
- ✅ Go dasturlash tilining asosiy va murakkab tushunchalarini o'rganish
- ✅ Amaliy loyihalar va mashqlar orqali bilimlarni mustahkamlash
- ✅ Professional darajada Go kod yozish ko'nikmalarini rivojlantirish
- ✅ Real-world loyihalar yaratish qobiliyatiga ega bo'lish

</div>

## 📊 Progress (Jarayon)

<div align="center">

**Umumiy progress**: **28/30 kun** (93.3% ✅)

```
████████████████████████████████████░ 93.3%
```

</div>

### 📈 Progress Breakdown

| Bo'lim | Progress | Status | Kunlar |
|-------|----------|--------|--------|
| **Kun 1-5** | ✅ 100% | Tugallandi | 5/5 |
| **Kun 6-10** | ✅ 100% | Tugallandi | 5/5 |
| **Kun 11-15** | ✅ 100% | Tugallandi | 5/5 |
| **Kun 16-20** | ✅ 100% | Tugallandi | 5/5 |
| **Kun 21-25** | ✅ 100% | Tugallandi | 5/5 |
| **Kun 26-30** | 🔄 60% | Jarayonda | 3/5 |

### 📝 Batafsil Progress

- ✅ **Kun 1-5**: Asosiy tushunchalar (100% - 5/5 kun)
  - ✅ Kun 1: Package, import, o'zgaruvchilar, ma'lumot turlari
  - ✅ Kun 2: Funksiyalar, parametrlar, variadic funksiyalar
  - ✅ Kun 3: Shartli operatorlar, Array va Slice
  - ✅ Kun 4: Map, For Loop
  - ✅ Kun 5: Struct, Kengaytirilgan funksiyalar

- ✅ **Kun 6-10**: Murakkab ma'lumot tuzilmalari va parallel dasturlash (100% - 5/5 kun)
  - ✅ Kun 6: Queue, Goroutine va Channel
  - ✅ Kun 7: WaitGroup va Channel (kengaytirilgan)
  - ✅ Kun 8: Context va Mutex
  - ✅ Kun 9: Atomic Operations va Error Handling
  - ✅ Kun 10: Pointers va Interfaces

- ✅ **Kun 11-15**: Funksiyalar va metodlar (100% - 5/5 kun)
  - ✅ **Kun 11**: Funksiyalar va Metodlar (Pointer, Interface)
  - ✅ **Kun 12**: Generics (Umumiy turlar)
  - ✅ **Kun 13**: Reflection (Refleksiya)
  - ✅ **Kun 14**: Packages va Modules (Paketlar va Modullar)
  - ✅ **Kun 15**: Packages & Modules (Advanced)

- ✅ **Kun 16-20**: Concurrency (100% - 5/5 kun)
  - ✅ **Kun 16**: Advanced Concurrency Patterns
  - ✅ **Kun 17**: Patterns (Worker Pool, Pub/Sub, Fan-Out/Fan-In)
  - ✅ **Kun 18**: Testing (Unit tests, Table-driven, Benchmarks)
  - ✅ **Kun 19**: JSON & I/O (Encoding, File operations)
  - ✅ **Kun 20**: HTTP Server (REST API, Handler, Middleware)

- ✅ **Kun 21-25**: HTTP Client, Middleware, Database (100% - 5/5 kun)
  - ✅ **Kun 21**: HTTP Client & API Testing (httptest)
  - ✅ **Kun 22**: HTTP Middleware & Error Handling
  - ✅ **Kun 23**: Database (SQL, connection pool)
  - ✅ **Kun 24**: Flags & CLI (Command-line arguments)
  - ✅ **Kun 25**: Logger (Structured logging)

- 🔄 **Kun 26-30**: Loyihalar va amaliyot (60% - 3/5 kun)
  - ✅ **Kun 26**: Config Management (Environment variables)
  - ✅ **Kun 27**: Graceful Shutdown & Signal Handling
  - ✅ **Kun 28**: WebSocket & Real-Time Communication
  - ⏳ **Kun 29**: TBD
  - ⏳ **Kun 30**: Yakuniy loyiha

## 🛠️ Texnologiyalar va Xususiyatlar (Technologies & Features)

### 💻 O'rganilgan mavzular (Topics Covered)

<table>
<tr>
<td width="50%">

#### ✅ Asosiy tushunchalar
- 📦 Package va Import
- 🔢 Variables va Data Types
- 🔄 Functions va Methods
- 🎯 Control Flow (if, switch, loops)

#### ✅ Ma'lumot tuzilmalari
- 📊 Arrays va Slices
- 🗺️ Maps (Xarita)
- 🏗️ Structs va Methods
- 🔗 Pointers

</td>
<td width="50%">

#### ⚡ Concurrency
- 🚀 Goroutines
- 📡 Channels
- ⏳ WaitGroup
- 🎛️ Context
- 🔒 Mutex va RWMutex
- ⚛️ Atomic Operations

#### 🎓 Advanced Features
- 🎨 Interfaces
- 🔧 Generics (Go 1.18+)
- 🔍 Reflection
- 🎭 Design Patterns

#### 🌐 Web & Network
- 🖥️ HTTP Server (net/http)
- 🌍 HTTP Client
- 🔌 WebSocket (Real-time)
- 🧪 API Testing (httptest)

</td>
</tr>
</table>

### 📦 Design Patterns

- 🏗️ **Builder Pattern** - Murakkab obyektlarni qurish
- 🎯 **Strategy Pattern** - Algoritmlarni almashtirish
- 👁️ **Observer Pattern** - O'zgarishlarni kuzatish

### 🔧 Go Versiyasi (Go Version)

| Versiya | Status | Xususiyatlar |
|---------|--------|--------------|
| **Go 1.18+** | ✅ Minimal | Generics qo'llab-quvvatlash |
| **Go 1.21+** | ⭐ Tavsiya | Eng so'nggi xususiyatlar |

### 📚 Qo'shimcha Resurslar

- 📖 [Go Documentation](https://golang.org/doc/)
- 🎓 [Go Tour](https://go.dev/tour/)
- 📝 [Effective Go](https://go.dev/doc/effective_go)
- 🐹 [Go by Example](https://gobyexample.com/)

---

<div align="center">

### 📅 Project Information

| Parametr | Qiymat |
|----------|--------|
| **🌍 Til (Language)** | O'zbek / English |
| **📅 Boshlanish sanasi** | 2026 |
| **📊 Status** | 🟢 Faol o'rganish jarayonida |
| **🔄 Oxirgi yangilanish** | Day 28 - WebSocket & Real-Time Communication |
| **📝 License** | MIT |

---

### ⭐ Star ⭐

Agar bu loyiha sizga foydali bo'lsa, yulduzcha qo'yib qo'llab-quvvatlang!

*If you find this project helpful, please give it a star!*


---

**Made with ❤️ and Go**

[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![GitHub](https://img.shields.io/badge/GitHub-Repository-black?style=for-the-badge&logo=github)](https://github.com)

</div>
