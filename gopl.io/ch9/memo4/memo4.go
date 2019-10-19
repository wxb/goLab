package memo4

import (
	"fmt"
	"sync"
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

// Memo 处理方法和数据缓存结构体
type Memo struct {
	f     handleFunc
	m     sync.Mutex
	cache map[string]*entry
}

// Get 从缓存或者请求结果中获取指定结果
func (m *Memo) Get(k string) (interface{}, error) {
	m.m.Lock()
	e := m.cache[k]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		m.cache[k] = e
		m.m.Unlock()

		e.res.value, e.res.err = m.f(k)
		fmt.Println("--", k, time.Now().UnixNano())
		close(e.ready)
	} else {
		m.m.Unlock()

		<-e.ready
		fmt.Println("++", k, time.Now().UnixNano())
	}

	return e.res.value, e.res.err
}

// New 生成一个Memo结构体指针
func New(f handleFunc) *Memo {
	return &Memo{
		f:     f,
		cache: map[string]*entry{},
	}
}
