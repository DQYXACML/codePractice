package main

func IsBinarySearchTree(root *TreeNode) bool {
	return processIsBst(root).isBSTree
}

// 最大值、最小值、是否是搜索二叉树
type Info struct {
	isBSTree bool
	maxVal   int
	minVal   int
}

func processIsBst(root *TreeNode) *Info {
	if root == nil {
		return &Info{
			isBSTree: true,
			maxVal:   0,
			minVal:   0,
		}
	}
	leftInfo := processIsBst(root.Left)
	rightInfo := processIsBst(root.Right)
	// 收集左右后，本层如何处理自己的数据
	min, max := root.Value, root.Value
	if leftInfo != nil {
		max = leftInfo.maxVal
		min = leftInfo.minVal
	}
	if rightInfo != nil {
		max = rightInfo.maxVal
		min = rightInfo.minVal
	}
	isBst := false
	leftIsBst := false
	rightIsBst := false
	leftMaxLessRoot := false
	rightMinMoreRoot := false
	if leftInfo == nil {
		leftIsBst = true
		leftMaxLessRoot = true
	} else {
		leftIsBst = leftInfo.isBSTree
		leftMaxLessRoot = leftInfo.maxVal < root.Value
	}
	if rightInfo == nil {
		rightIsBst = true
		rightMinMoreRoot = true
	} else {
		rightIsBst = rightInfo.isBSTree
		rightMinMoreRoot = rightInfo.minVal > root.Value
	}
	// 右数最小值
	if rightIsBst && leftIsBst && leftMaxLessRoot && rightMinMoreRoot {
		isBst = true
	}
	return &Info{
		isBSTree: isBst,
		maxVal:   max,
		minVal:   min,
	}
}
