package main

import (
	"sort"
	"fmt"
)

func main() {
	slicea := []int{3, 6, 1, 0, 2, 7, 5, 9, 4}
	sort.Ints(slicea)
	for _, v := range slicea {
		fmt.Println(v)
	}
}
