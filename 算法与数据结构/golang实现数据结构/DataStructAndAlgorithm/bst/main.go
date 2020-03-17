package main

import (
	"DataStructAndAlgorithm/bst/bst"
	"fmt"
)

func main() {
	b := bst.NewBst()
	nums := []int {5, 3, 6, 8, 4, 2}
	for _, i := range(nums) {
		b.Add(i)
	}
	b.PreOrder()
	b.PreOrderNR()
	b.LevelOrder()
	b.InOrder()
	b.PostOrder()
	fmt.Println(b)
}
