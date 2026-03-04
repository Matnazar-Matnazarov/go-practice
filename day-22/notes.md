# Kun 22: HTTP Middleware va Error Handling (net/http)

## Kirish

Kun 20–21 da HTTP server/client asoslarini ko‘rdik. Endi esa production’da juda ko‘p ishlatiladigan 2 ta mavzu:

- **Middleware** — handler’ni o‘rab, umumiy funksionallarni (logging, auth, metrics, timeout) qo‘shish.
- **Error handling** — noto‘g‘ri request, method, va **panic** holatlarida to‘g‘ri javob qaytarish.

---

## 1. Middleware nima?

Middleware odatda shunday ko‘rinishda bo‘ladi:

```go
func mw(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // before
        next.ServeHTTP(w, r)
        // after
    })
}
```

**G‘oya:** `next` — asl handler (yoki mux). Middleware esa so‘rov kelganda `next`ni chaqirishdan oldin/ketin ish bajaradi.

---

## 2. Logging middleware

Minimal logging:

- method/path
- request davomiyligi

Bu real loyihada tracing, request-id, structured logging (zap/logrus) bilan boyitiladi.

---

## 3. Recovery middleware (panic’ni ushlash)

Handler ichida `panic` bo‘lsa, server yiqilib ketmasligi kerak.

Recovery middleware:

- `defer` + `recover()` bilan panic’ni ushlaydi
- log qiladi
- klientga `500 Internal Server Error` qaytaradi

**Eslatma:** panic — normal control flow emas. Uni faqat kutilmagan holatlar uchun ishlating.

---

## 4. HTTP error handling (method / status)

Eng ko‘p uchraydigan pattern:

```go
if r.Method != http.MethodGet {
    http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    return
}
```

`http.Error` tez va toza yo‘l: status + oddiy text body.

---

## 5. Demo’ni ishga tushirish

Terminal:

```bash
cd day-22
./run.sh
```

Keyin 5 soniya ichida brauzer/curl orqali:

- `http://localhost:8080/api/health`
- `http://localhost:8080/api/panic`

---

## Xulosa

O‘rganildi:

- Middleware yozish: `func(next http.Handler) http.Handler`
- Logging middleware
- Recovery middleware (`recover()`)
- HTTP error handling (405/500)
