package main

// https://leetcode.com/problems/reverse-nodes-in-k-group
func ReverseNodesInKGroup(head *ListNode) *ListNode {
	return nil
}

func getKGroupEnd(head *ListNode, k int) *ListNode {
	k--
	for k != 0 && head != nil {
		head = head.Next
	}
	return head
}

func GetListLength(head *ListNode) int {
	if head == nil {
		return 0
	}
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	start := head
	end := getKGroupEnd(start, k)
	if end == nil {
		return head
	}
	head = end // 找到第一组待反转的节点，反转后的头即现在的尾
	reverse(start, end)
	lastEnd := start // 上一组的尾节点
	for lastEnd != nil {
		start = lastEnd.Next // 下一组的头节点
		end = getKGroupEnd(start, k)
		if end == nil {
			return head
		}
		reverse(start, end)
		lastEnd.Next = end // 上一组的尾节点的下一个需要连接上这一组的第一个，由于反转，end即为该组的头节点
		lastEnd = start    // 指向最后一个节点，供下一组使用
	}
	return head
}

func reverse(head, end *ListNode) {
	end = end.Next // 需要由head串起后面的节点
	pre := &ListNode{}
	cur := head
	next := &ListNode{}
	for cur != end {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	head.Next = end
}
