package main

import (
	"math"
)

func isMaxSubSearchTree(root *TreeNode) int {
	return maxSubSearchProcess(root).MaxSubSearchSize
}

type MaxSubSearchInfo struct {
	MaxSubSearchSize int
	AllSize          int
	Max              int
	Min              int
}

func maxSubSearchProcess(root *TreeNode) *MaxSubSearchInfo {
	if root == nil {
		return nil
	}
	leftInfo := maxSubSearchProcess(root.Left)
	rightInfo := maxSubSearchProcess(root.Right)
	var max int
	var min int
	var allsize int
	max = root.Val
	min = root.Val
	allsize = 1
	if leftInfo != nil {
		max = int(math.Max(float64(leftInfo.Max), float64(max)))
		min = int(math.Min(float64(leftInfo.Min), float64(min)))
		allsize += leftInfo.AllSize
	}
	if rightInfo != nil {
		max = int(math.Max(float64(rightInfo.Max), float64(max)))
		min = int(math.Min(float64(rightInfo.Min), float64(min)))
		allsize += rightInfo.AllSize
	}
	p1 := -1
	if leftInfo != nil {
		p1 = leftInfo.MaxSubSearchSize
	}
	p2 := -1
	if rightInfo != nil {
		p2 = rightInfo.MaxSubSearchSize
	}
	p3 := -1
	var leftBst bool
	var rightBst bool
	if leftInfo == nil {
		leftBst = true
	} else if leftInfo.MaxSubSearchSize == leftInfo.AllSize {
		leftBst = true
	}
	if rightInfo == nil {
		rightBst = true
	} else if rightInfo.MaxSubSearchSize == rightInfo.AllSize {
		rightBst = true
	}
	if leftBst && rightBst {
		var leftMaxLessX bool
		var rightMinMaxX bool
		if leftInfo == nil {
			leftMaxLessX = true
		} else if root.Val > leftInfo.Max {
			leftMaxLessX = true
		}
		if rightInfo == nil {
			rightMinMaxX = true
		} else if root.Val < rightInfo.Min {
			rightMinMaxX = true
		}
		if leftMaxLessX && rightMinMaxX {
			var leftSize int
			var rightSize int
			if leftInfo == nil {
				leftSize = 0
			} else {
				leftSize = leftInfo.AllSize
			}
			if rightInfo == nil {
				rightSize = 0
			} else {
				rightSize = rightInfo.AllSize
			}
			p3 = leftSize + rightSize + 1
		}

	}
	return &MaxSubSearchInfo{
		MaxSubSearchSize: int(math.Max(float64(p1), math.Max(float64(p2), float64(p3)))),
		AllSize:          allsize,
		Max:              max,
		Min:              min,
	}

}
