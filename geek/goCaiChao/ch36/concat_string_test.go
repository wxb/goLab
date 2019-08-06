package ch36

import (
	"bytes"
	"testing"
)

func BenchmarkConcatStringByAdd(b *testing.B) {
	elems := [...]string{"a", "b", "c", "d", "e"}
	ret := ""
	b.N = 50000
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range elems {
			ret += v
		}
	}
	b.StopTimer()
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {

	elems := [...]string{"a", "b", "c", "d", "e"}
	buf := bytes.Buffer{}
	b.N = 50000
	// var buf bytes.Buffer
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range elems {
			buf.WriteString(v)
		}
	}
	b.StopTimer()
}
