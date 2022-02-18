package main

import (
	"math"
)

// MaximumDepthOfBinaryTree 测试链接：https://leetcode.com/problems/maximum-depth-of-binary-tree
func MaximumDepthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(MaximumDepthOfBinaryTree(root.Left)), float64(MaximumDepthOfBinaryTree(root.Right))) + 1)
}
