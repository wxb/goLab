package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(time.Now(), t.Unix())
	for i := 0; i < 5000; i++ {
		go printHelloWorld(i)
	}
	fmt.Println(t, time.Now())

	time.Sleep(time.Millisecond)
	fmt.Println(time.Now())
}

func printHelloWorld(i int) {
	for {
		fmt.Printf("Hello World From goroutine %d \n", i)
	}
}
