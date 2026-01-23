# Kun 17: Concurrency Patterns II (Pipeline, Worker Pool, Rate Limit, ErrGroup)

## Maqsad

Day-16 dagi patternlardan keyin, bu kunda concurrency’ni real servislar va data-processing oqimlarida ko‘p ishlatiladigan **pipeline / worker-pool / rate-limit / errgroup** patternlari orqali mustahkamlaymiz.

## Mavzular

1. **Pipeline** (generator → stages → sink)
2. **Fan-out / Fan-in** (Worker pool)
3. **Rate limiting** (token bucket)
4. **ErrGroup-style**: birinchi xatoda context cancel va qolgan ishlarni to‘xtatish

## 1) Pipeline

Pipeline odatda:
- data oqimini bosqichma-bosqich qayta ishlash (decode → validate → transform → store)
- har bosqichda `context` orqali cancel qilish
- har bosqich `in <-chan T` qabul qilib `out <-chan U` qaytaradi

**Qoida:** Har stage goroutine’ida `select { case <-ctx.Done(): return }` bo‘lsin.

## 2) Worker pool (Fan-out/Fan-in)

Worker pool:
- input job’larni bir nechta worker goroutine parallel qayta ishlaydi (fan-out)
- natijalar bitta channel’da yig‘iladi (fan-in)

**Qoida:**
- producer `jobCh` ni yopadi (`close(jobCh)`)
- consumer’lar `range jobCh` yoki `j, ok := <-jobCh` bilan o‘qiydi
- natija channel’ini yopish odatda `wg.Wait()` dan keyin bo‘ladi

## 3) Rate limiting (token bucket)

Token bucket:
- har \(T\) vaqtda bitta token qo‘shiladi
- `burst` token to‘planishi mumkin (qisqa “spike”larni ko‘taradi)

Real ishlatilishi:
- API request limit
- DB query limit
- background job limit

## 4) ErrGroup-style cancel

Maqsad:
- parallel ishlar ichida **birinchi error** chiqqanda `cancel()` qilish
- qolgan goroutine’lar `ctx.Done()` orqali to‘xtashi

Go standart kutubxonasida `errgroup` yo‘q (u `golang.org/x/sync/errgroup`), lekin pattern’ni minimal ko‘rinishda o‘zingiz yozishingiz mumkin:
- `Group{wg, err, cancel}`
- `Go(fn)` ichida error bo‘lsa `err` ni set qilib `cancel()`

## Amaliy tavsiyalar

- **Har goroutine uchun stop sharti**: `context` yoki `done` channel.
- **Channel ownership**: kim yaratgan bo‘lsa, odatda o‘sha yopadi.
- **Buffered vs unbuffered**: buffering backpressure va throughput’ga ta’sir qiladi.
- **Rate limit**: burst’ni juda katta qilmang — spike’lar tizimni yiqitishi mumkin.

# Kun 17: Worker Pools, Pipelines va Rate Limiting (Concurrency Advanced)

## Kirish

Kun 16 da advanced concurrency patternlarni (bounded concurrency, cancellation, leak prevention, graceful shutdown) ko'rdik.

Bu kunda real loyihalarda juda ko'p ishlatiladigan 3 ta patternni ko'ramiz:

1. **Worker Pool** – ko'p ishlarni cheklangan ishchilar bilan bajarish
2. **Pipeline** – bosqichma-bosqich (stage) qayta ishlash
3. **Rate Limiting** – tashqi servis / resursni himoya qilish

---

## 1. Worker Pool (Ishchilar hovuzi)

### Nima?

**Worker pool** – bu pattern bo'lib, sizda:

- Ishlar (`jobs`) navbati bo'ladi
- Ma'lum miqdorda ishchilar (`workers`) bo'ladi
- Har bir worker navbatdan ish olib, qayta ishlaydi

### Nega kerak?

1. **Bounded concurrency**
   - Cheksiz goroutine ochmaslik
   - Resurslarni (CPU, DB, API) himoya qilish

2. **Throughput**
   - Bir vaqtning o'zida bir nechta ish bajarish

3. **Nazorat**
   - Worker sonini boshqarish oson

### Arxitektura

- `jobs` kanali – ishlar navbati
- `results` kanali – natijalar navbati
- `N` ta worker goroutine – har biri `jobs` kanalidan ish olib, `results` kanaliga natija yozadi

```go
type Job struct {
    ID int
    Value int
}

type Result struct {
    JobID int
    Output int
}
```

