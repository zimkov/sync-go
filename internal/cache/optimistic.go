package cache

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type OptimisticCache struct {
	mu      sync.Mutex
	data    map[string]string
	version int64
}

func NewOptimisticCache() *OptimisticCache {
	return &OptimisticCache{
		data: make(map[string]string),
	}
}

func (c *OptimisticCache) Get(key string) (string, bool) {
	// Чтение без блокировки
	startVersion := atomic.LoadInt64(&c.version)

	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	val, ok := c.data[key]

	// Проверка, что данные не изменились во время чтения
	if atomic.LoadInt64(&c.version) != startVersion {
		return "", false // Данные изменились, требуется повтор
	}
	return val, ok
}

func (c *OptimisticCache) Put(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	c.data[key] = value
	atomic.AddInt64(&c.version, 1)
}
