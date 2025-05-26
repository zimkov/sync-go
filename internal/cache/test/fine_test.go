package test

import (
	"sync"
	"testing"

	"sync-go/internal/cache"
)

func TestFineCache_ReadHeavyWorkload(t *testing.T) {
	c := cache.NewFineCache()
	c.Put("key", "initial")

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 90% чтений, 10% записей
			if i%10 == 0 {
				c.Put("key", "updated")
			} else {
				c.Get("key")
			}
		}()
	}

	wg.Wait()

	if val, _ := c.Get("key"); val != "updated" {
		t.Errorf("Final value mismatch: %s", val)
	}
}

func TestFineCache_MixedWorkload(t *testing.T) {
	c := cache.NewFineCache()
	var wg sync.WaitGroup

	// 50 писателей
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			c.Put("key", "value")
		}(i)
	}

	// 200 читателей
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Get("key")
		}()
	}

	wg.Wait()
}
