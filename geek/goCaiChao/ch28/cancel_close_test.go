package ch28_test

import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(ch chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}

func cancel1(ch chan struct{}) {
	ch <- struct{}{}
}

func cancel2(ch chan struct{}) {
	close(ch)
}

func TestCancel(t *testing.T) {
	cancelCh := make(chan struct{}, 0)

	for i := 0; i < 3; i++ {
		go func(i int, cc chan struct{}) {
			for {
				if isCancelled(cc) {
					break
				} else {
					// fmt.Println(i, "Not Yet")
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Canceled")
		}(i, cancelCh)
	}

	// time.Sleep(time.Millisecond * 100) // ???? 如果加上这一行, 就看不到Canceled的输出
	cancel1(cancelCh)
	// cancel2(cancelCh)
}
