package main

import (
	"math"
)

// https://leetcode.com/problems/diameter-of-binary-tree/

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return DiameterProcess(root).maxDistance
}

type DiameterInfo struct {
	height      int
	maxDistance int
}

func DiameterProcess(root *TreeNode) *DiameterInfo {
	if root == nil {
		return &DiameterInfo{
			height:      0,
			maxDistance: 0,
		}
	}
	leftInfo := DiameterProcess(root.Left)
	rightInfo := DiameterProcess(root.Right)
	height := int(math.Max(float64(leftInfo.height), float64(rightInfo.height))) + 1
	p1 := leftInfo.maxDistance
	p2 := rightInfo.maxDistance
	p3 := leftInfo.height + rightInfo.height + 1
	maxDistance := int(math.Max(float64(p3), math.Max(float64(p1), float64(p2))))
	return &DiameterInfo{
		height:      height,
		maxDistance: maxDistance,
	}
}
