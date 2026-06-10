# Kun 31: Rate Limiting

## Kirish

Rate limiting - bu serverga ma'lum vaqt ichida keladigan so'rovlar sonini cheklash usuli. U DDoS hujumlardan himoya qilish, resurslarni tejash va barcha mijozlarga adolatli xizmat ko'rsatish uchun kerak.

## 1. Asosiy Algoritmlar

1. **Token Bucket**: Buketga har $1/r$ soniyada token qo'shiladi. So'rov kelganda, token olinadi. Token yo'q bo'lsa, so'rov rad etiladi. Burst (birdaniga ko'p so'rov) qilish imkonini beradi.
2. **Leaky Bucket**: So'rovlar doimiy tezlikda qayta ishlanadi. Agar navbat to'lib qolsa, yangilari tashlab yuboriladi.
3. **Fixed Window**: Vaqt intervallariga bo'linadi (masalan, soat 1:00 dan 1:01 gacha 100 ta so'rov).
4. **Sliding Window Log**: Har bir so'rovning vaqti saqlanadi. Aniqroq, lekin xotirani ko'p yeydi.

## 2. Go'da Token Bucket

Go tilida `golang.org/x/time/rate` paketi rasmiy Token Bucket implementatsiyasi hisoblanadi.

```go
// Har soniyada 2 ta ruxsat, maksimal burst = 5 ta
limiter := rate.NewLimiter(2, 5)

if !limiter.Allow() {
    fmt.Println("Too many requests!")
}
```

## 3. IP Bo'yicha Rate Limiter

Serverga keladigan barcha so'rovlarni birgalikda cheklash o'rniga, har bir IP manzil uchun alohida limiter yaratish yaxshi amaliyotdir. Buni `map[string]*rate.Limiter` va `sync.RWMutex` orqali amalga oshiramiz.

## 4. HTTP Middleware

Rate limiting ko'pincha HTTP API'larda Middleware sifatida ishlatiladi:

```go
func rateLimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if !limiter.Allow() {
            http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
            return
        }
        next.ServeHTTP(w, r)
    }
}
```
