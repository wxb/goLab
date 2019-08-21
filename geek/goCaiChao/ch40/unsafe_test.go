package ch40_test

import "testing"

func TestConve(t *testing.T) {

	i := int8(1)

	ii := int32(i)

	t.Logf("%T", ii)
}
