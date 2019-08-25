package ch48_test

import (
	"fmt"
	"testing"
)

func TestArr(t *testing.T) {

	var a [5]int
	b := [5]int{}
	c := [...]int{1, 2, 3, 4, 5}
	// d := make([5]int)  不可以

	fmt.Println(a, b, c)
}
