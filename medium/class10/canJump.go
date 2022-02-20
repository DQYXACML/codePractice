package main

// https://leetcode.com/problems/jump-game/
func canJump(nums []int) bool {
	needNum := 1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] >= needNum {
			needNum = 1
		} else {
			needNum++
		}
		if i == 1 && needNum > 1 {
			return false
		}
	}
	return true
}
