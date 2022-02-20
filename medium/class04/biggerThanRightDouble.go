package main

// 测试链接 https://leetcode.com/problems/reverse-pairs/
func biggerThanRightDouble(nums []int) int {
	if nums == nil || len(nums) == 1 {
		return 0
	}
	return Process(nums, 0, len(nums)-1)
}

func Process(arr []int, L, R int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	mid := (L + R) / 2
	return Process(arr, L, mid) + Process(arr, mid+1, R) + Merge(arr, L, mid, R)
}

func Merge(arr []int, l, m, r int) int {
	help := make([]int, r-l+1)
	ans := 0
	windowR := m + 1
	for i := l; i <= m; i++ {
		for windowR <= r && arr[i] > arr[windowR]*2 {
			windowR++
		}
		ans += windowR - (m + 1)
	}
	index := 0
	p1 := l
	p2 := m + 1
	for p1 <= m && p2 <= r {
		if arr[p1] < arr[p2] {
			help[index] = arr[p1]
			index--
			p1--
		} else {
			help[index] = arr[p2]
			index--
			p2--
		}
	}
	for p1 <= m {
		help[index] = arr[p1]
		index--
		p1--
	}
	for p2 <= r {
		help[index] = arr[p2]
		index--
		p2--
	}
	for i := 0; i < len(help); i++ {
		arr[l+i] = help[i]
	}
	return ans
}
