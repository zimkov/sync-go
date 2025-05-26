package cache

import (
	"math/rand"
	"sync"
	"time"
)

type CoarseCache struct {
	mu   sync.Mutex
	data map[string]string
}

func NewCoarseCache() *CoarseCache {
	return &CoarseCache{
		data: make(map[string]string),
	}
}

func (c *CoarseCache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	val, ok := c.data[key]
	return val, ok
}

func (c *CoarseCache) Put(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	c.data[key] = value
}
