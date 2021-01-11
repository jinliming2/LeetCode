package main

// Runtime: 100 ms, faster than 11.25% of Go online submissions for Maximum Binary Tree.
// Memory Usage: 8.8 MB, less than 97.50% of Go online submissions for Maximum Binary Tree.
// https://leetcode.com/submissions/detail/441624892/
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	i, max := 0, nums[0]
	for j, n := range nums[1:] {
		if n > max {
			i, max = j+1, n
		}
	}
	return &TreeNode{
		Val:   max,
		Left:  constructMaximumBinaryTree(nums[:i]),
		Right: constructMaximumBinaryTree(nums[i+1:]),
	}
}

// TreeNode : Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
