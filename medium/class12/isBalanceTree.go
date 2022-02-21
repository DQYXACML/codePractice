package main

// https://leetcode.com/problems/balanced-binary-tree/

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type balanceInfo struct {
	maxHight      int
	isBalanceTree bool
}

func balanceProcess(root *TreeNode) *balanceInfo {
	if root == nil {
		return &balanceInfo{
			maxHight:      0,
			isBalanceTree: true,
		}
	}
	leftInfo := balanceProcess(root.Left)
	rightInfo := balanceProcess(root.Right)
	maxHigh := int(math.Max(float64(leftInfo.maxHight), float64(rightInfo.maxHight))) + 1
	isBalanceTree := false
	//if leftInfo.isBalanceTree && rightInfo.isBalanceTree {
	//	isBalanceTree = true
	//}
	if math.Abs(float64(leftInfo.maxHight-rightInfo.maxHight)) < 2 && leftInfo.isBalanceTree && rightInfo.isBalanceTree {
		isBalanceTree = true
	}
	return &balanceInfo{
		maxHight:      maxHigh,
		isBalanceTree: isBalanceTree,
	}
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return balanceProcess(root).isBalanceTree
}
