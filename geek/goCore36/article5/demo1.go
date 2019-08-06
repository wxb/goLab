package main

import "fmt"

var block string = "package"

func main() {
	fmt.Println(block, &block)
	block := "function"
	fmt.Println(block, &block)
	block, abc := "function-2", "abc"
	fmt.Println(block, &block, abc, &abc)

	{
		block = "function-2"
		fmt.Println(block, &block)
		block := true
		fmt.Println(block, &block)
		block = false
		fmt.Println(block, &block)
	}

	fmt.Println(block, &block)
}
