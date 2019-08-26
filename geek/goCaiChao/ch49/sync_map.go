package ch49

import "sync"

type SyncMap struct {
	m sync.Map
}

func NewSyncMap() *SyncMap {
	return &SyncMap{}
}

func (m *SyncMap) Get(k interface{}) (interface{}, bool) {
	return m.m.Load(k)
}

func (m *SyncMap) Set(k, v interface{}) {
	m.m.Store(k, v)
}

func (m *SyncMap) Del(k interface{}) {
	m.m.Delete(k)
}
