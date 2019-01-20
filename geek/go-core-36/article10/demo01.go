package main

import (
	"fmt"
	"time"
)

func main() {
	var ch0 chan string
	fmt.Println(ch0 == nil)

	ch := make(chan int, 1)
	fmt.Println(ch)
	go func(c chan int) {
		time.Sleep(1 * time.Second)
		fmt.Println(<-c)
	}(ch)
	ch <- 2
	fmt.Println("222")
	//ch <- 1
	//fmt.Println(<-ch)
}
