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
		fmt.Println(&stu)
		m[stu.Name] = &stu
	}
	fmt.Println(*m["zhou"], *(m["li"]))
}

func TestTopic003(t *testing.T) {
	pase_student()
}
