package main

// Runtime: 4 ms, faster than 97.10% of Go online submissions for Linked List Cycle II.
// Memory Usage: 3.8 MB, less than 34.78% of Go online submissions for Linked List Cycle II.
// https://leetcode.com/submissions/detail/449621776/

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = head
	for slow != fast {
		slow, fast = slow.Next, fast.Next
	}
	return slow
}

// ListNode : Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
