package main

// Runtime: 4 ms, faster than 90.46% of Go online submissions for Remove Duplicates from Sorted Array.
// Memory Usage: 3.1 MB, less than 100.00% of Go online submissions for Remove Duplicates from Sorted Array.
// https://leetcode.com/submissions/detail/452392164/

func deleteDuplicates(head *ListNode) *ListNode {
	left, right := head, head
	for right != nil {
		if right.Val != left.Val {
			left, left.Next = right, right
		}
		right = right.Next
	}
	if left != nil {
		left.Next = nil
	}
	return head
}

// ListNode : Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
