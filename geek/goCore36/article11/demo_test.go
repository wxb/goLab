package article11_test



import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type Persion interface {
	SayHi(name chan<- string)
}

type Student struct{}

func (s *Student) SayHi(name chan<- string) {

}

func TestDemo01(t *testing.T) {
	sWang := make(chan string, 1)
	sWang <- "wang xiaobo"
	li := &Student{}
	li.SayHi(sWang)
}

func TestDemo02(t *testing.T) {
	example1()
	example2()
}

// 示例1。
func example1() {
	// 准备好几个通道。
	intChannels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	// 随机选择一个通道，并向它发送元素值。
	rand.Seed(time.Now().Unix())
	index := rand.Intn(3)
	fmt.Printf("The index: %d\n", index)
	intChannels[index] <- index
	// 哪一个通道中有可取的元素值，哪个对应的分支就会被执行。
	var elem0 int
	select {
	case <-intChannels[0]:
		fmt.Println("The first candidate case is selected.")
	case elem0 = <-intChannels[1]:
		fmt.Println("The second candidate case is selected.", elem0)
	case elem := <-intChannels[2]:
		fmt.Printf("The third candidate case is selected, the element is %d.\n", elem)
	default:
		fmt.Println("No candidate case is selected!")
	}
}

// 示例2。
func example2() {
	intChan := make(chan int, 1)
	// 一秒后关闭通道。
	time.AfterFunc(time.Second, func() {
		close(intChan)
	})
	select {
	case _, ok := <-intChan:
		if !ok {
			fmt.Println("The candidate case is closed.")
			break
		}
		fmt.Println("The candidate case is selected.")
	}
}

var channels = [3]chan int{
	nil,
	make(chan int),
	nil,
}

var numbers = []int{1, 2, 3}

func TestDemo03(t *testing.T) {
	// ch := make(chan int, 2)
	select {
	case getChan(0) <- getNumber(0):
		fmt.Println("The first candidate case is selected.")
	case getChan(1) <- getNumber(1):
		fmt.Println("The second candidate case is selected.")
	case getChan(2) <- getNumber(2):
		fmt.Println("The third candidate case is selected")
	// case ch <- 1:
	// 	fmt.Println("The ch candidate case is selected")
	default:
		fmt.Println("No candidate case is selected!")
	}
}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]
}

func getChan(i int) chan int {
	fmt.Printf("channels[%d]\n", i)
	return channels[i]
}
