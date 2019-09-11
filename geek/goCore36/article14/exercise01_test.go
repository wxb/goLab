package article14_test


import (
	"fmt"
	"testing"
)

type Echo interface {
	Echo()
	String() string
}

type Man struct {
	Name string
}

func (m Man) String() string {
	return m.Name
}

func (m *Man) Echo() {
	fmt.Println("Name is:" + m.Name)
}

func TestExercise01(t *testing.T) {

	// var m1 *Man
	// var e1 Echo

	// e1 = m1

	// fmt.Println(m1)
	// if m1 == nil {
	// 	fmt.Println("111")
	// }

	// fmt.Println(e1)
	// if e1 == nil {
	// 	fmt.Println("222")
	// }

	// m1.Echo()
	// e1.Echo()

	// var m2 = &Man{}
	// // var m2 *Man
	// var e2 Echo
	// e2 = m2

	// e2.Echo()
	// fmt.Println(e2.String())

	var m3 *Man
	var e3 Echo

}
