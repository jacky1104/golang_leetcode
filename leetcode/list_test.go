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
