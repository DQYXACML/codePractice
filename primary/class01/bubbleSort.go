package main

import "fmt"

func bubbleSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	// 每轮确定最后一个最大的位置 i
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func main() {
	arr := []int{
		6, 2, 5, 4, 3, 2, 8, 3,
	}
	bubbleSort(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
	}
	fmt.Println()
	//PrintArr(arr)
}
