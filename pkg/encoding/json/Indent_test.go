package json_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestIndent(t *testing.T) {

	type Road struct {
		Name   string
		Number int
	}
	roads := []Road{
		{"Diamond Fork", 29},
		{"Sheep Creek", 51},
	}

	b, err := json.Marshal(roads)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	var out bytes.Buffer
	json.Indent(&out, b, "=", "\t")

	fmt.Println("====", out.String())
	out.WriteTo(os.Stdout)
}
