package ch27

import (
	"fmt"
	"sync"
)

// DataProducer 数据生产
func DataProducer(ch chan<- int, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
}

// DataCloseProducer close关闭通道
func DataCloseProducer(ch chan<- int, wg *sync.WaitGroup) {
	go func() {
		defer func() {
			// 1. 对一个已关闭的通道再赋值时会panic
			// 2. 对一个已关闭的通道再取值时会得到对应的类型零值
			close(ch)
			wg.Done()
		}()

		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
}

// DataReceiver 数据接收
func DataReceiver(ch <-chan int, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			data := <-ch
			fmt.Println(data)
		}
	}()
}

// DataForeceiver for循环if判断
func DataForeceiver(ch <-chan int, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for {
			data, ok := <-ch
			if !ok {
				break
			}
			fmt.Println(data)
		}
	}()
}

// DataRangeReceiver for循环➕range
func DataRangeReceiver(ch <-chan int, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for data := range ch {
			fmt.Println(data)
		}
	}()
}
