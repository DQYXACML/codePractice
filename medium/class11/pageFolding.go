package main

import "fmt"

// 从第i层开始，共N层
func process(i, N int, down bool) {
	if i > N {
		return
	}
	process(i+1, N, true) // 左孩子一定是凸
	if down {
		fmt.Print("凹")
	} else { // 右孩子一定是凹
		fmt.Print("凸")
	}
	process(i+1, N, false)
}
