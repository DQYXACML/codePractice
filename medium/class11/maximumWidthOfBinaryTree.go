package main

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
	for len(queue) > 0 {
		N := len(queue)
		for i := 0; i < N; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[N:]
	}
	return 0
}
