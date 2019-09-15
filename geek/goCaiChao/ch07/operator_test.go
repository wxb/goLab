package ch07_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}

	t.Log(a == b)
	// t.Log(a == c) 长度不同不可比较
	t.Log(c)
	t.Log(a == d)
}

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestRWE(t *testing.T) {
	a := 7 //

	// &^ 按位清零 运算符: 如果运算符(&^)右边对应位是1则左边变量对应该位结果为零，如果右边对应位是0则左边变量对应该位结果为左边该位值
	a = a &^ Readable
	a = a &^ Writable

	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
