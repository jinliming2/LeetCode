package main

// Runtime: 20 ms, faster than 13.72% of Go online submissions for Palindrome Linked List.
// Memory Usage: 6.5 MB, less than 44.59% of Go online submissions for Palindrome Linked List.
// https://leetcode.com/submissions/detail/441128357/
func isPalindromeV2(head *ListNode) bool {
	length, p := 0, head
	for p != nil {
		length++
		p = p.Next
	}
	half := length / 2
	left := make([]int, half)
	p = head
	for i := half - 1; i >= 0; i-- {
		left[i] = p.Val
		p = p.Next
	}
	if length%2 != 0 {
		p = p.Next
	}
	for _, num := range left {
		if num != p.Val {
			return false
		}
		p = p.Next
	}
	return true
}
