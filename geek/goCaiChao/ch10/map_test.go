package ch10_test

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2}

	m2 := map[string]int{}
	m2["three"] = 3

	m3 := make(map[string]int, 10) // 10 is cap, not len; 这里如果我们可以确定容量 有利于提高性能
	m3["four"] = 4

	var m4 map[string]int
	// m4["five"] = 5       // panic: assignment to entry in nil map
	m4 = map[string]int{"five": 5}

	t.Log(m1, m2, m3, m4)
	t.Log(len(m1), len(m2), len(m3), len(m4))
	// t.Log(cap(m1), cap(m2), cap(m3), cap(m4))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}

	// 在指定key不存在时，go并不会异常或者错误，而会返回该map值的类型零值
	t.Log(m1[1])

	m1[2] = 0
	t.Log(m1[2])

	// 利用这种形式来确认，map中是否存在某个key值
	if v, ok := m1[3]; ok {
		t.Log("key 3's value is ", v)
	} else {
		t.Log("key 3 is not exist")
	}
}

func TestTravelMap(t *testing.T) {
	m := map[int]string{1: "one", 2: "two", 3: "three"}
	// 利用for-range可以遍历出map的键值对
	for k, v := range m {
		t.Log(k, v)
	}
	// 输出键值对顺序，map底层使用的是hash存储，不是有序的
}
