package context_test

import (
	"context"
	"testing"
	"time"
)

func TestWithDeadline(t *testing.T) {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		t.Log("overslept")
	case <-ctx.Done():
		t.Log(ctx.Err())
	}
}
