package cache

import (
	"math/rand"
	"sync"
	"time"
)

type FineCache struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewFineCache() *FineCache {
	return &FineCache{
		data: make(map[string]string),
	}
}

func (c *FineCache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	val, ok := c.data[key]
	return val, ok
}

func (c *FineCache) Put(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	c.data[key] = value
}
