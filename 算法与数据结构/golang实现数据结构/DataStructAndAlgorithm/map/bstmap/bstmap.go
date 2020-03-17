package bstmap

import (
	"DataStructAndAlgorithm/utils"
	"errors"
	"fmt"
)

type Node struct {
	key interface{}
	value interface{}
	left *Node
	right *Node
}

type BstMap struct {
	root *Node
	size  int
}

// 新建Node节点对象
func NewNode(key interface{}, value interface{}, left *Node, right *Node) *Node {
	return &Node{key: key, value: value, left: left, right: right}
}

// 新建BstMap实例对象
func NewBstMap() *BstMap {
	return &BstMap{root: nil, size: 0}
}

// 获取关于k的节点,使用前序遍历查找
func (m *BstMap) getNode(node *Node, key interface{}) *Node {
	if node == nil {
		return nil
	}
	if utils.Compare(key, node.key) == 0 {
		return node
	} else if utils.Compare(key, node.key) < 0 {
		return m.getNode(node.left, key)
	} else {
		return m.getNode(node.right, key)
	}
}

// 映射新增键值对元素
func (m *BstMap) Add(key interface{}, value interface{}) {
	m.root = m.add(m.root, key, value)
}

// 向以Node为根节点的插入数据元素节点(key, value)，返回插入新节点后的二分搜索树的根，递归算法
func (m *BstMap) add(node *Node, key interface{}, value interface{}) *Node {
	if node == nil {
		m.size ++
		return NewNode(key, value, nil, nil)
	}
	if utils.Compare(key, node.key) < 0 {
		node.left = m.add(node.left, key, value)
	} else if utils.Compare(key, node.key) > 0 {
		node.right = m.add(node.right, key, value)
	} else {
		node.value = value
	}
	return  node
}

// 映射删除元素
func (m *BstMap) Remove(key interface{}) interface{} {
	node := m.getNode(m.root, key)
	if node != nil {
		m.root = m.remove(m.root, key)
		return node.value
	}
	return nil
}

// 删除掉以node为根的二分搜索树中键为key的节点，递归算法
// 返回删除节点后最新的二分搜索树的根
func (m *BstMap) remove(node *Node, key interface{}) *Node {
	if node == nil {
		return nil
	}
	if utils.Compare(key, node.key) < 0 {
		node.left = m.remove(node.left, key)
	} else if utils.Compare(key, node.key) > 0 {
		node.right = m.remove(node.right, key)
	} else {
		// 待删除节点的left子树为空的情况
		if node.left == nil {
			rightNode := node.right
			node.right = nil
			m.size --
			return rightNode
		}
		if node.right == nil {
			leftNode := node.left
			node.left = nil
			m.size --
			return leftNode
		}
		// 待删除节点的左右子树均不为空的情况
		// 找到比待删除节点大的最小节点，即待删除节点右子树的最小节点
		successor := m.minimum(node.right)
		successor.right = m.removeMin(node.right)
		successor.left = node.left
		node.left = nil
		node.right = nil
	}
	return node
}

// 找到bst树中的最小节点，即最左边的节点
func (m *BstMap) minimum(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return m.minimum(node.left)
}

// 删除bst树中的最小节点, 返回删除后的根节点
func (m *BstMap) removeMin(node *Node) *Node{
	if node.left == nil {
		rightNode := node.right
		node.right = nil
		m.size --
		return rightNode
	}
	node.left = m.removeMin(node.left)
	return node
}

// 映射是否包含元素k
func (m *BstMap) Contains(key interface{}) bool {
	return m.getNode(m.root, key) != nil
}

// 映射获取k的value值
func (m *BstMap) Get(key interface{}) interface{} {
	node := m.getNode(m.root, key)
	if node == nil {
		return nil
	} else {
		return node.value
	}
}

// 映射修改k的value值
func (m *BstMap) Set(key interface{}, value interface{}) (err error) {
	node := m.getNode(m.root, key)
	if node == nil {
		err = errors.New(fmt.Sprintf("%v doesn't exist", key))
	} else {
		node.value = value
	}
	return err
}

// 获取映射中的键值对个数
func (m *BstMap) GetSize() int {
	return m.size
}

// 判断映射是否为空
func (m *BstMap) IsEmpty() bool {
	return m.size == 0
}
