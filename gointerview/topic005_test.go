package gointerview_test

import (
	"fmt"
	"testing"
)

type People interface {
	Speak(string) string
}

type stu struct{}

func (s *stu) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func TestTopic005(t *testing.T) {
	var peo People = &stu{}
	think := "bitch"
	fmt.Println(peo.Speak(think))

}
