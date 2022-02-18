package main

// 测试链接：https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
// preOrder[L1..R1] inOrder[L2..R2]
func ConstructBinaryTreeFromPreorderAndInorderTraversal(preOrder, inorder []int, L1, R1, L2, R2 int) *TreeNode {
	if L1 > R1 {
		return nil
	}
	head := &TreeNode{
		Value: preOrder[L1],
		Left:  nil,
		Right: nil,
	}
	// 最后框出一个数
	if L1 == R1 {
		return head
	}
	// 在中序里找前序的字符，他的左面即为前序字符的左树，右面即为前序字符的右树
	find := L2
	for inorder[find] != preOrder[L1] {
		find++
	}
	// 左树
	// PreOrder: 1 2 4 5 3 6 7 L1=0, R1=6
	// InOrder:  4 2 5 1 6 3 7 L2=0, R2=6
	// find = 3
	// PreOrder: 1 [2 4 5] 3 6 7 L1=1, R1=3
	// InOrder:  [4 2 5] 1 6 3 7 L2=0, R2=2
	// find = 1
	// PreOrder: 1 2 [4] 5 3 6 7 L1=2, R1=2
	// InOrder:  [4] 2 5 1 6 3 7 L2=0, R2=0
	// 右树
	// PreOrder: 1 2 4 5 [3 6 7] L1=4, R1=6
	// InOrder:  4 2 5 1 [6 3 7] L2=4, R2=6
	// find = 5
	// PreOrder: 1 2 4 5 3 6 [7] L1=6, R1=6
	// InOrder:  4 2 5 1 6 3 [7] L2=6, R2=6
	head.Left = ConstructBinaryTreeFromPreorderAndInorderTraversal(preOrder, inorder, L1+1, L1+find-L2, L2, find-1)  // 左树
	head.Right = ConstructBinaryTreeFromPreorderAndInorderTraversal(preOrder, inorder, L1+find-L2+1, R1, find+1, R2) // 右树
	return head
}
