# Go 30-Day Journey

Go dasturlash tilini 0 dan o'rganish: 30 kunlik mashqlar va loyihalar.

Learning Go from scratch in 30 days. Daily notes, exercises, and mini-projects.

## 📁 Struktura (Structure)

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
  - 🔄 **Kun 14**: TBD
  - 🔄 **Kun 15**: TBD
- **Kun 16-20**: Concurrency (goroutine, channel) - Advanced concurrency patterns
- **Kun 21-25**: File I/O, error handling, testing - File operations, JSON/XML, testing frameworks
- **Kun 26-30**: Loyihalar va amaliyot - Real-world projects and applications

## 🌿 Branch strategiyasi (Branch Strategy)

- `main` → Stable progress / final working code
- `dev` → Daily exercises / experiments
- `day-01`, `day-02`, ... `day-30` → Har bir kun uchun alohida branch

## 📝 Eslatmalar (Notes)

- Har bir kunda `notes.md` faylida nazariya va qoidalar
- `code.go` faylida amaliy misollar va izohlar
- Barcha kodlar o'zbek tilida izohlar bilan yozilgan

## 🎯 Maqsad (Goal)

30 kun ichida Go dasturlash tilini o'rganib, professional darajada kod yozishni o'zlashtirish.

## 📊 Progress (Jarayon)

**Umumiy progress**: 13/30 kun (43.3% ✅)

- ✅ **Kun 1-5**: Asosiy tushunchalar (100% - 5/5 kun)
- ✅ **Kun 6-10**: Murakkab ma'lumot tuzilmalari va parallel dasturlash (100% - 5/5 kun)
- 🔄 **Kun 11-15**: Funksiyalar va metodlar (60% - 3/5 kun)
  - ✅ Kun 11: Funksiyalar va Metodlar (Pointer, Interface)
  - ✅ Kun 12: Generics (Umumiy turlar)
  - ✅ Kun 13: Reflection (Refleksiya)
  - 🔄 Kun 14: TBD
  - 🔄 Kun 15: TBD
- ⏳ **Kun 16-20**: Concurrency (0% - 0/5 kun)
- ⏳ **Kun 21-25**: File I/O, error handling, testing (0% - 0/5 kun)
- ⏳ **Kun 26-30**: Loyihalar va amaliyot (0% - 0/5 kun)

## 🛠️ Texnologiyalar va Xususiyatlar (Technologies & Features)

### O'rganilgan mavzular (Topics Covered)

- ✅ **Asosiy tushunchalar**: Variables, functions, data types, control flow
- ✅ **Ma'lumot tuzilmalari**: Arrays, slices, maps, structs
- ✅ **Concurrency**: Goroutines, channels, WaitGroup, Context, Mutex, Atomic operations
- ✅ **Advanced features**: Pointers, interfaces, generics, reflection
- ✅ **Error handling**: Error types, error wrapping, error handling patterns
- ✅ **Design patterns**: Builder, Strategy, Observer patterns

### Go Versiyasi (Go Version)

- **Minimal versiya**: Go 1.18+ (Generics qo'llab-quvvatlash uchun)
- **Tavsiya etilgan**: Go 1.21+ (eng so'nggi xususiyatlar uchun)

---

**Boshlanish sanasi**: 2026  
**Status**: 🟢 Faol o'rganish jarayonida  
**Oxirgi yangilanish**: Day 13 - Reflection (Refleksiya)
