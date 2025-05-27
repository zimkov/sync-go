package lazy

import (
	"sync"
	"testing"
)

func TestLazyInit(t *testing.T) {
	c := &LazyCache{}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			c.Set("key", 1)
			wg.Done()
		}()
	}
	wg.Wait()

	if c.data == nil {
		t.Error("Cache not initialized")
	}
}
