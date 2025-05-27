package optimistic

import (
	"sync"
)

type OptimisticCache struct {
	mu   sync.RWMutex
	data map[string]*entry
}

type entry struct {
	value   any
	version uint64
}

func NewOptimisticCache() *OptimisticCache {
	return &OptimisticCache{
		data: make(map[string]*entry),
	}
}

func (c *OptimisticCache) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = &entry{value: value, version: 0}
}

func (c *OptimisticCache) Get(key string) (any, uint64) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if e, ok := c.data[key]; ok {
		return e.value, e.version
	}
	return nil, 0
}

func (c *OptimisticCache) Update(key string, newValue any) bool {
	// Шаг 1: Читаем данные без блокировки
	c.mu.RLock()
	e, ok := c.data[key]
	if !ok {
		c.mu.RUnlock()
		return false
	}
	currentVersion := e.version
	c.mu.RUnlock()

	// Шаг 2: Подготовка новых данных (может быть долгой операцией)

	// Шаг 3: Проверка версии с блокировкой
	c.mu.Lock()
	defer c.mu.Unlock()

	// Перепроверяем существование записи
	if e, ok := c.data[key]; !ok || e.version != currentVersion {
		return false
	}

	c.data[key] = &entry{
		value:   newValue,
		version: currentVersion + 1,
	}
	return true
}
