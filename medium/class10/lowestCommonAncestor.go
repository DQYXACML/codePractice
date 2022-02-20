package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return process(root, p, q).ans
}

type Info struct {
	findA bool
	findB bool
	ans   *TreeNode
}

func process(root, p, q *TreeNode) *Info {
	if root == nil {
		return &Info{
			findA: false,
			findB: false,
			ans:   nil,
		}
	}
	leftInfo := process(root.Left, p, q)
	rightInfo := process(root.Right, p, q)
	findA := false
	findB := false
	if root == p || leftInfo.findA || rightInfo.findA {
		findA = true
	}
	if root == q || leftInfo.findB || rightInfo.findB {
		findB = true
	}

	var ans *TreeNode
	if leftInfo.ans != nil {
		ans = leftInfo.ans
	} else if rightInfo.ans != nil {
		ans = rightInfo.ans
	} else {
		if findB && findA {
			ans = root
		}
	}
	return &Info{
		findA: findA,
		findB: findB,
		ans:   ans,
	}
}
