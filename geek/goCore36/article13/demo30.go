package main

import "fmt"

type Cat struct {
	name           string
	scientificName string
	category       string
}

func (cat *Cat) SetName(name string) {
	cat.name = name
}

func main() {

	cat := Cat{}
	cat.SetName("miaomiao")
	fmt.Println(cat, &cat.name)

	(&cat).SetName("toto")
	fmt.Println(cat, &cat.name)

}
