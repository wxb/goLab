package memo3

import "sync"

type handleFunc func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// Memo 处理方法和数据缓存结构体
type Memo struct {
	f     handleFunc
	m     sync.Mutex
	cache map[string]result
}

// Get 从缓存或者请求结果中获取指定结果
func (m *Memo) Get(k string) (interface{}, error) {
	m.m.Lock()
	res, ok := m.cache[k]
	m.m.Unlock()
	if !ok {
		res.value, res.err = m.f(k)
		m.m.Lock()
		m.cache[k] = res
		m.m.Unlock()
	}

	return res.value, res.err
}

// New 生成一个Memo结构体指针
func New(f handleFunc) *Memo {
	return &Memo{
		f:     f,
		cache: map[string]result{},
	}
}
