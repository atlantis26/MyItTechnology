package priorityqueue

import "DataStructAndAlgorithm/maxheap/maxheap"

type PriorityQueue struct {
	heap *maxheap.MaxHeap
}

// 构造基于最大堆实现的优先队列实例
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{heap:maxheap.NewMaxHeap1()}
}

// 获取优先队列中元素个数
func (q *PriorityQueue) GetSize() int {
	return q.heap.GetSize()
}

// 判断优先队列是否为空
func (q *PriorityQueue) IsEmpty() bool {
	return q.heap.IsEmpty()
}

// 优先队列元素入队
func (q *PriorityQueue) Enqueue(e interface{}) {
	q.heap.Add(e)
}

// 优先队列元素出队
func (q *PriorityQueue) Dequeue() (interface{}, error) {
	return q.heap.ExtractMax()
}

// 查询优先队列队首元素
func (q *PriorityQueue) GetFront() (interface{}, error) {
	return q.heap.FindMax()
}