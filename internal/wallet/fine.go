package wallet

import "sync"

type FineWallet struct {
	mu      sync.RWMutex
	balance int
}

func (w *FineWallet) Deposit(amount int) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.balance += amount
}

func (w *FineWallet) Balance() int {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.balance
}
