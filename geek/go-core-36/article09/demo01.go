package main

import "fmt"

func main() {
	// var badMap2 = map[interface{}]int{
	// 	"1":      1,
	// 	[]int{2}: 2, // panic: runtime error: hash of unhashable type []int 键的类型错误，函数|映射|切片 不支持判等操作
	// 	3:        3,
	// }
	// fmt.Println(badMap2)

	var sliceParam []int
	//sliceParam[0] = 1 // panic: runtime error: index out of range
	fmt.Println(sliceParam, sliceParam == nil)

	var mapParam map[int]string
	//mapParam[1] = "ss" // panic: assignment to entry in nil map
	fmt.Println(mapParam, mapParam == nil)

	p01 := make([]int, 3, 6)
	fmt.Println(p01, p01 == nil)
	p02 := new([]int)
	fmt.Println(p02, *p02 == nil)

	p03 := make(map[int]string, 3)
	fmt.Println(p03, p03 == nil)

	p04 := new(map[int]string)
	//(*p04)[1] = "ss" // panic: assignment to entry in nil map
	fmt.Println(p04, *p04 == nil)

	p05 := new(int)
	//*p05 = 1
	fmt.Println(p05, *p05)

	type Persion struct {
		Name string
	}

	p06 := new(Persion)
	p06.Name = "wang xiaobo"
	(*p06).Name = "王晓勃"
	fmt.Println(*p06)
}

/*
在值为nil的字典上执行读操作会成功吗，那写操作呢？

对于slice、map、channel类型的某个nil变量，除了添加 键值对外，其他的任何操作都不会引起panic

极客时间版权所有: https://time.geekbang.org/column/article/14123
*/

/*

make用于内建类型（map、slice 和channel）的内存分配。new用于各种类型的内存分配。

内建函数new本质上说跟其它语言中的同名函数功能一样：new(T)分配了零值填充的T类型的内存空间，并且返回其地址，即一个*T类型的值。
用Go的术语说，它返回了一个指针，指向新分配的类型T的零值。有一点非常重要：new返回指针。


内建函数make(T, args)与new(T)有着不同的功能，make只能创建slice、map和channel，并且返回一个有初始值(非零)的T类型，而不是*T。

本质来讲，导致这三个类型有所不同的原因是指向数据结构的引用在使用前必须被初始化。
例如，一个slice，是一个包含指向数据（内部array）的指针、长度和容量的三项描述符；在这些项目被初始化之前，slice为nil。
对于slice、map和channel来说，make初始化了内部的数据结构，填充适当的值。make返回初始化后的（非零）值。

make 是 引用类型 初始化的方法。

*/
