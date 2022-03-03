package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return []int{head.Val}
	}

	res1, _ := reverse(head)

	res := make([]int, 0)
	for res1 != nil {
		res = append(res, res1.Val)
		res1 = res1.Next
	}
	return res
}

func reverse(head *ListNode) (*ListNode, int) {
	cur := head
	var pre *ListNode
	i := 0
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
		i++
	}
	return pre, i
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
