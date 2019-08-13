package ch29_test

import (
	"context"
	"testing"
	"time"
)

func TestWithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		t.Log("WithTimeout overslept")
	case <-ctx.Done():
		t.Log(ctx.Err())
	}
}
