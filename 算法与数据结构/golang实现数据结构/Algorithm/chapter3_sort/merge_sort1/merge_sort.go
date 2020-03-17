package merge_sort1

import "Algorithm/utils"

func MergeSort(arr []interface{}){
	mergeSort(arr, 0, len(arr)-1)
}

// 递归使用归并排序，对arr[l ,..., r]的范围进行排序
func mergeSort(arr []interface{}, l int, r int){
	//if l >= r { //表示区间区间表示错误，或一个元素都没有
	//	return
	//}
	// 优化，对于小数组不再使用归并方式的递归排序，而改为插入排序，插入排序在近似有序时，性能很高
	if r - l <= 15 {
		InsertSort(arr, l, r)
		return
	}
	mid := (l + r)/2
	mergeSort(arr, l , mid)  //左半区间递归排序
	mergeSort(arr, mid +1, r) // 右半区间递归排序
	if utils.Compare(arr[mid], arr[mid+1]) > 0 { //若小于等于0时，说明顺序已经满足排序，不需要归并操作
		Merge(arr, l, mid, r) // 将排序后的两个子区间元素进行归并操作
	}
}

// 将arr[l,...mid] 和arr[mid+1,...r]两部分进行归并
func Merge(arr []interface{}, l int, mid int, r int) {
	tmp := make([]interface{}, r - l + 1) //使用临时空间，将所有元素复制到零时空间
	for i:=l; i <= r; i++ {
		tmp[i-l] = arr[i]
	}
	i := l
	j := mid + 1
	for k:=l; k <=r; k++ {
		if i > mid {   //校验i索引的合法性
			arr[k] = tmp[j-l]
			j ++
		} else if j > r { //校验j索引的合法性
			arr[k] = tmp[i-l]
			i++
		} else if utils.Compare(tmp[i-l], tmp[j-l]) < 0 {
			arr[k] = tmp[i-l]
			i++
		} else {
			arr[k] = tmp[j-l]
			j++
		}
	}
}

func InsertSort(arr []interface{}, l int, r int){
	for i:=l+1; i <= r; i++ {
		// 寻找元素arr[i]合适的插入位置
		e := arr[i]
		var j int  // j保存元素e应该插入的位置
		for j=i; j > l; j-- {
			if utils.Compare(arr[j-1], e) > 0 {
				arr[j] = arr[j-1]
			} else {
				break
			}
		}
		arr[j] = e
	}
}
