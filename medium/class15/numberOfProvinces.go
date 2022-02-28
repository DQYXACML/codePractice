package main

import (
	"codePractice/tool/unionFind"
)

// https://leetcode.com/problems/number-of-provinces/

func findCircleNum(isConnected [][]int) int {
	uf := unionFind.NewUnionSet(len(isConnected))
	// 只需要遍历一半就可以了，是完全对称的
	for i := 0; i < len(isConnected); i++ {
		for j := 0; j <= i; j++ {
			if isConnected[i][j] == 1 {
				uf.Union(i, j)
			}
		}
	}
	return uf.TotalCount()
}
