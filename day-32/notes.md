# Kun 32: Profiling va Performance (pprof)

## Kirish

Go dasturlarini tezligini oshirish va xotira sarfini kamaytirish uchun ularni analiz qilish kerak bo'ladi. Buni Go'da **Profiling** deyiladi va asosiy asbob - `pprof`.

## 1. pprof Nima?

`pprof` (Profile) Go standart kutubxonasidagi asbob bo'lib, dasturning quyidagi jihatlarini analiz qiladi:
- **CPU**: Qaysi funksiyalar protsessorni eng ko'p band qilmoqda.
- **Heap (Xotira)**: Qaysi funksiyalar eng ko'p xotira ajratmoqda.
- **Goroutines**: Qaysi goroutine'lar ishlayapti va qayerda bloklangan.
- **Mutex**: Qulflashlar (locks) ustidagi raqobatni tahlil qiladi.

## 2. HTTP Serverda pprof ishlatish

HTTP serverda profileni yoqish juda oson:

```go
import _ "net/http/pprof"

func main() {
    // Profiler /debug/pprof/ endpointlarida avtomatik ishga tushadi
    http.ListenAndServe(":6060", nil)
}
```

## 3. Profileni O'qish

Terminaldan turib pprof asbobi orqali ma'lumotlarni yig'ish mumkin:

- **CPU profile (30 soniyalik)**: 
  `go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30`
- **Xotira (Heap) profile**: 
  `go tool pprof http://localhost:6060/debug/pprof/heap`

Interaktiv interfeys ochilgach, `top`, `list <funksiya>`, `web` kabi buyruqlarni ishlatish mumkin.

## 4. Benchmarklar

Kodning ma'lum bir qismini tezligini o'lchash uchun `testing` paketidagi benchmarklardan foydalaniladi.

```go
func BenchmarkMyFunc(b *testing.B) {
    for i := 0; i < b.N; i++ {
        MyFunc() // b.N avtomatik ortib boradi
    }
}
```
Ishga tushirish: `go test -bench=. -benchmem` (`-benchmem` xotira ajratish statistikasini ko'rsatadi).
