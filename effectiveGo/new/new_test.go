package new_test

import (
	"fmt"
	"testing"
)

func TestNewAllocat(t *testing.T) {

	var i1 int
	i2 := new(int)

	f1 := new(float32)

	fmt.Printf("%d %d %f", i1, *i2, *f1)
	fmt.Println("")
	fmt.Printf("%p %p", &i1, i2)
}

func TestNewSlice(t *testing.T) {
	p1 := new([]int)
	p2 := []int{}

	fmt.Printf("%p %p\n", *p1, p2)
	*p1 = append(*p1, 2)
	fmt.Printf("%p\n", *p1)

	fmt.Println(*p1 == nil, p2 == nil, len(*p1), cap(*p1), len(p2), cap(p2))

}

// func TestMakeSlice(t *testing.T) {
// 	p1 := make([]int, 10)

// }
