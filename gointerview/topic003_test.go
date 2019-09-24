package gointerview_test

import (
	"fmt"
	"testing"
)

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		fmt.Printf("%p\n", &stu)
		stu.Age = stu.Age + 10
		m[stu.Name] = &stu
	}
	fmt.Println(m)
	fmt.Printf("%v", m["li"])
}

func TestTopic003(t *testing.T) {
	pase_student()
}
