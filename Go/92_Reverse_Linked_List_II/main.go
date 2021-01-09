package main

var next *ListNode

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Linked List II.
// Memory Usage: 2.2 MB, less than 19.57% of Go online submissions for Reverse Linked List II.
// https://leetcode.com/submissions/detail/440718400/
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if n == 1 || head == nil || head.Next == nil {
		if head != nil {
			next = head.Next
		}
		return head
	}
	if m > 1 {
		head.Next = reverseBetween(head.Next, m-1, n-1)
		return head
	}
	last := reverseBetween(head.Next, m-1, n-1)
	head.Next.Next = head
	head.Next = next
	return last
}

// ListNode : Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
