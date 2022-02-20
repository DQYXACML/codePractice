package main

func process(arr []int, l, r int) int {
	if l == r {
		return 0
	}
	mid := (l + r) / 2
	return process(arr, l, mid) + process(arr, mid+1, r) + mergeSmallSum(arr, l, mid, r)
}

// 一个数组的每个元素左边比它小的数和，叫做小和
func mergeSmallSum(arr []int, L, M, R int) int {
	help := make([]int, 0)
	index := 0
	p1 := L
	p2 := M + 1
	ans := 0
	for p1 <= L && p2 <= R {
		if arr[p1] < arr[p2] {
			// 以左边的数为基准，右边有多少个比它大的
			ans += (R - p2 + 1) * arr[p1]
			help[index] = arr[p1]
			p1++
		} else {
			ans += 0
			help[index] = arr[p2]
			index++
			p2++
			index++
		}
	}
	for p1 <= L {
		help[index] = arr[p1]
		index++
		p1++
	}
	for p2 <= R {
		help[index] = arr[p2]
		index++
		p2++
	}
	for i := 0; i < len(help); i++ {
		arr[L+i] = help[i]
	}
	return ans
}
