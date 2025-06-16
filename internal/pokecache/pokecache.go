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
	cacheMap map[string]cacheEntry
	interval time.Duration
	mu       sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{
		cacheMap: make(map[string]cacheEntry),
		interval: interval,
	}

	go newCache.reapLoop()
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value := cacheEntry{createdAt: time.Now(), val: val}

	c.cacheMap[key] = value
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if value, exists := c.cacheMap[key]; exists {
		return value.val, true
	}

	return nil, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		<-ticker.C
		c.mu.Lock()

		for key, val := range c.cacheMap {
			if time.Since(val.createdAt) > c.interval {
				delete(c.cacheMap, key)
			}
		}

		c.mu.Unlock()
	}
}
