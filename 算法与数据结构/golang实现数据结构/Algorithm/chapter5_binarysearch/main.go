package chapter5_binarysearch

import "Algorithm/utils"

// 二分查找法，在有序数组arr中，查找target
// 如果找到target,返回相应的索引index,如果没有找到target,返回-1
func BinarySearch(arr[]interface{}, target interface{}) int {
	// 在arr[l,...r]之中查找target
	l := 0
	r := len(arr) -1
	for  l <= r {
		mid := l + (r - l) >> 1
		if utils.Compare(arr[mid], target) == 0 {
			return mid
		} else if utils.Compare(arr[mid], target) > 0  {
			// 在arr[l,...mid-1]之中查找target
			r = mid - 1
		} else {  // target > arr[mid]
			// 在arr[mid+1,...r]之中查找target
			l = mid + 1
		}
	}
	return -1
}
