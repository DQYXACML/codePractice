package main

// https://leetcode.com/problems/add-two-numbers/
type ListNode struct {
	Value int
	Next  *ListNode
}

func AddTwoListNumber(h1, h2 *ListNode) *ListNode {
	h1Len := GetListLength(h1)
	h2Len := GetListLength(h2)
	l := &ListNode{}
	s := &ListNode{}
	if h1Len > h2Len {
		l = h1
		s = h2
	} else {
		l = h2
		s = h1
	}
	curL := l
	curS := s
	curNum := 0
	carry := 0
	last := &ListNode{}
	for curS != nil {
		curNum = curS.Value + curL.Value + carry
		carry = curNum / 10
		curL.Value = curNum % 10
		last = curL
		curL = curL.Next
		curS = curS.Next
	}
	for curL != nil {
		curNum = curL.Value + carry
		curL.Value = curNum % 10
		carry = curNum / 10
		last = curL
		curL = curL.Next
	}
	if carry != 0 {
		last.Next = &ListNode{
			Value: 0,
			Next:  nil,
		}
	}
	return l
}
