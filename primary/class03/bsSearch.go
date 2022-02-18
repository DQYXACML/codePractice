package main

// 二分查找
func bsSearch(arr []int, num int) bool {
	if arr == nil || len(arr) == 0 {
		return false
	}
	L := 0
	R := len(arr) - 1
	for L <= R {
		mid := (L + R) / 2
		if arr[mid] == num {
			return true
		} else if arr[mid] > num {
			R = mid - 1
		} else if arr[mid] < num {
			L = mid + 1
		}
	}
	return false
}
