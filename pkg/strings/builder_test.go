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
