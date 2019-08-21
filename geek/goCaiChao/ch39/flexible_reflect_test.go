package ch39_test

import (
	"errors"
	"reflect"
	"testing"
)

type Employee struct {
	ID   string
	Name string `json:"name" bb:"name"`
	Age  int
}

type Customer struct {
	Name  string
	Age   int
	Price float64
}

func fillBySetting(st interface{}, set map[string]interface{}) error {

	stType := reflect.TypeOf(st)
	if stType.Kind() != reflect.Ptr {
		return errors.New("the first param should be a pointer to the struct type")
	}

	if set == nil {
		return errors.New("")
	}

	for k, v := range set {
		// field, ok := reflect.ValueOf(st).Elem().Type().FieldByName(k)
		field, ok := stType.Elem().FieldByName(k)
		if !ok {
			continue
		}

		if field.Type == reflect.TypeOf(v) {
			// panic: reflect: call of reflect.Value.FieldByName on ptr Value
			// reflect.ValueOf(st).FieldByName(k).Set(reflect.ValueOf(v))

			reflect.ValueOf(st).Elem().FieldByName(k).Set(reflect.ValueOf(v))
		}
	}

	return nil
}

func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name": "Mike", "Age": 33}

	e := Employee{}
	err := fillBySetting(&e, settings)
	if err != nil {
		t.Log("Employee", err)
	}
	t.Log("Employee", e)

	c := new(Customer)
	err = fillBySetting(c, settings)
	if err != nil {
		t.Log("Customer:", err)
	}
	t.Log("Customer:", c)
}
