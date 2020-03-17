package main

import (
	"Algorithm/chapter4_heap/heap_sort"
	"Algorithm/utils"
)

func main() {
	test1 := utils.SortTestHelper(100000, 1000000)
	utils.TestSort("HeapSort", heap_sort.HeapSort, test1)
	test2 := utils.SortTestHelper(100000, 1000000)
	utils.TestSort("HeapSort1", heap_sort.HeapSort1, test2)
}

