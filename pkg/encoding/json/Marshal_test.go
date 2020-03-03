package json_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestMarshalFieldTag(t *testing.T) {
	type person struct {
		// Field appears in JSON as key "myName".
		Name string `json:"myName"`

		// Field appears in JSON as key "myName" and
		// the field is omitted from the object if its value is empty,
		// as defined above.
		Age int `json:"myAge,omitempty"`

		// Field appears in JSON as key "Field" (the default), but
		// the field is skipped if empty.
		// Note the leading comma.
		Pass bool `json:",omitempty"`

		// Field is ignored by this package.
		PassWord int `json:"-"`

		// Field appears in JSON as key "-".
		Status int `json:"-,"`

		// It applies only to fields of string, floating point, integer, or boolean types.
		Int64String int64 `json:",string"`

		// The key name will be used if it's a non-empty string consisting of only Unicode letters, digits, and ASCII punctuation except quotation marks, backslash, and comma.
		Address string `json:"address'city"`
	}
	data := person{"Martina", 19, false, 1213, 0, 67890, "shanxi/xian"}

	b, err := json.Marshal(data)
	if err != nil {

	}
	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	io.Copy(os.Stdout, &out)
	out.WriteTo(os.Stderr)

	data1 := person{"Rihanna", 0, true, 56789, 2, 65, "Beijing/Changhe"}
	d, _ := json.MarshalIndent(data1, "", "\t")
	os.Stdout.Write(d)
}

func TestMarshalStruct(t *testing.T) {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}
