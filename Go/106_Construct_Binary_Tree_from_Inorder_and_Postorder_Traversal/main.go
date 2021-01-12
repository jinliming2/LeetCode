package main

// Runtime: 20 ms, faster than 67.63% of Go online submissions for Construct Binary Tree from Inorder and Postorder Traversal.
// Memory Usage: 9.7 MB, less than 49.64% of Go online submissions for Construct Binary Tree from Inorder and Postorder Traversal.
// https://leetcode.com/submissions/detail/442071312/
func buildTree(inorder []int, postorder []int) *TreeNode {
	inLength, postLength := len(inorder), len(postorder)
	if inLength == 0 {
		return nil
	}
	top := postorder[postLength-1]
	ii := 0
	for i, n := range inorder {
		if n == top {
			ii = i
			break
		}
	}
	return &TreeNode{
		Val:   top,
		Left:  buildTree(inorder[:ii], postorder[:ii]),
		Right: buildTree(inorder[ii+1:], postorder[ii:postLength-1]),
	}
}

// TreeNode : Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
