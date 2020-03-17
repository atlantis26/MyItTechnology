package main

import (
	"DataStructAndAlgorithm/linkedlist/linkedlist"
	"DataStructAndAlgorithm/linkedlist/linkedqueue"
	"DataStructAndAlgorithm/linkedlist/linkedstack"

	"fmt"
)

type Stack interface {
	Push(e interface{}) (err error)
	Pop() (e interface{}, err error)
	Peek() (e interface{}, err error)
	IsEmpty() bool
	GetSize() int
}

type Queue interface {
	Enqueue(e interface{}) (err error)
	Dequeue() (e interface{}, err error)
	GetFront() (e interface{}, err error)
	GetSize() int
	IsEmpty() bool
}

func main() {
	// 测试链表
	linkList := linkedlist.NewLinkedList()
	for i:=0; i < 5; i++ {
		linkList.AddFirst(i)
		fmt.Println(linkList)
	}
	linkList.Add(2, 666)
	fmt.Println(linkList)
	linkList.Remove(2)
	fmt.Println(linkList)
	linkList.RemoveFirst()
	fmt.Println(linkList)
	linkList.RemoveLast()
	fmt.Println(linkList)

	// 测试基于链表实现的栈
	var st  Stack
	linkStackInstance, _ := linkedstack.NewLinkedListStack()
	st = linkStackInstance
	for i:=0; i < 5; i++ {
		st.Push(i)
		fmt.Println(st)
	}
	st.Pop()
	fmt.Println(st)
	st.Pop()
	fmt.Println(st)

	// 测试基于链表实现的队列
	var q Queue
	linkQueueInstance := linkedqueue.NewLinkedQueue()
	q = linkQueueInstance
	for i:=0; i < 5; i++ {
		q.Enqueue(i)
		fmt.Println(q)
	}
	q.Dequeue()
	fmt.Println(q)
	q.Dequeue()
	fmt.Println(q)
}