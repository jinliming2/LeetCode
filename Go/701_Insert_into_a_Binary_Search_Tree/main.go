package main

// Runtime: 860 ms, faster than 8.00% of Go online submissions for Insert into a Binary Search Tree.
// Memory Usage: 62.7 MB, less than 100.00% of Go online submissions for Insert into a Binary Search Tree.
// https://leetcode.com/submissions/detail/444149799/
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

// TreeNode : Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
