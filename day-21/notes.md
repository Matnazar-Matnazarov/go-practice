# Kun 21: HTTP Client va API Testing (net/http, httptest)

## Kirish

Kun 20 da HTTP server va Web API yozishni ko‘rdik. Endi esa **HTTP klienti** (mijoz tomoni) va **API testlash** (httptest) ni professional darajada ko‘ramiz. Real loyihada server va klient bir-birini to‘ldiradi: server API beradi, klient esa shu API bilan ishlaydi; testing esa bularning sifatini kafolatlaydi.

---

## 1. HTTP Client Asoslari

### HTTP Client nima?

**HTTP Client** — boshqa serverlarga HTTP so‘rov yuborib, javobni o‘qiydigan kod. Go’da bu ishni **`http.Client`** va yordamchi funksiyalar (`http.Get`, `http.Post`, ...) bajaradi.

### Nima uchun kerak?

1. **External API lar** — to‘lov, SMS, email, third-party servislar
2. **Mikroxizmatlar o‘zaro muloqoti** — Service A → Service B REST chaqiriqlari
3. **Health check va monitoring** — ichki servislarni tekshirish

### http.Client vazifalari

1. **So‘rov yuborish** — URL, method, header, body
2. **Timeout boshqarish** — tarmoq muammolarida "osilib" qolmaslik
3. **Connection reuse** — connection pool va Keep-Alive
4. **Transport sozlash** — proxy, TLS, custom dialer

---

## 2. Oddiy GET va POST so‘rovlar

### http.Get va http.Post

```go
resp, err := http.Get("https://example.com")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()
body, err := io.ReadAll(resp.Body)
```

**Kamchilik:** `http.Get` default client ishlatadi; timeout yo‘q — production uchun xavfli.

### Professional yondashuv: custom http.Client

```go
client := &http.Client{Timeout: 10 * time.Second}
req, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
if err != nil { return }
req.Header.Set("Accept", "application/json")
resp, err := client.Do(req)
if err != nil { return }
defer resp.Body.Close()
```

**Qoida:** Production’da **har doim timeout** o‘rnating.

---

## 3. Context bilan ishlash

Context so‘rov umrini boshqarish uchun: timeout va cancellation.

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
resp, err := client.Do(req)
```

---

## 4. JSON API bilan ishlash

### JSON javobni o‘qish

```go
var u User
if err := json.NewDecoder(resp.Body).Decode(&u); err != nil { return }
```

### JSON body yuborish (POST)

```go
body, _ := json.Marshal(newUser)
req, err := http.NewRequest(http.MethodPost, baseURL+"/api/users", bytes.NewReader(body))
req.Header.Set("Content-Type", "application/json")
resp, err := client.Do(req)
```

---

## 5. API Testing: httptest

### Nima uchun httptest?

**net/http/httptest** — HTTP handler va klient kodini real port ochmasdan, tez va izolyatsiya qilingan tarzda test qilish.

### Handler’ni test qilish (Recorder)

```go
r := httptest.NewRequest(http.MethodGet, "/api/health", nil)
w := httptest.NewRecorder()
healthHandler(w, r)
resp := w.Result()
if resp.StatusCode != http.StatusOK { t.Fatal(...) }
```

### Test server (klient kodini test qilish)

```go
ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]any{"ok": true})
}))
defer ts.Close()
resp, err := http.Get(ts.URL)
```

---

## 6. Best Practices

- **Timeout majburiy** — `http.Client{Timeout: ...}`
- **Context** — uzun so‘rovlar uchun `NewRequestWithContext`
- **Error handling** — tarmoq xatosi, status kod, JSON xatolari
- **httptest** — handler va klientni real tarmoqqa chiqmasdan test qilish

---

## Xulosa

**HTTP Client** — boshqa serverlarga so‘rov yuboruvchi kod; REST API, external servislar, mikroxizmatlar muloqoti uchun.

**API Testing** — httptest bilan handler va klientni avtomatik tekshirish; xatolarni erta topish va refactoring’ni xavfsiz qilish.
