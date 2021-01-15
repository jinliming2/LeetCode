package main

import "math"

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Validate Binary Search Tree.
// Memory Usage: 5.5 MB, less than 80.56% of Go online submissions for Validate Binary Search Tree.
// https://leetcode.com/problems/validate-binary-search-tree/submissions/
func isValidBST(root *TreeNode) bool {
	return traverse(root, math.MinInt32-1, math.MaxInt32+1) // may failed in 32-bit OS
}

func traverse(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	return root.Val > min && root.Val < max && traverse(root.Left, min, root.Val) && traverse(root.Right, root.Val, max)
}

// TreeNode : Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
