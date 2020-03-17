package merge_sort2

import (
	"Algorithm/chapter3_sort/merge_sort1"
	"math"
)

// 自底向上的归并排序算法
func MergeSortBu(arr []interface{}) {
	for sz:=1; sz <= len(arr); sz += sz{
		for i:=0; i + sz < len(arr); i += sz +sz {
			// 对arr[i,...i+sz-1]和arr[i+sz,...i+2*sz-1]进行归并
			merge_sort1.Merge(arr, i, i +sz -1, int(math.Min(float64(i + sz + sz - 1), float64(len(arr) -1))))
		}
	}
}
