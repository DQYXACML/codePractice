package main

// 测试链接 https://leetcode.com/problems/sort-colors/
func sortColors(nums []int) {
	process(nums, 0, len(nums)-1)
}

func process(arr []int, l, R int) {
	if l >= R {
		return
	}
	equalE := partition(arr, l, R)
	process(arr, l, equalE[0]-1)
	process(arr, equalE[1], R)
}

// [lessR+1,moreL] 为等于区间
func partition(arr []int, l, r int) []int {
	if l == r {
		return []int{l, r}
	}
	if l > r {
		return []int{-1, -1}
	}
	lessR := l - 1
	moreL := r
	index := l
	for index < moreL {
		if arr[index] < arr[r] {
			lessR++
			arr[index], arr[lessR] = arr[lessR], arr[index]
			index++
		} else if arr[index] > arr[r] {
			moreL--
			arr[index], arr[moreL] = arr[moreL], arr[index]
		} else {
			index++
		}
	}
	arr[moreL], arr[r] = arr[r], arr[moreL]
	return []int{lessR + 1, moreL}
}
