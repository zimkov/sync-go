package main

import (
	"fmt"
	"sync-go/internal/cache_test"
	"testing"
)

type testRunner interface {
	RunTest(*testing.T)
}

func runTestSuite(name string, tests []testRunner) {
	fmt.Printf("\n=== Тестирование %s ===\n", name)
	for i, testCase := range tests {
		t := &capturingT{}
		testCase.RunTest(t)

		if t.Failed() {
			fmt.Printf("[%s] Тест %d: НЕ ПРОЙДЕН\n", name, i+1)
		} else {
			fmt.Printf("[%s] Тест %d: ПРОЙДЕН\n", name, i+1)
		}
	}
}

type capturingT struct {
	failed bool
}

func (t *capturingT) Error(args ...interface{}) {
	t.failed = true
}

func (t *capturingT) Errorf(format string, args ...interface{}) {
	t.failed = true
}

func (t *capturingT) FailNow() {
	t.failed = true
}

func (t *capturingT) Failed() bool {
	return t.failed
}

func main() {
	cacheTests := []testRunner{
		&cache_test.CoarseTest{},
		&cache_test.FineTest{},
		&cache_test.OptimisticTest{},
		&cache_test.LazyTest{},
		&cache_test.NonBlockingTest{},
	}
	runTestSuite("CACHE PATTERNS", cacheTests)
}
