package bstset

import "DataStructAndAlgorithm/bst/bst"

type BstSet struct {
	bst *bst.Bst
}

// 新建BstSet对象实例
func NewBstSet() *BstSet {
	return &BstSet{bst: bst.NewBst()}
}
// 添加集合元素
func (s BstSet) Add(e interface{}) {
	s.bst.Add(e)
}

// 移除集合元素
func (s BstSet) Remove(e interface{}){
	s.bst.Remove(e)
}

// 判断集合是否包含e元素
func (s BstSet) Contains(e interface{}) bool {
	return s.bst.Contains(e)
}

// 获取集合元素个数
func (s BstSet) GetSize() int {
	return s.bst.GetSize()
}

// 判断集合是否为空
func (s BstSet)	IsEmpty() bool {
	return s.bst.IsEmpty()
}
// 基于底层bst前序遍历实现的String(),格式化输出集合中的元素
func (s BstSet) String() string{
	s.bst.PreOrder()
	return ""
}
