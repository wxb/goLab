package article13_test

import (
	"fmt"
	"testing"
)

type MyInt int
type AliasInt = int

type MyStruct struct {
	Name string
}

func (my MyInt) String() {
	fmt.Println(my)
}

func (my MyStruct) String() {
	fmt.Println(my.Name)
}

func TestMy(t *testing.T) {

	var myI MyInt
	fmt.Printf("%T\n", myI)
	myI.String()

	MyStruct{}.String()

	var aa AliasInt
	fmt.Printf("%T\n", aa)
}
