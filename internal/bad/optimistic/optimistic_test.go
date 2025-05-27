package optimistic

import (
	"sync"
	"testing"
)

// Тесты
func TestOptimisticLock(t *testing.T) {
	c := NewOptimisticCache()
	c.Set("key", "value1")

	// Первое обновление
	if !c.Update("key", "value2") {
		t.Error("First update failed")
	}

	// Параллельные обновления
	const goroutines = 10
	var successCount int
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			if c.Update("key", "new_value") {
				mu.Lock()
				successCount++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()

	if successCount != 1 {
		t.Errorf("Expected 1 success, got %d", successCount)
	}

	val, ver := c.Get("key")
	t.Logf("Result: %s (v%d)", val, ver)
}
