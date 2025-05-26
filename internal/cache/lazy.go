package cache

import (
	"math/rand"
	"sync"
	"time"
)

type LazyCache struct {
	once sync.Once
	mu   sync.RWMutex
	data map[string]string
}

func NewLazyCache() *LazyCache {
	return &LazyCache{}
}

func (c *LazyCache) init() {
	c.data = make(map[string]string)
}

func (c *LazyCache) Get(key string) (string, bool) {
	c.once.Do(c.init)

	c.mu.RLock()
	defer c.mu.RUnlock()

	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	val, ok := c.data[key]
	return val, ok
}

func (c *LazyCache) Put(key, value string) {
	c.once.Do(c.init)

	c.mu.Lock()
	defer c.mu.Unlock()

	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	c.data[key] = value
}
