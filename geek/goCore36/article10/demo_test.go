package article10_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestDemo00(t *testing.T) {
	ch1 := make(chan int, 3)
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	elem1 := <-ch1
	fmt.Printf("The first element received from channel ch1: %v\n",
		elem1)
}

func TestDemo01(t *testing.T) {
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
}

func TestDemo02(t *testing.T) {
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

func TestDemo03(t *testing.T) {
	// 示例1。
	ch1 := make(chan int, 1)
	ch1 <- 1
	//ch1 <- 2 // 通道已满，因此这里会造成阻塞。

	// 示例2。
	ch2 := make(chan int, 1)
	//elem, ok := <-ch2 // 通道已空，因此这里会造成阻塞。
	//_, _ = elem, ok
	ch2 <- 1

	// 示例3。
	var ch3 chan int
	//ch3 <- 1 // 通道的值为nil，因此这里会造成永久的阻塞！
	//<-ch3 // 通道的值为nil，因此这里会造成永久的阻塞！
	_ = ch3
}
