package linkedmap

import (
	"errors"
	"fmt"
)

type Node struct {
	key interface{}
	value interface{}
	next *Node
}

type LinkedMap struct {
	dummyHead *Node
	size  int
}

// 新建Node节点对象
func NewNode(key interface{}, value interface{}, next *Node) *Node {
	return &Node{key: key, value: value, next: next}
}

// 新建LinkedMap实例对象
func NewLinkedMap() *LinkedMap {
	return &LinkedMap{dummyHead: NewNode(nil, nil, nil), size: 0}
}

// 获取关于k的节点
func (m *LinkedMap)  getNode(key interface{}) *Node {
	cur := m.dummyHead.next
	for cur != nil {
		if cur.key == key {
			return cur
		}
		cur = cur.next
	}
	return nil
}

// 映射新增键值对元素
func (m *LinkedMap) Add(key interface{}, value interface{}) {
	node := m.getNode(key)
	if node == nil {
		m.dummyHead.next = NewNode(key, value, m.dummyHead.next)
		m.size ++
	} else {
		node.value = value
	}
}

// 映射删除元素
func (m *LinkedMap) Remove(key interface{}) interface{} {
	prev := m.dummyHead
	for prev.next != nil {
		if prev.next.key == key {
			break
		}
		prev = prev.next
	}
	if prev.next != nil {
		delNode := prev.next
		prev.next = delNode.next
		delNode.next = nil
		m.size --
		return delNode.value
	}
	return nil
}

// 映射是否包含元素k
func (m *LinkedMap) Contains(key interface{}) bool {
	return m.getNode(key) != nil
}

// 映射获取k的value值
func (m *LinkedMap) Get(key interface{}) interface{} {
	node := m.getNode(key)
	if node == nil {
		return nil
	} else {
		return node.value
	}
}

// 映射修改k的value值
func (m *LinkedMap) Set(key interface{}, value interface{}) (err error) {
	node := m.getNode(key)
	if node == nil {
		err = errors.New(fmt.Sprintf("%v doesn't exist", key))
	} else {
		node.value = value
	}
	return err
}

// 获取映射中的键值对个数
func (m *LinkedMap) GetSize() int {
	return m.size
}

// 判断映射是否为空
func (m *LinkedMap) IsEmpty() bool {
	return m.size == 0
}


