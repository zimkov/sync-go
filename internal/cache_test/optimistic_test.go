package cache_test

import (
	"sync"
	"sync-go/internal/cache"
	"testing"
)

func TestOptimisticCache_Consistency(t *testing.T) {
	c := cache.NewOptimisticCache()
	var wg sync.WaitGroup

	// Параллельные записи
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			c.Put("key", "value")
		}(i)
	}

	// Параллельные чтения с проверкой
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if val, ok := c.Get("key"); ok && val != "value" {
				t.Errorf("Invalid value: %s", val)
			}
		}()
	}

	wg.Wait()
}

func TestOptimisticCache_VersionCheck(t *testing.T) {
	c := cache.NewOptimisticCache()
	c.Put("key", "value1")

	// Чтение с изменением данных во время операции
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		c.Get("key") // Должно обнаружить изменение
	}()

	go func() {
		defer wg.Done()
		c.Put("key", "value2")
	}()

	wg.Wait()
}
