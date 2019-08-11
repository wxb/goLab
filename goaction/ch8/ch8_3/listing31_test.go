package ch8_3_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMarshalIndent(t *testing.T) {
	c := make(map[string]interface{})
	c["name"] = "Gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home": "415.333.3333",
		"cell": "415.555.5555",
	}

	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		t.Fatal("Error:", err)
	}

	fmt.Println("Data:", string(data))
}

func TestMarshal(t *testing.T) {
	c := make(map[string]interface{})
	c["name"] = "Gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home": "415.333.3333",
		"cell": "415.555.5555",
	}

	data, err := json.Marshal(c)
	if err != nil {
		t.Fatal("Error:", err)
	}

	fmt.Println("Data:", string(data))
}
