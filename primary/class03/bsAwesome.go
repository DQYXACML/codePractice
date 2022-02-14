package main

// 局部最小值
// 相邻的数不想等
func bsAwesome(arr []int) int {
	if arr == nil || len(arr) ==0 {
		return -1
	}
	// 开头是否下降趋势
	if arr[1] > arr[0] {
		return 0
	}

	// 结尾是否上升趋势
	if arr[len(arr)-2] > arr[len(arr)-1] {
		return 0
	}
	L := 0
	R := len(arr)-1
	// 上面两个if不满足即肯定有局部最小
	for L < R - 1 {
		mid := (L+R) / 2
		if arr[mid] < arr[mid-1] && arr[mid] < arr[mid +1] {
			return mid
		} else if arr[mid] > arr[mid-1] {
			R = mid - 1
		} else {
			L = mid + 1
		}
	}
	if arr[L] < arr[R] {
		return L
	}else {
		return R
	}
}