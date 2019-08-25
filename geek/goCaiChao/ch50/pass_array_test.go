package ch50_test

import "testing"

const NumOfElems = 1000

type Content struct {
	Detail [10000]int
}

func withValue(arr [NumOfElems]Content) int {
	return 0
}

func withReference(arr *[NumOfElems]Content) int {
	return 0
}

func TestFn(t *testing.T) {
	var arr [NumOfElems]Content
	// t.Log(arr)

	withValue(arr)
	withReference(&arr)
}

func BenchmarkPassingArrayWithValue(b *testing.B) {

	var arr [NumOfElems]Content

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withValue(arr)
	}
	b.StopTimer()
}

func BenchmarkPassingArrayWithRef(b *testing.B) {

	var arr [NumOfElems]Content

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withReference(&arr)
	}
	b.StopTimer()
}

// $ go test -v -bench=. -run=^$ -benchmem
// Output:
// BenchmarkPassingArrayWithValue-4             100          16967367 ns/op        80003097 B/op          1 allocs/op
// BenchmarkPassingArrayWithRef-4          2000000000               0.32 ns/op            0 B/op          0 allocs/op
