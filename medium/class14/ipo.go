package main

import "container/heap"

// https://blog.csdn.net/dream_follower/article/details/89791991

type ipo struct {
	pay  int
	cost int
}

type minCost []ipo
type maxPay []ipo

// 支出小根堆
func (m minCost) Len() int {
	return len(m)
}

func (m minCost) Less(i, j int) bool {
	return m[i].cost < m[j].cost
}

func (m minCost) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minCost) Push(x interface{}) {
	*m = append(*m, x.(ipo))
}

func (m *minCost) Pop() interface{} {
	N := len(*m)
	item := (*m)[:N-1]
	*m = (*m)[:N-1]
	return item
}

func (m *minCost) Size() int {
	return len(*m)
}

func (m *minCost) Peek() interface{} {
	return (*m)[len(*m)-1]
}

// 收入大根堆
func (m maxPay) Len() int {
	return len(m)
}

func (m maxPay) Less(i, j int) bool {
	return m[i].pay > m[j].pay
}

func (m maxPay) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *maxPay) Push(x interface{}) {
	*m = append(*m, x.(ipo))
}

func (m *maxPay) Pop() interface{} {
	N := len(*m)
	item := (*m)[:N-1]
	*m = (*m)[:N-1]
	return item
}

func (m *maxPay) Size() int {
	return len(*m)
}

// k 买几次
// w 初始资金
// profits 所有项目的利润
// capitals 所有项目的花费
func ipoProcess(k, w int, profits, capitals []int) int {
	minC := make(minCost, len(capitals))
	maxP := make(maxPay, len(profits))
	heap.Init(&minC)
	heap.Init(&maxP)
	for i := 0; i < len(profits); i++ {
		minC = append(minC, ipo{
			pay:  profits[i],
			cost: capitals[i],
		})
	}
	for i := 0; i < k; i++ {
		for minC.Size() > 0 && w >= minC.Peek().(ipo).cost {
			maxP = append(maxP, minC.Pop().(ipo))
		}
		if maxP.Size() == 0 { // 提前结束
			return w
		}
		w += maxP.Pop().(ipo).pay
	}
	return w
}
