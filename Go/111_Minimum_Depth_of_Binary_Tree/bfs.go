package main

// Runtime: 244 ms, faster than 67.34% of Go online submissions for Minimum Depth of Binary Tree.
// Memory Usage: 19.7 MB, less than 73.79% of Go online submissions for Minimum Depth of Binary Tree.
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	current := []*TreeNode{root}
	depth := 1
	for len(current) > 0 {
		n := make([]*TreeNode, 0, 2*len(current))
		for _, node := range current {
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				n = append(n, node.Left)
			}
			if node.Right != nil {
				n = append(n, node.Right)
			}
		}
		depth++
		current = n
	}
	return depth
}

// TreeNode : Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
