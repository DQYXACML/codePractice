package main

// https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/

func findRepeatNumber(nums []int) int {
	m := make(map[int]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		if idx, ok := m[nums[i]]; ok {
			return idx
		}
		m[nums[i]] = nums[i]
	}
	return -1
}
