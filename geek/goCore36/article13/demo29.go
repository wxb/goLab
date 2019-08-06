package main

import "fmt"

type Tii struct {
	Name string
	Age  *int
	Cls  *[]string
}

func main() {
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
