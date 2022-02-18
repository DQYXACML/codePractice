package main

// 测试链接：https://leetcode.com/problems/symmetric-tree
func isSymmetricTree(t1, t2 *TreeNode) bool {
	if (t1 == nil) != (t2 == nil) {
		return false
	}
	if t1 == nil && t2 == nil {
		return true
	}
	if t1.Value == t2.Value {
		return isSymmetricTree(t1.Left, t2.Right) && isSymmetricTree(t1.Right, t2.Left)
	} else {
		return false
	}
}
