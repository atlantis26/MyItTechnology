package main

import (
	"DataStructAndAlgorithm/stack/stack"
	"fmt"
)

type Stack interface {
	Push(interface{}) error
	Pop() (interface{}, error)
	Peek() (interface{}, error)
	GetSize() int
	IsEmpty() bool
}

func main() {
	stackInstance := stack.NewArrayStack(10)
	var st Stack
	st = stackInstance
	for _, i := range ([]string{"测试stack1", "测试stack2", "测试stack3"}) {
		_ = st.Push(i)
	}
	fmt.Println(st)
	st.Pop()
	fmt.Println(st)
	st.Pop()
	fmt.Println(st)
}