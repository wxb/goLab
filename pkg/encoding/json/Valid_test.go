package json_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestValid(t *testing.T) {
	goodJSON := `{"example": 1}`
	badJSON := `{"example":2:]}}`

	fmt.Println(json.Valid([]byte(goodJSON)), json.Valid([]byte(badJSON)))
}
