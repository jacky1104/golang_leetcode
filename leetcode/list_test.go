package leetcode

/**
206. Reverse Linked List
https://leetcode.com/problems/reverse-linked-list/
*/
func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	newHead := head
	oldHead := head.Next
	newHead.Next = nil

	for oldHead != nil {
		temp := oldHead
		oldHead = oldHead.Next

		temp.Next = newHead
		newHead = temp
	}

	return newHead

}

/**
92. Reverse Linked List II
https://leetcode.com/problems/reverse-linked-list-ii/
*/
func reverseBetween(head *ListNode, m int, n int) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Val: 0, Next: head}

	pre := dummy
	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}

	start := pre.Next
	then := start.Next

	for i := 0; i < n-m; i++ {

		start.Next = then.Next
		then.Next = pre.Next
		pre.Next = then
		then = start.Next
	}

	return dummy.Next

}
