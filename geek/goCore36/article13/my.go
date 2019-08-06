package main

import "fmt"

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

func main() {

	var myI MyInt
	fmt.Printf("%T\n", myI)
	myI.String()

	MyStruct{}.String()

	var aa AliasInt
	fmt.Printf("%T\n", aa)
}
