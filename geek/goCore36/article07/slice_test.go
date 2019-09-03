package article07_test

import (
	"fmt"
	"testing"
)

func f1(ss []int) {
	ss[0] = 1
}

func f2(aa [5]int) {
	aa[0] = 1
}

func TestSliceAndArray(t *testing.T) {

	s := make([]int, 10)
	t.Log(s, len(s), cap(s))
	f1(s)
	t.Log(s, len(s), cap(s))

	a := [...]int{0, 0, 0, 0, 0}
	t.Log(a, len(a), cap(a))
	f2(a)
	t.Log(a, len(a), cap(a))

	// Output:
	// slice_test.go:16: [0 0 0 0 0 0 0 0 0 0] 10 10
	// slice_test.go:18: [1 0 0 0 0 0 0 0 0 0] 10 10
	// slice_test.go:21: [0 0 0 0 0] 5 5
	// slice_test.go:23: [0 0 0 0 0] 5 5
	// 如果传递的值是引用类型：slice、interface、func、map、chan；就是“传引用”
	// 如果传递的值是值类型，就是“传值”
}

func TestDemo15(t *testing.T) {
	// 示例 1。
	s1 := make([]int, 5)
	fmt.Printf("The length of s1: %d\n", len(s1))
	fmt.Printf("The capacity of s1: %d\n", cap(s1))
	fmt.Printf("The value of s1: %d\n", s1)
	s2 := make([]int, 5, 8)
	fmt.Printf("The length of s2: %d\n", len(s2))
	fmt.Printf("The capacity of s2: %d\n", cap(s2))
	fmt.Printf("The value of s2: %d\n", s2)
}
