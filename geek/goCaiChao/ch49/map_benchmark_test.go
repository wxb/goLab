package ch49_test

import (
	"strconv"
	"sync"
	"testing"

	"github.com/wxb/goLab/geek/goCaiChao/ch49"
)

const (
	NumOfReader = 10
	NumOfWriter = 100
)

type Map interface {
	Set(key interface{}, val interface{})
	Get(key interface{}) (interface{}, bool)
	Del(key interface{})
}

func benchmarkMap(b *testing.B, hm Map) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		// write
		for j := 0; j < NumOfWriter; j++ {
			wg.Add(1)
			go func() {
				for n := 0; n < 100; n++ {
					hm.Set(strconv.Itoa(n), n*n)
					hm.Set(strconv.Itoa(n), n*n)
					hm.Del(strconv.Itoa(n))
				}
				wg.Done()
			}()
		}

		// read
		for j := 0; j < NumOfReader; j++ {
			wg.Add(1)

			go func() {
				for n := 0; n < 100; n++ {
					hm.Get(strconv.Itoa(n))
				}
				wg.Done()
			}()
		}

		wg.Wait()
	}
}

func BenchmarkSyncMap(b *testing.B) {
	b.Run("map with RWLock", func(b *testing.B) {
		benchmarkMap(b, ch49.NewRWLockMap())
	})

	b.Run("sync.map", func(b *testing.B) {
		benchmarkMap(b, ch49.NewSyncMap())
	})

	b.Run("concurrent map", func(b *testing.B) {
		benchmarkMap(b, ch49.NewConcurrentMap(199))
	})

	b.Run("concurrent-map", func(b *testing.B) {
		benchmarkMap(b, ch49.NewCMap())
	})
}
