package main

// Runtime: 48 ms, faster than 6.47% of Go online submissions for Search in a Binary Search Tree.
// Memory Usage: 6.6 MB, less than 97.06% of Go online submissions for Search in a Binary Search Tree.
// https://leetcode.com/submissions/detail/444139732/
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if val < root.Val {
		return searchBST(root.Left, val)
	}
	return searchBST(root.Right, val)
}

// TreeNode : Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
