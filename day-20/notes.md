# Kun 20: HTTP Server va Web API (net/http)

## Kirish

Go dasturlash tilida **net/http** — standart kutubxona bo‘lib, HTTP server va klient yozish, REST API qurish va veb-xizmatlarni boshqarish uchun ishlatiladi. Production darajadagi veb-serverlar va mikroxizmatlar ko‘pincha faqat standart kutubxona bilan yoziladi.

---

## 1. HTTP Server Asoslari

### HTTP Server nima?

**HTTP Server** — HTTP so‘rovlarni (request) qabul qilib, javob (response) qaytaradigan dastur.

**Handler** — `http.Handler` interfeysini qondiradigan tip; har bir so‘rov uchun javobni belgilaydi.

### Nima uchun kerak?

1. **REST API** — JSON/XML orqali ma’lumot almashish
2. **Veb-ilovalar** — HTML, statik fayllar, SPA
3. **Mikroxizmatlar** — ichki va tashqi API endpoint’lar
4. **Health check** — `/health`, `/ready` kabi monitoring endpoint’lar

### net/http vazifalari

1. **Server ishga tushirish** — `ListenAndServe`, port va handler
2. **URL yo‘nalishlash (routing)** — path bo‘yicha handler tanlash
3. **Request/Response** — header, body, status kod
4. **Middleware** — logging, auth, timeout (handler o‘rash)

---

## 2. Handler va HandlerFunc

### http.Handler interfeysi

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

Har qanday tip `ServeHTTP(w http.ResponseWriter, r *http.Request)` metodiga ega bo‘lsa, u **Handler** hisoblanadi.

### http.HandlerFunc

**HandlerFunc** — oddiy funksiyani Handler ga aylantiradi.

```go
func greet(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Salom, %s!", r.URL.Path[1:])
}

// HandlerFunc(greet) endi http.Handler
http.HandleFunc("/greet/", greet)
```

**Qoida:** `HandleFunc` ichida funksiya `func(ResponseWriter, *Request)` bo‘lishi kerak.

### ResponseWriter va Request

- **http.ResponseWriter** — javob yozish: `Write`, `WriteHeader`, header o‘rnatish
- **http.Request** — so‘rov: `Method`, `URL`, `Header`, `Body`, `Form`, `PostForm`

```go
w.WriteHeader(http.StatusOK)           // 200
w.Header().Set("Content-Type", "application/json")
io.WriteString(w, `{"ok":true}`)
```

---

## 3. ServeMux (Default Router)

### Default ServeMux

**ServeMux** — URL path bo‘yicha handlerni tanlaydi. `http.Handle` va `http.HandleFunc` default ServeMux ga qo‘shadi.

```go
http.HandleFunc("/", homeHandler)
http.HandleFunc("/api/users", usersHandler)
http.ListenAndServe(":8080", nil)  // nil = default ServeMux
```

**Qoida:** `nil` o‘rniga boshqa mux berilsa, faqat o‘sha mux ishlatiladi.

### Path matching qoidalari

- **Trailing slash:** `/api/` — `/api` va `/api/`, `/api/foo` ni qamrab oladi (ichki path’lar uchun).
- **Uzun path ustun:** `/api/users` `/api/users/1` dan oldin tekshiriladi; eng uzun mos path tanlanadi.
- **Subpath:** `/api/` bilan ro‘yxatdan o‘tgan handler `/api/xyz` kabi ichki yo‘llarni ham qabul qiladi.

### Yangi ServeMux (best practice)

Default mux global o‘zgaruvchilar ishlatadi. Yaxshi amaliyot — yangi `http.NewServeMux()` yaratib, unga handlerni bog‘lash.

```go
mux := http.NewServeMux()
mux.HandleFunc("/", homeHandler)
mux.HandleFunc("/api/health", healthHandler)
server := &http.Server{Addr: ":8080", Handler: mux}
server.ListenAndServe()
```

**Afzallik:** Test qilish oson, bir nechta server (port) uchun alohida mux.

---

## 4. ListenAndServe va Server konfiguratsiyasi

### http.ListenAndServe

```go
err := http.ListenAndServe(":8080", handler)
// ":8080" — barcha interfeyslar, 8080 port
// handler — nil bo‘lsa default ServeMux
```

**Qoida:** Bloklaydi; xato bo‘lsa `err` qaytadi (graceful shutdown uchun Server ni alohida boshqarish kerak).

### http.Server — to‘liq nazorat

```go
server := &http.Server{
    Addr:         ":8080",
    Handler:      mux,
    ReadTimeout:  15 * time.Second,
    WriteTimeout: 15 * time.Second,
    IdleTimeout:  60 * time.Second,
}
go server.ListenAndServe()
// server.Shutdown(ctx) — graceful shutdown
```

**ReadTimeout / WriteTimeout** — slow client va response timeout; **IdleTimeout** — Keep-Alive bo‘sh vaqt. Production da timeout’lar majburiy.

---

## 5. Request ma’lumotlarini o‘qish

### Method, URL, Header

```go
method := r.Method                    // GET, POST, ...
path   := r.URL.Path                  // /api/users
query  := r.URL.Query()               // ?key=value
name   := query.Get("name")           // birinchi "name" parametri

contentType := r.Header.Get("Content-Type")
```

