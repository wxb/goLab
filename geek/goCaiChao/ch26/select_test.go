package ch26_test

import (
	"testing"
	"time"
)

func Service() <-chan bool {
	ok := make(chan bool)

	go func() {
		time.Sleep(time.Millisecond * 20)
		ok <- true
	}()

	return ok
}

// TestSelect select 多路选择与超时机制
func TestSelect(t *testing.T) {
	select {
	case ok := <-Service():
		t.Logf("service chan: %T", ok)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
