package cache

import (
	"sync"
	"time"
)

type Cache interface {
	BuildKey(keys ...string) string

	Get(key string) (interface{}, bool)

	Set(key string, value interface{}, ttl time.Duration)

	Delete(key string)

	Clear()
}

type cache struct {
	m sync.RWMutex
	v map[string]entryCache
}

type entryCache struct {
	value   interface{}
	expired time.Time
}

// MaxDuration is the maximum duration that can be represented by a Duration.
const maxDuration time.Duration = 1<<63 - 1

// Init Cache with ttl time.Duration
//   - If ttl == 0 then ttl = maxDuration
func InitCache(ttl time.Duration) Cache {
	values := make(map[string]entryCache)

	// Timer for clear expired cache
	if ttl == 0 {
		ttl = maxDuration
	}
	go func() {
		timer := time.NewTicker(ttl)
		defer timer.Stop()

		select {
		case <-timer.C:
			// Clear expired cache
			for k, v := range values {
				if time.Now().After(v.expired) {
					delete(values, k)
				}
			}

		}
	}()
	return &cache{
		v: values,
	}
}

// Function BuildKey by args ...string and return string
func (c *cache) BuildKey(keys ...string) string {
	key := ""
	for _, v := range keys {
		key += v + ":"
	}

	return key
}

// Function Get cache by key string
func (c *cache) Get(key string) (interface{}, bool) {
	c.m.RLock()
	defer c.m.RUnlock()

	if v, ok := c.v[key]; ok {
		if v.expired.Before(time.Now()) {
			return v.value, true
		}

		return "", false

	}
	return "", false
}

// Function Set cache by key string, value interface{}, ttl time.Duration
func (c *cache) Set(key string, value interface{}, ttl time.Duration) {
	c.m.Lock()
	defer c.m.Unlock()

	if ttl == 0 {
		ttl = maxDuration
	}

	c.v[key] = entryCache{
		value:   value,
		expired: time.Now().Add(ttl),
	}
}

// Function Delete cache by key string
func (c *cache) Delete(key string) {
	c.m.Lock()
	defer c.m.Unlock()

	delete(c.v, key)
}

// Function Clear all cache
func (c *cache) Clear() {
	c.m.Lock()
	defer c.m.Unlock()

	c.v = make(map[string]entryCache)
}
