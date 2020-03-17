package main

import (
	"DataStructAndAlgorithm/map/bstmap"
	"DataStructAndAlgorithm/map/linkedmap"
	"fmt"
)

type Map interface {
	Add(key interface{}, value interface{})
	Remove(key interface{}) interface{}
	Contains(key interface{}) bool
	Get(key interface{}) interface{}
	Set(key interface{}, value interface{}) (err error)
	GetSize() int
	IsEmpty() bool
}

func TestLinkedMap() {
	var m Map
	linkedMapInstance := linkedmap.NewLinkedMap()
	m = linkedMapInstance
	m.Add("1","测试1")
	m.Add("2", "测试2")
	fmt.Println(m.GetSize())
	fmt.Println(m.Get("1"))
	fmt.Println(m.Get("2"))
	fmt.Println(m.Get("3"))
}

func TestBstMap() {
	var m Map
	bstMapInstance := bstmap.NewBstMap()
	m = bstMapInstance
	m.Add("1","测试1")
	m.Add("2", "测试2")
	fmt.Println(m.GetSize())
	fmt.Println(m.Get("1"))
	fmt.Println(m.Get("2"))
	fmt.Println(m.Get("3"))
}

func main() {
	TestLinkedMap()
	TestBstMap()
}
