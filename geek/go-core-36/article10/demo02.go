package main

import (
	"fmt"
	"runtime"
)

func main() {
	//runtime.GOMAXPROCS(1)
	ch1 := make(chan int, 4)
	// 发送方。
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
			fmt.Printf("Sender: sending element %v...\n", i)
		}
		fmt.Println("Sender: close the channel...")
		close(ch1)
	}()
	fmt.Println(runtime.NumCPU(), runtime.NumGoroutine())

	//time.Sleep(1 * time.Second)
	// 接收方。
	for {
		elem, ok := <-ch1
		if !ok {
			fmt.Println("Receiver: closed channel")
			break
		}
		fmt.Printf("Receiver: received an element: %v\n", elem)
	}

	fmt.Println("End.")
}
