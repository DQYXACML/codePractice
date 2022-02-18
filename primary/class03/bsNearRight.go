package main

func bsNearRight(arr []int, num int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	L := 0
	R := len(arr) - 1
	index := -1
	for L <= R {
		mid := (L + R) / 2
		if arr[mid] <= num {
			index = mid
			L = mid - 1
		} else {
			R = mid + 1
		}
	}
	return index
}
