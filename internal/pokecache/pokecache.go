package pokecache

import (
	"time"
)

type Cache struct {
	Entries map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(duration time.Duration) Cache {
	return Cache{
		Entries: map[string]cacheEntry{},
	}
}

func (cache *Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	cache.Entries[key] = entry
}

func (cache *Cache) Get(key string) (val []byte, exists bool) {
	return cache.Entries[key].val, true
}