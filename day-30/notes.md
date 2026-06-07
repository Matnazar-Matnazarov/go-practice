# Kun 30: Cache Pattern (In-Memory va Redis)

## Kirish

Ilovalarni tezlashtirishning eng yaxshi usullaridan biri — ma'lumotlarni keshlash (caching). Ma'lumotlar bazasiga har safar so'rov yuborish o'rniga, tez-tez so'raladigan ma'lumotlarni operativ xotirada yoki maxsus kesh serverlarida (masalan, Redis, Memcached) saqlaymiz.

## 1. Keshlash Naqshlari (Caching Patterns)

- **Cache-Aside (Lazy Loading)**: Ilova oldin keshdan qidiradi. Topilmasa, bazadan o'qiydi va keshga yozib qo'yadi. (Eng ko'p ishlatiladigani)
- **Write-Through**: Ilova ma'lumotni ham keshga, ham bazaga bir vaqtda yozadi. Kesh doim yangi.
- **Write-Behind (Write-Back)**: Ilova ma'lumotni faqat keshga yozadi. Orqa fonda, boshqa jarayon keshdagi o'zgarishlarni bazaga sinxronlashtiradi (tezkor yozish uchun).

## 2. In-Memory Cache (Go'da)

Go'da in-memory kesh yaratish uchun odatda `map` va `sync.RWMutex` (parallel o'qish/yozishni boshqarish uchun) ishlatiladi.

```go
type item struct {
	value      string
	expiration int64
}

type InMemoryCache struct {
	items map[string]item
	mu    sync.RWMutex
}
```

**Kamchiliklari**: 
- Agar server o'chib qolsa, barcha kesh yo'qoladi.
- Bir nechta serverlar (instances) o'rtasida kesh ulashilmaydi.
- Garbage Collector'ga og'irlik tushishi mumkin (agar kesh juda katta bo'lsa).

## 3. Redis Bilan Keshlash

Redis — xotirada ishlovchi, kalit-qiymat (key-value) turidagi ma'lumotlar bazasi.

Afzalliklari:
- **Taqsimlangan (Distributed)**: Ko'plab serverlar bitta Redis klasteridan foydalanishi mumkin.
- **Persistance**: Ma'lumotlarni diskka saqlashi mumkin.
- **Data Structures**: Hash, List, Set, Sorted Set kabi turlarni qo'llab-quvvatlaydi.

Go'da Redis bilan ishlash uchun `github.com/redis/go-redis/v9` paketi ishlatiladi:

```go
import "github.com/redis/go-redis/v9"

client := redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "", // parolsiz
    DB:       0,  // default DB
})

err := client.Set(ctx, "key", "value", 1*time.Minute).Err()
val, err := client.Get(ctx, "key").Result()
```
