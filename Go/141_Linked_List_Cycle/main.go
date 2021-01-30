package main

// Runtime: 4 ms, faster than 95.56% of Go online submissions for Linked List Cycle.
// Memory Usage: 3.8 MB, less than 38.10% of Go online submissions for Linked List Cycle.
// https://leetcode.com/submissions/detail/449615599/

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != nil && fast != slow && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return fast != nil && fast.Next != nil
}

// ListNode : Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
