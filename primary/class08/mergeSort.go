package main

func process(arr []int, L, R int) {
	if L == R {
		return
	}
	mid := (L + R) / 2
	process(arr, L, mid)
	process(arr, mid+1, R)
	merge(arr, L, mid, R)
}

func merge(arr []int, L, M, R int) {
	// 辅助数组
	help := make([]int, R-L+1)
	i := 0
	p1 := L
	p2 := M + 1
	for p1 <= M && p2 <= R {
		if arr[p1] <= arr[p2] {
			help[i] = arr[p1]
			i++
			p1++
		} else {
			help[i] = arr[p2]
			i++
			p2++
		}
	}
	// 可能有某些没处理
	for p1 <= M {
		help[i] = arr[p1]
		i++
		p1++
	}
	for p2 <= R {
		help[i] = arr[p2]
		i++
		p2++
	}
	for index := 0; index < len(help); index++ {
		arr[L+index] = help[index]
	}
}
