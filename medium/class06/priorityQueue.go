package main

import (
	"container/heap"
	"fmt"
)

// Item 优先队列中的元素
type Item struct {
	value    string // 元素的值，可以是任意字符串。
	priority int    // 元素在队列中的优先级。
	index    int    // 元素在堆中的索引. 元素的索引可以用于更新操作，它由 heap.Interface 定义的方法维护
}

type PriorityQueue []*Item

func (p PriorityQueue) Len() int {
	return len(p)
}

// 我们希望 Pop 返回的是最大值而不是最小值，
// 因此这里使用大于号进行对比
func (p PriorityQueue) Less(i, j int) bool {
	return p[i].priority > p[j].priority
}
func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = i
	p[j].index = j
}

func (p *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*p)
	*p = append(*p, item)
}

func (p *PriorityQueue) Pop() interface{} {
	item := (*p)[len(*p)-1]
	item.index = -1
	*p = (*p)[:len(*p)-1]
	return item
}

func (p *PriorityQueue) Update(item *Item, value string, priority int) {

}

func main() {
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)
	heap.Push(&pq, &Item{
		value:    "orange",
		priority: 1,
	})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}
}
