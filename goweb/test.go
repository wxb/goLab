package main

import "fmt"
import "reflect"
import "strconv"

func main() {
	/*
		s := []string{"w", "a", "n", "g"}
		for _, v := range s {
			fmt.Println(v)
			for i := 0; i < 10; i++ {
				if i == 5 {
					continue LABEL1
				}
				if i == 8 {
					break LABEL1
				}
				fmt.Println(i)
			}
		}

		fmt.Println("out")
	*/

	/* 跳转语句
	LABEL1:

		for i := 0; i < 10; i++ {
			for  {
				fmt.Println(i)
				continue LABEL1
			}
		}

	*/

	/*
		  数组
		var arr_01 [2]int
		var arr_02 [2]string
		var arr_03 [3]bool
		arr_04 := [2]int{1, 2}
		var arr_05 = [20]int{19: 3}
		arr_06 := [...]string{"w", "a", "n", "g"}
		arr_07 := [...]string{4: "w", 6: "a", 2: "n", 9: "g"}

		fmt.Println(arr_01, arr_02, arr_03, arr_04, arr_05, arr_06, arr_07, arr_07[8], arr_07[2])

		x, y := 1, 2
		a := [...]*int{&x, &y}
		fmt.Println(a, *a[0])

		m := 10
		n := &m
		fmt.Println(n, *n)

		p := new([5]int)
		q := &arr_07
		fmt.Println(p, q)

		arr_08 := [...][2][3]int{
			{{1: 1}, {1: 1}},
			{{2: 2}, {2: 2}}}
		fmt.Println(arr_08)

		// 冒泡排序
		arr_09 := [...]int{6, 2, 8, 3, 1, 0, 56, 24, 99, 51}
		fmt.Println(arr_09)
		nums := len(arr_09)
		for i := 0; i < nums; i++ {
			for j := i + 1; j < nums; j++ {
				if arr_09[i] > arr_09[j] {
					temp := arr_09[i]
					arr_09[i] = arr_09[j]
					arr_09[j] = temp
				}
			}
		}
		fmt.Println(arr_09)
	*/

	/*
		 // slice

		a := [10]int{9: 3}
		b := a[5:]
		fmt.Println(a, b)

		s := make([]int, 3, 10)
		fmt.Println(s)

		arr01 := [10]int{1,2}
		s1 := arr01[:]
		fmt.Printf("%v %v \n", &arr01[0], &s1[0])
		fmt.Println(len(arr01), cap(arr01))
		fmt.Println(len(s1), cap(s1))
	*/

	/*
		// map
		var m map[int]string
		m = map[int]string{1: "wang", 2: "xiao", 3: "bo"}
		fmt.Println(m)

		mm := map[string]string{`name`: "wang xiaobo"}
		fmt.Println(mm)

		mmm := make(map[int]map[int]string)
		// mmm[1][1] = "ok" 内层map未初始化 会报错
		if v, ok := mmm[1][1]; ok {
			fmt.Println(v, ok)
		} else {
			mmm[1] = make(map[int]string)
			mmm[1][1] = "ok"
			fmt.Println(v, ok, mmm[1][1])
		}
	*/

	/*
		var fs = [4]func(){}
		for i:=0; i<4; i++{
			defer fmt.Println("defer i=", i)
			defer func() {fmt.Println("defer_closure i=", i)}()
			fs[i] = func() {
				fmt.Println("closure i=", i)
			}
		}

		for _,f := range fs{
			f()
		}
	*/

	/*
		type person struct {
			Name string
			Age  int
		}

		a := person{
			Name: "王小波",
		}
		a.Age = 28
		fmt.Println(a)

		b := struct {
			Name string
			Age  int
		}{
			Name: "wangxiao",
			Age:  29,
		}
		fmt.Println(b)
	*/

	c := make(chan bool)

	go func() {
		name := "wang xiaobo bo"
		for _, v := range name {
			tt := strconv.Itoa(666)
			fmt.Println(tt, reflect.TypeOf(v))
			//fmt.Println(strconv.Atoi(reflect.TypeOf(tt)))
		}
		c <- true
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
