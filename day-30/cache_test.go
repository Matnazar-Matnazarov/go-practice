package main

import (
	"context"
	"testing"
	"time"
)

func TestInMemoryCache_SetAndGet(t *testing.T) {
	cache := NewInMemoryCache()
	ctx := context.Background()

	err := cache.Set(ctx, "key1", "value1", 10*time.Minute)
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	val, err := cache.Get(ctx, "key1")
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}

	if val != "value1" {
		t.Errorf("Expected 'value1', got '%v'", val)
	}
}

func TestInMemoryCache_Expiration(t *testing.T) {
	cache := NewInMemoryCache()
	ctx := context.Background()

	// Kichik TTL bilan saqlaymiz
	err := cache.Set(ctx, "key2", "value2", 100*time.Millisecond)
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	// 150ms kutamiz (muddati o'tishi uchun)
	time.Sleep(150 * time.Millisecond)

	_, err = cache.Get(ctx, "key2")
	if err != ErrCacheMiss {
		t.Errorf("Expected ErrCacheMiss, got %v", err)
	}
}

func TestInMemoryCache_Miss(t *testing.T) {
	cache := NewInMemoryCache()
	ctx := context.Background()

	_, err := cache.Get(ctx, "non_existent")
	if err != ErrCacheMiss {
		t.Errorf("Expected ErrCacheMiss, got %v", err)
	}
}
