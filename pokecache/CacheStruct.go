package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entry map[string]cacheEntry
	mu    *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entry[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.entry[key]
	return entry.val, exists
}

func (c Cache) reapLoop(dur time.Duration) {
	ticker := time.NewTicker(dur)
	go func() {
		for {
			select {
			case t := <-ticker.C:
				c.mu.Lock()
				for key, val := range c.entry {
					if val.createdAt.Add(dur).Before(t) {
						delete(c.entry, key)
					}
				}
				c.mu.Unlock()
			}
		}
	}()
}
