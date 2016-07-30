package main

import "fmt"

func main() {
	//var number = []int{1, 2, 3, 4, 5}
	//var slice = number[1:4]
	//fmt.Printf("%d - %d - %v ", number[4], slice[2], 4==cap(slice))

	var numbers3 = [6]int{1, 2, 3, 4, 5, 6}
	//slice3 := numbers3[2:len(numbers3)] //[2,5)
	slice3 := numbers3[2:5:5] //[2,5)
	length := (3)
	capacity := (3)
	fmt.Printf("%v, %v\n", (length == len(slice3)), (capacity == cap(slice3)))
}
