package ch0107

import "encoding/json"

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Avatar string `json:"avatar"`
	Type   string `json:"type"`
}

type AgainPerson struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Avatar string `json:"avatar"`
	Type   string `json:"type"`
}

const MAX = 10000

func InitPerson() (persons []Person) {
	for i := 0; i < MAX; i++ {
		persons = append(persons, Person{
			Name:   "EDDYCJY",
			Age:    i,
			Avatar: "https://github.com/EDDYCJY",
			Type:   "Person",
		})
	}

	return
}

func ForStruct(p []Person, count int) {
	for i := 0; i < count; i++ {
		_, _ = i, p[i]
	}
}

func ForRangeStruct(p []Person) {
	for i, v := range p {
		_, _ = i, v
	}
}

func JsonToStruct(data []byte, againPerson []AgainPerson) ([]AgainPerson, error) {
	err := json.Unmarshal(data, &againPerson)
	return againPerson, err
}

func JsonIteratorToStruct(data []byte, againPerson []AgainPerson) ([]AgainPerson, error) {
	// var jsonIter = jsoniter.
	// err := json
	// return againPerson, err
}
