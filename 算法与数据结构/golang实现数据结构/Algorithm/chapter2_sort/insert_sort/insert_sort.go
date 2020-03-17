package insert_sort

import (
	"Algorithm/utils"
	"fmt"
)

func InsertSort(arr []interface{}){
	for i:=1; i < len(arr); i++ {
		// 寻找元素arr[i]合适的插入位置
		for j:=i; j > 0; j-- {
			if utils.Compare(arr[j], arr[j-1]) < 0 {
				arr[j], arr[j-1] = arr[i], arr[j]
			} else {
				break
			}
		}
	}
}

// 改进插入排序算法
func InsertSort1(arr []interface{}){
	for i:=1; i < len(arr); i++ {
		// 寻找元素arr[i]合适的插入位置
		e := arr[i]
		var j int  // j保存元素e应该插入的位置
		for j=i; j > 0; j-- {
			fmt.Println(j)
			if utils.Compare(arr[j-1], e) > 0 {
				arr[j] = arr[j-1]
			} else {
				break
			}
		}
		fmt.Println("j:",j)
		arr[j] = e
	}
}
