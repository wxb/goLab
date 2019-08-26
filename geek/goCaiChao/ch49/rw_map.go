package ch49

import "sync"

type RWLockMap struct {
	m  map[interface{}]interface{}
	mu sync.RWMutex
}

func NewRWLockMap() *RWLockMap {
	return &RWLockMap{
		m: make(map[interface{}]interface{}, 0),
	}
}

func (m *RWLockMap) Get(k interface{}) (interface{}, bool) {
	m.mu.RLock()
	v, has := m.m[k]
	m.mu.RUnlock()
	return v, has
}

func (m *RWLockMap) Set(k, v interface{}) {
	m.mu.Lock()
	m.m[k] = v
	m.mu.Unlock()
}

func (m *RWLockMap) Del(k interface{}) {
	m.mu.Lock()
	delete(m.m, k)
	m.mu.Unlock()
}
