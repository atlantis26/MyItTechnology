package linkedlist

import (
	"errors"
	"fmt"
	"DataStructAndAlgorithm/utils"
)

type LinkedList struct {
	dummyHead *Node // 虚拟头节点
	size int
}

type Node struct {
	e interface{}
	next *Node
}

// 新建节点,提供节点值
func NewNode(e interface{}, next *Node) *Node {
	return &Node{e: e, next: next}
}

// 新建节点,节点值默认为空值
func NewNode1() *Node {
	return &Node{e: nil, next: nil}
}

// 格式化输出节点信息
func (n *Node) String() string {
	return fmt.Sprintf("%v", n.e)
}

// 新建链表时，创建虚拟头节点dummyHead，dummyHead节点对用户不可见，不计入size
func NewLinkedList()  *LinkedList {
	return &LinkedList{dummyHead: NewNode(nil, nil), size:0}
}

// 获取链表中的元素个数
func (l * LinkedList) GetSize() int {
	return l.size
}

// 链表是否为空
func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

// 在链表的'index'位置添加元素;关键找到添加的节点的前一个节点；
func (l *LinkedList) Add(index int, e interface{}) (err error) {
	if index < 0 || index > l.size {
		err = errors.New("add failed, Illegal index")
		return err
	}
	prev := l.dummyHead
	for i:=0; i < index; i++ {
		prev = prev.next
	}
	// 新节点的next = prev节点的next; prev节点的next = 新节点
	prev.next = NewNode(e, prev.next)
	l.size ++
	return err
}

// 在链表头添加元素，(即在dummyHead节点后添加元素)
func (l *LinkedList) AddFirst(e interface{}) (err error) {
	err = l.Add(0, e)
	return err
}

// 在链表的末尾添加元素
func (l *LinkedList) AddLast(e interface{}) (err error) {
	err = l.Add(l.size, e)
	return err
}

// 查询链表的第index位置的元素的值
func (l *LinkedList) Get(index int) (e interface{}, err error) {
	if index < 0 || index >= l.size {
		err = errors.New("get failed, Illegal index")
		return nil, err
	}
	// 遍历链表，需知道链表的size值
	cur := l.dummyHead.next
	for i:=0; i < index; i++ {
		cur = cur.next
	}
	return cur.e, err
}

// 查询链表第一个元素
func (l *LinkedList) GetFirst() (e interface{}, err error) {
	return l.Get(0)
}

// 查询链表最后一个元素
func (l *LinkedList) GetLast() (e interface{}, err error) {
	return l.Get(l.size - 1)
}

// 修改链表第index位置的元素;
func (l *LinkedList) Set(index int, e interface{}) (err error) {
	if index < 0 || index > l.size {
		err = errors.New("set failed, Illegal index")
		return err
	}
	cur := l.dummyHead.next
	for i:=0; i < index; i++ {
		cur = cur.next
	}
	cur.e = e

	return err
}

// 查找链表是否存在元素e
func (l *LinkedList) Contains(e interface{}) bool {
	cur := l.dummyHead.next
	for cur != nil {
		if cur.e == e {
			return true
		}
		cur = cur.next
	}
	return false
}

// 删除链表中index位置的元素
func (l *LinkedList) Remove(index int) (data interface{}, err error) {
	if index < 0 || index > l.size - 1 {
		err = errors.New("remove failed, Illegal index")
		return nil, err
	}
	prev := l.dummyHead
	for i:=0; i < index; i++ {
		prev = prev.next
	}
	delNode := prev.next
	prev.next = delNode.next
	delNode.next = nil
	l.size --
	return delNode.e, err
}

// 删除链表第一个位置的元素
func (l *LinkedList) RemoveFirst() (data interface{}, err error) {
	return l.Remove(0)
}

// 删除链表末尾最后一个元素
func (l *LinkedList) RemoveLast() (data interface{}, err error) {
	return l.Remove(l.size - 1)
}

// 从链表中删除元素e
func (l *LinkedList) RemoveElement(e interface{}) {
	prev := l.dummyHead
	for prev.next != nil {
		if utils.Compare(prev.next.e, e) == 0 {
			break
		}
		prev = prev.next
	}
	if prev.next != nil {
		delNode := prev.next
		prev.next = delNode.next
		delNode.next = nil
	}
}

// 格式化输出链表信息
func (l *LinkedList) String() string {
	res := ""
	cur := l.dummyHead.next
	for cur != nil {
		res += fmt.Sprintf("%v->", cur.e)
		cur = cur.next
	}
	res += "nil"
	return res
}
