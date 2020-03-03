package json_test

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
)

func TestHTMLEscape(t *testing.T) {
	var out bytes.Buffer
	json.HTMLEscape(&out, []byte(`{"Name":"<b>HTML content</b>"}`))
	out.WriteTo(os.Stdout)
}
