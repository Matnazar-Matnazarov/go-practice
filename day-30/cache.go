package main

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var ErrCacheMiss = errors.New("cache miss")

type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string, expiration time.Duration) error
}

// InMemoryCache implementatsiyasi
type item struct {
	value      string
	expiration int64
}

type InMemoryCache struct {
	items map[string]item
	mu    sync.RWMutex
}

func NewInMemoryCache() *InMemoryCache {
	c := &InMemoryCache{
		items: make(map[string]item),
	}
	go c.cleanupLoop()
	return c
}

func (c *InMemoryCache) cleanupLoop() {
	ticker := time.NewTicker(1 * time.Minute)
	for range ticker.C {
		now := time.Now().UnixNano()
		c.mu.Lock()
		for k, v := range c.items {
			if v.expiration > 0 && now > v.expiration {
				delete(c.items, k)
			}
		}
		c.mu.Unlock()
	}
}

func (c *InMemoryCache) Get(ctx context.Context, key string) (string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	it, found := c.items[key]
	if !found {
		return "", ErrCacheMiss
	}

	if it.expiration > 0 && time.Now().UnixNano() > it.expiration {
		return "", ErrCacheMiss
	}

	return it.value, nil
}

func (c *InMemoryCache) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	var exp int64
	if expiration > 0 {
		exp = time.Now().Add(expiration).UnixNano()
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = item{
		value:      value,
		expiration: exp,
	}
	return nil
}

// RedisCache implementatsiyasi
type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(addr string) *RedisCache {
	return &RedisCache{
		client: redis.NewClient(&redis.Options{
			Addr: addr,
		}),
	}
}

func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", ErrCacheMiss
	}
	return val, err
}

func (r *RedisCache) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}
