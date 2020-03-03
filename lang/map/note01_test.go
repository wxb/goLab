package map_test

import (
	"fmt"
	"testing"
)

// 禁止对map元素取址的原因是map可能随着元素数量的增长而重新分配更大的内存空间，从而可能导致之前的地址无效
func TestMapRef(t *testing.T) {
	var age map[string]int
	fmt.Println(age == nil)

	// map上的大部分操作，包括查找、删除、len和range循环都可以安全工作在nil值的map上，它们的行为和一个空的map类似。但是向一个nil值的map存入元素将导致一个panic异常
	fmt.Println(age["name"], len(age))
	delete(age, "name")
	for k, v := range age {
		fmt.Println(k, v)
	}
	// panic: assignment to entry in nil map
	// age["name"] = 0

	ages := make(map[string]int)
	fmt.Println(ages == nil)
	fmt.Printf("ages %p\n", ages)

	ages["alice"] = 31
	ages["charlie"] = 34

	// cannot take address of ages["charlie"] (map index expression of type int)
	// _ = &ages["charlie"]

	sala := 19
	bob := 20
	newAges := make(map[string]*int)
	newAges["sala"] = &sala
	newAges["bob"] = &bob
	_ = newAges["bob"]
	fmt.Printf("sala %p\n", newAges["sala"])
}

func TestMapEqual(t *testing.T) {
	mapA := map[string]int{}
	mapB := map[string]int{}

	// slice一样，map之间也不能进行相等比较
	// map can only be compared to nil
	fmt.Println(mapA == mapB)

}
