package ch27_test

import (
	"sync"
	"testing"

	"github.com/wxb/goLab/geek/goCaiChao/ch27"
)

type FiveFunc func(input string) string

func TestChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	ch27.DataProducer(ch, &wg)

	wg.Add(1)
	ch27.DataReceiver(ch, &wg)

	wg.Wait()
}

func TestChannelClose(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	ch27.DataCloseProducer(ch, &wg)

	wg.Add(1)
	ch27.DataForeceiver(ch, &wg)

	wg.Add(1)
	ch27.DataRangeReceiver(ch, &wg)

	wg.Wait()

	// 1. 对一个已关闭的通道再赋值时会panic
	// 2. 对一个已关闭的通道再取值时会得到对应的类型零值
	d := <-ch
	t.Log(d)

}

func TestFiveReferType(t *testing.T) {
	var m map[string]string
	mm := make(map[string]string)

	t.Log(m, mm, m == nil, mm == nil)

	var a [1]int
	t.Log(a)

	var s []string
	ss := make([]string, 0)
	t.Log(s, ss, s == nil, ss == nil, &s, &ss)

	var ch chan int
	chh := make(chan int)
	t.Log(ch, chh, ch == nil, chh == nil)

	var f FiveFunc
	ff := func(i string) string { return i }
	t.Log(f, f == nil, ff == nil)
}
