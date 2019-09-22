package time_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func cancelTaskAfter(interval time.Duration, cancel context.CancelFunc) {
	go func(cancel context.CancelFunc) {
		time.Sleep(interval)
		fmt.Println("Cancell task", time.Now())
		cancel()
		wg.Done()
	}(cancel)
}

func TestTimerTask1(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	cancelTaskAfter(time.Second*25, cancel)
	wg.Add(1)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Cancelled", time.Now())
				wg.Done()
				return
			default:
				time.Sleep(time.Second * 10)
				fmt.Println("Invoked", time.Now())
			}
		}
	}(ctx)
	wg.Wait()
}

func TestTimerTask2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	cancelTaskAfter(time.Second*25, cancel)
	wg.Add(1)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Cancelled", time.Now())
				wg.Done()
				return
			case <-time.After(time.Second * 10):
				fmt.Println("Invoked", time.Now())
			}
		}
	}(ctx)
	wg.Wait()
}
