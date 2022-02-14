package main

import "fmt"

func selectionSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	N := len(arr)
	minIndex := 0
	for i := 0; i < N; i++ {
		minIndex = i // 每一轮找最小，然后和i替换
		for j := i + 1; j < N; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		arr[minIndex], arr[i] = arr[i], arr[minIndex]
	}
}

func PrintArr(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
	}
	fmt.Println()
}

func main() {
	arr := []int{
		6, 2, 5, 4, 3, 2, 8, 3,
	}
	selectionSort(arr)
	PrintArr(arr)
}
