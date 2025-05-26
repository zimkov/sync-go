package test

import (
	"sync"
	"testing"

	"sync-go/internal/cache"
)

func TestLazyCache_InitOnce(t *testing.T) {
	c := cache.NewLazyCache()
	var wg sync.WaitGroup

	// 100 операций до инициализации
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Put("key", "value")
		}()
	}

	wg.Wait()

	if val, _ := c.Get("key"); val != "value" {
		t.Errorf("Initialization failed: %s", val)
	}
}

func TestLazyCache_Uninitialized(t *testing.T) {
	c := cache.NewLazyCache()

	// Чтение без инициализации
	if _, ok := c.Get("key"); ok {
		t.Error("Unexpected data in uninitialized cache")
	}
}
