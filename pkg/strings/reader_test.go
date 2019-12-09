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

func TestReadAt(t *testing.T) {
	r := strings.NewReader("abc的大写为ABC")
	t.Log(r.Len(), r.Size())

	s := []byte{1, 2, 3}
	n, err := r.ReadAt(s, 2)
	t.Log(n, err, string(s))

	ss := []byte{1, 2, 3, 4}
	nn, errr := r.ReadAt(ss, 2)
	t.Log(nn, errr, string(ss))

	sss := []byte{1, 2, 3, 4, 5}
	nnn, eer := r.ReadAt(sss, 0)
	t.Log(nnn, eer, string(sss))
}
