package main

import "fmt"

type Persion interface {
	SayHi(name chan<- string)
}

type Student struct{}

func (s *Student) SayHi(name chan<- string) {

}

func main() {
	sWang := make(chan string, 1)
	sWang <- "wang xiaobo"
	sLI := Student{}
	say := sLI.SayHi(sWang)
	fmt.Println(<-say)
}
