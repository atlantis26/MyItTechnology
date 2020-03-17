package linkedqueue

import (
	"errors"
	"fmt"
)

type Node struct {
	e interface{}
	next *Node
}

// linklist有dummyHead，但是没有尾结点，导致出队需要遍历，导致O(n)，这里重新定义，添加尾结点tail，而不复用
type LinkedQueue struct {
	head *Node
	tail *Node
	size int
}

// 新建节点对象
func NewNode(e interface{}, next *Node) *Node {
	return &Node{e: e, next: next}
}

// 获取队列元素个数
func (l *LinkedQueue) GetSize() int {
	return l.size
}

// 判断队列是否为空
func (l *LinkedQueue) IsEmpty() bool {
	return l.size == 0
}

// 新建链表队列对象
func NewLinkedQueue() * LinkedQueue {
	queue := LinkedQueue{}
	return &queue
}

// 元素入队，从链表尾部进行
func (l *LinkedQueue) Enqueue(e interface{}) (err error) {
	if l.tail == nil {
		l.tail = NewNode(e, nil)
		l.head = l.tail
	} else{
		newNode := NewNode(e, nil)
		l.tail.next = newNode
		l.tail = newNode
	}
	l.size ++
	return err
}

// 元素出队, 从链表头部出队
func (l *LinkedQueue) Dequeue() (e interface{}, err error) {
	if l.IsEmpty() {
		return nil, errors.New("cannot dequeue from an empty queue")
	}
	retNode := l.head
	l.head = l.head.next
	retNode.next = nil
	if l.head == nil{
		l.tail = nil
	}
	l.size --
	return  retNode.e, err
}

// 查询队首元素信息
func (l *LinkedQueue) GetFront() (e interface{}, err error){
	if l.IsEmpty() {
		err = errors.New("queue is empty")
		return nil, err
	}
	return l.head.e, err
}

// 格式化输出队列信息
func (l *LinkedQueue) String() string {
	res := "Queue: front "
	cur := l.head
	for cur != nil {
		res += fmt.Sprintf("%v->", cur.e)
		cur = cur.next
	}
	res += "nil tail"
	return res
}
