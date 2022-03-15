package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

// https://www.4async.com/2021/08/golang-117-generics/

func mapFunc[T any, M any](a []T, f func(T) M) []M {
    n := make([]M, len(a), cap(a))
    for i, e := range a {
        n[i] = f(e)
    }
    return n
}

func filterFunc[T any](a []T, f func(T) bool) []T {
	var n []T

	for _, e := range a {
		if f(e) {
			n = append(n, e)
		}
	}

	return n
}

func main() {
    vi := []int{1,2,3,4,5,6}
    vs := mapFunc(vi, func(v int) int {
        return v*v
    })
    fmt.Println(vs)



    vss := mapFunc(vi, func(v int) string {
		return "<"+fmt.Sprint(v)+">"
    })
    fmt.Println(vss)


	vii := filterFunc(mapFunc([]int{1,2,3,4,5,6,7}, func (v int) int {
		return v * v
	}), func (v int) bool {
		return v < 40
	})
	fmt.Println(vii)


	vsss := filterFunc(mapFunc([]string{"a", "b", "c", "d", "e"}, func(v string) string{
		n, _ := rand.Int(rand.Reader, big.NewInt(5))

		i := int(n.Int64()) + 1
		return strings.Repeat(v,i)
	}), func (v string) bool {
		return len(v) > 3
	})
	fmt.Println(vsss)
}
