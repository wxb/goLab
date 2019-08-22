package ch44

import (
	"encoding/json"
	"testing"
)

func TestMarshalEasyJSON(t *testing.T) {
	e := Employee{BasicInfo: BasicInfo{"王晓勃", 123}}
	b, err := e.MarshalJSON()
	t.Log("MarshalJSON:", string(b), err)
}

func TestUnmarshalEasyJSON(t *testing.T) {

	bi := BasicInfo{}
	err := bi.UnmarshalJSON([]byte(`{"name":"王晓勃","age":123}`))

	t.Log("UnmarshalJSON:", bi, err)
}

func BenchmarkMarshalGoJSON(b *testing.B) {
	e := Employee{BasicInfo: BasicInfo{"王晓勃", 123}}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		bytes, err := json.Marshal(e)
		if err != nil {
			b.Log("MarshalGoJSON:", err)
		}

		ee := Employee{}
		json.Unmarshal(bytes, &ee)
	}

	b.StopTimer()
}

func BenchmarkMarshalEasyJSON(b *testing.B) {
	e := Employee{BasicInfo: BasicInfo{"王晓勃", 123}}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		bytes, err := e.MarshalJSON()
		if err != nil {
			b.Log("MarshalEasyJSON:", err)
		}

		ee := Employee{}
		ee.UnmarshalJSON(bytes)
	}

	b.StopTimer()
}
