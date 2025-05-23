package wallet

import (
	"sync"
	"testing"
)

func TestFineWallet_ConcurrentReads(t *testing.T) {
	w := &FineWallet{balance: 100}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = w.Balance()
		}()
	}

	wg.Wait()
}

func TestFineWallet_DepositAndRead(t *testing.T) {
	w := &FineWallet{}
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			w.Deposit(1)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			_ = w.Balance()
		}
	}()

	wg.Wait()

	if w.Balance() != 1000 {
		t.Errorf("Expected 1000, got %d", w.Balance())
	}
}
