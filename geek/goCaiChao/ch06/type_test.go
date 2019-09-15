package ch06_test

import (
	"math"
	"testing"
)

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64
	// go语言不支持隐式类型转换
	//  b=1
	b = int64(a)

	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestMathMax(t *testing.T) {
	// 对于一些常用的数值类型的最大最小值在math这个包中都有预定义的常量
	t.Log(math.MaxUint8, math.MaxInt64, math.MaxFloat64, math.MinInt64)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	t.Logf("%T, %T", a, aPtr)

	// go语言不支持直接操作内存，so这里的变量地址加一不能完成
	// aPtr = aPtr + 1
}

func TestString(t *testing.T) {
	str := "在对一个字符串应用range语句时，请非常注意k和v所对应的值"
	for k, v := range str {
		t.Logf("%v, %c, %q, %s", k, v, v, string(v))
	}

}
