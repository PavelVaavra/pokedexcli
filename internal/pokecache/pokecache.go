package pokecache

import (
	"time"
	"sync"
	"fmt"
)

type Cache struct {
	Entries map[string]cacheEntry
	Mux *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(duration time.Duration) Cache {
	cache := Cache{
		Entries: map[string]cacheEntry{},
		Mux: &sync.Mutex{},
	}
	go cache.reapLoop(duration)
	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	cache.Mux.Lock()
	fmt.Println("Adding: ", key)
	cache.Entries[key] = entry
	cache.Mux.Unlock()
}

func (cache *Cache) Get(key string) (val []byte, exists bool) {
	cache.Mux.Lock()
	entry, ok := cache.Entries[key]
	cache.Mux.Unlock()
	if ok {
		fmt.Println("Getting: ", key)
		return entry.val, true
	}
	return nil, false
}

func (cache *Cache) reapLoop(duration time.Duration) {
	ticker := time.NewTicker(duration)
	for {
		<-ticker.C
		cache.Mux.Lock()
		for key, value := range cache.Entries {
			if time.Since(value.createdAt) > duration {
				fmt.Println("Deleting: ", key)
				delete(cache.Entries, key)
			}
		}
		cache.Mux.Unlock()
	}
}