### Qoidalar

- Worker soni – CPU, IO, tashqi API ga qarab tanlanadi
- Har doim kanallarni yopish (close) haqida o'ylang
- Context yoki `done` kanallari bilan to'xtatish imkoniyatini qoldiring

---

## 2. Pipeline Pattern

### Nima?

**Pipeline** – ma'lumot bir nechta bosqichdan (`stage`) ketma-ket o'tishi:

```text
source -> stage1 -> stage2 -> stage3 -> sink
```

Har bir stage:

- Kirish kanali (input)
- Chiqish kanali (output)
- Ichida 1 yoki bir nechta worker bo'lishi mumkin

### Nega kerak?

1. **Responsibility separation**
   - Har bir stage bitta ish uchun mas'ul

2. **Parallelism**
   - Har bir stage alohida goroutine / worker pool bo'lishi mumkin

3. **Composition**
   - Yangi pipeline larni modul sifatida yasash oson

### Misol

- Stage 1: sonlar generatsiyasi
- Stage 2: kvadrat olish
- Stage 3: filtr (faqat juft kvadratlar)

---

## 3. Rate Limiting

### Nima?

**Rate limiting** – tashqi servis yoki resursni ma'lum tezlikda chaqirish.

Masalan:

- Telegram / Discord API: 30 req/sekundan ko'p bo'lmasin
- DB: 50 query/sekundan ko'p bo'lmasin

### Oddiy usullar

1. **`time.Ticker` bilan**
2. **Token bucket** pattern

### Qoidalar

- Tashqi servis dokumentatsiyasini o'qing
- Limitga yaqinlashganda log yozing
- Limit oshsa – kutish yoki backoff qilishingiz kerak

---

## 4. Worker Pool Dizayn Detallari

### Kanal dizayni

- `jobs := make(chan Job)`
- `results := make(chan Result)`

Worker:

```go
func worker(id int, jobs <-chan Job, results chan<- Result) {
    for job := range jobs {
        // ishni bajarish
        // natijani results kanaliga yozish
    }
}
```

Dispatcher:

- Workerlarni ishga tushiradi
- Ishlarni `jobs` kanaliga yuboradi
- `jobs` kanalini yopadi

### Xatolar bilan ishlash

Variantlar:

- `Result` struct ichida `Err error` maydoni
- Alohida `errors` kanali

---

## 5. Pipeline Dizayn Detallari

### Stage funksiyasi

Har bir stage quyidagicha bo'lishi mumkin:

```go
func stage(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for v := range in {
            // qayta ishlash
            out <- v
        }
    }()
    return out
}
```

### Qoidalar

- Har bir stage chiqish kanalini yopishi kerak (`defer close(out)`).
- Kiritilgan kanal (`in`) ni stage yopmaydi.
- Context yoki `done` kanali bilan pipeline ni to'xtatish mumkin.

---

## 6. Rate Limiting Detallari

### `time.Ticker` bilan oddiy limiter

```go
ticker := time.NewTicker(time.Second)
defer ticker.Stop()

for _, req := range requests {
    <-ticker.C // har 1 sekundda bitta ish
    doRequest(req)
}
```

### Token bucket g'oyasi

- Har 100ms da 1 token qo'shiladi
- Har request 1 token ishlatadi
- Token bo'lmasa – kutish

---

## 7. Best Practices (Concurrency)

1. **Goroutine leak'lar yo'qligiga ishonch hosil qiling**
   - Har bir goroutine uchun chiqish yo'li bo'lsin

2. **Context ishlating**
   - `context.WithCancel`, `context.WithTimeout`

3. **`select` dan foydalaning**
   - `case <-ctx.Done()` har bir bloklovchi operatsiyada bo'lsin

4. **Kanallarni ehtiyotkorlik bilan yoping**
   - Faqat yozuvchi kanalni yopadi

5. **Bounded concurrency**
   - Worker pools / semaphore pattern

---

## Xulosa

Bu kunda ko'rilgan patternlar real loyihalarda juda ko'p ishlatiladi:

- **Worker Pool** – ko'p ishlarni cheklangan ishchilar bilan bajarish
- **Pipeline** – bosqichma-bosqich qayta ishlash
- **Rate Limiting** – tashqi servis va resurslarni himoya qilish

Keyingi qadamlar:

- Bu patternlarni kichik loyihada qo'llash
- HTTP server bilan integratsiya qilish
- DB / external API bilan birga ishlatish

