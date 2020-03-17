package linkedset

import (
	"DataStructAndAlgorithm/linkedlist/linkedlist"
)

type LinkedSet struct {
	link *linkedlist.LinkedList
}

// 新建LinkedSet对象实例
func NewLinkedSet() *LinkedSet {
	return &LinkedSet{link: linkedlist.NewLinkedList()}
}
// 添加集合元素
func (s LinkedSet) Add(e interface{}) {
	if ! s.link.Contains(e) {
		s.link.AddFirst(e)
	}
}

// 移除集合元素
func (s LinkedSet) Remove(e interface{}){
	s.link.RemoveElement(e)
}

// 判断集合是否包含e元素
func (s LinkedSet) Contains(e interface{}) bool {
	return s.link.Contains(e)
}

// 获取集合元素个数
func (s LinkedSet) GetSize() int {
	return s.link.GetSize()
}

// 判断集合是否为空
func (s LinkedSet)	IsEmpty() bool {
	return s.link.IsEmpty()
}
// 格式化输出集合中的元素
func (s LinkedSet) String() string{
	return s.link.String()
}