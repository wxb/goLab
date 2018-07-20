package main

import "fmt"

func main() {
	// >0 有缓存，=0 无缓存，阻塞
	c := make(chan int, 1) //修改2为1就报错，修改2为3可以正常运行
	c <- 1
	fmt.Println(<-c)
	c <- 2
	fmt.Println(<-c)
}

//修改为1报如下的错误:
//fatal error: all goroutines are asleep - deadlock!
