package main

type Node struct {
	Value int
	Next  *Node
}

func ReverseList(head *Node) *Node {
	pre := &Node{
		Value: 0,
		Next:  nil,
	}
	next := pre
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
