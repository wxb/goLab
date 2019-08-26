package ch49

import (
	"github.com/easierway/concurrent_map"
)

type ConcurrentMap struct {
	cm *concurrent_map.ConcurrentMap
}

func NewConcurrentMap(numOfPartitions int) *ConcurrentMap {
	conMap := concurrent_map.CreateConcurrentMap(numOfPartitions)
	return &ConcurrentMap{conMap}
}

func (m *ConcurrentMap) Get(k interface{}) (interface{}, bool) {
	return m.cm.Get(concurrent_map.StrKey(k.(string)))
}

func (m *ConcurrentMap) Set(k, v interface{}) {
	m.cm.Set(concurrent_map.StrKey(k.(string)), v)
}

func (m *ConcurrentMap) Del(k interface{}) {
	m.cm.Del(concurrent_map.StrKey(k.(string)))
}
