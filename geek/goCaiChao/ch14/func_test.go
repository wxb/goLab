package ch14_test

import (
	"fmt"
	"testing"
)

// 可变长参数
func sum(args ...int) (total int) {
	for _, p := range args {
		total += p
	}
	return
}

func TestVarParams(t *testing.T) {
	t.Log(sum(1, 2, 3))
	t.Log(sum(1, 2, 3, 4))
	t.Log(sum([]int{1, 2, 3, 4}...))
	t.Log([]interface{}{1, 2, 3, 4}...)
}

func TestDefer(t *testing.T) {
	defer func() {
		fmt.Println("defer")
	}()

	fmt.Println("normal")
	// panic("panic")   // 会执行defer
	// os.Exit(1)       // 不会执行defer
}
