package ch03_test

import (
	"fmt"
	"sync"
	"testing"
)

func fooReentryMutex(l sync.Locker) {
	fmt.Println("in fooReentryMutex")
	l.Lock()
	barReentryMutex(l)
	l.Unlock()
}

func barReentryMutex(l sync.Locker) {
	l.Lock()
	fmt.Println("in bar")
	l.Unlock()
}
func TestReentryMutex(t *testing.T) {
	l := &sync.Mutex{}
	fooReentryMutex(l)
}
