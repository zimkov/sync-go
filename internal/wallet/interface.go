package wallet

type WalletInterface interface {
	Deposit(int)
	Balance() int
}

// Интерфейс для atomic-реализации
type AtomicWallet interface {
	WalletInterface
	AtomicBalance() int64
}
