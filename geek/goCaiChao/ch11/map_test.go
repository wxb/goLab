package ch11_test

import (
	"testing"
)

// map与工厂模式

// 函数在go里面是一等公民
// map的值可以是一个函数或者说方法
//

type f1 func(i int) int

func TestMapWithFunValue(t *testing.T) {
	m := map[int]f1{}
	m[1] = func(i int) int { return i }
	m[2] = func(ii int) int { return ii * ii }
	m[3] = func(iii int) int { return iii * iii * iii }

	t.Log(m[1](2), m[2](2), m[3](2))
}

// 实现Set(一组key的集合,key不能重复;所以，在set中，没有重复的key。)
// go内置没有实现set,可以利用 map[type]bool 来实现
// 1. 元素的唯一性
// 2. 基本操作：a.添加元素；b.删除元素；c.判断元素是否存在；d.返回元素个数

// type set struct {
// 	t reflect.Kind
// 	m map[reflect.Kind]bool
// }

// func NewSet(list []interface{}) {
// 	s := set{}
// 	for _, v := range list {
// 		t := reflect.TypeOf(v).Kind()

// 	}
// }

// func Add() {

// }

// func Del() {

// }

// func Num() {

// }

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] {
		t.Log(n, "is existing")
	} else {
		t.Log(n, "is not existing")
	}

	mySet[3] = true
	t.Log(mySet)

	delete(mySet, 1)
	if mySet[n] {
		t.Log(n, "is existing")
	} else {
		t.Log(n, "is not existing")
	}

	s := []int{1, 2, 3}
	delete(s, 1)
}
