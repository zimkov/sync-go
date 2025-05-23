package main

import (
	"fmt"
	"sync"
	"wallet/internal/wallet"
)

func runTest(walletType string, w wallet.WalletInterface) {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			w.Deposit(1)
		}()
	}

	wg.Wait()
	fmt.Printf("[%-12s] Balance: %d\n", walletType, w.Balance())
}

func main() {
	runTest("Unsafe", &wallet.UnsafeWallet{})
	runTest("Coarse", &wallet.CoarseWallet{})
	runTest("Fine", &wallet.FineWallet{})
	runTest("Optimistic", &wallet.OptimisticWallet{})
	runTest("Lazy", &wallet.LazyWallet{})

	// Пример с неблокирующим кошельком
	nbWallet := &wallet.NonBlockingWallet{}
	runTest("NonBlocking", nbWallet)
	fmt.Printf("[%-12s] Atomic Balance: %d\n", "NonBlocking", nbWallet.Balance())
}
