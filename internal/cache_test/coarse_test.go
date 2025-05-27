package cache_test

import (
	"sync"
	"sync-go/internal/cache"
	"testing"
)

type CoarseTest struct{}

func (t CoarseTest) RunTest(testT *testing.T) {
	c := cache.NewCoarseCache()

	// Тест 1: Конкурентные записи
	testT.Run("ConcurrentWrites", func(t *testing.T) {
		var wg sync.WaitGroup
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.Put("key", "value")
			}()
		}
		wg.Wait()
	})

	// Тест 2: Чтение после записи
	testT.Run("ReadAfterWrite", func(t *testing.T) {
		if val, _ := c.Get("key"); val != "value" {
			t.Error("Несоответствие данных")
		}
	})
}
func TestCoarseCache_Concurrency(t *testing.T) {
	c := cache.NewCoarseCache()
	var wg sync.WaitGroup

	// 100 писателей
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			c.Put("key", "value")
		}(i)
	}

	// 100 читателей
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Get("key")
		}()
	}

	wg.Wait()

	if val, _ := c.Get("key"); val != "value" {
		t.Errorf("Invalid value: %s", val)
	}
}
