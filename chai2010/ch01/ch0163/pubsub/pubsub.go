package pubsub

import (
	"errors"
	"log"
	"sync"
	"time"
)

var (
	errSendTimeOut = errors.New("Send Topic Timeout")
)

type (
	subscriber chan interface{}         // 订阅者为一个管道
	topicFunc  func(v interface{}) bool // 主题过滤器
)

// Publisher 发布者
type Publisher struct {
	m          sync.Mutex
	buffer     int
	timeout    time.Duration
	subscriber map[subscriber]topicFunc
}

// NewPublisher 构建一个发布者对象
func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:     buffer,
		timeout:    publishTimeout,
		subscriber: make(map[subscriber]topicFunc),
	}
}

// SubscribeTopic 通过一个订阅器订阅一个主题
func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscriber[ch] = topic
	p.m.Unlock()

	return ch
}

// Subscribe 定义全部主题
func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

// Evict 退出订阅
func (p *Publisher) Evict(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	delete(p.subscriber, sub)
	close(sub)
}

// Publish 发布一个主题
func (p *Publisher) Publish(v interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	var wg sync.WaitGroup
	for sub, topic := range p.subscriber {
		wg.Add(1)
		go func(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
			defer wg.Done()
			err := p.sendTopic(sub, topic, v)
			if err == errSendTimeOut {
				log.Println(err.Error())
			}
		}(sub, topic, v, &wg)
	}

	wg.Wait()
}

// Close 关闭发布者
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscriber {
		delete(p.subscriber, sub)
		close(sub)
	}
}

// 发送主题
func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v interface{}) error {
	if topic != nil && !topic(v) {
		return nil
	}

	select {
	case sub <- v:
		return nil
	case <-time.After(p.timeout):
		return errSendTimeOut
	}
}
