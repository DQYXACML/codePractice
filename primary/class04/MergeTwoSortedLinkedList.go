package main

// https://leetcode.com/problems/merge-two-sorted-lists
func MergeTwoSortedList(h1, h2 *ListNode) *ListNode {
	if h1 == nil || h2 == nil {
		return nil
	}
	head := &ListNode{
		Value: 0,
		Next:  nil,
	}
	cur := &ListNode{} // 两个链表选出的头节点的另外一个链表的头节点
	if h1.Value > h2.Value {
		head = h2
		cur = h1
	} else {
		head = h1
		cur = h2
	}
	cur1 := head.Next // 两个链表选出的头节点的下一个节点
	pre := head       // 前一个节点，用于串起整个链表
	for cur != nil && cur1 != nil {
		if cur.Value <= cur1.Value {
			pre.Next = cur
			cur = cur.Next
		} else {
			pre.Next = cur1
			cur1 = cur1.Next
		}
		pre = pre.Next
	}
	// 剩余的直接串起来即可
	if cur == nil {
		pre.Next = cur1
	} else {
		pre.Next = cur
	}
	return head
}

// 更容易理解
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	if l1.Value > l2.Value {
		head = &ListNode{
			Value: 0,
			Next:  l1,
		}
	} else {
		head = &ListNode{
			Value: 0,
			Next:  l2,
		}
	}
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Value > l2.Value {
			cur.Next = l2
			l2 = l2.Next
		} else {
			cur.Next = l1
			l1 = l1.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return head.Next
}
