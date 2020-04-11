package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			select {
			case <-time.After(10 * time.Second):
				fmt.Println(time.Now().Unix())
			}
		}
	}()

	time.Sleep(50 * time.Second)
}
