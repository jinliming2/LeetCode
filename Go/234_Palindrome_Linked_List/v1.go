package main

var left *ListNode

// Runtime: 28 ms, faster than 6.07% of Go online submissions for Palindrome Linked List.
// Memory Usage: 9.9 MB, less than 5.54% of Go online submissions for Palindrome Linked List.
// https://leetcode.com/submissions/detail/441125176/
func isPalindromeV1(head *ListNode) bool {
	left = head
	return traverse(head)
}

func traverse(head *ListNode) bool {
	if head == nil {
		return true
	}
	res := traverse(head.Next) && left.Val == head.Val
	left = left.Next
	return res
}

// ListNode : Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
