package queue

import (
	"errors"
	"fmt"
)

// 循环队列,不再复用自定义的动态数组
type LoopQueue struct {
	data []interface{}
	front int
	tail int
	size int
}

// 新建循环队列对象，指定容量
func NewLoopQueue (capacity int) *LoopQueue {
	queue := LoopQueue{data: make([]interface{}, capacity+1),  front: 0, tail: 0, size: 0}
	return  &queue
}

// 新建循环队列对象，使用默认容量
func NewLoopQueue1() *LoopQueue {
	return NewLoopQueue(10)
}

// 查询队列容量
func (q *LoopQueue) GetCapacity() int {
	return len(q.data) - 1
}

// 查询队列元素个数
func (q *LoopQueue) GetSize() int {
	return q.size
}

// 扩容
func (q *LoopQueue) resize(capacity int) {
	newData := make([]interface{}, capacity + 1)
	for i:=0; i < q.size; i++ {
		newData[i] = q.data[(i+ q.front) % len(q.data)]
	}
	q.data = newData
	q.front = 0
	q.tail = q.size
}

// 查询队列是否为空
func (q *LoopQueue) IsEmpty() bool {
	return q.front == q.tail
}

// 元素入队
func (q *LoopQueue) Enqueue(e interface{}) error {
	// 循环队列有意浪费一个位置，如果tail + 1 % len(q.data) == q.front，则表明队列已满，不能再存元素，需要扩容
	if (q.tail+1) % len(q.data) == q.front {
		q.resize(q.GetCapacity() << 1)
	}
	q.data[q.tail] = e
	q.tail = (q.tail + 1) % len(q.data)
	q.size ++
	return nil
}

// 元素出队
func (q *LoopQueue) Dequeue () (data interface{}, err error) {
	if q.IsEmpty() {
		err = errors.New("cannot dequeue from an empty queue")
		return nil, err
	}
	data = q.data[q.front]
	q.data[q.front] = nil
	q.front = (q.front + 1) % len(q.data)
	q.size --
	// 缩容
	if q.size == q.GetCapacity() >> 2 && q.GetCapacity() >> 1 != 0 {
		q.resize(q.GetCapacity() >> 1)
	}
	return data, err
}

// 查询队首元素
func (q *LoopQueue) GetFront() (data interface{}, err error){
	if q.IsEmpty() {
		err = errors.New("queue is empty")
		return nil, err
	}
	return q.data[q.front], err
}

// 格式化输出队列元素信息
func (q *LoopQueue) String() string {
	res := fmt.Sprintf("Queue: size =%d, capacity = %d\nfront [", q.size, q.GetCapacity())
	for i:=0; i < q.size; i++ {
		res += fmt.Sprintf("%v", q.data[(i+q.front) % len(q.data)])
		if (i+ q.front) % len(q.data) != q.tail {
			res += ", "
		}
	}
	res += "] tail"
	return res
}