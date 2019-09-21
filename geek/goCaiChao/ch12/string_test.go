package ch12_test

import (
	"strconv"
	"strings"
	"testing"
)

// 字符串

// 1. string是数据类型，不是引用类型或指针类型
// 2. string是只读的不可变的byte slice，len可以得到它所包含的byte字节数
// 3. string的byte slice不仅可以存储可见字符，也可以存储不属于可见字符的byte
// 4. for-range遍历string时，k按照byte赋值，而k是按照字符赋值，这点在多字节字符时将会很关键

func TestMultiByteString(t *testing.T) {
	s := "I have a dream,我的梦想是世界和平！🤣"
	for k, v := range s {
		t.Logf("%d: %[2]q,%[2]d", k, v)
	}
}

func TestString(t *testing.T) {
	var s string
	t.Log(s)

	s = "hello"
	t.Log(len(s))

	// s[1] = "3" // string 是不可变的byte slice

	s = "\xE4\xB8\xA5" // 可以存储任何二级制数据
	t.Log(s, len(s))

	s = "中"
	t.Log(len(s))

	c := []rune(s)
	t.Log(len(c))
	t.Logf("中 Unicode: dec(%[1]v), hex(%[1]x)", c[0])
	t.Logf("中 UTF-8: %x", s)
}

func TestStringsFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for i, v := range parts {
		t.Log(i, v)
	}

	str := strings.Join(parts, "|")
	t.Log(str)
}

func TestStrConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str:" + s)

	num, err := strconv.Atoi("10")
	if err == nil {
		t.Log(10 + num)
	}
}
