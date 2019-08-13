package ch29_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 30; i++ {
		go func(i int) {
			for {
				if isCancelled(ctx) {
					break
				} else {
					// fmt.Println(i, "Not Yet")
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Canceled")
		}(i)
	}

	cancel()
	time.Sleep(time.Millisecond * 100)
	t.Logf(ctx.Err().Error())
}
