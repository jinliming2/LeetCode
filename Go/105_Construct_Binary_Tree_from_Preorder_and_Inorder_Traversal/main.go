package main

// Runtime: 48 ms, faster than 5.45% of Go online submissions for Construct Binary Tree from Preorder and Inorder Traversal.
// Memory Usage: 26.5 MB, less than 5.00% of Go online submissions for Construct Binary Tree from Preorder and Inorder Traversal.
// https://leetcode.com/submissions/detail/441637053/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	pi := 0
	for i, n := range inorder {
		if n == preorder[0] {
			pi = i
			break
		}
	}
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:pi+1], inorder[:pi]),
		Right: buildTree(preorder[pi+1:], inorder[pi+1:]),
	}
}

// TreeNode : Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
