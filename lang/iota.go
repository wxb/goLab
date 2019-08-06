package main

import (
	"fmt"
	"reflect"
)

const (
	a = iota
	b
	c
	d
	_
	e
)

const f = iota

type MyInt64 int64

const (
	x MyInt64 = iota
	y
	_
	z
)

type Allergen int

const (
	IgEggs         Allergen = 1 << iota // 1 << 0 which is 00000001
	IgChocolate                         // 1 << 1 which is 00000010
	IgNuts                              // 1 << 2 which is 00000100
	IgStrawberries                      // 1 << 3 which is 00001000
	IgShellfish                         // 1 << 4 which is 00010000
)

type ByteSize float64

const (
	_           = iota             // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
	MB                             // 1 << (10*2)
	GB                             // 1 << (10*3)
	TB                             // 1 << (10*4)
	PB                             // 1 << (10*5)
	EB                             // 1 << (10*6)
	ZB                             // 1 << (10*7)
	YB                             // 1 << (10*8)
)

const (
	Apple, Banana = iota + 1, iota + 2
	Cherimoya, Durian
	Elderberry, Fig
)

const (
	i = iota
	j = 3.14
	k = iota
	l
)

func main() {

	fmt.Println(a, e, f)
	t := reflect.TypeOf(a)
	fmt.Println(t)

	fmt.Println(x, y, z)
	t = reflect.TypeOf(x)
	fmt.Println(t)

	fmt.Println(IgShellfish)
	fmt.Println(IgEggs | IgChocolate | IgShellfish)

	fmt.Println(KB, MB, GB, TB, PB)

	fmt.Println(Apple, Banana, Cherimoya, Durian, Elderberry, Fig)

	fmt.Println(i, j, k, l)
}
