package sct2_test

import (
	"fmt"
	"testing"
)

type Duck interface {
	Walk()
	Quack()
}

type Cat struct{}

// func (c *Cat) Walk() {

// }

// func (c *Cat) Quack() {

// }

// func (c Cat) Walk() {

// }

// func (c Cat) Quack() {

// }

func TestInitReceiver(t *testing.T) {
	// var d Duck = Cat{}
	d := Cat{}
	if interface{}(d) == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("non-nil")
	}
	// var dd Duck = &d
	// dd.Quack()
}
