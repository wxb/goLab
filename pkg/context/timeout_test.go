package context

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

func TestWithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		t.Log("WithTimeout overslept")
	case <-ctx.Done():
		t.Log("ctx.Done():", ctx.Err())
	}
}

// ctx.Done的查找链
func TestWithTimeoutChain(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	go func(ctx context.Context) {
		cctx, _ := context.WithCancel(ctx)
		go func(cctx context.Context) {
			ccctx, cancelF := context.WithCancel(cctx)
			go func(ccctx context.Context) {
				for {
					select {
					case <-ccctx.Done():
						fmt.Println("\t\tccctx done:", ccctx.Err(), checkMark)
						return
					default:
						fmt.Println("\t\tccctx: none", ballotX)
					}
				}
			}(ccctx)
			time.Sleep(1 * time.Millisecond)
			cancelF()

			for {
				select {
				case <-cctx.Done():
					fmt.Println("\tcctx done:", cctx.Err(), checkMark)
					return
				default:
					time.Sleep(10 * time.Millisecond)
					fmt.Println("\tcctx: none", ballotX)
				}
			}
		}(cctx)

		select {
		case <-time.After(500 * time.Millisecond):
			fmt.Println("WithTimeout overslept")
		case <-ctx.Done():
			fmt.Println("ctx:", ctx.Err(), checkMark)
		}
	}(ctx)

	time.Sleep(1 * time.Second)
}
