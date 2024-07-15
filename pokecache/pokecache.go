package pokecache

import (
	"sync"
	"time"
)

func NewCache(dur time.Duration) Cache {
	cache := Cache{
		entry: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}
	cache.reapLoop(dur)
	return cache
}
