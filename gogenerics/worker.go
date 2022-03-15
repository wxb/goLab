package main

import (
	"fmt"
	"strconv"
	"sync"
)

// https://www.4async.com/2021/08/golang-117-generics/

type T1 interface{}
type T2 interface{}

func ParallelMap(parallelism int, in []T1, f func(T1) (T2, error)) ([]T2, error) {
	var wg sync.WaitGroup
	defer wg.Wait()

	inc, outc, errc := make(chan T1), make(chan T2), make(chan error)

	donec := make(chan struct{})
	defer close(donec)

	wg.Add(parallelism)
	for i := 0; i < parallelism; i++ {
		go func() {
			defer wg.Done()
			for x := range inc {
				y, err := f(x)
				if err != nil {
					select {
					case errc <- err:
					case <-donec:
					}
					return
				}
				select {
				case outc <- y:
				case <-donec:
					return
				}
			}
			select {
			case errc <- nil:
			case <-donec:
			}
		}()
	}

	go func() {
		for _, x := range in {
			inc <- x
		}
		close(inc)
	}()

	out := make([]T2, 0, len(in))
	for rem := parallelism; rem > 0; {
		select {
		case err := <-errc:
			if err != nil {
				return nil, err
			}
			rem--
		case y := <-outc:
			out = append(out, y)
		}
	}
	return out, nil
}

func main() {
	in := []T1{"1", "2", "3", "4", "5"}
	out, err := ParallelMap(4, in, func(x T1) (T2, error) {
		return strconv.Atoi(x.(string))
	})
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(out)

	in2 := []T1{1, 2, 3, 4, 5}
	out2, err := ParallelMap(4, in2, func(x T1) (T2, error) {
		return fmt.Sprintf("<%d>", x), nil
	})
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(out2)
}
