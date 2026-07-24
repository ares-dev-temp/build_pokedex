package main

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu    sync.Mutex
	thing map[string]cacheEntry
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{sync.Mutex{}, make(map[string]cacheEntry)}
	go cache.readLoop(interval)
	return cache
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	entry := cacheEntry{time.Now(), val}
	c.thing[key] = entry
	c.mu.Unlock()
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	entry, ok := c.thing[key]
	c.mu.Unlock()

	if ok {
		return entry.val, true
	}

	return nil, false
}

func (c Cache) readLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.thing {
			if time.Since(entry.createdAt) > interval {
				delete(c.thing, key)
			}
		}
		c.mu.Unlock()
	}

}
