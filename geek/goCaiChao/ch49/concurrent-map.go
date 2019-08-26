package ch49

import cmap "github.com/orcaman/concurrent-map"

type CMap struct {
	m cmap.ConcurrentMap
}

func NewCMap() *CMap {
	return &CMap{
		m: cmap.New(),
	}
}

func (m *CMap) Get(k interface{}) (interface{}, bool) {
	return m.m.Get(k.(string))
}

func (m *CMap) Set(k, v interface{}) {
	m.m.Set(k.(string), v)
}

func (m *CMap) Del(k interface{}) {
	m.m.Remove(k.(string))
}
