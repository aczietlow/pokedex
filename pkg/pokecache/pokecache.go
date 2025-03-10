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
		mux: &sync.Mutex{},
	}

	cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	if _, exists := c.Entry[key]; exists == false {
		c.Entry[key] = CacheEntry{
			createdAt: time.Now(),
			val:       val,
		}
	}
	c.mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	if entry, exists := c.Entry[key]; exists == true {
		c.mux.Unlock()
		return entry.val, exists
	}

	c.mux.Unlock()
	return nil, false

}

func (c *Cache) reapLoop() {
	c.mux.Lock()
	for key, ce := range c.Entry {
		if time.Now().Sub(ce.createdAt) > c.ttl {
			delete(c.Entry, key)
		}
	}
	c.mux.Unlock()
}
