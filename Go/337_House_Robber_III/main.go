package main

// Runtime: 4 ms, faster than 98.74% of Go online submissions for House Robber III.
// Memory Usage: 5.2 MB, less than 100.00% of Go online submissions for House Robber III.
// https://leetcode.com/submissions/detail/438030027/
func rob(root *TreeNode) int {
	dp0, dp1 := traverse(root)
	if dp0 > dp1 {
		return dp0
	}
	return dp1
}

func traverse(root *TreeNode) (dp0, dp1 int) {
	if root == nil {
		return 0, 0
	}
	left0, left1 := traverse(root.Left)
	right0, right1 := traverse(root.Right)
	leftMax := left0
	if left1 > left0 {
		leftMax = left1
	}
	rightMax := right0
	if right1 > right0 {
		rightMax = right1
	}
	dp0, dp1 = leftMax+rightMax, left0+right0+root.Val
	return
}

// TreeNode : Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
