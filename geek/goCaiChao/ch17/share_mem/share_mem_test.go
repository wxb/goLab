package share_mem_test

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {

	counter := 0

	for i := 0; i < 50; i++ {
		go func() { counter++ }()
	}

	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterThreadSafe(t *testing.T) {

	counter := 0
	var m sync.Mutex

	for i := 0; i < 500000; i++ {
		go func() {
			m.Lock()
			defer m.Unlock()
			counter++
		}()
	}

	time.Sleep(1 * time.Microsecond)
	t.Logf("counter = %d", counter)

}

func TestCounterWaitGroup(t *testing.T) {

	counter := 0
	var m sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 50000; i++ {
		wg.Add(1)
		go func() {
			m.Lock()
			defer func() {
				m.Unlock()
				wg.Done()
			}()
			counter++
		}()
	}

	wg.Wait()
	t.Logf("counter = %d", counter)

}
