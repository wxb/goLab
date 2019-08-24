package op

import (
	"bytes"
	"math/rand"
	"strconv"
)

// CPUProfile 测试操作
func CPUProfile() error {
	max := 10000000
	var buf bytes.Buffer

	for i := 0; i < max; i++ {
		num := rand.Int63n(int64(max))
		str := strconv.FormatInt(num, 10)
		buf.WriteString(str)
	}
	_ = buf.String()
	return nil
}
