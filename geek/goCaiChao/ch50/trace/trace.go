package main

import (
	"os"
	"runtime/trace"
	"strconv"
)

func main() {

	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	var str string

	for i := 0; i < 10000; i++ {
		str += strconv.Itoa(i)
	}
}

// go tool trace trace.out
