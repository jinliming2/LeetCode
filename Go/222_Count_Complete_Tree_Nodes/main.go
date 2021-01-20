package main

// Runtime: 16 ms, faster than 79.46% of Go online submissions for Count Complete Tree Nodes.
// Memory Usage: 6.5 MB, less than 47.03% of Go online submissions for Count Complete Tree Nodes.
// https://leetcode.com/submissions/detail/445496560/
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	hl, hr := 1, 1
	p := root
	for p.Left != nil {
		p = p.Left
		hl++
	}
	p = root
	for p.Right != nil {
		p = p.Right
		hr++
	}
	if hl == hr {
		return (1 << hl) - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

// TreeNode : Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
