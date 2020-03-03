package json_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestRawMessageMarshal(t *testing.T) {
	h := json.RawMessage(`{"precomputed": true}`)

	c := struct {
		Header *json.RawMessage `json:"header"`
		Body   string           `json:"body"`
	}{Header: &h, Body: "Hello Gophers!"}
	b, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		t.Error("error:", err)
	}
	os.Stdout.Write(b)
}

func TestRawMessageUnmarshal(t *testing.T) {
	type Color struct {
		Space string
		Point json.RawMessage // delay parsing until we know the color space
	}
	type RGB struct {
		R uint8
		G uint8
		B uint8
	}
	type YCbCr struct {
		Y  uint8
		Cb int8
		Cr int8
	}

	var j = []byte(`[
	{"Space": "YCbCr", "Point": {"Y": 255, "Cb": 0, "Cr": -10}},
	{"Space": "RGB",   "Point": {"R": 98, "G": 218, "B": 255}}
	]`)

	var colors []Color
	err := json.Unmarshal(j, &colors)
	if err != nil {
		t.Fatal("error:", err)
	}
	for _, c := range colors {
		var dst interface{}
		switch c.Space {
		case "RGB":
			dst = new(RGB)
		case "YCbCr":
			dst = new(YCbCr)
		}
		err := json.Unmarshal(c.Point, dst)
		if err != nil {
			t.Fatal("error:", err)
		}
		fmt.Println(c.Space, dst)
	}
}
