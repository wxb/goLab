package main

import (
	"flag"
	"fmt"
	"runtime"
)

var host string

func init() {
	flag.StringVar(&host, "host", "127.0.0.1", "Remote Host")
}

func main() {
	var port = flag.Int("port", 9909, "Remote Port")
	flag.Parse()
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	fmt.Println(host, *port, runtime.NumGoroutine())
}