### Body o‘qish (POST, PUT, PATCH)

```go
body, err := io.ReadAll(r.Body)
defer r.Body.Close()
if err != nil {
    http.Error(w, "Bad Request", http.StatusBadRequest)
    return
}
// JSON bo‘lsa: json.Unmarshal(body, &structVar)
```

**Qoida:** `r.Body` faqat bir marta o‘qiladi; keyin yopish (`Close`) yoki `io.ReadAll` kabi to‘liq o‘qish kerak.

### Form va PostForm

```go
err := r.ParseForm()
if err != nil { ... }
// application/x-www-form-urlencoded
email := r.FormValue("email")
name  := r.PostFormValue("name")  // faqat POST body
```

**ParseMultipartForm** — `multipart/form-data` (file upload) uchun.

---

## 6. Response yozish

### Status kod va header

```go
w.WriteHeader(http.StatusCreated)   // 201
w.Header().Set("Content-Type", "application/json; charset=utf-8")
w.Header().Set("X-Request-Id", requestID)
```

**Qoida:** `WriteHeader` faqat bir marta chaqiriladi; undan oldin header’lar o‘rnatilishi kerak.

### Body yozish

```go
fmt.Fprintf(w, "Hello, %s", name)
io.WriteString(w, "text")
json.NewEncoder(w).Encode(data)
w.Write([]byte("raw bytes"))
```

**http.Error** — status va text body uchun qisqa yo‘l:

```go
http.Error(w, "Unauthorized", http.StatusUnauthorized)
```

---

## 7. JSON API pattern

### JSON javob

```go
func usersHandler(w http.ResponseWriter, r *http.Request) {
    users := []User{{ID: 1, Name: "Ali"}, {ID: 2, Name: "Vali"}}
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(users); err != nil {
        // WriteHeader allaqachon 200 — xato log qilish kerak
    }
}
```

### JSON so‘rov body (POST)

```go
var input CreateUserRequest
if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
    http.Error(w, "Invalid JSON", http.StatusBadRequest)
    return
}
defer r.Body.Close()
```

**Best practice:** Request uchun struct, validation (masalan, required fieldlar), keyin business logic.

---

## 8. Middleware (Handler o‘rash)

Middleware — bitta Handler ni qabul qilib, yangi Handler qaytaradi; logging, auth, panic recovery va hokazo qo‘shish uchun.

### Oddiy middleware

```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
    })
}

mux := http.NewServeMux()
mux.HandleFunc("/", homeHandler)
wrapped := loggingMiddleware(mux)
http.ListenAndServe(":8080", wrapped)
```

### Ketma-ket middleware

```go
chain := loggingMiddleware(authMiddleware(mux))
```

**Qoida:** Ichki handler birinchi chaqiriladi; tashqi middleware so‘rovni oldin ko‘radi, javobni keyin.

---

## 9. Graceful Shutdown

Server ni to‘xtatishda mavjud so‘rovlar tugashini kutish — **graceful shutdown**.

```go
server := &http.Server{Addr: ":8080", Handler: mux}
go server.ListenAndServe()

sigChan := make(chan os.Signal, 1)
signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
<-sigChan

ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
if err := server.Shutdown(ctx); err != nil {
    log.Fatal(err)
}
```

**Shutdown** — yangi so‘rovlarni qabul qilmaydi, mavjud handler’lar tugaguncha (yoki timeout gacha) kutadi.

---

## 10. Muhim qoidalar

### Server

1. **Timeout** — `ReadTimeout`, `WriteTimeout`, `IdleTimeout` o‘rnating.
2. **Graceful shutdown** — `Shutdown(ctx)` bilan to‘xtating.
3. **Yangi ServeMux** — default mux o‘rniga `http.NewServeMux()` ishlatish yaxshiroq.

### Handler

1. **Content-Type** — JSON bo‘lsa `application/json` o‘rnating.
2. **Error handling** — `r.Body` o‘qish/xato bo‘lsa 400/500 va `http.Error` yoki JSON error body.
3. **Method tekshirish** — faqat kerakli method’larni qabul qiling: `if r.Method != http.MethodPost { ... }`.

### Xavfsizlik

1. **Path traversal** — `r.URL.Path` ni sanitize qiling; fayl sistemaga to‘g‘ridan-to‘g‘ri path bermang.
2. **Rate limiting** — production da middleware orqali so‘rovlar sonini cheklang.
3. **Header’lar** — `X-Content-Type-Options`, CORS kerak bo‘lsa to‘g‘ri sozlang.

---

## Xulosa

### HTTP Server

**Nima:** HTTP so‘rovlarni qabul qiluvchi va javob qaytaruvchi dastur.

**Vazifasi:**
- Handler va HandlerFunc
- ServeMux (routing)
- Request/Response boshqaruvi
- JSON API, middleware, graceful shutdown

**Afzalliklari:**
- Faqat standart kutubxona — qo‘shimcha framework shart emas
- Tez va kam xotira
- Production uchun yetarli

### Keyingi qadamlar

- Chiqarilgan router (gorilla/mux, echo, gin) — kerak bo‘lsa
- HTTPS (TLS), HTTP/2
- Middleware zanjiri, recovery, CORS
- Integration testlar (`httptest`)
