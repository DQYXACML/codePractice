package main

import "math"

// https://leetcode.com/problems/validate-binary-search-tree/

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return SearchProcess(root).isSearch
}

type SearchInfo struct {
	Max      int
	Min      int
	isSearch bool
}

func SearchProcess(root *TreeNode) *SearchInfo {
	if root == nil {
		return nil
	}

	leftInfo := SearchProcess(root.Left)
	rightInfo := SearchProcess(root.Right)
	Max := root.Val
	Min := root.Val
	if leftInfo != nil {
		Max = int(math.Max(float64(leftInfo.Max), float64(Max)))
		Min = int(math.Min(float64(leftInfo.Min), float64(Min)))
	}
	if rightInfo != nil {
		Max = int(math.Max(float64(rightInfo.Max), float64(Max)))
		Min = int(math.Min(float64(rightInfo.Min), float64(Min)))
	}

	isSearch := true
	if leftInfo != nil && !leftInfo.isSearch || rightInfo != nil && !rightInfo.isSearch {
		isSearch = false
	}
	if leftInfo != nil && root.Val <= leftInfo.Max || rightInfo != nil && root.Val >= rightInfo.Min {
		isSearch = false
	}
	return &SearchInfo{
		Max:      int(Max),
		Min:      int(Min),
		isSearch: isSearch,
	}
}
