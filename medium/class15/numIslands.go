package main

import (
	"codePractice/tool/unionFind"
)

// https://leetcode.com/problems/number-of-islands/

func numIslands(grid [][]byte) int {
	row := len(grid)
	col := len(grid[0])
	uf := unionFind.NewUnionSetByDoubleList(grid)
	// 第一行
	for i := 1; i < col; i++ {
		if grid[0][i]-1 == '1' && grid[0][i] == '1' {
			uf.UnionByDouble(0, i, 0, i-1)
		}
	}
	// 第一列
	for i := 0; i < row; i++ {
		if grid[i][0] == '1' && grid[i-1][0] == '1' {
			uf.UnionByDouble(i, 0, i-1, 0)
		}
	}
	// 剩余
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			// 左和上
			if grid[i][j-1] == '1' {
				uf.UnionByDouble(i, j-1, i, j)
			}
			if grid[i-1][j] == '1' {
				uf.UnionByDouble(i-1, j, i, j)
			}
		}
	}
	return uf.TotalCount()
}
