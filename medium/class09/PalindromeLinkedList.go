package main

// 回文链表 https://leetcode.com/problems/palindrome-linked-list/

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

// 空间复杂度O(n)
func isPalindrome(head *ListNode) bool {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	for i, j := 0, len(arr)-1; i < j; {
		if arr[i] != arr[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// 空间复杂度O(1)
func isPalindrome1(head *ListNode) bool {
	// 找到中点
	slow := findMiddle(head)
	// 以中点开始反转后续链表
	p2 := reverseHalf(slow)
	p1 := head
	return compareTwoList(p1, p2)
}

func findMiddle(head *ListNode) *ListNode {
	p1 := head // fast
	p2 := head // slow
	for p1 != nil && p1.Next != nil {
		p1 = p1.Next.Next
		p2 = p2.Next
	}
	return p2
}

func reverseHalf(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func compareTwoList(p1, p2 *ListNode) bool {
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}
