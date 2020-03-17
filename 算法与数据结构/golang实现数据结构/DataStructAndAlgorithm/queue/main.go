package main

import (
	"DataStructAndAlgorithm/queue/queue"
	"fmt"
	"time"
)

type Queue interface {
	Enqueue(interface{}) error
	Dequeue() (interface{}, error)
	GetFront() (interface{}, error)
	GetSize() int
	IsEmpty() bool
}
// 测试数组队列与循环队列的性能差异的方法
func TestQueue(q Queue, opCount int) float64 {
	startTime := time.Now()
	for i:=0; i < opCount; i++ {
		q.Enqueue(i)
	}
	for i:=0; i < opCount; i++ {
		q.Dequeue()
	}
	endTime := time.Now()
	subM := endTime.Sub(startTime).Seconds()
	return subM

}

func main() {
	var q Queue
	// 测试数组队列的性能耗时；出队操作O(n)时间复杂度
	queueInstance := queue.NewArrayQueue(10)
	q = queueInstance
	fmt.Println(TestQueue(q, 100000))
	// 测试循环队列的性能耗时；出队操作0(1)时间复杂度
	loopQueueInstance := queue.NewLoopQueue(10)
	q = loopQueueInstance
	fmt.Println(TestQueue(q, 100000))
	//两种队列的性能差异很大，导致两者差异的主要原因就在出队操作上
}