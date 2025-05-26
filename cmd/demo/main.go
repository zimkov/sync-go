package main

import (
	"fmt"
	"sync"
	"sync-go/internal/cache"
	"sync-go/internal/wallet"
	"time"
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

func benchmarkCache(c cache.Cache, name string) {
	start := time.Now()
	var wg sync.WaitGroup

	// 1000 операций записи
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			c.Put(fmt.Sprintf("key%d", i), "value")
		}(i)
	}

	// 1000 операций чтения
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			c.Get(fmt.Sprintf("key%d", i))
		}(i)
	}

	wg.Wait()
	fmt.Printf("[%-12s] Total time: %v\n", name, time.Since(start))
}

func main() {
	fmt.Printf("Тесты для примера с кошельком:\n")
	runTest("Unsafe", &wallet.UnsafeWallet{})
	runTest("Coarse", &wallet.CoarseWallet{})
	runTest("Fine", &wallet.FineWallet{})
	runTest("Optimistic", &wallet.OptimisticWallet{})
	runTest("Lazy", &wallet.LazyWallet{})
	runTest("NonBlocking", &wallet.NonBlockingWallet{})

	fmt.Printf("\n\nТесты для задания с кэшем:\n")
	// Пример с кэшем
	implementations := []struct {
		name     string
		instance cache.Cache
	}{
		{"Coarse", cache.NewCoarseCache()},
		{"Fine", cache.NewFineCache()},
		{"Optimistic", cache.NewOptimisticCache()},
		{"Lazy", cache.NewLazyCache()},
		{"NonBlocking", cache.NewNonBlockingCache()},
	}

	for _, impl := range implementations {
		benchmarkCache(impl.instance, impl.name)
		time.Sleep(2 * time.Second) // Сброс состояния
	}
}
