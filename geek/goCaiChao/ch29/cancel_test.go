package ch29_test

import (
	"context"
	"testing"
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
