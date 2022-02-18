package main

// 返回等于区的元素下标, 以最后一个元素arr[R]作为划分值
func partition(arr []int, L, R int) []int {
	lessR := L - 1 // 小于区初始在第一个元素的前一个
	moreL := R     // 大于区初始在最后一个
	index := L
	for index < moreL { // 当下标碰到大于区时返回
		if arr[index] < arr[R] {
			lessR++
			arr[lessR], arr[index] = arr[index], arr[lessR]
			index++
		} else if arr[index] > arr[R] {
			moreL--
			arr[moreL], arr[R] = arr[R], arr[moreL]
			index++
		} else {
			index++
		}
	}
	arr[moreL], arr[R] = arr[R], arr[moreL]
	return []int{lessR + 1, moreL}
}

func processQuick(arr []int, L, R int) {
	if L >= R {
		return
	}
	equalE := partition(arr, L, R)
	processQuick(arr, L, equalE[0]-1)
	processQuick(arr, equalE[1], R)
}
