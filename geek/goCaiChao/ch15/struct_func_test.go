package ch15_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type user struct {
	name string
	age  int
}

// 使用值接受者声明方法，调用时会用这个值的一个副本来执行
func (u user) Name() string {
	fmt.Printf("Name address:%p %x\n", &u, unsafe.Pointer(&u.name))
	return u.name
}

// 使用指针接受者声明方法时，这个方法会共享调用方法时接受者所指向的值
func (u *user) Age() int {
	fmt.Printf("Age address:%p %x %x\n", u, unsafe.Pointer(&u.name), unsafe.Pointer(&u.age))
	return u.age
}

func TestUserFunc(t *testing.T) {
	u1 := user{}
	fmt.Printf("u1 address: %p %x %x\n", &u1, unsafe.Pointer(&u1.name), unsafe.Pointer(&u1.age))
	u1.Name()
	u1.Age()
	(&u1).Age()

	u2 := &user{}
	fmt.Printf("u2 address: %p %x %x\n", u2, unsafe.Pointer(&(*u2).name), unsafe.Pointer(&u2.age))
	u2.Name()
	(*u2).Name()
	u2.Age()
}

// 通过这节课纠正了自己对值接受者和指针接受者一个误区：

// 以前：
// 两者区别：1.值接受者声明的方法，调用时使用这个值的副本，指针接受者声明的方法，调用时共享这个值；2. 指针值可以调用值接受者声明方法，值不能调用指针接受者声明方法
// 现在：
// 两者区别：值接受者声明的方法，调用时使用这个值的副本，指针接受者声明的方法，调用时共享这个值

// 仔细验证发现不知道自己以前是在那里学到的第二点错误的知识，哎
