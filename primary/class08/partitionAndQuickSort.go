package main

import "fmt"

// 返回等于区的元素下标, 以最后一个元素arr[R]作为划分值
func partition(arr []int, L, R int) []int {
	if L == R {
		return []int{L, R}
	}
	if L > R {
		return []int{-1, -1}
	}
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
			arr[moreL], arr[index] = arr[index], arr[moreL]
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

func main() {
	arr := []int{
		2, 0, 2, 1, 1, 0,
	}
	processQuick(arr, 0, len(arr)-1)
	for _, i := range arr {
		fmt.Print(i)
	}
	fmt.Println()
}
