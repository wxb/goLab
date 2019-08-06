package ch27

import "sync"

// DataProducer 数据生产
func DataProducer(ch chan<- int, wg *sync.WaitGroup) {

}

// DataReceiver 数据接收
func DataReceiver(ch <-chan int, wg *sync.WaitGroup) {

}
