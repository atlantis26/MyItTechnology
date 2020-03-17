package utils

import (
	"DataStructAndAlgorithm/utils"
	"fmt"
	"math/rand"
	"time"
)

// 生成n个元素的随机数组，每个元素的随机范围为[0, rangeR]
func SortTestHelper(n int, rangeR int) (arr []interface{}) {
	if rangeR <= 0{
		panic("rangeR should bigger than zero")
	}
	arr = make([]interface{}, n)
	for i:=0; i < n; i++ {
		arr[i] = rand.Intn(rangeR)
	}
	return arr
}

//
func IsSorted(arr []interface{}) bool {
	for i:=0 ; i < len(arr) - 1; i++{
		if utils.Compare(arr[i], arr[i + 1]) > 0 {
			return false
		}
	}
	return true
}

//
func TestSort(sortName string, f func([]interface{}), arr []interface{}) {
	startTime := time.Now()
	f(arr)
	endTime := time.Now()
	subM := endTime.Sub(startTime).Seconds()
	sorted := IsSorted(arr)
	fmt.Printf("Sort method '%s' execute result: time=%f Seconds, IsSorted=%v\n", sortName, subM, sorted)

}
