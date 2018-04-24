package main

import (
	"fmt"
	"time"
)

var c chan string

func SayHi() {
	i := 1
	for {
		fmt.Println("From SayHi: Hi,", <-c, i)
		c <- fmt.Sprint(i)
		i++
	}
}

func main() {
	c = make(chan string)
	names := map[int]string{1: "wang xiaobo", 2: "王晓勃"}
	for _, v := range names {
		go SayHi()
		c <- v
		fmt.Println("From main: ", <-c)
	}

	time.Sleep(2*time.Second)
}
