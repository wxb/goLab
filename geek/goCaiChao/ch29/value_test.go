package ch29_test

import (
	"context"
	"fmt"
	"testing"
)

func TestWithValue(t *testing.T) {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			t.Log("found value:", v)
			return
		}
		t.Log("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "go")

	f(ctx, k)
	f(ctx, favContextKey("color"))

	fmt.Println(context.Background(), context.TODO())
}
