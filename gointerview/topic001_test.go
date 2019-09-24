package gointerview_test

import "testing"

func TestTopic001(t *testing.T) {
	i := GetValue()

	switch interface{}(i).(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}

}

func GetValue() int {
	return 1
}
