# Kun 26: Configuration va Environment (12-factor)

## Kirish

Kun 24–25 da CLI va structured logging ni ko‘rdik. Production dasturlar sozlamalarni **environment** va **config fayl** dan olishi kerak — kod ichida port, API key yoki DB URL bo‘lmasligi kerak. **12-factor** metodologiyasi va Go’da buni **`os.Getenv`**, default qiymatlar va (ixtiyoriy) config fayl bilan professional qilish.

---

## 1. Nima uchun environment?

- **Xavfsizlik** — parol, API key kodda bo‘lmasin; deploy vaqtida env orqali beriladi
- **Muhitlar** — dev, staging, prod har xil `DATABASE_URL`, `LOG_LEVEL`
- **Container/Docker** — env orqali inject qilish standart

---

## 2. os.Getenv va default

```go
port := os.Getenv("PORT")
if port == "" {
    port = "8080"
}
```

**Qoida:** `os.Getenv` bo‘sh string qaytaradi — key mavjud emas bo‘lsa. Default qiymatni o‘zingiz belgilang.

### Helper: getEnv(key, default)

```go
func getEnv(key, defaultValue string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return defaultValue
}
port := getEnv("PORT", "8080")
```

---

## 3. Config struct — bitta joyda sozlamalar

Barcha sozlamalarni bir **struct** da to‘plash — type-safe va testda mock qilish oson.

```go
type Config struct {
    Env      string
    Port     int
    LogLevel string
    DBURL    string
}
```

**Yuklash:** env’dan o‘qib, struct’ni to‘ldirish; kerak bo‘lsa config fayl (JSON/YAML) qo‘shiladi.

---

## 4. Validation — majburiy sozlamalar

Ba’zi sozlamalar bo‘lmasa dastur ishlamasligi kerak (masalan `DATABASE_URL`).

```go
if cfg.DBURL == "" {
    return nil, fmt.Errorf("DATABASE_URL is required")
}
if cfg.Port < 1 || cfg.Port > 65535 {
    return nil, fmt.Errorf("PORT must be 1-65535")
}
```

**Best practice:** config yuklangach darhol validate qiling; xato bo‘lsa tez fail qiling.

---

## 5. Config fayl (ixtiyoriy)

Env ustun, lekin murakkab sozlamalar uchun fayl ham qo‘llanadi:

- **JSON** — `encoding/json`, fayl yo‘li env’dan: `CONFIG_PATH`
- **YAML** — third-party library
- **Priority:** env > config fayl > default (env har doim override qiladi)

---

## 6. 12-factor qisqacha

1. **Config env’da** — kodda hardcode yo‘q
2. **Muhitlar** — dev/staging/prod env’lar bilan ajratiladi
3. **Credentials** — env’da, faylga commit qilinmasligi kerak

---

## 7. Best practices

- **Default’lar** — development’da ishlash uchun mantiqiy default
- **Validate early** — dastur boshida config tekshirish
- **Sensitive** — parol/token’larni log’ga chiqarmang
- **Document** — qaysi env o‘zgaruvchilar kerak, README yoki env.example da

---

## Xulosa

**Configuration** — env + (ixtiyoriy) config fayl orqali sozlamalarni boshqarish; struct, default va validation bilan production-ready.

**Keyingi qadamlar:** Viper/cobra, secret manager (Vault), feature flags.
