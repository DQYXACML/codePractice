package main

import (
	"container/heap"
	"math"
)

// 线段重合问题
// 重叠区间个数，一般按照线段的start排序，遍历每个线段，将end放入小根堆排序
// 当某个线段的start 大于小根堆的最小值时，说明该线段和最小值线段不会有交集，弹出小根堆里小于该start的数，之后小根堆里的数量即为重合的段数

type Line struct {
	start, end int
}

func maxCover(arr [][]int) int {
	// 将二维数组转换为一维结构体数组，便于小根堆排序
	lines := make(Lines, len(arr))
	for i := 0; i < len(arr); i++ {
		lines[i] = Line{
			start: arr[i][0],
			end:   arr[i][1],
		}
	}
	heap.Init(&lines)
	max := 0
	for i := 0; i < len(lines); i++ {
		for !lines.IsEmpty() && lines.Peek() <= lines[i].start {
			heap.Pop(&lines)
		}
		heap.Push(&lines, lines[i].end)
	}
	max = int(math.Max(float64(max), float64(lines.Size())))
	return max
}

type Lines []Line

func (l Lines) Less(i, j int) bool {
	return l[i].end < l[j].end
}

func (l Lines) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l Lines) Len() int {
	return len(l)
}

func (l *Lines) Push(x interface{}) {
	*l = append(*l, x.(Line))
}

func (l *Lines) Pop() interface{} {
	item := (*l)[len(*l)-1]
	*l = (*l)[:len(*l)-1]
	return item
}

func (l *Lines) Peek() int {
	return (*l)[0].end
}

func (l *Lines) IsEmpty() bool {
	n := len(*l)
	if n == 0 {
		return true
	} else {
		return false
	}
}

func (l *Lines) Size() int {
	return len(*l)
}
