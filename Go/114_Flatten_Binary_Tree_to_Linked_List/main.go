package main

// Runtime: 8 ms, faster than 8.62% of Go online submissions for Flatten Binary Tree to Linked List.
// Memory Usage: 6.8 MB, less than 96.55% of Go online submissions for Flatten Binary Tree to Linked List.
// https://leetcode.com/submissions/detail/439852054/
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	right := root.Right
	root.Left, root.Right = nil, root.Left
	for root.Right != nil {
		root = root.Right
	}
	root.Right = right
}

// TreeNode : Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
