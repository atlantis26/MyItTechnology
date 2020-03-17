package linkedstack

import (
	"DataStructAndAlgorithm/linkedlist/linkedlist"
	"errors"
	"fmt"
)

// 使用链表实现栈数据结构
type LinkedListStack struct {
	link *linkedlist.LinkedList
	capacity int  // capacity等于0则标示不限容量，大于0则标示限容
}

// 新建LinkedListStack对象,不限制栈容量
func NewLinkedListStack() (st *LinkedListStack, err error){
	stack := LinkedListStack{link: linkedlist.NewLinkedList(), capacity: 0}
	return &stack, err
}

// 新建LinkedListStack对象,限制栈容量
func NewLinkedListStack1(capacity int) (st *LinkedListStack, err error){
	if capacity < 0 {
		err = errors.New("capacity must bigger than 0")
		return st, err
	}
	stack := LinkedListStack{link: linkedlist.NewLinkedList(), capacity: capacity}
	return &stack, err
}

// 入栈
func (st * LinkedListStack) Push(e interface{}) (err error) {

	if st.capacity != 0  &&  st.link.GetSize() == st.capacity{
		err = errors.New("push failed, stack already full")
		return err
	}
	return st.link.AddFirst(e)
}

// 出栈
func (st *LinkedListStack) Pop() (e interface{}, err error) {
	return st.link.RemoveFirst()
}

// 获取栈顶元素信息
func (st *LinkedListStack) Peek() (e interface{}, err error) {
	return st.link.GetFirst()
}

// 获取栈是否为空
func (st *LinkedListStack) IsEmpty() bool {
	return st.link.IsEmpty()
}

// 获取栈元素的个数
func (st *LinkedListStack) GetSize() int {
	return st.link.GetSize()
}

// 格式化输出栈信息
func (st *LinkedListStack) String() string {
	res := "stack: top ["
	for i:=0; i < st.GetSize(); i++ {
		e, _ := st.link.Get(i)
		res += fmt.Sprintf("%v", e)
		if i != st.GetSize() - 1 {
			res += ", "
		}
	}
	res += "]"
	return res
}
