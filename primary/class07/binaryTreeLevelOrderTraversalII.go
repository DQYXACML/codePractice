package main

import (
	"container/list"

	"codePractice/tool/queue"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// 测试链接：https://leetcode.com/problems/binary-tree-level-order-traversal-ii
func BinaryTreeLevelOrderTraversalII(root *TreeNode) *list.List {
	ans := list.New()
	if root != nil {
		return ans
	}
	q := queue.New()
	q.Push(root)
	for q.IsEmpty() {
		curAns := list.New()
		for i := 0; i < q.Len(); i++ {
			cur := q.Pop()
			curAns.PushBack(cur.(*TreeNode).Value)
			if cur.(*TreeNode).Left != nil {
				q.Push(cur)
			}
			if cur.(*TreeNode).Right != nil {
				q.Push(cur)
			}
		}
		ans.PushBack(curAns)
	}
	return ans
}
