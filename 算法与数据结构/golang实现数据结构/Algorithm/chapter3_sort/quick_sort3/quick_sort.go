package quick_sort3

import (
	"Algorithm/chapter3_sort/merge_sort1"
	"Algorithm/utils"
	"math/rand"
)

// 三路快速排序
func QuickSort3(arr []interface{}) {
	quickSort3(arr, 0, len(arr) - 1)
}

// 三路快速排序处理arr[l,...r]
// 将arr[l,...r]分为<v;==v;>v三部分
// 之后递归对<v;>v的两部分继续进行三路快速排序；==v的部分不用排序处理，在等于v的元素很多时，就很有用
func quickSort3(arr []interface{}, l int, r int) {
	//对于小数组的排序使用插入排序将更有效率
	if r - l <= 15 {
		merge_sort1.InsertSort(arr, l, r)
		return
	}

	// partition
	p := rand.Intn(r-l+1)+l
	arr[l], arr[p] = arr[p], arr[l]
	v := arr[l]

	lt := l  // arr[l+1,...lt] < v
	i := l+1    // arr[lt+1,...i] == v
	gt := r+1   // arr[gt,...r] > v
	for i < gt {
		if utils.Compare(arr[i], v) < 0 {
			arr[i], arr[lt+1] = arr[lt+1], arr[i]
			lt ++
			i ++
		} else if utils.Compare(arr[i], v) > 0 {
			arr[i], arr[gt-1] = arr[gt -1], arr[i]
			gt --
		} else {  // arr[i] ==v
			i ++
		}
	}
	arr[l], arr[lt] = arr[lt], arr[l]
	quickSort3(arr, l, lt-1)
	quickSort3(arr, gt, r)
}

