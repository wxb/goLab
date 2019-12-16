package main

type gate chan bool

func (g gate) enter() { g <- true }

func (g gate) leave() { <-g }

type tree struct {
	path string
	gate
}

// Request 执行网页请求抓取
func Request() {

}

func main() {

}
