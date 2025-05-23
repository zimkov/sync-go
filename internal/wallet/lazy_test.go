package wallet

import (
	"sync"
	"testing"
)

func TestLazyWallet_InitOnce(t *testing.T) {
	w := &LazyWallet{}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			w.Deposit(1)
		}()
	}

	wg.Wait()

	if w.Balance() != 100 {
		t.Errorf("Expected 100, got %d", w.Balance())
	}
}

func TestLazyWallet_Uninitialized(t *testing.T) {
	w := &LazyWallet{}

	if w.Balance() != 0 {
		t.Errorf("Expected 0 for uninitialized, got %d", w.Balance())
	}
}
