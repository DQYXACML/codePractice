package main

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 每一条二叉树路径
func binaryTreeNode(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	res := []string{}
	// 处理叶子节点
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}
	// 左树
	tmpLeft := binaryTreeNode(root.Left)
	for i := 0; i < len(tmpLeft); i++ {
		res = append(res, strconv.Itoa(root.Val)+"->"+tmpLeft[i])
	}
	// 右树
	tmpRight := binaryTreeNode(root.Right)
	for i := 0; i < len(tmpRight); i++ {
		res = append(res, strconv.Itoa(root.Val)+"->"+tmpRight[i])
	}
	return res
}
