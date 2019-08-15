package cond_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCondMailBox(t *testing.T) {
	var mailbox uint8
	var lock sync.RWMutex
	sendCond := sync.NewCond(&lock)
	recvCond := sync.NewCond(lock.RLocker())

	go func() {
		lock.Lock()
		for mailbox == 1 {
			sendCond.Wait()
		}
		fmt.Println("MailBox before send:", mailbox)
		mailbox = 1
		lock.Unlock()
		recvCond.Signal()
	}()

	go func() {
		lock.RLock()
		for mailbox == 0 {
			recvCond.Wait()
		}
		fmt.Println("MailBox before recver:", mailbox)
		mailbox = 0
		lock.RUnlock()
		sendCond.Signal()
	}()
}

func TestCondBroadcast(t *testing.T) {
	var m sync.Mutex
	c := sync.NewCond(&m)
	n := 2

	for i := 0; i < n; i++ {
		go func(i int) {
			m.Lock()
			fmt.Println("inner")
			c.Wait()
			fmt.Println("continue")
			m.Unlock()
		}(i)
	}

	fmt.Println("start")
	time.Sleep(100 * time.Millisecond)
	c.Broadcast()
}
