package strings_test

import (
	"strings"
	"testing"
)

func TestBuilder(t *testing.T) {
	var builder strings.Builder
	builder.Grow(1)
	// builder1 := builder
	// builder1.Grow(1)
	t.Logf("%p", builder)

	builder.Reset()
	t.Logf("%p", builder)
	builder2 := builder
	builder2.Grow(1)
	t.Logf("%p", builder2)

}

func TestBuilderLen(t *testing.T) {
	var b strings.Builder
	b.WriteString("ABC的小写为abc")
	t.Log(b.Len() == len(b.String()))
}

// The len built-in function returns the length of v, according to its type:
// 	Array: the number of elements in v.
// 	Pointer to array: the number of elements in *v (even if v is nil).
// 	Slice, or map: the number of elements in v; if v is nil, len(v) is zero.
// 	String: the number of bytes in v.
// 	Channel: the number of elements queued (unread) in the channel buffer;
// 			if v is nil, len(v) is zero.
func TestLen(t *testing.T) {
	s := "ABC的小写为abc"
	// 对于字符串，len函数统计的是底层字节切片的长度
	t.Log(len(s))
	println(len(s))
}
