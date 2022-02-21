package main

// https://leetcode.com/problems/binary-tree-level-order-traversal-ii/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := make([][]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		N := len(queue)
		curAns := make([]int, 0, N)
		for i := 0; i < N; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			curAns = append(curAns, queue[i].Val)
		}
		queue = queue[N:]
		ans = append(ans, curAns)
	}
	return ans
}

func levelOrderBottom(root *TreeNode) [][]int {
	res := levelOrder(root)
	if res == nil {
		return nil
	}
	ans := make([][]int, 0, len(res))
	for i := len(res) - 1; i >= 0; i-- {
		ans = append(ans, res[i])
	}
	return ans
}
