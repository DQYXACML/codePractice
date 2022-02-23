package main

// https://leetcode.com/problems/reverse-nodes-in-k-group

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

func getKGroupEnd(head *ListNode, k int) *ListNode {
	k-- // 需要减去当前节点的个数，即如果k=2，即2个元素反转，从第i个开始已经算一个了，此处只要找剩下的k-1个即可
	for k != 0 && head != nil {
		head = head.Next
		k--
	}
	return head
}

// [1,2,3,4,5]
// 2

func reverseKGroup(head *ListNode, k int) *ListNode {
	start := head
	end := getKGroupEnd(start, k)
	if end == nil { // 够k个一组进行反转
		return head
	}
	head = end          // 找到第一组待反转的节点，反转后的头即现在的尾
	reverse(start, end) // start 经过这一步start已经将链表串了起来
	lastEnd := start    // 上一组的尾节点，即反转前的首节点
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
	end = end.Next // 获取到待反转的后一个节点，以便于串起整个节点
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
