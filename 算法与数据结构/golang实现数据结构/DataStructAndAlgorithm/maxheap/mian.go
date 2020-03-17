package main

import (
	"DataStructAndAlgorithm/maxheap/maxheap"
	"DataStructAndAlgorithm/maxheap/priorityqueue"
	"fmt"
)

type MaxHeap interface {
	GetSize() int
	IsEmpty() bool
	Parent(index int) (int, error)
	LeftChild(index int) int
	RightChild(index int) int
	Add(e interface{})
	ExtractMax() (ret interface{}, err error)
	Replace(e interface{}) interface{}
}

type Queue interface {
	GetSize() int
	IsEmpty() bool
	Enqueue(e interface{})
	Dequeue() (interface{}, error)
	GetFront() (interface{}, error)
}

func TestArrayMaxHeap() {
	var heap MaxHeap
	n := 10
	arrayMaxHeapInstance := maxheap.NewMaxHeap(n)
	heap = arrayMaxHeapInstance
	for i:=0; i < n; i++ {
		heap.Add(i)
	}
	arr := make([]int, n)
	for i:=0; i < n; i++ {
		t, err := heap.ExtractMax()
		fmt.Println(t)
		if nil != err {
			fmt.Println(err)
		}
		arr[i] = t.(int)
	}
	for i:=1; i < n -1; i++ {
		if arr[i-1] < arr[i] {
			fmt.Println("error")
		}
	}
	fmt.Println("test maxHeap completed")
}


func TestHeapifyMaxHeap() {
	var heap MaxHeap
	n := 10
	testArr := []interface{} {1, 3, 9, 7, 5, 9, 2, 4, 6, 0}
	arrayMaxHeapInstance := maxheap.HeapifyNewMaxHeap(testArr)
	heap = arrayMaxHeapInstance

	arr := make([]int, n)
	for i:=0; i < n; i++ {
		t, err := heap.ExtractMax()
		fmt.Println(t)
		if nil != err {
			fmt.Println(err)
		}
		arr[i] = t.(int)
	}
	for i:=1; i < n -1; i++ {
		if arr[i-1] < arr[i] {
			fmt.Println("error")
		}
	}
	fmt.Println(arr)
	fmt.Println("test maxHeap completed")
}

func TestPriorityQueue() {
	var q Queue
	priorityQueueInstance := priorityqueue.NewPriorityQueue()
	q = priorityQueueInstance
	testArr := []interface{} {1, 3, 9, 7, 5, 9, 2, 4, 6, 0}
	for i:=0; i < len(testArr); i++ {
		q.Enqueue(testArr[i])
	}
	for i:=0; i < len(testArr); i++ {
		e,err := q.Dequeue()
		fmt.Println(e, err)
	}

}

func main() {
	TestArrayMaxHeap()
	TestHeapifyMaxHeap()
	TestPriorityQueue()
}
