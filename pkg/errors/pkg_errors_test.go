package errors_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/pkg/errors"
)

// 这个错误处理库的错误类型，都实现了Formatter接口，我们可以通过fmt.Printf函数输出对应的错误信息。
// %s,%v //功能一样，输出错误信息，不包含堆栈
// %q //输出的错误信息带引号，不包含堆栈
// %+v //输出错误信息和堆栈

func isOpen(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "can't open the path file")
	}
	defer f.Close()

	_, err = ioutil.ReadAll(f)
	if err != nil {
		return errors.Wrap(err, "can't read the path file")
	}

	return nil
}

func TestPkgErrorsNew(t *testing.T) {
	err := errors.New("this is pkg/errors instance")
	t.Logf("%s", err)
	t.Logf("%v", err)
	t.Logf("%q", err)
	t.Logf("%+v", err)
}

func TestPkgErrorsWrap(t *testing.T) {
	err := isOpen("")
	if err != nil {
		t.Errorf("%s", err)
		t.Errorf("%v", err)
		t.Errorf("%+v", err)

	}
}

func TestPkgErrorsCause(t *testing.T) {
	open := func(path string) error {
		err := isOpen(path)
		if err != nil {
			return errors.Wrapf(err, "check[func:%s] is fail!", "isOpen")
		}

		f, _ := os.Open(path)
		defer f.Close()

		_, err = ioutil.ReadAll(f)
		if err != nil {
			return errors.Wrapf(err, "Read[func:%s %d] is fail", "ReadAll", 61)
		}

		return nil
	}
	err := open("")
	if err != nil {
		t.Errorf("%s", err)
		t.Errorf("%v", err)
		t.Errorf("%+v", err)

		e := errors.Cause(err)
		t.Errorf("%s", e)
		t.Errorf("%v", e)
		t.Errorf("%+v", e)
	}
}

func TestPkgErrorsErrorf(t *testing.T) {
	err := errors.Errorf("whoops: %s", "foo")
	t.Logf("%+v", err)
}
