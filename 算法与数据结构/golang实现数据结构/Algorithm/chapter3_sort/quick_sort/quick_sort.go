package quick_sort

import (
	"Algorithm/utils"
	"math/rand"
)

func QuickSort(arr []interface{}) {
	quickSort(arr,0, len(arr) -1)
}


// 对arr[l,...r]部分进行快速排序
func quickSort(arr []interface{}, l int, r int) {
	if l >= r {
		return
	}
	p := partition(arr, l, r)
	quickSort(arr, l, p-1)
	quickSort(arr, p+1, r)
}

// 对arr[l,...r]部分进行partition操作
// 返回p, 使得arr[l,...p-1] < arr[p]; arr[p+1,...r] > arr[p]
func partition(arr []interface{}, l int, r int) int{
	// 随机化快速排序特定元素为非首元素，确保在近似有序的数组时，不会退化为O(n^2)
	p := rand.Intn(r-l) + l
	arr[p], arr[l] = arr[l], arr[p]
	v := arr[l]
	// arr[l+1,...j] < v; arr[j+1,...i] > v
	j := l
	for i := l + 1; i <= r; i++{
		if utils.Compare(arr[i], v) < 0 {
			arr[j+1], arr[i] = arr[i], arr[j+1]
			j++
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}