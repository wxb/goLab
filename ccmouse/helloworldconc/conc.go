package main

import (
	"fmt"
)

func printHelloWorld(num int, ch chan string) {
	for {
		ch <- fmt.Sprintf("%d Hello, World!\n", num)

	}
}

func main() {
	ch := make(chan string)
	for i := 0; i < 5; i++ {
		// goroutine
		go printHelloWorld(i, ch)

	}

	for {
		msg := <-ch
		fmt.Println(msg)
	}

	//time.Sleep(time.Millisecond)
}
