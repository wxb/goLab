package ch30_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

var singletonInstance *Singleton
var once sync.Once

func SingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Singleton Object")
		singletonInstance = new(Singleton)
	})

	return singletonInstance
}

func TestSingletonObj(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			obj := SingletonObj()
			fmt.Printf("%o %v\n", obj, unsafe.Pointer(obj))
		}()
	}

	wg.Wait()
}

func TestUnsafe(t *testing.T) {
	type Human struct {
		sex  bool
		age  uint8
		min  int
		name string
	}

	h := Human{
		true,
		30,
		1,
		"hello",
	}
	i := unsafe.Sizeof(h)
	j := unsafe.Alignof(h.age)
	k := unsafe.Offsetof(h.name)
	fmt.Println(i, j, k)
	fmt.Printf("%p\n", &h)
	var p unsafe.Pointer
	p = unsafe.Pointer(&h)
	fmt.Println(p)
}
