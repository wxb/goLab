package ch38_test

import (
	"fmt"
	"reflect"
	"testing"
)

func checkType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int8, reflect.Int64, reflect.Int16:
		fmt.Println("Int")
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Bool:
		fmt.Println("Boolean")
	default:
		fmt.Println("Unkown")
	}
}

func TestCheckType(t *testing.T) {
	checkType(false)
}

func TestBasicType(t *testing.T) {
	var f float64 = 65

	checkType(f)
	checkType(&f)
}

// 通过反射灵活调用类型方法或属性

type Employee struct {
	ID   string
	Name string `json:"name" bb:"name"`
	Age  int
}

func (e *Employee) UpdateAge(val int) {
	e.Age = val
}

// 演示如何动态调用方法和属性
func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}

	t.Logf("Name: value(%[1]v), type(%[1]T)", reflect.ValueOf(*e).FieldByName("Name"))

	nameField, ok := reflect.TypeOf(*e).FieldByName("Name")
	if !ok {
		t.Error("Failed to get 'Name' field")
	} else {
		t.Logf("json:%s, bb:%s", nameField.Tag.Get("json"), nameField.Tag.Get("bb"))
	}

	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(12)})

	t.Log("Update Age:", e)
}
