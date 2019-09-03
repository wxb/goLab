package strings_test

import (
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	r := strings.NewReader("0123456789")

	// b := make([]byte, 20)
	b := []byte{0, 0, 0, 0, 0, 0, 0}
	n, err := r.Read(b)
	t.Log(n, err, b, r.Len())

	c := make([]byte, 2)
	nn, err := r.ReadAt(c, 0)
	t.Log(nn, err, c, r.Len())

}
