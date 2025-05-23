package wallet

import "sync"

type CoarseWallet struct {
	mu      sync.Mutex
	balance int
}

func (w *CoarseWallet) Deposit(amount int) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.balance += amount
}

func (w *CoarseWallet) Balance() int {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.balance
}
