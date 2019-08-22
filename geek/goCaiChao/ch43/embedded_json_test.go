package ch43_test

import (
	"encoding/json"
	"testing"
)

type BasicInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type JobInfo struct {
	Skills []string `json:"skills"`
}

type Employee struct {
	BasicInfo BasicInfo `json:"basic_info"`
	JobInfo   JobInfo   `json:"job_info"`
}

var jsonStr string = `{"basic_infl":{}, "job_info":{"skills":["php", "golang", "java"]}}`

func TestEmbeddJSON(t *testing.T) {

	e := new(Employee)

	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}

	t.Log(*e)

	bytes, err := json.Marshal(e)
	if err != nil {
		t.Error(err)
	}

	t.Log(string(bytes))
	t.Log(string(bytes) == jsonStr)
}
