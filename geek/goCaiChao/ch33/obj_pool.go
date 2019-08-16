package ch33

import (
	"errors"
	"time"
)

// ReusableObj 池链接
type ReusableObj struct {
}

// ObjPool 连接池
type ObjPool struct {
	bufChan chan *ReusableObj
}

// NewObjPool 获取一个连接池
func NewObjPool(numObj int) *ObjPool {
	pool := ObjPool{
		bufChan: make(chan *ReusableObj, numObj),
	}

	for i := 0; i < numObj; i++ {
		pool.bufChan <- &ReusableObj{}
	}

	return &pool
}

// Obj 从连接池取得一个连接
func (pool *ObjPool) Obj(timeout time.Duration) (*ReusableObj, error) {

	select {
	case obj := <-pool.bufChan:
		return obj, nil
	case <-time.After(timeout):
		return nil, errors.New("Time out")
	}
}

// Release 释放一个连接
func (pool *ObjPool) Release(obj *ReusableObj) error {
	select {
	case pool.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}
