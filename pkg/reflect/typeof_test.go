package reflect_test

import (
	"reflect"
	"testing"
)

func TestTypeOf(t *testing.T) {

	type User struct {
		Name string
		Age  int
	}

	u := User{"张三", 20}
	tt := reflect.TypeOf(u)
	vv := reflect.ValueOf(u)

	t.Log(tt, vv, tt.Kind())
}
