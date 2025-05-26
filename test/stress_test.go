package test

import (
	"sync-go/internal/wallet"
	"testing"
)

func BenchmarkUnsafe(b *testing.B) {
	w := &wallet.UnsafeWallet{}
	benchmarkWallet(b, w)
}

func BenchmarkCoarse(b *testing.B) {
	w := &wallet.CoarseWallet{}
	benchmarkWallet(b, w)
}

func BenchmarkFine(b *testing.B) {
	w := &wallet.FineWallet{}
	benchmarkWallet(b, w)
}

func BenchmarkLazy(b *testing.B) {
	w := &wallet.LazyWallet{}
	benchmarkWallet(b, w)
}

func BenchmarkNonBlocking(b *testing.B) {
	w := &wallet.NonBlockingWallet{}
	benchmarkWallet(b, w)
}

func benchmarkWallet(b *testing.B, w wallet.WalletInterface) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			w.Deposit(1)
			_ = w.Balance()
		}
	})
}
