package test

import (
	"fmt"
	"sync"
	"sync-go/internal/cache"
	"testing"
	"time"
)

func TestNonBlockingCache_ConcurrentWrites(t *testing.T) {
	c := cache.NewNonBlockingCache()
	var wg sync.WaitGroup

	keys := []string{"a", "b", "c", "d", "e"}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			key := keys[idx%5]
			c.Put(key, "value")
		}(i)
	}

	wg.Wait()

	for _, key := range keys {
		if val, _ := c.Get(key); val != "value" {
			t.Errorf("Key %s has invalid value: %s", key, val)
		}
	}
}

func TestNonBlockingCache_DataConsistency(t *testing.T) {
	c := cache.NewNonBlockingCache()
	var wg sync.WaitGroup

	// 1000 операций записи
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			c.Put(fmt.Sprintf("key%d", i), "value")
		}(i)
	}

	wg.Wait()

	// Проверяем все записанные значения
	for i := 0; i < 1000; i++ {
		if val, _ := c.Get(fmt.Sprintf("key%d", i)); val != "value" {
			t.Errorf("Key key%d has invalid value: %s", i, val)
		}
	}
}

func TestNonBlocking_ProgressGuarantee(t *testing.T) {
	c := cache.NewNonBlockingCache()
	var wg sync.WaitGroup
	done := make(chan struct{})

	// Горутина, которая может "застрять"
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; ; i++ {
			select {
			case <-done:
				return
			default:
				c.Put("key", "value")
			}
		}
	}()

	// Главная горутина должна иметь возможность читать
	for i := 0; i < 1000; i++ {
		if _, ok := c.Get("key"); !ok {
			t.Fatal("Read failed while write stuck")
		}
	}

	close(done)
	wg.Wait()
}

func TestNonBlocking_NoSpinning(t *testing.T) {
	c := cache.NewNonBlockingCache()
	start := time.Now()

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Put("key", "value")
		}()
	}

	wg.Wait()

	if time.Since(start) > 100*time.Millisecond {
		t.Errorf("Too much time spent: %v", time.Since(start))
	}
}

func TestNonBlocking_PauseTolerance(t *testing.T) {
	c := cache.NewNonBlockingCache()
	var wg sync.WaitGroup

	// Горутина с долгой операцией
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.Put("key", "old") // Имитация долгой записи
		time.Sleep(1 * time.Second)
	}()

	// Должны получить актуальные данные, несмотря на паузу
	time.Sleep(100 * time.Millisecond)
	c.Put("key", "new")

	val, _ := c.Get("key")
	if val != "new" {
		t.Errorf("Update failed: got %s, want 'new'", val)
	}

	wg.Wait()
}
