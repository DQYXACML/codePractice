package main

// 从右往左的归并
func reversePairs(nums []int) int {
	if nums == nil || len(nums) == 1 {
		return 0
	}
	return reverseProcess(nums, 0, len(nums)-1)
}

func reverseProcess(arr []int, L, R int) int {
	if L == R {
		return 0
	}
	mid := (L + R) / 2
	return reverseProcess(arr, L, mid) + reverseProcess(arr, mid+1, R) + merge(arr, L, mid, R)
}

func merge(arr []int, L, M, R int) int {
	help := make([]int, R-L+1)
	index := len(help) - 1
	p1 := M
	p2 := R
	ans := 0
	for p1 >= L && p2 > M {
		if arr[p1] > arr[p2] {
			ans += p2 - M
			help[index] = arr[p1]
			index--
			p1--
		} else {
			ans += 0
			help[index] = arr[p2]
			index--
			p2--
		}
	}
	for p1 >= L {
		help[index] = arr[p1]
		p1--
		index--
	}
	for p2 > M {
		help[index] = arr[p2]
		index--
		p2--
	}
	for i := 0; i < len(help); i++ {
		arr[L+i] = help[i]
	}
	return ans
}
