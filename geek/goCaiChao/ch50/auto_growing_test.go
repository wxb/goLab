package ch50_test

import (
	"testing"
)

const (
	numOfElems = 100000
	times      = 1000
)

func TestAutoGrow(t *testing.T) {
	for i := 0; i < times; i++ {
		s := []int{}
		for j := 0; j < numOfElems; j++ {
			s = append(s, j)
		}
	}
}

func TestProperInit(t *testing.T) {
	for i := 0; i < times; i++ {
		s := make([]int, 0, 100000)
		for j := 0; j < numOfElems; j++ {
			s = append(s, j)
		}
	}
}

func TestOverSizeInit(t *testing.T) {
	for i := 0; i < times; i++ {
		s := make([]int, 0, numOfElems*8)
		for j := 0; j < numOfElems; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkAutoGrow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []int{}
		for j := 0; j < numOfElems; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkProperInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, 100000)
		for j := 0; j < numOfElems; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkOverSizeInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, numOfElems*8)
		for j := 0; j < numOfElems; j++ {
			s = append(s, j)
		}
	}
}

// $ go test -v -bench=. -run=^$ -benchmem
// goos: darwin
// goarch: amd64
// pkg: github.com/wxb/goLab/geek/goCaiChao/ch50
// BenchmarkAutoGrow-4                         2000           1294646 ns/op         4654339 B/op         30 allocs/op
// BenchmarkProperInit-4                       5000            271827 ns/op          802818 B/op          1 allocs/op
// BenchmarkOverSizeInit-4                     1000           1338662 ns/op         6406144 B/op          1 allocs/op

// 可见切片可变的长度在底层实现：容量不够-> 新申请内存 -> 复制旧值 -> 存储新值  这个操作在性能上不够高
// 同时提早申请过大的容量，也会造成存储时操作耗时
