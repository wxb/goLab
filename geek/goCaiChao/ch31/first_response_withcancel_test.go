package ch31_test

import (
	"context"
	"fmt"
	"testing"
)

func doSlowRequest(id int) string {
	fmt.Println("===", id)
	return fmt.Sprintf("The request is from %d", id)
}

func FirstResponseWithCancel() string {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	numOfRunner := 10
	ch := make(chan string)
	for i := 0; i < numOfRunner; i++ {
		go func(i int, ctx context.Context) {

			select {
			case <-ctx.Done():
				fmt.Println("--", i)
				return
			default:
				ret := doSlowRequest(i)
				ch <- ret
			}

		}(i, ctx)
	}

	return <-ch
}

func TestFirstResponseWithCancel(t *testing.T) {
	fmt.Println(FirstResponseWithCancel())
}
