package article12_test

import (
	"fmt"
	"testing"
)

type Stu struct {
	Name string
	Age  int
	Cls  map[string]string
}

type FF func(i int) error

func TestDemo00(t *testing.T) {
	var a FF
	var b interface{}
	var c []int
	var d chan<- int
	fmt.Println(a, b == nil, c == nil, d)

	s := Stu{Name: "wangxiaob", Cls: map[string]string{"s": "s"}}
	ss := s
	fmt.Println(s, ss, &s == &ss, &s.Cls == &ss.Cls)

	ss.Age = 29

	ss.Cls["s"] = "ss"
	fmt.Println(s, ss, &s == &ss, &s.Cls == &ss.Cls)

	ss.Cls = map[string]string{"1": "2"}
	fmt.Println(s, ss, &s == &ss, &s.Cls == &ss.Cls)

	c = append(c, 1)
	fmt.Println(c)
	// var p FF
	p := uu
	p(123)
}

func uu(j int) error {
	fmt.Println(j)
	return nil
}
