package ch123_test

import (
	"fmt"
	"reflect"
)

func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid", path)
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			display(fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name), v.Field(i))
		}
	case reflect.Map:
		// for _, key := range v.MapKeys() {
		// 	display(fmt.Sprintf("%s[%s]", path))
		// }
	case reflect.Ptr:

	case reflect.Interface:

	default:
	}
}

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}
