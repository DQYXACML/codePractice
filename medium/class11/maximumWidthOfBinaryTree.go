package main

import "math"

// https://leetcode.com/problems/maximum-width-of-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	maxCurrent := 0
	var curTail *TreeNode
	var nextTail *TreeNode
	curTail = root
	for len(queue) > 0 {
		current := 0
		N := len(queue)
		for i := 0; i < N; i++ {
			if queue[i] == nil {
				current++
				continue
			}
			if queue[i].Left != nil && queue[i].Right != nil {
				queue = append(queue, queue[i].Left)
				queue = append(queue, queue[i].Right)
				nextTail = queue[i].Right
			} else if queue[i].Left == nil && queue[i].Right != nil {
				queue = append(queue, nil)
				queue = append(queue, queue[i].Right)
				nextTail = queue[i].Right
			} else if queue[i].Left != nil && queue[i].Right == nil {
				queue = append(queue, queue[i].Left)
				queue = append(queue, nil)
				nextTail = queue[i].Left
			}
			if queue[i] != curTail && queue[i] != nil {
				current++
			} else {
				current++
				maxCurrent = int(math.Max(float64(maxCurrent), float64(current)))
				current = 0
				curTail = nextTail
			}
		}
		queue = queue[N:]
	}
	return maxCurrent
}
