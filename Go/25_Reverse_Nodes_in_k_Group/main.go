package main

// Runtime: 8 ms, faster than 9.23% of Go online submissions for Reverse Nodes in k-Group.
// Memory Usage: 3.9 MB, less than 8.46% of Go online submissions for Reverse Nodes in k-Group.
// https://leetcode.com/submissions/detail/441109052/
func reverseKGroup(head *ListNode, k int) *ListNode {
	var ret, last *ListNode
	currentHead := head
	for {
		current := currentHead
		// Check length
		for i := 1; i < k || current == nil; i++ {
			if current == nil {
				if ret == nil {
					return head
				}
				return ret
			}
			current = current.Next
		}
		// reverse
		current, next := currentHead, currentHead.Next
		for i := 1; i < k && next != nil; i++ {
			next.Next, current, next = current, next, next.Next
		}
		// update head
		currentHead.Next = next
		if last != nil {
			last.Next = current
		}
		if ret == nil {
			ret = current
		}
		last = currentHead
		currentHead = next
	}
}

// ListNode : Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
