package pokecache_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeldiRium/learning-go-pokedex/pokecache"
)

func TestCache(t *testing.T) {
	t.Run("should store and return an entry", func(t *testing.T) {
		cache := pokecache.Cache{}

		cache.AddEntry("test-key", []byte("test-value"))
		entry, ok := cache.GetEntry("test-key")

		assert.True(t, ok)
		assert.Equal(t, "test-value", string(entry))
	})

	t.Run("should remove entries once the size limit is reached", func(t *testing.T) {
		cache := pokecache.Cache{}
		cache.SetCapacity(2)

		cache.AddEntry("test-key-1", []byte("test-value"))
		cache.AddEntry("test-key-2", []byte("test-value"))
		cache.AddEntry("test-key-3", []byte("test-value"))
		_, ok := cache.GetEntry("test-key-1")
		assert.False(t, ok, "oldest key was contained in cache although it should have been pruned")
		_, ok = cache.GetEntry("test-key-2")
		assert.True(t, ok)
		_, ok = cache.GetEntry("test-key-3")
		assert.True(t, ok)
	})

	t.Run("setting the size should prune the cache to the new size", func(t *testing.T) {
		cache := pokecache.Cache{}

		cache.AddEntry("test-key-1", []byte("test-value"))
		cache.AddEntry("test-key-2", []byte("test-value"))
		cache.AddEntry("test-key-3", []byte("test-value"))

		_, ok := cache.GetEntry("test-key-1")
		assert.True(t, ok)
		_, ok = cache.GetEntry("test-key-2")
		assert.True(t, ok)
		_, ok = cache.GetEntry("test-key-3")
		assert.True(t, ok)

		cache.SetCapacity(2)

		_, ok = cache.GetEntry("test-key-1")
		assert.False(t, ok)
		_, ok = cache.GetEntry("test-key-2")
		assert.True(t, ok)
		_, ok = cache.GetEntry("test-key-3")
		assert.True(t, ok)
	})

	t.Run("pruning considers the time a key was last used", func(t *testing.T) {
		cache := pokecache.Cache{}

		cache.AddEntry("test-key-1", []byte("test-value"))
		cache.AddEntry("test-key-2", []byte("test-value"))
		cache.AddEntry("test-key-3", []byte("test-value"))

		cache.GetEntry("test-key-1")

		cache.SetCapacity(2)

		_, ok := cache.GetEntry("test-key-1")
		assert.True(t, ok)
		_, ok = cache.GetEntry("test-key-2")
		assert.False(t, ok)
		_, ok = cache.GetEntry("test-key-3")
		assert.True(t, ok)
	})
}
