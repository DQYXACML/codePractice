package main

import (
	"math"
)

// 测试链接：https://leetcode.com/problems/balanced-binary-tree
func isBalancedBinaryTree(root *TreeNode) bool {
	return process(root).isBalanced
}

type info struct {
	height     int
	isBalanced bool
}

func process(root *TreeNode) *info {
	if root == nil {
		return &info{0, true}
	}
	leftInfo := process(root.Left)
	rightInfo := process(root.Right)
	height := math.Max(float64(leftInfo.height), float64(rightInfo.height)) + 1
	isBalanced := leftInfo.isBalanced && rightInfo.isBalanced && math.Abs(float64(leftInfo.height-rightInfo.height)) < 2
	return &info{
		height:     int(height),
		isBalanced: isBalanced,
	}
}
