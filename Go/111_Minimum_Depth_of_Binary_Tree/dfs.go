package main

import "math"

// Runtime: 208 ms, faster than 99.59% of Go online submissions for Minimum Depth of Binary Tree.
// Memory Usage: 20.5 MB, less than 45.12% of Go online submissions for Minimum Depth of Binary Tree.
func minDepthDFS(root *TreeNode) int {
	return sub(root, 1, math.MaxInt32)
}

func sub(root *TreeNode, depth, max int) int {
	if root == nil {
		return depth - 1
	}
	if depth == max {
		return depth
	}

	left := sub(root.Left, depth+1, max)
	if left != depth && left < max {
		max = left
	}
	right := sub(root.Right, depth+1, max)

	if left == depth {
		return right
	}
	if right == depth {
		return left
	}
	if left >= right {
		return right
	}
	return left
}
