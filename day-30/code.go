package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("🚀 Caching Patterns: In-Memory va Redis Demo")

	ctx := context.Background()
	
	// 1. In-Memory Cache demo
	memCache := NewInMemoryCache()
	
	fmt.Println("\n--- In-Memory Cache ---")
	memCache.Set(ctx, "user:1", "Matnazar", 2*time.Second)
	
	val, err := memCache.Get(ctx, "user:1")
	if err != nil {
		log.Printf("Xato: %v\n", err)
	} else {
		fmt.Printf("Topildi: %s\n", val)
	}

	fmt.Println("Kutish (2.5 soniya)...")
	time.Sleep(2500 * time.Millisecond)

	val, err = memCache.Get(ctx, "user:1")
	if err == ErrCacheMiss {
		fmt.Println("Cache Miss: Ma'lumot eskirgan yoki topilmadi")
	} else if err != nil {
		log.Printf("Xato: %v\n", err)
	} else {
		fmt.Printf("Topildi: %s\n", val)
	}

	// 2. Redis Cache demo (faqat ulangan bo'lsa ishlaydi, shuning uchun izohda qoldirilgan)
	/*
	redisCache := NewRedisCache("localhost:6379")
	err = redisCache.Set(ctx, "session:123", "active", 1*time.Minute)
	if err != nil {
		log.Printf("Redis Set Error: %v", err)
	}
	val, err = redisCache.Get(ctx, "session:123")
	fmt.Printf("Redis Get: %s, err: %v\n", val, err)
	*/
}
