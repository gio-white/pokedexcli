package pokecache

import (
	"sync"
	"time"
)


type cacheEntry struct {
	createdAt 	time.Time
	val			[]byte
}


type Cache struct {
	data 		map[string]cacheEntry
	mu			sync.Mutex
	interval 	time.Duration
}



func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		data: make(map[string]cacheEntry),
		interval: interval,
	}
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			c.reapLoop()
		}
	}()

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.mu.Unlock()
}


func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, ok := c.data[key]
	if ok {
		return item.val, true
	}
	return nil, false
}


func (c *Cache) reapLoop() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.data {
		if v.createdAt.Before(time.Now().UTC().Add(-c.interval)) {
			delete(c.data, k)
		}
	}
}