package cond_test

import (
	"sync"
	"testing"
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
		t.Log("send:", mailbox)
		mailbox = 1
		lock.Unlock()
		recvCond.Signal()
	}()

	go func() {
		lock.RLock()
		for mailbox == 0 {
			recvCond.Wait()
		}
		t.Log("recver:", mailbox)
		mailbox = 0
		lock.RUnlock()
		sendCond.Signal()
	}()

}
