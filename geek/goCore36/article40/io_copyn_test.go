package article40_test

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestIOCopyN(t *testing.T) {
	src := strings.NewReader(
		"CopyN copies n bytes (or until an error) from src to dst. " +
			"It returns the number of bytes copied and " +
			"the earliest error encountered while copying.")
	// dst := new(strings.Builder)
	var dst strings.Builder

	written, err := io.CopyN(&dst, src, 58)
	if err != nil {
		fmt.Printf("error:%v\n", err)
	} else {
		fmt.Printf("Written(%d): %q", written, (&dst).String())
	}
}

func TestIOCopyBuffer(t *testing.T) {
	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	// buf := make([]byte, 8)
	// buf := make([]byte, 0)    // buf!=nil, len(buf)==0   👇 panic
	var buf []byte // buf==nil, len(buf)==0

	t.Log(buf, len(buf), buf == nil)

	// buf is used here...
	written, err := io.CopyBuffer(os.Stdout, r1, buf)
	t.Log(written, err, buf, len(buf))

	// ... reused here also. No need to allocate an extra buffer.
	written, err = io.CopyBuffer(os.Stdout, r2, buf)
	t.Log(written, err, buf, len(buf))
}
