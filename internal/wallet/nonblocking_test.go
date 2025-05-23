package wallet

import (
	"sync"
	"testing"
)

func TestNonBlockingWallet_Transfer(t *testing.T) {
	w1 := &NonBlockingWallet{}
	w2 := &NonBlockingWallet{}

	w1.Deposit(100)

	if !w1.Transfer(w2, 50) {
		t.Error("Transfer failed")
	}

	if w1.Balance() != 50 || w2.Balance() != 50 {
		t.Errorf("Balances incorrect: w1=%d, w2=%d", w1.Balance(), w2.Balance())
	}
}

func TestNonBlockingWallet_ConcurrentTransfers(t *testing.T) {
	w1 := &NonBlockingWallet{}
	w2 := &NonBlockingWallet{}
	w1.Deposit(1000)

	var wg sync.WaitGroup
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			w1.Transfer(w2, 2)
		}()
	}

	wg.Wait()

	total := w1.Balance() + w2.Balance()
	if total != 1000 {
		t.Errorf("Total balance mismatch: %d", total)
	}
}
