package ch44

import "testing"

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

func BenchmarkGoJSON(b *testing.B) {

}

func BenchmarkEasyJSON(b *testing.B) {

}
