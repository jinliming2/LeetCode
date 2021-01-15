package main

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Search Tree to Greater Sum Tree.
// Memory Usage: 2.5 MB, less than 12.00% of Go online submissions for Binary Search Tree to Greater Sum Tree.
// https://leetcode.com/submissions/detail/443324466/
func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	traverse(root, &sum)
	return root
}

func traverse(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	traverse(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	traverse(root.Left, sum)
}

// TreeNode : Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
