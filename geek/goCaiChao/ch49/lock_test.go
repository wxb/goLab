package ch49_test

import (
	"fmt"
	"sync"
	"testing"
)

const NUM_OF_READER int = 40
const READ_TIMES = 100000

var cache map[string]string

func init() {
	cache = make(map[string]string)

	cache["a"] = "aa"
	cache["b"] = "bb"
}

func lockFreeAccess() {
	var wg sync.WaitGroup
	wg.Add(NUM_OF_READER)
	for i := 0; i < NUM_OF_READER; i++ {
		go func() {
			for j := 0; j < READ_TIMES; j++ {
				_, has := cache["a"]
				if !has {
					fmt.Println("Nothing")
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func lockAccess() {
	var wg sync.WaitGroup
	var mu sync.RWMutex

	wg.Add(NUM_OF_READER)
	for i := 0; i < NUM_OF_READER; i++ {
		go func() {
			for j := 0; j < READ_TIMES; j++ {
				mu.RLock()
				_, has := cache["b"]
				mu.RUnlock()
				if !has {
					fmt.Println("Nothing")
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkLockFree(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lockFreeAccess()
	}
}

func BenchmarkLock(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lockAccess()
	}
}

// Output:
// $ go test -v -bench=. -run=^$ -benchmem
// goos: darwin
// goarch: amd64
// pkg: github.com/wxb/goLab/geek/goCaiChao/ch49
// BenchmarkLockFree-4          100          11193404 ns/op             650 B/op          2 allocs/op
// BenchmarkLock-4               10         167910457 ns/op              57 B/op          2 allocs/op
// PASS
// ok      github.com/wxb/goLab/geek/goCaiChao/ch49        2.984s
