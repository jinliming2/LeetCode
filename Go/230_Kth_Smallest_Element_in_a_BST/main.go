package main

// Runtime: 8 ms, faster than 93.45% of Go online submissions for Kth Smallest Element in a BST.
// Memory Usage: 6 MB, less than 87.77% of Go online submissions for Kth Smallest Element in a BST.
// https://leetcode.com/submissions/detail/442914349/
func kthSmallest(root *TreeNode, k int) int {
	count := 0
	res := 0
	traverse(root, k, &count, &res)
	return res
}

func traverse(root *TreeNode, k int, n, res *int) {
	if root == nil {
		return
	}
	traverse(root.Left, k, n, res)
	*n++
	if *n == k {
		*res = root.Val
		return
	}
	traverse(root.Right, k, n, res)
}

// TreeNode : Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
