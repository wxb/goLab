package memo5

import (
	"fmt"
	"time"
)

type handleFunc func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

type request struct {
	key      string
	response chan<- result
}

// Memo 请求池
type Memo struct {
	RequestPool chan request
}

var (
	maxReqPool = 20
)

// NewMemo 实例化一个Memo结构体指针实例
func NewMemo(f handleFunc) *Memo {
	m := &Memo{
		RequestPool: make(chan request),
		// RequestPool: make(chan request, maxReqPool),
	}

	go m.server(f)
	return m
}

// Get 获取结果
func (m *Memo) Get(k string) (interface{}, error) {
	rspCh := make(chan result)
	start := time.Now()
	m.RequestPool <- request{key: k, response: rspCh}

	fmt.Printf("--1 %s\n", time.Since(start))

	start2 := time.Now()
	res := <-rspCh
	fmt.Printf("--2 %s\n", time.Since(start2))
	return res.value, res.err
}

func (m *Memo) server(f handleFunc) {
	cache := make(map[string]*entry)
	for req := range m.RequestPool {
		time.Sleep(500 * time.Millisecond)
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(req.key, f)
		}

		go e.deliver(req.response)
	}
}

// Close 关闭请求池
func (m *Memo) Close() {
	close(m.RequestPool)
}

func (e *entry) call(k string, f handleFunc) {
	e.res.value, e.res.err = f(k)
	close(e.ready)
}

func (e *entry) deliver(rsp chan<- result) {
	<-e.ready
	rsp <- e.res
}
