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

/*
203. Remove Linked List Elements
https://leetcode.com/problems/remove-linked-list-elements/
*/
func removeElements(head *ListNode, val int) *ListNode {

	if head == nil {
		return head
	}

	for head != nil && head.Val == val {
		head = head.Next
	}

	temp := head
	for temp != nil && temp.Next != nil {
		if temp.Next.Val == val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}

	return head
}

/**

82. Remove Duplicates from Sorted List II
https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/

 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/
func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	virHead := &ListNode{Next: head}
	preNode, currentNode := virHead, virHead.Next

	for currentNode != nil && currentNode.Next != nil {

		if currentNode.Val != currentNode.Next.Val {
			preNode = currentNode
			currentNode = currentNode.Next
		} else {
			temp := currentNode.Next
			for temp != nil && temp.Val == currentNode.Val {
				temp = temp.Next
			}
			preNode.Next = temp
			currentNode = preNode.Next
		}

	}

	return virHead.Next

}

/**
83. Remove Duplicates from Sorted List
https://leetcode.com/problems/remove-duplicates-from-sorted-list/
*/
func deleteDuplicates1(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	temp := head
	for temp != nil && temp.Next != nil {
		if temp.Val == temp.Next.Val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}
	return head

}

/**
19. Remove Nth Node From End of List
https://leetcode.com/problems/remove-nth-node-from-end-of-list/
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if head == nil || n <= 0 {
		return head
	}
	dummyHead := &ListNode{Next: head}
	first, second := dummyHead, dummyHead

	for i := 0; i < n && first != nil; i++ {
		first = first.Next
	}

	for first.Next != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next
	return dummyHead.Next
}
