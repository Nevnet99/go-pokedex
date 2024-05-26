package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu       sync.Mutex
	entries  map[string]cacheEntry
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
	}

	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {
			cache.reapLoop()
		}
	}()

	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	if ok {
		return entry.val, true
	}

	return nil, false
}

func (c *Cache) reapLoop() {
	c.mu.Lock()
	defer c.mu.Unlock()

	expirationTime := time.Now().Add(-c.interval)

	for key, entry := range c.entries {
		if entry.createdAt.Before(expirationTime) {
			delete(c.entries, key)
		}
	}

}
