package ch13_test

import (
	"fmt"
	"testing"
	"time"
)

// 函数时一等公民
// 1. 可以有多个返回值
// 2. 所有参数都是值传递
// 3. 函数可以作为变量的值
// 4. 函数可以作为参数和返回值

func TestType(t *testing.T) {
	arr1 := [...][]int{
		[]int{1, 2, 3},
		[]int{4, 5},
	}

	arr2 := arr1
	t.Logf("%p, %p", &arr1, &arr2)     // arr2 内存地址 不同于 arr1
	t.Logf("%p, %p", arr1[0], arr2[0]) // 但是其中的切片值的指针属性指向同一个内存

	arr2[0][1] = 1
	t.Log(arr1, arr2) // 所以 arr2 的修改会影响arr1

	f1 := func(s []int) {
		s[0] = 0
		t.Logf("%p", &s)
	}
	s := []int{1, 2, 3, 4}
	f1(s)
	t.Logf("%p %v", &s, s)
}

type fn func(op int) int

func spentTime(f fn) fn {
	return func(op int) int {
		start := time.Now()
		res := f(op)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return res
	}
}

func doSlow(op int) int {
	time.Sleep(2)
	return op
}

func TestFn(t *testing.T) {
	// 函数作为参数，同时函数作为返回值
	fn := spentTime(doSlow)
	t.Log(fn(10))
}
