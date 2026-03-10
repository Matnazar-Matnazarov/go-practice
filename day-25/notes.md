# Kun 25: Structured Logging (log/slog)

## Kirish

Kun 20–24 da HTTP server, client, middleware, database va CLI ni ko‘rdik. Production’da barcha qatlamlarda **strukturli logging** kerak: faqat matn emas, balki **key-value** ma’lumotlar (request_id, user_id, duration, error) — buning uchun Go 1.21 dan beri standart kutubxonada **`log/slog`** mavjud.

---

## 1. slog nima?

**slog** — strukturli logging uchun standart package. Log har bir yozuv **level** (Debug, Info, Warn, Error) va ixtiyoriy **attribute** lar (key-value) dan iborat.

### Nima uchun strukturli?

- **O‘qilishi oson** — odam uchun text, tizim uchun JSON
- **Filter va qidiruv** — log aggregator’da (Loki, ELK) field bo‘yicha qidirish
- **Metrikalar** — xatoliklar soni, latency distribution

---

## 2. Asosiy tushunchalar

| Komponent | Vazifasi |
|-----------|----------|
| **Logger** | Log yozuvlarini yuboradi: `Info`, `Error`, `Debug`, `Warn` |
| **Handler** | Loglarni qayta ishlaydi: format (JSON/text), output (io.Writer), level filter |
| **Attr** | Bir juft key-value: `slog.String("key", "value")` |
| **Record** | Bir log yozuvi: time, level, message, attrs |

---

## 3. Oddiy ishlatish

```go
import "log/slog"

slog.Info("server started", "port", 8080)
slog.Error("connection failed", "err", err, "host", host)
slog.Debug("cache hit", "key", key)
```

**Qoida:** Default logger `os.Stderr` ga text formatda yozadi. Production’da odatda JSON va level sozlanadi.

---

## 4. Handler turlari

### JSONHandler

Loglarni JSON qatorlariga yozadi. Production va log aggregator’lar uchun ideal.

```go
h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
    Level: slog.LevelInfo,
    AddSource: true, // fayl:qator — debug uchun
})
logger := slog.New(h)
logger.Info("event", "user_id", 123, "action", "login")
// {"time":"...","level":"INFO","msg":"event","user_id":123,"action":"login"}
```

### TextHandler

O‘qilishi oson text format. Development uchun qulay.

```go
h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
logger := slog.New(h)
```

---

## 5. Level va HandlerOptions

**Levels:** `LevelDebug`, `LevelInfo`, `LevelWarn`, `LevelError`. Pastroq level’lar (masalan Debug) yuqoridagilarni ham chiqaradi.

```go
opts := &slog.HandlerOptions{
    Level: slog.LevelInfo,  // Debug chiqmasin
    AddSource: false,       // qayerdan log — dev da true
    ReplaceAttr: func(g []string, a slog.Attr) slog.Attr {
        // key’larni o‘zgartirish yoki filtrlash
        return a
    },
}
```

---

## 6. WithAttrs va WithGroup

**WithAttrs** — barcha keyingi log’larga qo‘shimcha attribute’lar qo‘shadi (masalan request_id).

```go
log := logger.With("request_id", "abc-123")
log.Info("handler started")
log.Info("handler finished")
// ikkalasida ham request_id=abc-123 bo‘ladi
```

**WithGroup** — attribute’larni guruhga joylash (JSON’da nested object).

```go
log := logger.WithGroup("http")
log.Info("request", "method", "GET", "path", "/api")
// "http": {"method":"GET","path":"/api"}
```

---

## 7. Context bilan ishlash

Request-scoped logger’ni context’da saqlash — middleware’da `WithAttrs("request_id", id)` yaratib, handler’ga context orqali uzatish. Keyin `slog.FromContext(ctx)` orqali olish (yoki o‘zingiz context’ga logger qo‘yasiz).

---

## 8. SetDefault va global logger

`slog.SetDefault(logger)` — paketlarda `slog.Info(...)` chaqirganda shu logger ishlatiladi. Dastur boshida bitta marta sozlab, hamma joyda ishlatish mumkin.

---

## 9. Best practices

1. **Production’da JSON** — log aggregator va parsing uchun
2. **Level’ni env’dan** — development’da Debug, production’da Info yoki Warn
3. **WithAttrs** — request_id, user_id kabi umumiy field’lar uchun
4. **Sensitive ma’lumot** — parol, token’larni log’ga yozmaslik; ReplaceAttr da mask qilish

---

## Xulosa

**slog** — strukturli logging uchun standart yechim; level, handler (JSON/text), WithAttrs/WithGroup va context bilan production’da to‘liq qo‘llanadi.

**Keyingi qadamlar:** Tracing (OpenTelemetry), metrikalar (Prometheus), centralized logging (Loki, ELK).
