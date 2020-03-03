package json_test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestMarshalIndenet(t *testing.T) {

	data := map[string]int{
		"a": 1,
		"b": 2,
	}

	json, err := json.MarshalIndent(data, "<prefix>", "<indent>")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json))
}
