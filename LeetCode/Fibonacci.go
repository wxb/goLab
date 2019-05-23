package main

import "fmt"

type Fib func(x int) int

func fib(n int) int {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	f := fib
	for k := range [10]uint{} {
		fmt.Println(fib(k))
		fmt.Println(f(k))
	}
}
