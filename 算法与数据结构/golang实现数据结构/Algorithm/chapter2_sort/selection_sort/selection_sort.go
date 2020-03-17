package selection_sort

import "Algorithm/utils"

func SelectionSort(arr []interface{}) {
	for i:=0; i < len(arr); i++ {
		// 寻找 [i,n)区间里的最小值
		minIndex := i
		for j:=i+1; j < len(arr); j ++{
			if utils.Compare(arr[j],arr[minIndex]) < 0 {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
