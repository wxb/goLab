package main

import (
	"fmt"
)

func funcOfSomeType() interface{} {

	return 65
}

func main() {

	var t interface{}

	t = funcOfSomeType()

	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T", t) // %T 输出 t 是什么类型
	case bool:
		fmt.Printf("boolean %t\n", t) // t 是 bool 类型
	case int:
		fmt.Printf("integer %d\n", t) // t 是 int 类型
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t 是 *bool 类型
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t 是 *int 类型
	}
}
