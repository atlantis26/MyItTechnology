package main

import (
	"DataStructAndAlgorithm/set/bstset"
	"DataStructAndAlgorithm/set/linkedset"
	"fmt"
)

type Set interface {
	Add(e interface{})
	Remove(e interface{})
	Contains(e interface{}) bool
	GetSize() int
	IsEmpty() bool
}

func TestBstSet() {
	words := []string {"1", "1", "2", "A", "B", "D","测试"}
	var s Set
	bstSetInstance := bstset.NewBstSet()
	s = bstSetInstance
	for _, i := range(words) {
		s.Add(i)
	}
	fmt.Println(s.IsEmpty())
	fmt.Println(s.GetSize())
	fmt.Println(s.Contains("A"))
	fmt.Println(s)
}

func TestLinkedSet() {
	words := []string {"1", "1", "2", "A", "B", "D","测试"}
	var s Set
	linkSetInstance := linkedset.NewLinkedSet()
	s = linkSetInstance
	for _, i := range(words) {
		s.Add(i)
	}
	fmt.Println(s.IsEmpty())
	fmt.Println(s.GetSize())
	fmt.Println(s.Contains("A"))
	fmt.Println(s)
}

func main() {
	TestBstSet()
	TestLinkedSet()
}