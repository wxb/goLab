package main

import "fmt"

func main() {
	// 	var aa []string
	// 	fmt.Println(aa)

	// Loop:
	// 	for {
	// 		for i := 0; i < 10; i++ {
	// 			switch {
	// 			case len(aa) > 30:
	// 				break Loop
	// 			default:
	// 				aa = append(aa, "a")
	// 			}
	// 		}
	// 	}

	// 	fmt.Println("finish", aa)

	// i := new([]int)
	// *i = append(*i, 5)

	// fmt.Println(*i)

	var i []int
	fmt.Println(i == nil)
	fmt.Println(i)

	ii := []int{}
	fmt.Println(ii == nil)
	fmt.Println(ii)

	iii := new([]int)
	fmt.Println(*iii)
}
