package mapstructure_test

import (
	"fmt"
	"testing"

	"github.com/mitchellh/mapstructure"
)

func TestDecode(t *testing.T) {
	type Person struct {
		Name   string
		Age    int
		Emails []string
		Extra  map[string]string
	}

	// This input can come from anywhere, but typically comes from
	// something like decoding JSON where we're not quite sure of the
	// struct initially.
	input := map[string]interface{}{
		"name":   "Mitchell",
		"age":    91,
		"emails": []string{"one", "two", "three"},
		"extra": map[string]string{
			"twitter": "mitchellh",
		},
	}

	var result Person
	err := mapstructure.Decode(input, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", result)
}

func TestDecodeEmbeddedStruct(t *testing.T) {
	// Squashing multiple embedded structs is allowed using the squash tag.
	// This is demonstrated by creating a composite struct of multiple types
	// and decoding into it. In this case, a person can carry with it both
	// a Family and a Location, as well as their own FirstName.
	type Family struct {
		LastName string
	}
	type Location struct {
		City string
	}
	type Person struct {
		Family    `mapstructure:",squash"`
		Location  `mapstructure:",squash"`
		FirstName string
	}

	input := map[string]interface{}{
		"FirstName": "Mitchell",
		"LastName":  "Hashimoto",
		"City":      "San Francisco",
	}

	var result Person
	err := mapstructure.Decode(input, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s %s, %s", result.FirstName, result.LastName, result.City)

}

func TestDecodeErrors(t *testing.T) {
	type Person struct {
		Name   string
		Age    int
		Emails []string
		Extra  map[string]string
	}

	// This input can come from anywhere, but typically comes from
	// something like decoding JSON where we're not quite sure of the
	// struct initially.
	input := map[string]interface{}{
		"name":   123,
		"age":    "bad value",
		"emails": []int{1, 2, 3},
	}

	var result Person
	err := mapstructure.Decode(input, &result)
	if err == nil {
		panic("should have an error")
	}

	fmt.Println(err.Error())
}

func TestDecodeMetaData(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	// This input can come from anywhere, but typically comes from
	// something like decoding JSON where we're not quite sure of the
	// struct initially.
	input := map[string]interface{}{
		"name":  "Mitchell",
		"age":   91,
		"email": "foo@bar.com",
	}

	// For metadata, we make a more advanced DecoderConfig so we can
	// more finely configure the decoder that is used. In this case, we
	// just tell the decoder we want to track metadata.
	var md mapstructure.Metadata
	var result Person
	config := &mapstructure.DecoderConfig{
		Metadata: &md,
		Result:   &result,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		panic(err)
	}

	if err := decoder.Decode(input); err != nil {
		panic(err)
	}

	fmt.Printf("Unused keys: %#v %#v", md.Keys, md.Unused)
}

func TestDecodeTags(t *testing.T) {
	// Note that the mapstructure tags defined in the struct type
	// can indicate which fields the values are mapped to.
	type Person struct {
		Name string `mapstructure:"person_name"`
		Age  int    `mapstructure:"person_age"`
	}

	input := map[string]interface{}{
		"person_name": "Mitchell",
		"person_age":  91,
	}

	var result Person
	err := mapstructure.Decode(input, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", result)
}

func TestDecodeWeaklyTypedInput(t *testing.T) {
	type Person struct {
		Name   string
		Age    int
		Emails []string
	}

	// This input can come from anywhere, but typically comes from
	// something like decoding JSON, generated by a weakly typed language
	// such as PHP.
	input := map[string]interface{}{
		"name":   123,                      // number => string
		"age":    "42",                     // string => number
		"emails": map[string]interface{}{}, // empty map => empty array
	}

	var result Person
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &result,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		panic(err)
	}

	err = decoder.Decode(input)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", result)
}

func TestDecodeString(t *testing.T) {
	// type Person struct {
	// 	Name   string
	// 	Age    int
	// 	Emails []string
	// }

	input := "0"
	var result bool
	err := mapstructure.WeakDecode(input, &result)

	fmt.Println(err)
	fmt.Printf("%#v", result)
}
