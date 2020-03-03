package json_test

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
)

func TestCompact(t *testing.T) {
	data := map[string]interface{}{
		"name": "Lady Gaga",
		"song": []string{
			"Poker Face",
			"The Cure",
			"Bad Romance",
		},
	}
	src, _ := json.MarshalIndent(data, "", "\t")
	t.Log(string(src))

	// Compact 将src进行压缩删除无用的空格后追加到dst
	var dst bytes.Buffer
	err := json.Compact(&dst, src)
	if err != nil {
		t.Fatal("compact fail", err)
	}

	dst.WriteTo(os.Stdout)
}
