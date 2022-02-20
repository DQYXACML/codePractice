package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

// 实现less、swap、len方法，来实现sort接口
func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h IntHeap) Len() int {
	return len(h)
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	item := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return item
}

// 大根堆选topk小值 -> 前k个元素进大根堆，从k+1开始，如果比堆顶小，则弹出堆顶，然后进堆，最后一定会剩下k个最小值在堆里
// 小根堆选topk大值 -> 前k个元素进小根堆，从k+1开始，如果比堆顶大，则弹出堆顶，然后进堆，最后一定会剩下k个最大值在堆里
func main() {
	k := 3
	arr := []int{3, 5, 6, 8, 2, 14, 5, 6, 9, 8, 7, 5, 65}
	h := make(IntHeap, k)
	copy(h, IntHeap(arr[:k+1]))
	heap.Init(&h)
	for i := k; i < len(arr); i++ {
		if arr[i] < h[0] {
			heap.Pop(&h)
			heap.Push(&h, arr[i])
		}
	}
	for i, i2 := range h {
		fmt.Printf("%d %d", i, i2)
		fmt.Println()
	}
}
