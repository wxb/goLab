package memo1

type handleFunc func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// Memo 处理方法和数据缓存结构体
type Memo struct {
	f     handleFunc
	cache map[string]result
}

// Get 从缓存或者请求结果中获取指定结果
func (m *Memo) Get(k string) (interface{}, error) {
	res, ok := m.cache[k]
	if !ok {
		res.value, res.err = m.f(k)
		m.cache[k] = res
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
