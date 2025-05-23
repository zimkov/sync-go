package wallet

type UnsafeWallet struct {
	balance int
}

func (w *UnsafeWallet) Deposit(amount int) {
	w.balance += amount
}

func (w *UnsafeWallet) Balance() int {
	return w.balance
}
