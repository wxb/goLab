package pool_test

import (
	"bytes"
	"io"
	"os"
	"sync"
	"testing"
	"time"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer([]byte(""))
	},
}

func timeNow() time.Time {
	return time.Unix(1136214245, 0)
}

func log(w io.Writer, key, val string) {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	// Replace this with time.Now() in a real logger.
	b.WriteString(timeNow().UTC().Format(time.RFC3339))
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	w.Write(b.Bytes())
	bufPool.Put(b)

	bb := bufPool.Get().(*bytes.Buffer)
	w.Write(bb.Bytes())
}

func TestLog(t *testing.T) {
	log(os.Stdout, "path", "/search?q=flowers")
}
