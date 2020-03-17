package main

import (
	"Algorithm/chapter3_sort/quick_sort"
	"Algorithm/chapter3_sort/quick_sort2"
	"Algorithm/chapter3_sort/quick_sort3"
	"Algorithm/utils"
)

func main() {
	test1 := utils.SortTestHelper(100000, 100000000)
	utils.TestSort("QuickSort", quick_sort.QuickSort, test1)
	//fmt.Println(test1)
	test2 := utils.SortTestHelper(100000, 100000000)
	utils.TestSort("QuickSort2", quick_sort2.QuickSort2, test2)
	//fmt.Println(test2)
	test3 := utils.SortTestHelper(100000, 100000000)
	utils.TestSort("QuickSort3", quick_sort3.QuickSort3, test3)
	//fmt.Println(test3)

}
