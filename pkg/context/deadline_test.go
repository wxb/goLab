package context_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestWithDeadline(t *testing.T) {
	d := time.Now().Add(1001 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		// t.Log("overslept")
		fmt.Println("overslept")
	case <-ctx.Done():
		// t.Log(ctx.Err())
		fmt.Println(ctx.Err())
		// default:
		// 	fmt.Println("default")
	}
	fmt.Println("select finish")
}
