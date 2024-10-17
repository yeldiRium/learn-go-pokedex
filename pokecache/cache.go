package pokecache

import (
	"time"
)

type cacheEntry struct {
	lastUsed time.Time
	value    []byte
}

type Cache struct {
	capacity int
	entries  map[string]*cacheEntry
}

// If this reduces the capacity, it has runtime O(n^2) where n = oldCapacity
func (cache *Cache) SetCapacity(capacity int) {
	cache.capacity = capacity
	cache.prune()
}

func (cache *Cache) Len() int {
	return len(cache.entries)
}

func (cache *Cache) prune() {
	if cache.capacity == 0 {
		return
	}

	for cache.Len() > cache.capacity {
		oldestTimestamp := time.Now()
		oldestKey := ""
		for key, entry := range cache.entries {
			if entry.lastUsed.Before(oldestTimestamp) {
				oldestTimestamp = entry.lastUsed
				oldestKey = key
			}
		}
		delete(cache.entries, oldestKey)
	}
}

func (cache *Cache) AddEntry(key string, value []byte) {
	if cache.entries == nil {
		cache.entries = make(map[string]*cacheEntry)
	}

	newEntry := cacheEntry{
		value:    value,
		lastUsed: time.Now(),
	}

	cache.entries[key] = &newEntry

	cache.prune()
}

func (cache *Cache) GetEntry(key string) ([]byte, bool) {
	entry, ok := cache.entries[key]
	if !ok {
		return []byte{}, false
	}

	entry.lastUsed = time.Now()

	return entry.value, true
}
