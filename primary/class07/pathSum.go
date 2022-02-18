package main

// 测试链接：https://leetcode.com/problems/path-sum

func pathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	// 叶子节点
	if root.Left == nil && root.Right == nil {
		if root.Value == sum {
			return true
		}
	}
	return pathSum(root.Left, sum-root.Value) || pathSum(root.Right, sum-root.Value)
}
