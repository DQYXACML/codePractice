package main

import "fmt"

func insertionSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			arr[j+1], arr[j] = arr[j], arr[j+1]
		}
	}
}

func main() {
	arr := []int{
		6, 2, 5, 4, 3, 2, 8, 3,
	}
	insertionSort(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
	}
	fmt.Println()
	//PrintArr(arr)
}
