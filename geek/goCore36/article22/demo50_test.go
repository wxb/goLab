package article22_test

import (
	"fmt"
	"testing"
)

func TestDemo50(t *testing.T) {
	defer fmt.Println("first defer")
	for i := 0; i < 3; i++ {
		defer fmt.Printf("defer in for [%d]\n", i)
	}
	defer fmt.Println("last defer")
}
