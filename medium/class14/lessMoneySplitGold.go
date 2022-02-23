package main

import "container/heap"

// https://ac.nowcoder.com/questionTerminal/418d2fcdf7f24d6f8f4202e23951c0da
// 哈夫曼编码问题，小根堆解决
func lessMoney(arr []int) int {
	pQ := make(SmallSort, len(arr))
	copy(pQ, arr)
	heap.Init(&pQ)
	var sum int
	for pQ.Size() > 0 {
		cur := pQ.Pop().(int) + pQ.Pop().(int)
		sum += cur
		pQ.Push(cur)
	}
	return sum
}

type SmallSort []int

func (s SmallSort) Len() int {
	return len(s)
}

func (s SmallSort) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s SmallSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *SmallSort) Push(x interface{}) {
	*s = append(*s, x.(int))
}

func (s *SmallSort) Pop() interface{} {
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}

func (s *SmallSort) Size() int {
	return len(*s)
}
