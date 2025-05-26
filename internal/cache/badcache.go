package cache

import (
	"math/rand"
	"time"
)

type BadCache struct {
	data map[string]string
}

func NewBadCache() *BadCache {
	return &BadCache{
		data: make(map[string]string),
	}
}

func (c *BadCache) Get(key string) (string, bool) {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	val, ok := c.data[key]
	return val, ok
}

func (c *BadCache) Put(key, value string) {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	c.data[key] = value
}
