package main

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Nth Node From End of List.
// Memory Usage: 2.2 MB, less than 22.29% of Go online submissions for Remove Nth Node From End of List.
// https://leetcode.com/submissions/detail/449639341/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	last, slow, fast := (*ListNode)(nil), head, head
	for ; n > 0; n-- {
		if fast == nil {
			return head
		}
		fast = fast.Next
	}
	for fast != nil {
		last, slow, fast = slow, slow.Next, fast.Next
	}
	if last == nil {
		return slow.Next
	}
	last.Next = slow.Next
	return head
}

// ListNode : Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
