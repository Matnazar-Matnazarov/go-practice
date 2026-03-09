# Kun 24: CLI va Command-Line (flag, os.Args)

## Kirish

Kun 20–23 da HTTP server, client, middleware va database ni ko‘rdik. Ko‘p dasturlar **command-line** orqali ishga tushadi: parametrlar, flag’lar va subcommand’lar. Go’da buni **`flag`** va **`os.Args`** bilan professional darajada qilish mumkin.

---

## 1. CLI nima?

**CLI** (Command-Line Interface) — foydalanuvchi dasturni terminalda ishga tushirib, argumentlar va flag’lar orqali boshqaradi.

**Misol:** `go build -o app .` — bu yerda `-o app` flag, `.` argument.

### Nima uchun kerak?

1. **Server/daemon** — port, config fayl, log level
2. **Utility** — input/output fayl, format, dry-run
3. **Script’lar** — avtomatlashtirish, CI/CD

---

## 2. os.Args — oddiy argumentlar

**os.Args** — dastur ishga tushganda berilgan barcha argumentlar (string slice).

- `os.Args[0]` — dastur nomi (yo‘l bilan)
- `os.Args[1]`, `os.Args[2]`, ... — foydalanuvchi argumentlari

```go
// Misol: go run . a b c
fmt.Println(os.Args)   // [path/to/binary, a, b, c]
fmt.Println(os.Args[1]) // a
```

**Qoida:** `len(os.Args)` tekshiring — argument berilmaganda xato chiqmasin.

---

## 3. flag package — standart flag’lar

**flag** — `-flag value` yoki `-flag=value` ko‘rinishidagi flag’larni parse qiladi.

### Flag turlari

| Funksiya | Qaytaradi | Misol |
|----------|-----------|--------|
| `flag.String(name, default, usage)` | *string | `-name=Ali` |
| `flag.Int(name, default, usage)` | *int | `-port=8080` |
| `flag.Bool(name, default, usage)` | *bool | `-verbose` |
| `flag.Duration(name, default, usage)` | *time.Duration | `-timeout=5s` |

### Ishlatish tartibi

1. Flag’larni e’lon qilish (masalan, `port := flag.Int("port", 8080, "server port")`)
2. `flag.Parse()` — `os.Args` ni parse qiladi
3. Pointer orqali qiymat olish (`*port`)

```go
var port = flag.Int("port", 8080, "HTTP server port")
var verbose = flag.Bool("v", false, "verbose output")
func init() { flag.Parse() }
func main() {
    fmt.Println("Port:", *port)
    fmt.Println("Verbose:", *verbose)
}
```

**Qoida:** `flag.Parse()` odatda `main()` boshida yoki `init()` da chaqiladi.

---

## 4. Flag’ni o‘zgaruvchiga bog‘lash (Var)

`flag.StringVar`, `flag.IntVar` — mavjud o‘zgaruvchiga bog‘laydi.

```go
var configPath string
flag.StringVar(&configPath, "config", "config.json", "config file path")
flag.Parse()
// configPath endi -config=... dan to‘ldiriladi
```

**Afzallik:** struct yoki global o‘zgaruvchilarni to‘ldirishda qulay.

---

## 5. Qolgan argumentlar (Args)

`flag.Parse()` dan keyin:

- **flag.Args()** — flag’lar bo‘lmagan barcha argumentlar (slice)
- **flag.NArg()** — ularning soni

```go
flag.Parse()
for _, arg := range flag.Args() {
    fmt.Println("Arg:", arg)
}
```

**Misol:** `app -port=80 file1.txt file2.txt` — `flag.Args()` = `["file1.txt", "file2.txt"]`.

---

## 6. Help va Usage

- **flag.Usage** — `-h` / `--help` da chiqadigan matn (o‘zgartirish mumkin)
- **flag.PrintDefaults()** — barcha flag’lar va ularning default qiymatlari

```go
flag.Usage = func() {
    fmt.Fprintf(os.Stderr, "Ishlatish: %s [options] <file>\n", os.Args[0])
    flag.PrintDefaults()
}
```

---

## 7. Best practices

1. **Default qiymatlar** — har bir flag uchun mantiqiy default belgilang
2. **Usage matni** — qisqa va tushunarli
3. **flag.Parse()** — bitta marta, dastur boshida
4. **os.Args tekshirish** — ixtiyoriy argumentlar uchun `len(flag.Args())` yoki `flag.NArg()`

---

## Xulosa

**os.Args** — oddiy argumentlar ro‘yxati; tezkor skriptlar uchun yetarli.

**flag** — production’da flag’lar uchun standart yechim: tipi aniq, help avtomatik, default qulay.

**Keyingi qadamlar:** Subcommand’lar (cobra, urfave/cli), environment (os.Getenv), config fayllar.
