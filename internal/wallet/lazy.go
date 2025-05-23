package wallet

import "sync"

type LazyWallet struct {
	once    sync.Once
	mu      sync.Mutex
	balance *int
}

func (w *LazyWallet) init() {
	w.balance = new(int)
}

func (w *LazyWallet) Deposit(amount int) {
	w.once.Do(w.init)

	w.mu.Lock()
	defer w.mu.Unlock()
	*w.balance += amount
}

func (w *LazyWallet) Balance() int {
	w.once.Do(w.init)

	w.mu.Lock()
	defer w.mu.Unlock()
	return *w.balance
}
