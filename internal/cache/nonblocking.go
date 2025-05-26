package cache

import (
	"math/rand"
	"sync/atomic"
	"time"
)

// Обертка для map, чтобы избежать сравнения несравнимых типов
type cacheSnapshot struct {
	data map[string]string
}

type NonBlockingCache struct {
	snapshot atomic.Value // Хранит *cacheSnapshot
}

func NewNonBlockingCache() *NonBlockingCache {
	c := &NonBlockingCache{}
	c.snapshot.Store(&cacheSnapshot{
		data: make(map[string]string),
	})
	return c
}

func (c *NonBlockingCache) Get(key string) (string, bool) {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	snap := c.snapshot.Load().(*cacheSnapshot)
	val, ok := snap.data[key]
	return val, ok
}

func (c *NonBlockingCache) Put(key, value string) {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)

	for {
		oldSnap := c.snapshot.Load().(*cacheSnapshot)
		newData := make(map[string]string, len(oldSnap.data)+1)

		// Копируем данные
		for k, v := range oldSnap.data {
			newData[k] = v
		}
		newData[key] = value

		newSnap := &cacheSnapshot{data: newData}

		// Сравниваем указатели, а не map
		if c.snapshot.CompareAndSwap(oldSnap, newSnap) {
			return
		}
	}
}
