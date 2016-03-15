package main

import "fmt"

func swap(x, y string) (string, string) { return y, x }
func main() {
	a, b := swap("hello", "Go语言")
	fmt.Println(a, b)
}
