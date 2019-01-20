package main

import (
	"fmt"
	"reflect"
)

type Pet interface {
	Name() string
	Category() string
}

type Dog struct {
	name string // 名字。
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	// 示例1。
	var dog1 *Dog
	fmt.Println("The first dog is nil.", dog1)
	fmt.Printf("The first dog type is %T \n", dog1)
	dog2 := dog1
	fmt.Println("The second dog is nil.", dog2)
	var pet Pet
	fmt.Printf("%T \n", pet)
	pet = nil
	fmt.Printf("%T \n", pet)
	pet = dog1
	fmt.Printf("%T \n", pet)
	if pet == nil {
		fmt.Println("The pet is nil.", pet)
	} else {
		fmt.Println("The pet is not nil.", pet)
	}

	fmt.Printf("The type of pet is %T.\n", pet)
	fmt.Printf("The type of pet is %s.\n", reflect.TypeOf(pet).String())
	fmt.Printf("The type of second dog is %T.\n", dog2)
	fmt.Println()

	// 示例2。
	wrap := func(dog *Dog) Pet {
		if dog == nil {
			return nil
		}
		return dog
	}
	pet = wrap(dog2)
	if pet == nil {
		fmt.Println("The pet is nil.")
	} else {
		fmt.Println("The pet is not nil.")
	}

	// 示例3
	var x *[]int
	fmt.Printf("x is %v type is %T \n", x, x)
	type Pett interface{}
	var pett Pett
	pett = x
	fmt.Println("pett:", pett, pett == nil)
	fmt.Printf("%v %T %T", pett, pett, nil)
}
