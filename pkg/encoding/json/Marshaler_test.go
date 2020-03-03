package json_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
)

type Staff string

func (s Staff) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	n, err := fmt.Fprintf(buf, `{"staff":"%s"}`, s)
	if err != nil {
		return nil, err
	}
	if n-12 != len(s) {
		return nil, errors.New("EOF")
	}

	return buf.Bytes(), nil
}

func TestMarshaler(t *testing.T) {
	staff := map[string]interface{}{
		"Name": Staff("Bobi"),
	}
	res, err := json.Marshal(staff)
	fmt.Println(string(res), err)
}
