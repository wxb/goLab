package ch31_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The request is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}

	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine()) // 2

	t.Log(FirstResponse())
	time.Sleep(10 * time.Millisecond)

	t.Log("After:", runtime.NumGoroutine()) // 11 FirstResponse函数启动的10个协程只有一个结果写入ch后退出，其他9个还在goroutine中阻塞
}

func FirstResponseWithCapChan() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner) // 有缓冲通道
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}

	return <-ch
}

func TestFirstResponseWithCapChan(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine()) // 2

	t.Log(FirstResponseWithCapChan())
	time.Sleep(10 * time.Millisecond)

	t.Log("After:", runtime.NumGoroutine()) // 11 FirstResponse函数启动的10个协程只有一个结果写入ch后退出，其他9个还在goroutine中阻塞
}
