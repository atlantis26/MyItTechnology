package quick_sort2

import (
	"Algorithm/chapter3_sort/merge_sort1"
	"Algorithm/utils"
	"math/rand"
)

// 双路快速排序
func QuickSort2(arr []interface{}) {
	quickSort2(arr, 0, len(arr)-1)
}

// 对arr[l,...r]部分进行快速排序
func quickSort2(arr []interface{}, l int, r int) {
	if r - l <= 15 {
		merge_sort1.InsertSort(arr, l, r)
		return
	}
	p := partition2(arr, l, r)
	quickSort2(arr, l, p-1)
	quickSort2(arr, p+1, r)
}

func partition2(arr[]interface{}, l int, r int) int {
	p := rand.Intn(r-l+1)+l
	arr[l] , arr[p] = arr[p], arr[l]
	v := arr[l]
	// arr[l+1,...i) <= v; arr[j,...r] >= v
	i := l+1
	j := r
	for {
		for i <= r && utils.Compare(arr[i], v) < 0 {
			i++
		}
		for j >= l+1 && utils.Compare(arr[j], v) > 0 {
			j--
		}
		if i > j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i ++
		j --
	}
	arr[l], arr[j] = arr[j], arr[l]

	return j
}
