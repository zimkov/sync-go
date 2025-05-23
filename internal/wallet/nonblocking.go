package wallet

import "sync/atomic"

type NonBlockingWallet struct {
	balance atomic.Int64
}

func (w *NonBlockingWallet) Deposit(amount int) {
	w.balance.Add(int64(amount))
}

func (w *NonBlockingWallet) Balance() int {
	return int(w.balance.Load())
}

func (w *NonBlockingWallet) Transfer(to *NonBlockingWallet, amount int) bool {
	for {
		currentFrom := w.balance.Load()
		if currentFrom < int64(amount) {
			return false
		}

		currentTo := to.balance.Load()

		if w.balance.CompareAndSwap(currentFrom, currentFrom-int64(amount)) &&
			to.balance.CompareAndSwap(currentTo, currentTo+int64(amount)) {
			return true
		}
	}
}
