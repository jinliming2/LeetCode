package main

// Runtime: 376 ms, faster than 6.67% of Go online submissions for Delete Node in a BST.
// Memory Usage: 234.9 MB, less than 61.33% of Go online submissions for Delete Node in a BST.
// https://leetcode.com/submissions/detail/444593095/
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if root.Left == nil && root.Right == nil {
		return nil
	} else if root.Left == nil {
		return root.Right
	} else if root.Right == nil {
		return root.Left
	} else {
		left := root.Left
		for left.Right != nil {
			left = left.Right
		}
		left.Left = deleteNode(root.Left, left.Val)
		left.Right = root.Right
		return left
	}
	return root
}

// TreeNode : Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
