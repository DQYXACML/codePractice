package main

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// 测试链接：https://leetcode.com/problems/same-tree
func isSameTree(t1, t2 *TreeNode) bool {
	// t1 t2 任何有一个节点为空则不相同
	if (t1 == nil) != (t2 == nil) {
		return false
	}
	if t1 == nil && t2 == nil {
		return true
	}
	if t1.Value == t2.Value {
		return isSameTree(t1.Left, t2.Left) && isSameTree(t1.Right, t2.Right)
	} else {
		return false
	}
}
