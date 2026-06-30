# Kun 27: Graceful Shutdown va Signal Handling

## Kirish

Real loyihalarda server faqat ishga tushishi yetmaydi — **to‘g‘ri to‘xtatish** ham kerak. Container (Docker/K8s) yoki systemd server’ni to‘xtatganda **SIGTERM** yuboradi; agar dastur darhol o‘lsa, mavjud so‘rovlar kesiladi va ma’lumot yo‘qolishi mumkin. **Graceful shutdown** — signal kelganda yangi so‘rovlarni qabul qilmaslik, mavjud so‘rovlar tugashini kutish, keyin chiqish.

---

## 1. Nima uchun graceful shutdown?

- **So‘rovlar tugashi** — HTTP handler’lar, DB transaction’lar tugaguncha kutish
- **Connection drain** — klientlar javob oladi, 502/504 kamayadi
- **Deploy va scaling** — K8s pod’ni evict qilganda in-flight request’lar xavfsiz tugaydi

---

## 2. Signal’lar

| Signal | Qachon keladi | Odatiy harakat |
|--------|----------------|----------------|
| **SIGINT**  | Ctrl+C | Dasturni to‘xtatish |
| **SIGTERM** | kill, container stop, systemctl stop | Dasturni to‘xtatish (graceful uchun ideal) |

Go’da **`os/signal`** package — signal’larni qabul qilish va channel orqali boshqarish.

```go
sigChan := make(chan os.Signal, 1)
signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
sig := <-sigChan  // bloklaydi, signal kelguncha
```

**Qoida:** Buffered channel (1) — signal handler bloklamasligi uchun.

---

## 3. context.WithTimeout — shutdown deadline

Shutdown cheksiz kutmasligi kerak — ba’zi so‘rovlar osilib qolishi mumkin. **Context with timeout** — ma’lum vaqtdan keyin majburiy tugatish.

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
if err := server.Shutdown(ctx); err != nil {
    log.Printf("shutdown: %v", err)
}
```

---

## 4. http.Server.Shutdown

**Shutdown** — yangi so‘rovlarni qabul qilmaydi, mavjud handler’lar tugaguncha (yoki context cancel/timeout gacha) kutadi. Listener’ni yopadi.

```go
if err := server.Shutdown(ctx); err != nil {
    // context.DeadlineExceeded — timeout; yoki boshqa xato
}
```

**ListenAndServe** esa Shutdown chaqilganda **ErrServerClosed** qaytaradi — bu normal, xato emas.

---

## 5. Umumiy tartib

1. **Server’ni goroutine’da ishga tushirish** — `go server.ListenAndServe()`
2. **Signal kutish** — `<-sigChan`
3. **Shutdown context** — `context.WithTimeout(..., 10*time.Second)`
4. **server.Shutdown(ctx)** — drain
5. **Qo‘shimcha resurslar** — DB connection pool yopish, cache flush va hokazo

---

## 6. Best practices

- **SIGTERM va SIGINT** ikkalasini ham qabul qiling
- **Timeout** — 10–30 soniya; production’da env’dan o‘qing
- **Shutdown xatosini log’lang** — timeout bo‘lsa diqqat qiling
- **ListenAndServe err** — `err == http.ErrServerClosed` bo‘lsa ignore qiling

---

## Xulosa

**Graceful shutdown** — signal kelganda server’ni tuzatib to‘xtatish; yangi so‘rov yo‘q, mavjudlar tugaydi, keyin dastur chiqadi. Real serverlar va container muhitida majburiy.

**Keyingi qadamlar:** Health endpoint (readiness/liveness), drain period, multiple server (gRPC + HTTP) shutdown.
