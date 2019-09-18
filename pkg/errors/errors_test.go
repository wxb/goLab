package errors_test

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestErrorsWrap(t *testing.T) {
	err := errors.New("这是一条错误")

	// go errors包中并没有提供wrap这样的方法，而是通过在fmt.Errorf中增加一个w动词
	e := fmt.Errorf("嵌套一条错误：%w", err)
	if w, ok := e.(interface {
		Unwrap() error
	}); ok {
		fmt.Println("这是一条嵌套error:", w)
		fmt.Println("嵌套的error为:", w.Unwrap())
	}

}

func TestErrorsUnwrap(t *testing.T) {
	err := fmt.Errorf("父Error：%w", errors.New("子Error"))

	w := errors.Unwrap(err)
	if w != nil {
		fmt.Println("内部嵌套的Error是:", w)
	} else {
		fmt.Println("内部没有嵌套的Error, 错误为:", err)
	}
}

func TestErrorsIs(t *testing.T) {
	err := fmt.Errorf("父Error: %w", os.ErrExist)
	if errors.Is(err, os.ErrExist) {
		t.Log("err is os.ErrExist type")
	} else {
		t.Error("err is not os.ErrExist type")
	}

	e := fmt.Errorf("父Error:%w", errors.New("子ERROR"))
	if errors.Is(e, os.ErrExist) {
		t.Log("e is os.ErrExist type")
	} else {
		t.Error("e is not os.ErrExist type")
	}
}

func TestErrorsAs(t *testing.T) {
	if _, err := os.Open("non-existing"); err != nil {
		var pathError *os.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}
}

func TestFmtErrorf(t *testing.T) {
	const name, id = "bueller", 17

	// 1. 用法
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	fmt.Println(err.Error())

	// 2. %w 用法
	w := fmt.Errorf("parent wrap: %w", err)
	fmt.Println("2.1 ", w)
	if ww, ok := w.(interface {
		Unwrap() error
	}); ok {
		e := ww.Unwrap()
		fmt.Println("2.2 ", e.Error())
	}

	// 3. 多个 %w， %w 相当于%v
	fmt.Println(
		fmt.Errorf("3.1 parent wrap: %w %w", err, errors.New("lt one w")),
		fmt.Errorf("3.1 parent wrap: %w %v", err, errors.New("lt one w")),
	)

	// 4. 提供一个未实现error接口的参数, %w 相当于%v
	fmt.Println(
		fmt.Errorf("4. this is number %w", 123),
	)
}
