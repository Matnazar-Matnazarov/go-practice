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
```

## 📚 O'rganish rejasi (Learning Plan)

- **Kun 1-5**: Asosiy tushunchalar (o'zgaruvchilar, funksiyalar, ma'lumot turlari)
  - ✅ **Kun 1**: Package, import, o'zgaruvchilar, ma'lumot turlari, print funksiyalari
  - ✅ **Kun 2**: Funksiyalar, parametrlar, qaytarish qiymatlari, variadic funksiyalar
  - ✅ **Kun 3**: Shartli operatorlar (if, switch), Array va Slice
  - ✅ **Kun 4**: Map (xarita), For Loop (tsikllar)
- **Kun 6-10**: Murakkab ma'lumot tuzilmalari (struct, metodlar, interface)
- **Kun 11-15**: Funksiyalar va metodlar (pointer, interface)
- **Kun 16-20**: Concurrency (goroutine, channel)
- **Kun 21-25**: File I/O, error handling, testing
- **Kun 26-30**: Loyihalar va amaliyot

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

---

**Boshlanish sanasi**: 2026
**Status**: 🟢 Faol o'rganish jarayonida
