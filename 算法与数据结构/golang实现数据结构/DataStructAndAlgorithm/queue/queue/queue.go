package queue

import (
	"DataStructAndAlgorithm/queue/array"
	"fmt"
)

//在已实现的动态数组基础上通过接口实现队列
type ArrayQueue struct {
	array *array.Array
}

// 新建ArrayQueue对象实例，指定容量
func NewArrayQueue(capacity int) *ArrayQueue {
	arr := array.NewArray(capacity)
	queue := ArrayQueue{array: arr}
	return &queue
}

// 新建ArrayQueue对象实例，使用默认容量
func NewArrayQueue1() *ArrayQueue {
	arr := array.NewArray1()
	queue := ArrayQueue{array: arr}
	return &queue
}

// 元素入队
func (q *ArrayQueue) Enqueue(e interface{}) error {
	return q.array.AddLast(e)
}

// 元素出队
func (q *ArrayQueue) Dequeue() (interface{}, error) {
	return q.array.RemoveFirst()
}

// 查询队首元素
func (q *ArrayQueue) GetFront() (interface{}, error) {
	return q.array.GetFirst()
}

// 获取队列中元素数量
func (q *ArrayQueue) GetSize() int {
	return q.array.GetSize()
}

// 获取队列容量
func (q *ArrayQueue) GetCapacity() int {
	return q.array.GetCapacity()
}

// 判断队列是否为空
func (q *ArrayQueue) IsEmpty() bool {
	return q.array.IsEmpty()
}

// 格式化输出队列元素信息
func (q *ArrayQueue) String() string {
	res := fmt.Sprintf("Queue: front [")
	for i:=0; i<q.GetSize(); i++ {
		data, _ := q.array.Get(i)
		res += fmt.Sprintf("%v", data)
		if i != q.GetSize()-1 {
			res += ", "
		}
	}
	res += "] tail"
	return res
}