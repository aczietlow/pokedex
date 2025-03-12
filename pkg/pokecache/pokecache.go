package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entry map[string]CacheEntry
	ttl   time.Duration
	mux   *sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		mux:   &sync.Mutex{},
		Entry: make(map[string]CacheEntry),
	}

	go cache.cacheCleanup(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	if _, exists := c.Entry[key]; exists == false {
		c.Entry[key] = CacheEntry{
			createdAt: time.Now(),
			val:       val,
		}
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	if entry, exists := c.Entry[key]; exists == true {
		return entry.val, exists
	}

	return nil, false
}

func (c *Cache) reapLoop() {
	c.mux.Lock()
	defer c.mux.Unlock()
	for key, ce := range c.Entry {
		if time.Now().Sub(ce.createdAt) > c.ttl {
			delete(c.Entry, key)
		}
	}
}

func (c *Cache) cacheCleanup(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// Everytime ticker.C sends a single do a thing
	for range ticker.C {
		c.reapLoop()
	}
}
