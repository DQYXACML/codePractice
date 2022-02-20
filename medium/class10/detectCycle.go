package main

// https://leetcode.com/problems/linked-list-cycle-ii/
// 快慢指针，当相遇时，快指针回头，然后两指针一步一步走，相遇即相交点
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head
	isCycle := false
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			fast = head
			isCycle = true
		}
	}
	if !isCycle {
		return nil
	} else {
		for fast != slow {
			fast = fast.Next
			slow = slow.Next
		}
	}
	return slow
}
