package main

// Runtime: 12 ms, faster than 80.84% of Go online submissions for Lowest Common Ancestor of a Binary Tree.
// Memory Usage: 7.1 MB, less than 91.19% of Go online submissions for Lowest Common Ancestor of a Binary Tree.
// https://leetcode.com/submissions/detail/445487999/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left, right := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil && right == nil {
		return nil
	}
	if left == nil {
		return right
	}
	return left
}

// TreeNode : Definition for TreeNode.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
