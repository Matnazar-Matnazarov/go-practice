# Kun 16: Advanced Concurrency Patterns

## Maqsad

Bu kunda concurrency'ni faqat `goroutine/channel` darajasida emas, balki **amaliy pattern** sifatida ishlatishni o'rganamiz.

## Mavzular

1. Bounded concurrency (semaphore pattern)
2. Context cancellation (fan-out ishlar)
3. Goroutine leak'larni oldini olish (done channel, select)
4. Graceful shutdown (signal + context)

## Amaliy tavsiyalar

- Limit qo'ymasdan parallel ishlash (API/DB) tizimni yiqitishi mumkin.
- Xato chiqqanda qolgan ishlarni bekor qilish resursni tejaydi.
- Har bir goroutine uchun stop sharti bo'lsin (`context` yoki `done`).
