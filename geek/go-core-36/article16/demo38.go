package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("--->", i, &i)
		go func() {
			fmt.Println("===>", i, &i)
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("***>", i)
		}()
	}
}
