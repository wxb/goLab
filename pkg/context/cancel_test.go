package context_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

func TestWithCancel(t *testing.T) {
	gen := func(ctx context.Context) <-chan int {
		n := 1
		dst := make(chan int)
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for n := range gen(ctx) {
		t.Log(n)
		if n == 5 {
			break
		}
	}
}

func TestWithCancelSub(t *testing.T) {
	// 由context.TODO()根ctx衍生出一级子ctx
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	for i := 0; i < 3; i++ {
		go func(i int, ctx context.Context) {
			// 由一级子ctx衍生出二级timeCtx，因测试由一级ctx取消传递，这里的二级cancel函数不接受
			timeCtx, _ := context.WithCancel(ctx)
			go func(i int, timeCtx context.Context) {
				for {
					select {
					// 二级timeCtx的done收到了一级ctx取消传递的信息
					case <-timeCtx.Done():
						fmt.Printf("\ttimeoutCtx goroutine[%d] get cancel signal %v\n ", i, checkMark)
						return
					default:
						fmt.Printf("\ttimeoutCtx goroutine[%d] get cancel signal %v\n ", i, ballotX)
					}
				}
			}(i, timeCtx)

			for {
				select {
				case <-ctx.Done():
					fmt.Printf("goroutine[%d] get cancel signal %v\n ", i, checkMark)
					return
				default:
					fmt.Printf("goroutine[%d] get cancel signal %v\n ", i, ballotX)
				}
			}
		}(i, ctx)
	}

	time.Sleep(100 * time.Millisecond)
}
