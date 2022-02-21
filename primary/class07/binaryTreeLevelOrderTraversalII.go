package main

type TreeNode struct {
	Val  int
	Left *TreeNode
	Right *TreeNode
}

// 测试链接：https://leetcode.com/problems/binary-tree-level-order-traversal
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := make([][]int,0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		N := len(queue)
		curAns := make([]int,0,N)
		for i := 0; i <N; i++ {
			if queue[i].Left != nil {
				queue = append(queue,queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue,queue[i].Right)
			}
			curAns = append(curAns,queue[i].Val)
		}
		queue = queue[N:]
		ans = append(ans,curAns)
	}
	return ans
}
