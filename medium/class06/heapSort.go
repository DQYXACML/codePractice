package main

import "fmt"

func heapSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	// 建堆
	for i := 0; i < len(arr); i++ {
		heapInsert(arr, i)
	}
	heapSize := len(arr) - 1
	arr[heapSize], arr[0] = arr[0], arr[heapSize] // 建立完堆后，最大值和最后一位替换
	for heapSize > 0 {
		heapify(arr, 0, heapSize)
		heapSize-- // 逻辑上删除最后一个元素
		arr[0], arr[heapSize] = arr[heapSize], arr[0]
	}
}

// 从下往上调整堆结构
func heapInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] {
		arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
		index = (index - 1) / 2
	}
}

// 从上往下调整堆结构
func heapify(arr []int, index int, heapSize int) {
	left := 2*index + 1
	for left < heapSize {
		largest := left
		if left+1 < heapSize && arr[left+1] > arr[left] {
			largest = left + 1
		}
		if arr[largest] < arr[index] {
			largest = index
		}
		if largest == index {
			break // 当前值即为最大值，不需要再调整
		}
		// 存在孩子比当前值还要大，则交换，并以孩子节点作为下一次循环的开始
		arr[largest], arr[index] = arr[index], arr[largest]
		index = largest
		left = 2*left + 1
	}
}

func main() {
	arr := []int{39, 100, 1, 99, 0, 34, 2, 35}
	heapSort(arr)
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
