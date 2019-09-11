package article13_test

import (
	"fmt"
	"testing"
)

type Person struct {
	Name    string
	Age     int
	Phone   string
	Address string
}

type Company struct {
	No string
}

type Engineer struct {
	Person
	*Company
	Skill string
}

type PersonName interface {
	EchoName()
	SetName(n string)
}

func (p Person) EchoName() {
	fmt.Println(p.Name)
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func TestMy01(t *testing.T) {

	p := Person{}
	// 值类型 未实现PersonName接口
	if _, ok := interface{}(p).(PersonName); ok {
		fmt.Println("p impletement PersonName")
	} else {
		fmt.Println("p is not impletement PersonName")
	}

	// 引用类型 实现了PersonName接口
	if _, ok := interface{}(&p).(PersonName); ok {
		fmt.Println("&p impletement PersonName")
	} else {
		fmt.Println("&p is not impletement PersonName")
	}

	staff := Engineer{}
	// 值类型 未实现PersonName接口
	if _, ok := interface{}(staff).(PersonName); ok {
		fmt.Println("staff impletement PersonName")
	} else {
		fmt.Println("staff is not impletement PersonName")
	}

	// 引用类型 实现了PersonName接口
	if _, ok := interface{}(&staff).(PersonName); ok {
		fmt.Println("&staff impletement PersonName")
	} else {
		fmt.Println("&staff is not impletement PersonName")
	}

	staff.Name = "wangxb"
	fmt.Println(staff.Name)
	staff.Company = &Company{"123"}
	staff.Company.No = "123123"
	fmt.Println(staff.No)
}
