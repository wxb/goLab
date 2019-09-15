package ch09_test

import "testing"

func TestSliceGrowing(t *testing.T) {
	s := []int{}

	// slice的增长在1024内时是成倍增长
	for i := 0; i < 1024; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}

	// 超过1024以后是25%的增长
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}

	// 关于slice增长的策略，可以参看go源码 src/runtime/slice.go 文件 76行growslice函数：
	// func growslice(et *_type, old slice, cap int) slice
}

func TestSliceCompare(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}

	// invalid operation: s1 == s2 (slice can only be compared to nil);
	// so slice是除了nil都不可比较的
	if s1 == s2 {
		t.Log(("s1==s2"))
	}
}
