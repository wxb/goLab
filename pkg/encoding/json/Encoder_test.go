package json_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	type Message struct {
		Name, Text string
	}

	msg := Message{"QQ", "join in"}
	buf := bytes.NewBuffer([]byte{})
	enc := json.NewEncoder(buf)

	err := enc.Encode(msg)
	fmt.Println(err, buf.String(), buf.Cap(), buf.Len())

	r, n, e := buf.ReadRune()
	fmt.Println(string(r), n, e, buf.Cap(), buf.Len())

	line, e := buf.ReadString(':')
	fmt.Println(line, e)

	buf.Truncate(5)
	fmt.Println(buf.String(), buf.Cap(), buf.Len())
}

func TestEncoderSetEscapeHTML(t *testing.T) {

	type Message struct {
		Name, Text string
	}

	msg := Message{"QQ", "<h1>Slogan</h1> join in "}

	enc := json.NewEncoder(os.Stdout)

	enc.Encode(msg) // {"Name":"QQ","Text":"\u003ch1\u003eSlogan\u003c/h1\u003e join in "}

	enc.SetEscapeHTML(false)
	enc.Encode(msg) // {"Name":"QQ","Text":"<h1>Slogan</h1> join in "}
}

func TestEncoderSetIndent(t *testing.T) {
	type Message struct {
		Name, Text string
	}

	msg := Message{"QQ", "<h1>Slogan</h1> join in "}

	enc := json.NewEncoder(os.Stdout)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "\t")
	enc.Encode(msg)
}
