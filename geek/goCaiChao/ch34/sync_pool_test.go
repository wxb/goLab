package ch34_test

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object")
			return 100
		},
	}

	v := (&pool).Get().(int)
	fmt.Println(v)

	pool.Put(3)
	// runtime.GC() // GC 函数会出发go垃圾回收，这里就是清理了临时对象缓存
	v1 := pool.Get().(int)
	fmt.Println(v1)

	v2 := pool.Get().(int)
	fmt.Println(v2)
}

func TestSyncPoolInMultiGoroutine(t *testing.T) {

	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(101)
	pool.Put(102)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i, pool.Get())
		}(i)
	}

	wg.Wait()
}
