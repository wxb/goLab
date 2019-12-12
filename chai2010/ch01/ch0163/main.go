package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/wxb/goLab/chai2010/ch01/ch0163/pubsub"
)

func main() {
	p := pubsub.NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	// 订阅所有主题
	all := p.Subscribe()
	go func() {
		for msg := range all {
			fmt.Println("all:", msg)
		}
	}()

	// 订阅golang主题
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})
	go func() {
		for msg := range golang {
			fmt.Println("golang:", msg)
		}
	}()

	// 发布主题
	go func() {
		rand.Seed(time.Now().UnixNano())
		content := [...]string{"c", "world", "golang", "php", "java", "python"}
		str := strings.Builder{}
		for {
			str.WriteString("hello,")
			str.WriteString(content[rand.Intn(len(content))])
			p.Publish(str.String())
			str.Reset()

			time.Sleep(1 * time.Second)
		}
	}()

	// 信号控制执行
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	s := <-ch
	fmt.Println("Got signal:", s)
}
