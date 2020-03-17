package main

import (
	"Algorithm/chapter2_sort/insert_sort"
	"Algorithm/utils"
	"fmt"
)

func main() {
	//test1 := utils.SortTestHelper(6, 100)
	//utils.TestSort("SelectionSort", selection_sort.SelectionSort, test1)
	//test2 := utils.SortTestHelper(6, 100)
	//utils.TestSort("InsertSort", insert_sort.InsertSort, test2)
	test3 := utils.SortTestHelper(6, 100)
	utils.TestSort("InsertSort1", insert_sort.InsertSort1, test3)
	fmt.Println(test3)

}

