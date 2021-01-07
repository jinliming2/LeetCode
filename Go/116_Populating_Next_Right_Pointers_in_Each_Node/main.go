package main

// Runtime: 8 ms, faster than 15.90% of Go online submissions for Populating Next Right Pointers in Each Node.
// Memory Usage: 6.1 MB, less than 33.68% of Go online submissions for Populating Next Right Pointers in Each Node.
// https://leetcode.com/submissions/detail/439860733/
func connect(root *Node) *Node {
	if root != nil {
		connectChildren(root.Left, root.Right)
	}
	return root
}

func connectChildren(left, right *Node) {
	if left == nil || right == nil {
		return
	}
	left.Next = right
	connectChildren(left.Left, left.Right)
	connectChildren(left.Right, right.Left)
	connectChildren(right.Left, right.Right)
}

// Node : Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
