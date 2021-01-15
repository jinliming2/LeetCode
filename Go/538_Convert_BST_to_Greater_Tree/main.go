package main

// Runtime: 208 ms, faster than 72.22% of Go online submissions for Convert BST to Greater Tree.
// Memory Usage: 199.9 MB, less than 55.56% of Go online submissions for Convert BST to Greater Tree.
// https://leetcode.com/submissions/detail/443320222/
func convertBST(root *TreeNode) *TreeNode {
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
