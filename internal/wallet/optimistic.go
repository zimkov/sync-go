package wallet

import "sync/atomic"

type OptimisticWallet struct {
	balance atomic.Int64
}

func (w *OptimisticWallet) Deposit(amount int) {
	w.balance.Add(int64(amount))
}

func (w *OptimisticWallet) Balance() int {
	return int(w.balance.Load())
}

func (w *OptimisticWallet) AtomicBalance() int64 {
	return w.balance.Load()
}
