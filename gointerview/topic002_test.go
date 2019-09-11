package gointerview_test


import (
	"fmt"
	"testing"
)

func TestTopic002(t *testing.T) {
	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}
