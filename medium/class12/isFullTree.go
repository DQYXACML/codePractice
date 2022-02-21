package main

import (
	"math"
)

type fullInfo struct {
	height int
	isFull bool
}

func process(root *TreeNode) *fullInfo {
	if root == nil {
		return &fullInfo{
			height: 0,
			isFull: true,
		}
	}
	leftInfo := process(root.Left)
	rightInfo := process(root.Right)
	height := int(math.Max(float64(leftInfo.height), float64(rightInfo.height))) + 1
	isFull := false

	if leftInfo.isFull && rightInfo.isFull && rightInfo.height == leftInfo.height {
		isFull = true
	}

	return &fullInfo{
		height: height,
		isFull: isFull,
	}
}

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return process(root).isFull
}
