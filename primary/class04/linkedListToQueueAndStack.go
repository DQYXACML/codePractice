package main

import (
	"math"
)

// MyQueue 链表实现队列和栈
type MyQueue struct {
	Head *Node
	Tail *Node
	Size int
}

func (q *MyQueue) NewQueue(size int) *MyQueue {
	return &MyQueue{
		Head: nil,
		Tail: nil,
		Size: size,
	}
}

func (q *MyQueue) Offer(v int) {
	cur := &Node{
		Value: v,
		Next:  nil,
	}
	if q.Tail == nil {
		q.Head = cur
		q.Tail = cur
	} else {
		q.Tail.Next = cur
		q.Tail = cur
	}
	q.Size++
}

func (q *MyQueue) Poll() int {
	ans := math.MaxInt
	if q.Head != nil {
		ans = q.Head.Value
		q.Head = q.Head.Next
		q.Size--
	}
	if q.Head == nil {
		q.Tail = nil
	}
	return ans
}

func (q *MyQueue) Peek() int {
	ans := math.MaxInt
	if q.Size == 0 {
		return ans
	} else {
		ans = q.Head.Value
		return ans
	}
}

// MyStack 链表实现栈
// 单头节点，头插即可模拟栈
type MyStack struct {
	Size int
	Head *Node
}

func (s *MyStack) NewStack(size int) *MyStack {
	return &MyStack{
		Size: size,
		Head: nil,
	}
}

func (s *MyStack) push(v int) {
	cur := &Node{
		Value: v,
		Next:  nil,
	}
	if s.Head == nil {
		s.Head = cur
	} else {
		cur.Next = s.Head
		s.Head = cur
	}
	s.Size++
}

func (s *MyStack) pop() int {
	ans := math.MaxInt
	if s.Head != nil {
		ans = s.Head.Value
		s.Head = s.Head.Next
		s.Size--
	}
	return ans
}

func (s *MyStack) peek() int {
	if s.Head == nil {
		return math.MaxInt
	} else {
		return s.Head.Value
	}
}
