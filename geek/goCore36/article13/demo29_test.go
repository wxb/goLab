package article13_test

import (
	"fmt"
	"testing"
)

type Tii struct {
	Name string
	Age  *int
	Cls  *[]string
}

func TestDemo29(t *testing.T) {
	var aa *bool

	a := new(int)
	fmt.Println(a, aa)

	bb := 1
	var b *int
	b = &bb

	fmt.Println(b, *b)

	c := Tii{
		Name: "tii",
		Age:  b,
		Cls:  &[]string{"111"},
	}
	fmt.Println(c, c, *c.Age, *(c.Age))

	var d *[]int
	d = &[]int{1, 2}
	fmt.Println(d)
	var f *Tii
	fmt.Println(f)
}
