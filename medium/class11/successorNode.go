package main

// 提供父节点指针
// 存在右子树，找该节点右树的最左节点
// 不存在右子树，向上找，判断父节点的右孩子是不是自己，若是，继续
// 若不是，则该节点即为后继

type STreeNode struct {
	val    int
	Left   *STreeNode
	Right  *STreeNode
	parent *STreeNode
}

func findSuccessorNode(root *STreeNode) *STreeNode {
	if root == nil {
		return nil
	}
	if root.Right != nil {
		return findMostLeft(root.Right)
	} else {
		parent := root.parent
		for parent != nil && parent.Right != root {
			root = parent
			parent = root.parent
		}
		return parent
	}
}

func findMostLeft(root *STreeNode) *STreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}
