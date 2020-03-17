package stack

import (
	"DataStructAndAlgorithm/stack/array"
	"fmt"
)

//在已实现的动态数组基础上通过接口实现栈
 type ArrayStack struct {
	 array *array.Array
 }

// 新建ArrayStack对象实例，指定容量
func NewArrayStack(capacity int) *ArrayStack {
	arr := array.NewArray(capacity)
	stack := ArrayStack{array: arr}
	return &stack
}

// 新建ArrayStack对象实例，使用默认容量
func NewArrayStack1() *ArrayStack {
	arr := array.NewArray1()
	stack := ArrayStack{array: arr}
	return &stack
}

// 向栈中加入新元素
func (st *ArrayStack) Push(e interface{}) error {
	return st.array.AddLast(e)
}

// 从栈中取出栈顶元素
func (st *ArrayStack) Pop() (interface{}, error) {
	return st.array.RemoveLast()
}

// 查询栈顶元素
func (st *ArrayStack) Peek() (interface{}, error){
	return st.array.GetLast()
}

// 查询栈元素个数
func (st *ArrayStack) GetSize() int {
	return st.array.GetSize()
}

// 查询栈的容量
func (st *ArrayStack) GetCapacity() int {
	return st.array.GetCapacity()
}

// 判断栈是否为空
func (st *ArrayStack) IsEmpty() bool {
	return st.array.IsEmpty()
}

// 格式化输出栈信息
func (st *ArrayStack) String() string{
	res := fmt.Sprintf("Stack: [")
	for i:=0; i < st.array.GetSize(); i++ {
		data, _ := st.array.Get(i)
		res += fmt.Sprintf("%v", data)
		if i != st.array.GetSize() - 1{
			res += ", "
		}
	}
	res += "] top"
	return  res
}
