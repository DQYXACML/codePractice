package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type MyListNode struct {
	Val  int
	Next *MyListNode
}

// 测试链接：https://leetcode.com/problems/merge-k-sorted-lists/
func mergeKLists(lists []*MyListNode) *MyListNode {
	length := len(lists)
	if lists == nil {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	num := length / 2
	left := mergeKLists(lists[:num])  // [0..num-1]
	right := mergeKLists(lists[num:]) // [num..length]
	return mergeTwoLists1(left, right)
}

// 合并两个有序链表的递归版本
func mergeTwoLists1(l1, l2 *MyListNode) *MyListNode {
	head := &MyListNode{}
	if l1.Val > l2.Val {
		head = &MyListNode{
			Val:  0,
			Next: l1,
		}
	} else {
		head = &MyListNode{
			Val:  0,
			Next: l2,
		}
	}
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
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
