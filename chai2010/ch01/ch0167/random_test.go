package ch0167_test

import (
	"strings"
	"testing"
)

func random(n int) string {
	ch := make(chan string)

	go func() {
		for {
			select {
			case ch <- "0123456789":
			case ch <- "abcdefghijklmnopqrstuvwxyz":
			case ch <- "ABCDEFGHIJKLMNOPQRSTUVWXYZ":
			case ch <- "?!@#$^&*":
			}
		}
	}()

	r := strings.Builder{}
	r.Grow(n)
	for i := 0; i < n; i++ {
		r.WriteString(<-ch)
	}

	return r.String()
}

func TestRandom(t *testing.T) {
	l := 30
	if len(random(l)) != l {
		t.Error("")
	}
}

func BenchmarkRandom(b *testing.B) {

	for i := 0; i < b.N; i++ {
		// fmt.Println()
		random(30)
	}
}
