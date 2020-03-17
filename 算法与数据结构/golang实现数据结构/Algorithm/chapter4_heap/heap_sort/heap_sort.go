package heap_sort

import (
	"Algorithm/utils"
	"DataStructAndAlgorithm/maxheap/maxheap"
)

// 最大堆排序
func HeapSort(arr []interface{}) {
	heap := maxheap.HeapifyNewMaxHeap(arr)
	tmp := make([]interface{}, len(arr))
	for i:= len(arr) -1; i >= 0; i -- {
		t ,_ := heap.ExtractMax()
		tmp[i] = t
	}
	for i:= 0; i <len(arr); i++ {
		arr[i] = tmp[i]
	}
}

// 原址堆排序，不需要开额外的存储空间
func HeapSort1(arr []interface{}) {
	n := len(arr)
	for i:=(n-1)/2; i >=0; i -- {
		__shiftDown(arr, n, i)
	}
	for i:= n-1; i>0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		__shiftDown(arr, i, 0)
	}
}

func __shiftDown(arr []interface{}, n int, k int) {
	for 2*k+1 < n {
		j := 2*k+1 //在此轮训中，arr[k] 和arr[j]交换位置
		if j+1 < n && utils.Compare(arr[j+1], arr[j]) > 0 {
			j += 1
		}
		if utils.Compare(arr[k], arr[j]) >= 0 {
			break
		}
		arr[k], arr[j] = arr[j], arr[k]
		k = j
	}
}