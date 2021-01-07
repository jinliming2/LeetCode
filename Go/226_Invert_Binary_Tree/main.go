package main

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Invert Binary Tree.
// Memory Usage: 2.2 MB, less than 100.00% of Go online submissions for Invert Binary Tree.
// https://leetcode.com/submissions/detail/439411800/
func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = root.Right, root.Left
		invertTree(root.Left)
		invertTree(root.Right)
	}
	return root
}

// TreeNode : Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
