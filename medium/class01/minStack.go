package main

// 测试链接 https://leetcode.com/problems/min-stack/
type MinStack struct {
	elements, min []int
	l             int
}

func Constructor() MinStack {
	return MinStack{
		elements: make([]int, 0),
		min:      make([]int, 0),
		l:        0,
	}
}

func (this *MinStack) Push(val int) {
	this.elements = append(this.elements, val) // 正常入栈
	if this.l == 0 {                           // 第一次直接放入min栈
		this.min = append(this.min, val)
	} else { // 如果当前值比min小，当前值入栈，否则最小值入栈
		min := this.GetMin()
		if val < min {
			this.min = append(this.min, val)
		} else {
			this.min = append(this.min, min)
		}
	}
	this.l++
}

func (this *MinStack) Pop() {
	this.l--
	this.min = this.min[:this.l]
	this.elements = this.elements[:this.l]
}

func (this *MinStack) Top() int {
	return this.elements[this.l-1]
}

// 获取，不弹出，类似于peek
func (this *MinStack) GetMin() int {
	return this.min[this.l-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
