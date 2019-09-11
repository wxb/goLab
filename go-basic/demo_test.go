package basic_test

import (
	"fmt"
	"testing"
	"time"
)

func TestDemo5_1(t *testing.T) {
	var numbers = [3]int{1, 2, 3}
	_ = numbers
	//fmt.Printf("%d length is %d", numbers[0], len(numbers))
	var numbers2 [5]int
	numbers2[0] = 2
	numbers2[3] = numbers2[0] - 3
	numbers2[1] = numbers2[2] + 5
	numbers2[4] = len(numbers2)
	sum := (11)
	// “==”用于两个值的相等性判断
	fmt.Printf("%v\n", (sum == numbers2[0]+numbers2[1]+numbers2[2]+numbers2[3]+numbers2[4]))
}

func TestDemo5_2(t *testing.T) {
	//var number = []int{1, 2, 3, 4, 5}
	//var slice = number[1:4]
	//fmt.Printf("%d - %d - %v ", number[4], slice[2], 4==cap(slice))

	var numbers3 = [6]int{1, 2, 3, 4, 5, 6}
	//slice3 := numbers3[2:len(numbers3)] //[2,5)
	slice3 := numbers3[2:5:5] //[2,5)
	length := (3)
	capacity := (3)
	fmt.Printf("%v, %v\n", (length == len(slice3)), (capacity == cap(slice3)))
}

func TestDemo5_3(t *testing.T) {
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice5 := numbers4[4:6:8]
	length := (2)
	capacity := (4)
	fmt.Printf("%v, %v\n", length == len(slice5), capacity == cap(slice5))
	slice5 = slice5[:cap(slice5)]
	slice5 = append(slice5, 11, 12, 13)
	length = (7)
	fmt.Printf("%v\n", length == len(slice5))
	slice6 := []int{0, 0, 0}
	copy(slice5, slice6)
	e2 := (0)
	e3 := (8)
	e4 := (11)
	fmt.Printf("%v, %v, %v\n", e2 == slice5[2], e3 == slice5[3], e4 == slice5[4])
	fmt.Printf("%v\n", numbers4[9])
	fmt.Print("hello world\n")
	fmt.Println("hello world")
}

func TestDemo5_4(t *testing.T) {
	var str string = "hello golang"
	mm2 := map[string]int{"golang": 42, "java": 1, "python": 8}
	mm2["scala"] = 25
	mm2["erlang"] = 50
	mm2["python"] = 0
	fmt.Printf("%d, %d, %d \n", mm2["scala"], mm2["erlang"], mm2["python"])
	fmt.Printf("%v", str)
}

func TestDemo5_5(t *testing.T) {
	ch2 := make(chan string, 1)
	// 下面就是传说中的通过启用一个Goroutine来并发的执行代码块的方法。
	// 关键字 go 后跟的就是需要被并发执行的代码块，它由一个匿名函数代表。
	// 对于 go 关键字以及函数编写方法，我们后面再做专门介绍。
	// 在这里，我们只要知道在花括号中的就是将要被并发执行的代码就可以了。
	go func() {
		ch2 <- ("已达到！")
	}()
	var value string = "数据"
	value = value + (<-ch2)
	fmt.Println(value)
}

type Sender chan<- int

type Receiver <-chan int

func TestDemo5_6(t *testing.T) {
	var myChannel = make(chan int, (0))
	var number = 6
	go func() {
		var sender Sender = myChannel
		sender <- number
		fmt.Println("Sent!")
	}()
	go func() {
		var receiver Receiver = myChannel
		fmt.Println("Received!", <-receiver)
	}()
	// 让main函数执行结束的时间延迟1秒，
	// 以使上面两个代码块有机会被执行。
	time.Sleep(time.Second)
}
